// package grpchandler gibs a set of gRPC handlers for ya almost same as the http handlers except
// the original shorten/post url handlers were combined into a single handler since they do the same
// and the difference only exists in the HTTP-style API implementation
package grpchandler

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net"
	"strconv"

	pb "github.com/T-V-N/gourlshortener/internal/grpcHandler/proto"
	"github.com/T-V-N/gourlshortener/internal/storage"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func isValidToken(c string, key string) (bool, error) {
	value, err := hex.DecodeString(c)
	if err != nil {
		return false, err
	}
	signature := value[:32]

	h := hmac.New(sha256.New, []byte(key))
	h.Write(value[32:])
	dst := h.Sum(nil)

	return hmac.Equal(dst, signature), nil
}

func generateToken(key string) (token string, err error) {
	uid, err := auth.GenerateRandom(4)
	if err != nil {
		return "", err
	}

	h := hmac.New(sha256.New, []byte(key))
	h.Write(uid)
	signature := h.Sum(nil)

	cookieVal := append(signature, uid...)

	return hex.EncodeToString(cookieVal), nil
}

func exctractUIDFromCtx(ctx context.Context) (string, error) {
	var uid string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("uid")
		if len(values) > 0 {
			uid = values[0]
			return uid, nil
		}
	}
	return "", status.Error(codes.Internal, "wrong uid")
}

// RPC server manages RPC connections
type RPCServer struct {
	pb.UnimplementedURLShortenerServer
	app *app.App
}

// AuthInterceptor parses auth_token from incoming message and if it's ok allows requests. In case there is no token, creates and returns it to the client.
func InitAuthInterceptor(cfg *config.Config) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var token string
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Internal, "cant parse metadata")
		}

		values := md.Get("auth_token")

		if len(values) == 0 {
			token, err := generateToken(cfg.SecretKey)
			if err != nil {
				return nil, status.Error(codes.Internal, "cant generate token")
			}

			md.Set("auth_token", token)
			grpc.SetHeader(ctx, md)

			md.Set("uid", token[32:])
			return handler(metadata.NewIncomingContext(ctx, md), req)
		}

		token = values[0]
		valid, err := isValidToken(token, cfg.SecretKey)

		if err != nil {
			return nil, status.Error(codes.Internal, "cant parse metadata")
		}

		if !valid {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		return handler(ctx, req)
	}
}

// ShortenURL gets an URL from the message and saves it.
func (rh *RPCServer) ShortenURL(ctx context.Context, in *pb.ShortenURLRequest) (*pb.ShortenURLResponse, error) {
	uid, err := exctractUIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	url, err := rh.app.SaveURL(ctx, in.Url, uid)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, status.Error(codes.AlreadyExists, "already exists")
		}
		return nil, status.Error(codes.Internal, "cant save url")
	}

	var r pb.ShortenURLResponse
	r.Url = url
	return &r, nil
}

// ShortenBatchURL saves a list of URLs
func (rh *RPCServer) ShortenBatchURL(ctx context.Context, in *pb.ShortenBatchURLRequest) (*pb.ShortenBatchURLResponse, error) {
	uid, err := exctractUIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	urls := []storage.BatchURL{}
	for _, u := range in.Urls {
		urls = append(urls, storage.BatchURL{OriginalURL: u.OriginalUrl, CorrelationID: u.CorrelationHash, ShortURL: u.ShortUrl})
	}

	urls, err = rh.app.BatchSaveURL(ctx, urls, uid)

	if err != nil {
		return nil, status.Error(codes.Internal, "cant save url")
	}

	var r pb.ShortenBatchURLResponse
	for _, u := range urls {
		r.Urls = append(r.Urls, &pb.ShortenURL{CorrelationHash: u.CorrelationID, ShortUrl: u.ShortURL})
	}
	return &r, nil
}

// GetURL returns url by hash
func (rh *RPCServer) GetURL(ctx context.Context, in *pb.GetURLRequest) (*pb.GetURLResponse, error) {
	url, err := rh.app.GetURL(ctx, in.UrlHash)

	if err != nil {
		return nil, status.Error(codes.NotFound, "no such url")
	}

	var r pb.GetURLResponse
	r.Url = url.URL
	return &r, nil
}

// DeleteURL schedules deletion of a url
func (rh *RPCServer) DeleteURL(ctx context.Context, in *pb.DeleteListURLRequest) (*pb.Empty, error) {
	uid, err := exctractUIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = rh.app.DeleteListURL(ctx, in.Hashes, uid)
	if err != nil {
		return nil, status.Error(codes.Internal, "deletion error")
	}

	var r pb.Empty
	return &r, nil
}

// DeleteURL schedules deletion of a url
func (rh *RPCServer) DeleteListURL(ctx context.Context, in *pb.DeleteListURLRequest) (*pb.Empty, error) {
	uid, err := exctractUIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = rh.app.DeleteListURL(ctx, in.Hashes, uid)
	if err != nil {
		return nil, status.Error(codes.Internal, "deletion error")
	}

	var r pb.Empty
	return &r, nil
}

// Ping pings storage
func (rh *RPCServer) Ping(ctx context.Context, in *pb.Empty) (*pb.PingResponse, error) {
	err := rh.app.PingStorage(ctx)

	if err != nil {
		return nil, status.Error(codes.Internal, "no such url")
	}

	var r pb.PingResponse
	r.Status = "OK"
	return &r, nil
}

// GetListURL lists url based on the requester uid
func (rh *RPCServer) GetListURL(ctx context.Context, in *pb.Empty) (*pb.GetListURLResponse, error) {
	uid, err := exctractUIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	urls, err := rh.app.GetURLByUID(uid, ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "error getting urls")
	}

	if len(urls) == 0 {
		return nil, status.Error(codes.NotFound, "no urls found")
	}

	var r pb.GetListURLResponse
	for _, u := range urls {
		r.Urls = append(r.Urls, &pb.URLForList{ShortUrl: u.ShortURL, OriginalUrl: u.URL})
	}
	return &r, nil
}

// GetStats returns server stats
func (rh *RPCServer) GetStats(ctx context.Context, in *pb.Empty) (*pb.GetStatsResponse, error) {
	_, trustedNet, err := net.ParseCIDR(rh.app.Config.TrustedSubnet)
	if err != nil {
		return nil, status.Error(codes.Internal, "can't parse cidr")
	}

	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "can't parse ip")
	}

	host, _, err := net.SplitHostPort(p.Addr.String())
	if err != nil {
		return nil, status.Error(codes.Internal, "can't parse ip")
	}

	ip := net.ParseIP(host)
	if ip == nil || !trustedNet.Contains(ip) {
		return nil, status.Error(codes.Unauthenticated, "not from trusted ip")
	}

	users, urls, err := rh.app.GetStats(ctx)

	if err != nil {
		return nil, status.Error(codes.Internal, "can't return stats")
	}
	var r pb.GetStatsResponse
	r.Urls = strconv.Itoa(urls)
	r.Users = strconv.Itoa(users)
	return &r, nil
}

// InitHandler creates handlers for an app
func InitRPCServer(cfg *config.Config, a *app.App) *RPCServer {
	return &RPCServer{pb.UnimplementedURLShortenerServer{}, a}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proto/handler.proto

package handler

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlHash string `protobuf:"bytes,1,opt,name=url_hash,json=urlHash,proto3" json:"url_hash,omitempty"`
}

func (x *GetURLRequest) Reset() {
	*x = GetURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLRequest) ProtoMessage() {}

func (x *GetURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLRequest.ProtoReflect.Descriptor instead.
func (*GetURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{0}
}

func (x *GetURLRequest) GetUrlHash() string {
	if x != nil {
		return x.UrlHash
	}
	return ""
}

type GetURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetURLResponse) Reset() {
	*x = GetURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLResponse) ProtoMessage() {}

func (x *GetURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLResponse.ProtoReflect.Descriptor instead.
func (*GetURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{1}
}

func (x *GetURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetURLResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ShortenURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortenURLRequest) Reset() {
	*x = ShortenURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLRequest) ProtoMessage() {}

func (x *ShortenURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLRequest.ProtoReflect.Descriptor instead.
func (*ShortenURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{2}
}

func (x *ShortenURLRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ShortenURLResponse) Reset() {
	*x = ShortenURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLResponse) ProtoMessage() {}

func (x *ShortenURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLResponse.ProtoReflect.Descriptor instead.
func (*ShortenURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{3}
}

func (x *ShortenURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ShortenURLResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{4}
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterResponse) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl     string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	CorrelationHash string `protobuf:"bytes,2,opt,name=correlation_hash,json=correlationHash,proto3" json:"correlation_hash,omitempty"`
	ShortUrl        string `protobuf:"bytes,3,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{6}
}

func (x *URL) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

func (x *URL) GetCorrelationHash() string {
	if x != nil {
		return x.CorrelationHash
	}
	return ""
}

func (x *URL) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type URLForList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl    string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *URLForList) Reset() {
	*x = URLForList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLForList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLForList) ProtoMessage() {}

func (x *URLForList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLForList.ProtoReflect.Descriptor instead.
func (*URLForList) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{7}
}

func (x *URLForList) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *URLForList) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type ShortenURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationHash string `protobuf:"bytes,1,opt,name=correlation_hash,json=correlationHash,proto3" json:"correlation_hash,omitempty"`
	ShortUrl        string `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *ShortenURL) Reset() {
	*x = ShortenURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURL) ProtoMessage() {}

func (x *ShortenURL) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURL.ProtoReflect.Descriptor instead.
func (*ShortenURL) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{8}
}

func (x *ShortenURL) GetCorrelationHash() string {
	if x != nil {
		return x.CorrelationHash
	}
	return ""
}

func (x *ShortenURL) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type ShortenBatchURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*URL `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *ShortenBatchURLRequest) Reset() {
	*x = ShortenBatchURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenBatchURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenBatchURLRequest) ProtoMessage() {}

func (x *ShortenBatchURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenBatchURLRequest.ProtoReflect.Descriptor instead.
func (*ShortenBatchURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{9}
}

func (x *ShortenBatchURLRequest) GetUrls() []*URL {
	if x != nil {
		return x.Urls
	}
	return nil
}

type ShortenBatchURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls  []*ShortenURL `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	Error string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ShortenBatchURLResponse) Reset() {
	*x = ShortenBatchURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenBatchURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenBatchURLResponse) ProtoMessage() {}

func (x *ShortenBatchURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenBatchURLResponse.ProtoReflect.Descriptor instead.
func (*ShortenBatchURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{10}
}

func (x *ShortenBatchURLResponse) GetUrls() []*ShortenURL {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *ShortenBatchURLResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{11}
}

func (x *PingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users string `protobuf:"bytes,1,opt,name=users,proto3" json:"users,omitempty"`
	Urls  string `protobuf:"bytes,2,opt,name=urls,proto3" json:"urls,omitempty"`
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{12}
}

func (x *GetStatsResponse) GetUsers() string {
	if x != nil {
		return x.Users
	}
	return ""
}

func (x *GetStatsResponse) GetUrls() string {
	if x != nil {
		return x.Urls
	}
	return ""
}

type DeleteListURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashes []string `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *DeleteListURLRequest) Reset() {
	*x = DeleteListURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteListURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteListURLRequest) ProtoMessage() {}

func (x *DeleteListURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteListURLRequest.ProtoReflect.Descriptor instead.
func (*DeleteListURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteListURLRequest) GetHashes() []string {
	if x != nil {
		return x.Hashes
	}
	return nil
}

type GetListURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls  []*URLForList `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	Error string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetListURLResponse) Reset() {
	*x = GetListURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListURLResponse) ProtoMessage() {}

func (x *GetListURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListURLResponse.ProtoReflect.Descriptor instead.
func (*GetListURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{14}
}

func (x *GetListURLResponse) GetUrls() []*URLForList {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *GetListURLResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_handler_proto protoreflect.FileDescriptor

var file_proto_handler_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x22, 0x2a,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x48, 0x61, 0x73, 0x68, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x25, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x3c, 0x0a, 0x12, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x31, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x70, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x21, 0x0a, 0x0c,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x4c, 0x0a, 0x0a, 0x55, 0x52, 0x4c, 0x46, 0x6f,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x54, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x55, 0x52, 0x4c, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x3a, 0x0a, 0x16, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x52,
	0x4c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22, 0x58, 0x0a, 0x17, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x26, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22, 0x2e, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x52, 0x4c, 0x46, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xc7, 0x03, 0x0a,
	0x0c, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x39, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x12, 0x1a, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x54, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55,
	0x52, 0x4c, 0x12, 0x1f, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x2e,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x0e, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x19, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x0e, 0x2e, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x1d, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_handler_proto_rawDescOnce sync.Once
	file_proto_handler_proto_rawDescData = file_proto_handler_proto_rawDesc
)

func file_proto_handler_proto_rawDescGZIP() []byte {
	file_proto_handler_proto_rawDescOnce.Do(func() {
		file_proto_handler_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_handler_proto_rawDescData)
	})
	return file_proto_handler_proto_rawDescData
}

var file_proto_handler_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_handler_proto_goTypes = []interface{}{
	(*GetURLRequest)(nil),           // 0: handler.GetURLRequest
	(*GetURLResponse)(nil),          // 1: handler.GetURLResponse
	(*ShortenURLRequest)(nil),       // 2: handler.ShortenURLRequest
	(*ShortenURLResponse)(nil),      // 3: handler.ShortenURLResponse
	(*Empty)(nil),                   // 4: handler.Empty
	(*RegisterResponse)(nil),        // 5: handler.RegisterResponse
	(*URL)(nil),                     // 6: handler.URL
	(*URLForList)(nil),              // 7: handler.URLForList
	(*ShortenURL)(nil),              // 8: handler.ShortenURL
	(*ShortenBatchURLRequest)(nil),  // 9: handler.ShortenBatchURLRequest
	(*ShortenBatchURLResponse)(nil), // 10: handler.ShortenBatchURLResponse
	(*PingResponse)(nil),            // 11: handler.PingResponse
	(*GetStatsResponse)(nil),        // 12: handler.GetStatsResponse
	(*DeleteListURLRequest)(nil),    // 13: handler.DeleteListURLRequest
	(*GetListURLResponse)(nil),      // 14: handler.GetListURLResponse
}
var file_proto_handler_proto_depIdxs = []int32{
	6,  // 0: handler.ShortenBatchURLRequest.urls:type_name -> handler.URL
	8,  // 1: handler.ShortenBatchURLResponse.urls:type_name -> handler.ShortenURL
	7,  // 2: handler.GetListURLResponse.urls:type_name -> handler.URLForList
	0,  // 3: handler.URLShortener.GetURL:input_type -> handler.GetURLRequest
	2,  // 4: handler.URLShortener.ShortenURL:input_type -> handler.ShortenURLRequest
	9,  // 5: handler.URLShortener.ShortenBatchURL:input_type -> handler.ShortenBatchURLRequest
	4,  // 6: handler.URLShortener.Ping:input_type -> handler.Empty
	4,  // 7: handler.URLShortener.GetStats:input_type -> handler.Empty
	4,  // 8: handler.URLShortener.GetListURL:input_type -> handler.Empty
	13, // 9: handler.URLShortener.DeleteListURL:input_type -> handler.DeleteListURLRequest
	1,  // 10: handler.URLShortener.GetURL:output_type -> handler.GetURLResponse
	3,  // 11: handler.URLShortener.ShortenURL:output_type -> handler.ShortenURLResponse
	10, // 12: handler.URLShortener.ShortenBatchURL:output_type -> handler.ShortenBatchURLResponse
	11, // 13: handler.URLShortener.Ping:output_type -> handler.PingResponse
	12, // 14: handler.URLShortener.GetStats:output_type -> handler.GetStatsResponse
	14, // 15: handler.URLShortener.GetListURL:output_type -> handler.GetListURLResponse
	4,  // 16: handler.URLShortener.DeleteListURL:output_type -> handler.Empty
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_handler_proto_init() }
func file_proto_handler_proto_init() {
	if File_proto_handler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_handler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLForList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURL); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenBatchURLRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenBatchURLResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteListURLRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_handler_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListURLResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_handler_proto_goTypes,
		DependencyIndexes: file_proto_handler_proto_depIdxs,
		MessageInfos:      file_proto_handler_proto_msgTypes,
	}.Build()
	File_proto_handler_proto = out.File
	file_proto_handler_proto_rawDesc = nil
	file_proto_handler_proto_goTypes = nil
	file_proto_handler_proto_depIdxs = nil
}
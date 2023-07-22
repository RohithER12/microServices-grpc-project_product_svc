// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: pkg/pb/product.proto

package pb

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

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Stock int64  `protobuf:"varint,2,opt,name=stock,proto3" json:"stock,omitempty"`
	Price int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CreateProductRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id     int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateProductResponse) Reset() {
	*x = CreateProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResponse) ProtoMessage() {}

func (x *CreateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResponse.ProtoReflect.Descriptor instead.
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateProductResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateProductResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindOneData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Stock int64  `protobuf:"varint,3,opt,name=stock,proto3" json:"stock,omitempty"`
	Price int64  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *FindOneData) Reset() {
	*x = FindOneData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneData) ProtoMessage() {}

func (x *FindOneData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneData.ProtoReflect.Descriptor instead.
func (*FindOneData) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{2}
}

func (x *FindOneData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindOneData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindOneData) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *FindOneData) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type FindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneRequest) Reset() {
	*x = FindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneRequest) ProtoMessage() {}

func (x *FindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneRequest.ProtoReflect.Descriptor instead.
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{3}
}

func (x *FindOneRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   *FindOneData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FindOneResponse) Reset() {
	*x = FindOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneResponse) ProtoMessage() {}

func (x *FindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneResponse.ProtoReflect.Descriptor instead.
func (*FindOneResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindOneResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindOneResponse) GetData() *FindOneData {
	if x != nil {
		return x.Data
	}
	return nil
}

type DecreaseStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId  int64 `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Quantity int64 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *DecreaseStockRequest) Reset() {
	*x = DecreaseStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecreaseStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseStockRequest) ProtoMessage() {}

func (x *DecreaseStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseStockRequest.ProtoReflect.Descriptor instead.
func (*DecreaseStockRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{5}
}

func (x *DecreaseStockRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DecreaseStockRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *DecreaseStockRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type DecreaseStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DecreaseStockResponse) Reset() {
	*x = DecreaseStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecreaseStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseStockResponse) ProtoMessage() {}

func (x *DecreaseStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseStockResponse.ProtoReflect.Descriptor instead.
func (*DecreaseStockResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{6}
}

func (x *DecreaseStockResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DecreaseStockResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ListProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProductsRequest) Reset() {
	*x = ListProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsRequest) ProtoMessage() {}

func (x *ListProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsRequest.ProtoReflect.Descriptor instead.
func (*ListProductsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{7}
}

type ListProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64          `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   []*FindOneData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListProductsResponse) Reset() {
	*x = ListProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResponse) ProtoMessage() {}

func (x *ListProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResponse.ProtoReflect.Descriptor instead.
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{8}
}

func (x *ListProductsResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ListProductsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ListProductsResponse) GetData() []*FindOneData {
	if x != nil {
		return x.Data
	}
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{9}
}

func (x *SearchRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type SearchReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64          `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   []*FindOneData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SearchReponse) Reset() {
	*x = SearchReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReponse) ProtoMessage() {}

func (x *SearchReponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReponse.ProtoReflect.Descriptor instead.
func (*SearchReponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{10}
}

func (x *SearchReponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SearchReponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SearchReponse) GetData() []*FindOneData {
	if x != nil {
		return x.Data
	}
	return nil
}

type SortByPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sort string `protobuf:"bytes,1,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *SortByPriceRequest) Reset() {
	*x = SortByPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortByPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortByPriceRequest) ProtoMessage() {}

func (x *SortByPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortByPriceRequest.ProtoReflect.Descriptor instead.
func (*SortByPriceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{11}
}

func (x *SortByPriceRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type SortByPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64          `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   []*FindOneData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SortByPriceResponse) Reset() {
	*x = SortByPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortByPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortByPriceResponse) ProtoMessage() {}

func (x *SortByPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortByPriceResponse.ProtoReflect.Descriptor instead.
func (*SortByPriceResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{12}
}

func (x *SortByPriceResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SortByPriceResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SortByPriceResponse) GetData() []*FindOneData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_pb_product_proto protoreflect.FileDescriptor

var file_pkg_pb_product_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0x56, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x55, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d,
	0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x20, 0x0a,
	0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x69, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5c, 0x0a, 0x14, 0x44, 0x65,
	0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x45, 0x0a, 0x15, 0x44, 0x65, 0x63, 0x72,
	0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6e, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22,
	0x67, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x28,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x12, 0x53, 0x6f, 0x72, 0x74,
	0x42, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x22, 0x6d, 0x0a, 0x13, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xcb, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e,
	0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61,
	0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x6f, 0x72,
	0x74, 0x42, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_product_proto_rawDescOnce sync.Once
	file_pkg_pb_product_proto_rawDescData = file_pkg_pb_product_proto_rawDesc
)

func file_pkg_pb_product_proto_rawDescGZIP() []byte {
	file_pkg_pb_product_proto_rawDescOnce.Do(func() {
		file_pkg_pb_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_product_proto_rawDescData)
	})
	return file_pkg_pb_product_proto_rawDescData
}

var file_pkg_pb_product_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pkg_pb_product_proto_goTypes = []interface{}{
	(*CreateProductRequest)(nil),  // 0: product.CreateProductRequest
	(*CreateProductResponse)(nil), // 1: product.CreateProductResponse
	(*FindOneData)(nil),           // 2: product.FindOneData
	(*FindOneRequest)(nil),        // 3: product.FindOneRequest
	(*FindOneResponse)(nil),       // 4: product.FindOneResponse
	(*DecreaseStockRequest)(nil),  // 5: product.DecreaseStockRequest
	(*DecreaseStockResponse)(nil), // 6: product.DecreaseStockResponse
	(*ListProductsRequest)(nil),   // 7: product.ListProductsRequest
	(*ListProductsResponse)(nil),  // 8: product.ListProductsResponse
	(*SearchRequest)(nil),         // 9: product.SearchRequest
	(*SearchReponse)(nil),         // 10: product.SearchReponse
	(*SortByPriceRequest)(nil),    // 11: product.SortByPriceRequest
	(*SortByPriceResponse)(nil),   // 12: product.SortByPriceResponse
}
var file_pkg_pb_product_proto_depIdxs = []int32{
	2,  // 0: product.FindOneResponse.data:type_name -> product.FindOneData
	2,  // 1: product.ListProductsResponse.data:type_name -> product.FindOneData
	2,  // 2: product.SearchReponse.data:type_name -> product.FindOneData
	2,  // 3: product.SortByPriceResponse.data:type_name -> product.FindOneData
	0,  // 4: product.ProductService.CreateProduct:input_type -> product.CreateProductRequest
	3,  // 5: product.ProductService.FindOne:input_type -> product.FindOneRequest
	5,  // 6: product.ProductService.DecreaseStock:input_type -> product.DecreaseStockRequest
	7,  // 7: product.ProductService.ListProducts:input_type -> product.ListProductsRequest
	9,  // 8: product.ProductService.Search:input_type -> product.SearchRequest
	11, // 9: product.ProductService.SortByPrice:input_type -> product.SortByPriceRequest
	1,  // 10: product.ProductService.CreateProduct:output_type -> product.CreateProductResponse
	4,  // 11: product.ProductService.FindOne:output_type -> product.FindOneResponse
	6,  // 12: product.ProductService.DecreaseStock:output_type -> product.DecreaseStockResponse
	8,  // 13: product.ProductService.ListProducts:output_type -> product.ListProductsResponse
	10, // 14: product.ProductService.Search:output_type -> product.SearchReponse
	12, // 15: product.ProductService.SortByPrice:output_type -> product.SortByPriceResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_pb_product_proto_init() }
func file_pkg_pb_product_proto_init() {
	if File_pkg_pb_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
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
		file_pkg_pb_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductResponse); i {
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
		file_pkg_pb_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneData); i {
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
		file_pkg_pb_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneRequest); i {
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
		file_pkg_pb_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneResponse); i {
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
		file_pkg_pb_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecreaseStockRequest); i {
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
		file_pkg_pb_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecreaseStockResponse); i {
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
		file_pkg_pb_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductsRequest); i {
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
		file_pkg_pb_product_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductsResponse); i {
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
		file_pkg_pb_product_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_pkg_pb_product_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReponse); i {
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
		file_pkg_pb_product_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortByPriceRequest); i {
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
		file_pkg_pb_product_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortByPriceResponse); i {
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
			RawDescriptor: file_pkg_pb_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_product_proto_goTypes,
		DependencyIndexes: file_pkg_pb_product_proto_depIdxs,
		MessageInfos:      file_pkg_pb_product_proto_msgTypes,
	}.Build()
	File_pkg_pb_product_proto = out.File
	file_pkg_pb_product_proto_rawDesc = nil
	file_pkg_pb_product_proto_goTypes = nil
	file_pkg_pb_product_proto_depIdxs = nil
}

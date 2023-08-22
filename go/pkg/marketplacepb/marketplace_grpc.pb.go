// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package marketplacepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MarketplaceServiceClient is the client API for MarketplaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketplaceServiceClient interface {
	Collections(ctx context.Context, in *CollectionsRequest, opts ...grpc.CallOption) (MarketplaceService_CollectionsClient, error)
	CollectionStats(ctx context.Context, in *CollectionStatsRequest, opts ...grpc.CallOption) (*CollectionStatsResponse, error)
	NFTs(ctx context.Context, in *NFTsRequest, opts ...grpc.CallOption) (MarketplaceService_NFTsClient, error)
	NFTCollectionAttributes(ctx context.Context, in *NFTCollectionAttributesRequest, opts ...grpc.CallOption) (MarketplaceService_NFTCollectionAttributesClient, error)
	Quests(ctx context.Context, in *QuestsRequest, opts ...grpc.CallOption) (MarketplaceService_QuestsClient, error)
	Activity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (MarketplaceService_ActivityClient, error)
	NFTPriceHistory(ctx context.Context, in *NFTPriceHistoryRequest, opts ...grpc.CallOption) (*NFTPriceHistoryResponse, error)
	Banners(ctx context.Context, in *BannersRequest, opts ...grpc.CallOption) (*BannersResponse, error)
	News(ctx context.Context, in *NewsRequest, opts ...grpc.CallOption) (*NewsResponse, error)
	DApps(ctx context.Context, in *DAppsRequest, opts ...grpc.CallOption) (*DAppsResponse, error)
	DAppsGroups(ctx context.Context, in *DAppsGroupsRequest, opts ...grpc.CallOption) (*DAppsGroupsResponse, error)
	SearchNames(ctx context.Context, in *SearchNamesRequest, opts ...grpc.CallOption) (*SearchNamesResponse, error)
	SearchCollections(ctx context.Context, in *SearchCollectionsRequest, opts ...grpc.CallOption) (*SearchCollectionsResponse, error)
}

type marketplaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketplaceServiceClient(cc grpc.ClientConnInterface) MarketplaceServiceClient {
	return &marketplaceServiceClient{cc}
}

func (c *marketplaceServiceClient) Collections(ctx context.Context, in *CollectionsRequest, opts ...grpc.CallOption) (MarketplaceService_CollectionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MarketplaceService_ServiceDesc.Streams[0], "/marketplace.v1.MarketplaceService/Collections", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketplaceServiceCollectionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MarketplaceService_CollectionsClient interface {
	Recv() (*CollectionsResponse, error)
	grpc.ClientStream
}

type marketplaceServiceCollectionsClient struct {
	grpc.ClientStream
}

func (x *marketplaceServiceCollectionsClient) Recv() (*CollectionsResponse, error) {
	m := new(CollectionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *marketplaceServiceClient) CollectionStats(ctx context.Context, in *CollectionStatsRequest, opts ...grpc.CallOption) (*CollectionStatsResponse, error) {
	out := new(CollectionStatsResponse)
	err := c.cc.Invoke(ctx, "/marketplace.v1.MarketplaceService/CollectionStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) NFTs(ctx context.Context, in *NFTsRequest, opts ...grpc.CallOption) (MarketplaceService_NFTsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MarketplaceService_ServiceDesc.Streams[1], "/marketplace.v1.MarketplaceService/NFTs", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketplaceServiceNFTsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MarketplaceService_NFTsClient interface {
	Recv() (*NFTsResponse, error)
	grpc.ClientStream
}

type marketplaceServiceNFTsClient struct {
	grpc.ClientStream
}

func (x *marketplaceServiceNFTsClient) Recv() (*NFTsResponse, error) {
	m := new(NFTsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *marketplaceServiceClient) NFTCollectionAttributes(ctx context.Context, in *NFTCollectionAttributesRequest, opts ...grpc.CallOption) (MarketplaceService_NFTCollectionAttributesClient, error) {
	stream, err := c.cc.NewStream(ctx, &MarketplaceService_ServiceDesc.Streams[2], "/marketplace.v1.MarketplaceService/NFTCollectionAttributes", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketplaceServiceNFTCollectionAttributesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MarketplaceService_NFTCollectionAttributesClient interface {
	Recv() (*NFTCollectionAttributesResponse, error)
	grpc.ClientStream
}

type marketplaceServiceNFTCollectionAttributesClient struct {
	grpc.ClientStream
}

func (x *marketplaceServiceNFTCollectionAttributesClient) Recv() (*NFTCollectionAttributesResponse, error) {
	m := new(NFTCollectionAttributesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *marketplaceServiceClient) Quests(ctx context.Context, in *QuestsRequest, opts ...grpc.CallOption) (MarketplaceService_QuestsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MarketplaceService_ServiceDesc.Streams[3], "/marketplace.v1.MarketplaceService/Quests", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketplaceServiceQuestsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MarketplaceService_QuestsClient interface {
	Recv() (*QuestsResponse, error)
	grpc.ClientStream
}

type marketplaceServiceQuestsClient struct {
	grpc.ClientStream
}

func (x *marketplaceServiceQuestsClient) Recv() (*QuestsResponse, error) {
	m := new(QuestsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *marketplaceServiceClient) Activity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (MarketplaceService_ActivityClient, error) {
	stream, err := c.cc.NewStream(ctx, &MarketplaceService_ServiceDesc.Streams[4], "/marketplace.v1.MarketplaceService/Activity", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketplaceServiceActivityClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MarketplaceService_ActivityClient interface {
	Recv() (*ActivityResponse, error)
	grpc.ClientStream
}

type marketplaceServiceActivityClient struct {
	grpc.ClientStream
}

func (x *marketplaceServiceActivityClient) Recv() (*ActivityResponse, error) {
	m := new(ActivityResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *marketplaceServiceClient) NFTPriceHistory(ctx context.Context, in *NFTPriceHistoryRequest, opts ...grpc.CallOption) (*NFTPriceHistoryResponse, error) {
	out := new(NFTPriceHistoryResponse)
	err := c.cc.Invoke(ctx, "/marketplace.v1.MarketplaceService/NFTPriceHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) Banners(ctx context.Context, in *BannersRequest, opts ...grpc.CallOption) (*BannersResponse, error) {
	out := new(BannersResponse)
	err := c.cc.Invoke(ctx, "/marketplace.v1.MarketplaceService/Banners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) News(ctx context.Context, in *NewsRequest, opts ...grpc.CallOption) (*NewsResponse, error) {
	out := new(NewsResponse)
	err := c.cc.Invoke(ctx, "/marketplace.v1.MarketplaceService/News", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) DApps(ctx context.Context, in *DAppsRequest, opts ...grpc.CallOption) (*DAppsResponse, error) {
	out := new(DAppsResponse)
	err := c.cc.Invoke(ctx, "/marketplace.v1.MarketplaceService/DApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) DAppsGroups(ctx context.Context, in *DAppsGroupsRequest, opts ...grpc.CallOption) (*DAppsGroupsResponse, error) {
	out := new(DAppsGroupsResponse)
	err := c.cc.Invoke(ctx, "/marketplace.v1.MarketplaceService/DAppsGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) SearchNames(ctx context.Context, in *SearchNamesRequest, opts ...grpc.CallOption) (*SearchNamesResponse, error) {
	out := new(SearchNamesResponse)
	err := c.cc.Invoke(ctx, "/marketplace.v1.MarketplaceService/SearchNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) SearchCollections(ctx context.Context, in *SearchCollectionsRequest, opts ...grpc.CallOption) (*SearchCollectionsResponse, error) {
	out := new(SearchCollectionsResponse)
	err := c.cc.Invoke(ctx, "/marketplace.v1.MarketplaceService/SearchCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketplaceServiceServer is the server API for MarketplaceService service.
// All implementations must embed UnimplementedMarketplaceServiceServer
// for forward compatibility
type MarketplaceServiceServer interface {
	Collections(*CollectionsRequest, MarketplaceService_CollectionsServer) error
	CollectionStats(context.Context, *CollectionStatsRequest) (*CollectionStatsResponse, error)
	NFTs(*NFTsRequest, MarketplaceService_NFTsServer) error
	NFTCollectionAttributes(*NFTCollectionAttributesRequest, MarketplaceService_NFTCollectionAttributesServer) error
	Quests(*QuestsRequest, MarketplaceService_QuestsServer) error
	Activity(*ActivityRequest, MarketplaceService_ActivityServer) error
	NFTPriceHistory(context.Context, *NFTPriceHistoryRequest) (*NFTPriceHistoryResponse, error)
	Banners(context.Context, *BannersRequest) (*BannersResponse, error)
	News(context.Context, *NewsRequest) (*NewsResponse, error)
	DApps(context.Context, *DAppsRequest) (*DAppsResponse, error)
	DAppsGroups(context.Context, *DAppsGroupsRequest) (*DAppsGroupsResponse, error)
	SearchNames(context.Context, *SearchNamesRequest) (*SearchNamesResponse, error)
	SearchCollections(context.Context, *SearchCollectionsRequest) (*SearchCollectionsResponse, error)
	mustEmbedUnimplementedMarketplaceServiceServer()
}

// UnimplementedMarketplaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMarketplaceServiceServer struct {
}

func (UnimplementedMarketplaceServiceServer) Collections(*CollectionsRequest, MarketplaceService_CollectionsServer) error {
	return status.Errorf(codes.Unimplemented, "method Collections not implemented")
}
func (UnimplementedMarketplaceServiceServer) CollectionStats(context.Context, *CollectionStatsRequest) (*CollectionStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionStats not implemented")
}
func (UnimplementedMarketplaceServiceServer) NFTs(*NFTsRequest, MarketplaceService_NFTsServer) error {
	return status.Errorf(codes.Unimplemented, "method NFTs not implemented")
}
func (UnimplementedMarketplaceServiceServer) NFTCollectionAttributes(*NFTCollectionAttributesRequest, MarketplaceService_NFTCollectionAttributesServer) error {
	return status.Errorf(codes.Unimplemented, "method NFTCollectionAttributes not implemented")
}
func (UnimplementedMarketplaceServiceServer) Quests(*QuestsRequest, MarketplaceService_QuestsServer) error {
	return status.Errorf(codes.Unimplemented, "method Quests not implemented")
}
func (UnimplementedMarketplaceServiceServer) Activity(*ActivityRequest, MarketplaceService_ActivityServer) error {
	return status.Errorf(codes.Unimplemented, "method Activity not implemented")
}
func (UnimplementedMarketplaceServiceServer) NFTPriceHistory(context.Context, *NFTPriceHistoryRequest) (*NFTPriceHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NFTPriceHistory not implemented")
}
func (UnimplementedMarketplaceServiceServer) Banners(context.Context, *BannersRequest) (*BannersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Banners not implemented")
}
func (UnimplementedMarketplaceServiceServer) News(context.Context, *NewsRequest) (*NewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method News not implemented")
}
func (UnimplementedMarketplaceServiceServer) DApps(context.Context, *DAppsRequest) (*DAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DApps not implemented")
}
func (UnimplementedMarketplaceServiceServer) DAppsGroups(context.Context, *DAppsGroupsRequest) (*DAppsGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DAppsGroups not implemented")
}
func (UnimplementedMarketplaceServiceServer) SearchNames(context.Context, *SearchNamesRequest) (*SearchNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchNames not implemented")
}
func (UnimplementedMarketplaceServiceServer) SearchCollections(context.Context, *SearchCollectionsRequest) (*SearchCollectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCollections not implemented")
}
func (UnimplementedMarketplaceServiceServer) mustEmbedUnimplementedMarketplaceServiceServer() {}

// UnsafeMarketplaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketplaceServiceServer will
// result in compilation errors.
type UnsafeMarketplaceServiceServer interface {
	mustEmbedUnimplementedMarketplaceServiceServer()
}

func RegisterMarketplaceServiceServer(s grpc.ServiceRegistrar, srv MarketplaceServiceServer) {
	s.RegisterService(&MarketplaceService_ServiceDesc, srv)
}

func _MarketplaceService_Collections_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CollectionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketplaceServiceServer).Collections(m, &marketplaceServiceCollectionsServer{stream})
}

type MarketplaceService_CollectionsServer interface {
	Send(*CollectionsResponse) error
	grpc.ServerStream
}

type marketplaceServiceCollectionsServer struct {
	grpc.ServerStream
}

func (x *marketplaceServiceCollectionsServer) Send(m *CollectionsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MarketplaceService_CollectionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).CollectionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplace.v1.MarketplaceService/CollectionStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).CollectionStats(ctx, req.(*CollectionStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_NFTs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NFTsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketplaceServiceServer).NFTs(m, &marketplaceServiceNFTsServer{stream})
}

type MarketplaceService_NFTsServer interface {
	Send(*NFTsResponse) error
	grpc.ServerStream
}

type marketplaceServiceNFTsServer struct {
	grpc.ServerStream
}

func (x *marketplaceServiceNFTsServer) Send(m *NFTsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MarketplaceService_NFTCollectionAttributes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NFTCollectionAttributesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketplaceServiceServer).NFTCollectionAttributes(m, &marketplaceServiceNFTCollectionAttributesServer{stream})
}

type MarketplaceService_NFTCollectionAttributesServer interface {
	Send(*NFTCollectionAttributesResponse) error
	grpc.ServerStream
}

type marketplaceServiceNFTCollectionAttributesServer struct {
	grpc.ServerStream
}

func (x *marketplaceServiceNFTCollectionAttributesServer) Send(m *NFTCollectionAttributesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MarketplaceService_Quests_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QuestsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketplaceServiceServer).Quests(m, &marketplaceServiceQuestsServer{stream})
}

type MarketplaceService_QuestsServer interface {
	Send(*QuestsResponse) error
	grpc.ServerStream
}

type marketplaceServiceQuestsServer struct {
	grpc.ServerStream
}

func (x *marketplaceServiceQuestsServer) Send(m *QuestsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MarketplaceService_Activity_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ActivityRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketplaceServiceServer).Activity(m, &marketplaceServiceActivityServer{stream})
}

type MarketplaceService_ActivityServer interface {
	Send(*ActivityResponse) error
	grpc.ServerStream
}

type marketplaceServiceActivityServer struct {
	grpc.ServerStream
}

func (x *marketplaceServiceActivityServer) Send(m *ActivityResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MarketplaceService_NFTPriceHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NFTPriceHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).NFTPriceHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplace.v1.MarketplaceService/NFTPriceHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).NFTPriceHistory(ctx, req.(*NFTPriceHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_Banners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).Banners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplace.v1.MarketplaceService/Banners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).Banners(ctx, req.(*BannersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_News_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).News(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplace.v1.MarketplaceService/News",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).News(ctx, req.(*NewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_DApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).DApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplace.v1.MarketplaceService/DApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).DApps(ctx, req.(*DAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_DAppsGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DAppsGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).DAppsGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplace.v1.MarketplaceService/DAppsGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).DAppsGroups(ctx, req.(*DAppsGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_SearchNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).SearchNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplace.v1.MarketplaceService/SearchNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).SearchNames(ctx, req.(*SearchNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_SearchCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCollectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).SearchCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketplace.v1.MarketplaceService/SearchCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).SearchCollections(ctx, req.(*SearchCollectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MarketplaceService_ServiceDesc is the grpc.ServiceDesc for MarketplaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketplaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "marketplace.v1.MarketplaceService",
	HandlerType: (*MarketplaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollectionStats",
			Handler:    _MarketplaceService_CollectionStats_Handler,
		},
		{
			MethodName: "NFTPriceHistory",
			Handler:    _MarketplaceService_NFTPriceHistory_Handler,
		},
		{
			MethodName: "Banners",
			Handler:    _MarketplaceService_Banners_Handler,
		},
		{
			MethodName: "News",
			Handler:    _MarketplaceService_News_Handler,
		},
		{
			MethodName: "DApps",
			Handler:    _MarketplaceService_DApps_Handler,
		},
		{
			MethodName: "DAppsGroups",
			Handler:    _MarketplaceService_DAppsGroups_Handler,
		},
		{
			MethodName: "SearchNames",
			Handler:    _MarketplaceService_SearchNames_Handler,
		},
		{
			MethodName: "SearchCollections",
			Handler:    _MarketplaceService_SearchCollections_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Collections",
			Handler:       _MarketplaceService_Collections_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NFTs",
			Handler:       _MarketplaceService_NFTs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NFTCollectionAttributes",
			Handler:       _MarketplaceService_NFTCollectionAttributes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Quests",
			Handler:       _MarketplaceService_Quests_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Activity",
			Handler:       _MarketplaceService_Activity_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "marketplace/v1/marketplace.proto",
}

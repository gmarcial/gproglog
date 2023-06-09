package log_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogClient interface {
	Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error)
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error)
	ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Log_ConsumeStreamClient, error)
	ProduceStream(ctx context.Context, opts ...grpc.CallOption) (Log_ProduceStreamClient, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

var logProduceStreamDesc = &grpc.StreamDesc{
	StreamName: "Produce",
}

func (c *logClient) Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error) {
	out := new(ProduceResponse)
	err := c.cc.Invoke(ctx, "/log.v1.Log/Produce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var logConsumeStreamDesc = &grpc.StreamDesc{
	StreamName: "Consume",
}

func (c *logClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error) {
	out := new(ConsumeResponse)
	err := c.cc.Invoke(ctx, "/log.v1.Log/Consume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var logConsumeStreamStreamDesc = &grpc.StreamDesc{
	StreamName:    "ConsumeStream",
	ServerStreams: true,
}

func (c *logClient) ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Log_ConsumeStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, logConsumeStreamStreamDesc, "/log.v1.Log/ConsumeStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &logConsumeStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Log_ConsumeStreamClient interface {
	Recv() (*ConsumeResponse, error)
	grpc.ClientStream
}

type logConsumeStreamClient struct {
	grpc.ClientStream
}

func (x *logConsumeStreamClient) Recv() (*ConsumeResponse, error) {
	m := new(ConsumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var logProduceStreamStreamDesc = &grpc.StreamDesc{
	StreamName:    "ProduceStream",
	ServerStreams: true,
	ClientStreams: true,
}

func (c *logClient) ProduceStream(ctx context.Context, opts ...grpc.CallOption) (Log_ProduceStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, logProduceStreamStreamDesc, "/log.v1.Log/ProduceStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &logProduceStreamClient{stream}
	return x, nil
}

type Log_ProduceStreamClient interface {
	Send(*ProduceRequest) error
	Recv() (*ProduceResponse, error)
	grpc.ClientStream
}

type logProduceStreamClient struct {
	grpc.ClientStream
}

func (x *logProduceStreamClient) Send(m *ProduceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logProduceStreamClient) Recv() (*ProduceResponse, error) {
	m := new(ProduceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogService is the service API for Log service.
// Fields should be assigned to their respective handler implementations only before
// RegisterLogService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type LogService struct {
	Produce       func(context.Context, *ProduceRequest) (*ProduceResponse, error)
	Consume       func(context.Context, *ConsumeRequest) (*ConsumeResponse, error)
	ConsumeStream func(*ConsumeRequest, Log_ConsumeStreamServer) error
	ProduceStream func(Log_ProduceStreamServer) error
}

func (s *LogService) produce(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProduceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Produce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/log.v1.Log/Produce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Produce(ctx, req.(*ProduceRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *LogService) consume(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Consume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/log.v1.Log/Consume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Consume(ctx, req.(*ConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *LogService) consumeStream(_ interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.ConsumeStream(m, &logConsumeStreamServer{stream})
}
func (s *LogService) produceStream(_ interface{}, stream grpc.ServerStream) error {
	return s.ProduceStream(&logProduceStreamServer{stream})
}

type Log_ConsumeStreamServer interface {
	Send(*ConsumeResponse) error
	grpc.ServerStream
}

type logConsumeStreamServer struct {
	grpc.ServerStream
}

func (x *logConsumeStreamServer) Send(m *ConsumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

type Log_ProduceStreamServer interface {
	Send(*ProduceResponse) error
	Recv() (*ProduceRequest, error)
	grpc.ServerStream
}

type logProduceStreamServer struct {
	grpc.ServerStream
}

func (x *logProduceStreamServer) Send(m *ProduceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logProduceStreamServer) Recv() (*ProduceRequest, error) {
	m := new(ProduceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterLogService registers a service implementation with a gRPC server.
func RegisterLogService(s grpc.ServiceRegistrar, srv *LogService) {
	srvCopy := *srv
	if srvCopy.Produce == nil {
		srvCopy.Produce = func(context.Context, *ProduceRequest) (*ProduceResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Produce not implemented")
		}
	}
	if srvCopy.Consume == nil {
		srvCopy.Consume = func(context.Context, *ConsumeRequest) (*ConsumeResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Consume not implemented")
		}
	}
	if srvCopy.ConsumeStream == nil {
		srvCopy.ConsumeStream = func(*ConsumeRequest, Log_ConsumeStreamServer) error {
			return status.Errorf(codes.Unimplemented, "method ConsumeStream not implemented")
		}
	}
	if srvCopy.ProduceStream == nil {
		srvCopy.ProduceStream = func(Log_ProduceStreamServer) error {
			return status.Errorf(codes.Unimplemented, "method ProduceStream not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "log.v1.Log",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Produce",
				Handler:    srvCopy.produce,
			},
			{
				MethodName: "Consume",
				Handler:    srvCopy.consume,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "ConsumeStream",
				Handler:       srvCopy.consumeStream,
				ServerStreams: true,
			},
			{
				StreamName:    "ProduceStream",
				Handler:       srvCopy.produceStream,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "api/v1/log.proto",
	}

	s.RegisterService(&sd, nil)
}
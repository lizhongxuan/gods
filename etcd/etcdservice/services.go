package etcdservice

import (
	"encoding/hex"
	"fmt"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var clientKeepaliveParams = keepalive.ClientParameters{
	Time:                time.Second * 30,
	Timeout:             time.Second * 20,
	PermitWithoutStream: true,
}

var enforcementPolicy = keepalive.EnforcementPolicy{
	MinTime:             time.Second * 30,
	PermitWithoutStream: true,
}

//服务器的名字
const (
	WPS        = "wps"
	WORDMASTER = "wordmaster"
	ET         = "webet"
	WORD       = "webword"
	WAL        = "wal"
	DATA       = "data"
	NOTIFY     = "notify"
	WECHAT     = "wechat"
)

// grpc收发大小限制
const (
	GrpcRecvMsgMaxSize = 512 * 1024 * 1024
	GrpcSendMsgMaxSize = 512 * 1024 * 1024
)

func UUID() string {
	uuid, _ := uuid.NewV4()
	return hex.EncodeToString(uuid[:])
}

// Dial 创建一个ClientConn，如果服务不可用，返回错误
func Dial(serviceName string) (*grpc.ClientConn, error) {
	balancer := grpc.RoundRobin(defaultResolver)
	conn, err := grpc.Dial(serviceName, grpc.WithInsecure(), grpc.WithBalancer(balancer),
		grpc.WithBlock(), grpc.WithTimeout(5*time.Second), grpc.WithKeepaliveParams(clientKeepaliveParams))
	return conn, err
}

// NewClientConn 创建一个ClientConn
func NewClientConn(serviceName string, opts ...grpc.DialOption) *grpc.ClientConn {
	balancer := grpc.RoundRobin(defaultResolver)

	if len(opts) == 0 {
		opts = make([]grpc.DialOption, 0, 3)
	}
	opts = append(opts,
		grpc.WithInsecure(),
		grpc.WithBalancer(balancer),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(GrpcRecvMsgMaxSize), grpc.MaxCallSendMsgSize(GrpcSendMsgMaxSize)),
		grpc.WithKeepaliveParams(clientKeepaliveParams))

	fmt.Println("serviceName:", serviceName)
	conn, err := grpc.Dial(serviceName, opts...)
	if err != nil {
		fmt.Printf("can't create client conn for %s: %s", serviceName, err)
		panic(err)
	}
	return conn
}

//func WithServerReqID() grpc.ServerOption {
//	var interceptor grpc.UnaryServerInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//		ctx = requestid.FromIncomingContext(ctx)
//		return handler(ctx, req)
//	}
//	return grpc.UnaryInterceptor(interceptor)
//}

func NewServer() *grpc.Server {
	return grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(enforcementPolicy),
		grpc.MaxRecvMsgSize(GrpcRecvMsgMaxSize),
		grpc.MaxSendMsgSize(GrpcSendMsgMaxSize),
	//WithServerReqID()
	)
}

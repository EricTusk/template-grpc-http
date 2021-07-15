package service

import (
	"golang.org/x/net/context"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor is a gRPC server-side interceptor that provides detailed logs for Unary RPCs.
func UnaryServerInterceptor() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if logrus.GetLevel() != logrus.DebugLevel {
			return handler(ctx, req)
		}

		logrus.Debug("===request: ", info.FullMethod, req.(proto.Message).String())
		resp, err := handler(ctx, req)
		// if err != nil {
		// 	logrus.Debug("===response: nil, ", err)
		// 	return resp, err
		// }
		// 	logrus.Debug("===response: ", responseString(resp))
		return resp, err
	}
}

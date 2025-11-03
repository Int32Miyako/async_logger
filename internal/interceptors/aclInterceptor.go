package interceptors

import (
	"async_logger/internal/acl"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AclInterceptor возвращает интерсептор с замыканием на aclData
func AclInterceptor(aclData map[string][]string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		err := checkACL(aclData, ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		// разрешаем выполнение
		return handler(ctx, req)
	}
}

func AclStreamInterceptor(aclData map[string][]string) grpc.StreamServerInterceptor {
	return func(
		srv any,
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		if err := checkACL(aclData, stream.Context(), info.FullMethod); err != nil {
			return err
		}
		return handler(srv, stream)
	}
}

func checkACL(aclData map[string][]string, ctx context.Context, fullMethod string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "no metadata provided")
	}

	consumers := md.Get("consumer")
	if len(consumers) == 0 {
		return status.Error(codes.Unauthenticated, "no consumer provided")
	}

	user := consumers[0]
	if len(user) == 0 {
		return status.Error(codes.Unauthenticated, "no user header provided")
	}

	if !acl.IsUserAllowedForMethod(aclData, user, fullMethod) {
		return status.Errorf(codes.Unauthenticated, "user %s not allowed for %s", user, fullMethod)
	}

	return nil
}

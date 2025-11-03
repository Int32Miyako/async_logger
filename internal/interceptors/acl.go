package interceptors

import (
	"async_logger/internal/acl"
	"context"
	"fmt"

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

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "no metadata provided")
		}

		consumers := md.Get("consumer")
		if len(consumers) == 0 {
			return nil, status.Error(codes.Unauthenticated, "no consumer provided")
		}
		user := consumers[0]

		if len(user) == 0 {
			return nil, status.Error(codes.Unauthenticated, "no user header provided")
		}

		method := info.FullMethod

		if !acl.IsUserAllowedForMethod(aclData, user, method) {
			return nil, status.Errorf(codes.Unauthenticated, "user %s not allowed for %s", user, method)
		}

		fmt.Printf("[ACL] user=%s -> %s ✅\n", user, method)

		// разрешаем выполнение
		return handler(ctx, req)
	}
}

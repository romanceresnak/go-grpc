package interceptors

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/romanceresnak/go-grpc/repos"
	"github.com/romanceresnak/go-grpc/utils"
	"google.golang.org/grpc"
)

func globalRepoInjector(db *xorm.Engine) grpc.UnaryServerInterceptor {
	return grpc.UnaryInterceptor(
		func(
			ctx context.Context,
			req interface{},
			info *grpc.UnaryServerInfo,
			handler grpc.UnaryHandler) (resp interface{}, err error) {
			globalRepo := repos.GlobalRepo(db)
			newContext := utils.SetGlobalRepoOnContext(ctx, globalRepo)

			//Before the request

			//Make the actual request
			res, err := handler(newContext, req)

			//After the request
			return res, err
		})
}

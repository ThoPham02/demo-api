// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"demo-api/service/account/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/login",
				Handler: AccountLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/register",
				Handler: AccountRegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AccessLogMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/info",
					Handler: AccountGetInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/delete",
					Handler: AccountDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/update",
					Handler: AccountUpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}

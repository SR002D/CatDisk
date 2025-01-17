// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"CatDisk/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/mail/code/send/register",
				Handler: MailCodeSendRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/share/basic/detail",
				Handler: ShareBasicDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/detail",
				Handler: UserDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: UserRegisterHandler(serverCtx),
			},
		},
	)
}

package auth

import (
	"log/slog"
	"net/http"
)

type Handler struct {
	Logger  *slog.Logger
	Service *Service
}
type HandlerDeps struct {
	Logger  *slog.Logger
	Service *Service
}

func NewHandler(deps HandlerDeps) *Handler {
	return &Handler{
		Logger:  deps.Logger,
		Service: deps.Service,
	}
}

// login 用户名密码登录
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	
}

// todo 查询
// func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
// 	// 1. 分页请求参数校验
// 	var pageQuery request.PageQuery
// 	pageQuery, err := request.CheckPageQuery(r)
// 	if err != nil {
// 		h.Logger.Error("参数错误", "error", err)
// 		response.Fail(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	// 2.调用 service 层的关于查询todo的业务逻辑
// 	res, err := h.Service.List(r.Context(), ListTodoReq{
// 		Page: pageQuery.Page,
// 		Size: pageQuery.Size,
// 	})
// 	if err != nil {
// 		h.Logger.Error("查询错误", "error", err)
// 		response.Fail(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	// 3. 成功的话返回查询数据
// 	response.Page(w, res.List, pageQuery.Page, pageQuery.Size, res.Total)
// }

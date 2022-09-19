package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gctree/internal/entity"
	"gctree/internal/usecase"
	"gctree/pkg/logger"
)

type appRoutes struct {
	t usecase.CTree
	l logger.Interface
}

func newAppRoutes(handler *gin.RouterGroup, t usecase.CTree, l logger.Interface) {
	r := &appRoutes{t, l}

	h := handler.Group("tree")
	{
		h.POST("/save", r.save)
		h.GET("/list", r.list)
		h.GET("/detail", r.detail)
		h.POST("/delete", r.delete)
	}
}

type appResponse struct {
	ErrCode int         `json:"errcode" example:"0"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

type doSaveRequest struct {
	TreeNode entity.TreeNode `json:"tree_node" bind:"required"`
}

type appSaveData struct {
	MetaIDList []string `json:"meta_id_list"`
}

// @Summary     Save tree
// @Description.markdown save
// @ID          save
// @Tags  	    tree
// @Accept      json
// @Produce     json
// @Param       param body doSaveRequest true "tree to be saved"
// @Success     200 {object} appResponse
// @Failure     400 {object} entity.HTTPError
// @Failure     404 {object} entity.HTTPError
// @Failure     500 {object} entity.HTTPError
// @Router      /tree/save [post]
func (r *appRoutes) save(c *gin.Context) {
	var request doSaveRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - save")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	appIDList, errcode, err := r.t.Save(c, request.TreeNode)
	message := "success"
	if err != nil {
		message = err.Error()
	}

	c.JSON(http.StatusOK, appResponse{ErrCode: errcode, Message: message, Data: appSaveData{MetaIDList: appIDList}})
}

type doListRequest struct {
	MasterID string `form:"master_id" binding:"required" example:"fcb3883b-7847-40d1-a89c-03e70db89fff"`
	PID      string `form:"pid"  example:"fcb3883b-7847-40d1-a89c-03e70db89ccc"`
}

// @Summary     List tree
// @Description.markdown list
// @ID          list
// @Tags  	    tree
// @Accept      json
// @Produce     json
// @Param       master_id query string true "tree master id"
// @Param       pid query string false "tree parent id"
// @Success     200 {object} appResponse
// @Failure     400 {object} entity.HTTPError
// @Failure     404 {object} entity.HTTPError
// @Failure     500 {object} entity.HTTPError
// @Router      /tree/list [get]
func (r *appRoutes) list(c *gin.Context) {
	var request doListRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		r.l.Error(err, "http - v1 - doList")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}
	metaNode, errcode, err := r.t.List(c, request.MasterID, request.PID)
	message := "success"
	if err != nil {
		message = err.Error()
	}

	c.JSON(http.StatusOK, appResponse{ErrCode: errcode, Message: message, Data: metaNode})
}

type doDetailRequest struct {
	MetaID string `form:"meta_id" binding:"required" example:"fcb3883b-7847-40d1-a89c-03e70db89sss"`
}

// @Summary     Show tree detail info
// @Description.markdown detail
// @ID          detail
// @Tags  	    tree
// @Accept      json
// @Produce     json
// @Param       meta_id query string true "tree meta id"
// @Success     200 {object} appResponse
// @Failure     400 {object} entity.HTTPError
// @Failure     404 {object} entity.HTTPError
// @Failure     500 {object} entity.HTTPError
// @Router      /tree/detail [get]
func (r *appRoutes) detail(c *gin.Context) {
	var request doDetailRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		r.l.Error(err, "http - v1 - doDetail")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	meta, errcode, err := r.t.Detail(c, request.MetaID)
	message := "success"
	if err != nil {
		message = err.Error()
	}

	c.JSON(http.StatusOK, appResponse{ErrCode: errcode, Message: message, Data: meta})
}

type doDeleteRequest struct {
	MetaIDList []string `json:"meta_id_list" binding:"required" example:"fcb3883b-7847-40d1-a89c-03e70db89111,fcb3883b-7847-40d1-a89c-03e70db89222,fcb3883b-7847-40d1-a89c-03e70db89333"`
}

// @Summary     Delete tree
// @Description.markdown delete
// @ID          delete
// @Tags  	    tree
// @Accept      json
// @Produce     json
// @Param       param body doDeleteRequest true "tree delete by meta id list"
// @Success     200 {object} appResponse
// @Failure     400 {object} entity.HTTPError
// @Failure     404 {object} entity.HTTPError
// @Failure     500 {object} entity.HTTPError
// @Router      /tree/delete [post]
func (r *appRoutes) delete(c *gin.Context) {
	var request doDeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doDelete")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}
	errcode, err := r.t.Delete(c, request.MetaIDList)
	message := "success"
	if err != nil {
		message = err.Error()
	}

	c.JSON(http.StatusOK, appResponse{ErrCode: errcode, Message: message})
}

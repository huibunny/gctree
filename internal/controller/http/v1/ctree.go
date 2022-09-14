package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"gctree/internal/entity"
	"gctree/internal/usecase"
	"gctree/pkg/logger"
)

type appRoutes struct {
	t usecase.CTree
	l logger.Interface
}

func newAppRoutes(handler *gin.RouterGroup, t usecase.CTree, l logger.Interface, serviceName string) {
	r := &appRoutes{t, l}

	h := handler.Group(strings.Join([]string{"/", serviceName}, ""))
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
	AppIDList []string `json:"app_id_list"`
}

// @Summary     Save
// @Description Save system
// @ID          Save
// @Tags  	    save
// @Accept      json
// @Produce     json
// @Param       request body doSaveRequest true "Save System"
// @Success     200 {object} appResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /user/save [post]
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

	c.JSON(http.StatusOK, appResponse{ErrCode: errcode, Message: message, Data: appSaveData{AppIDList: appIDList}})
}

type doListRequest struct {
	MasterID string `form:"master_id" binding:"required" example:"ksjdflksdjflksdjf"`
	PID      string `form:"pid"  example:"ksjdflksdjflksdjf"`
}

type appListData struct {
	MetaNode entity.MetaNode `json:"meta_node"`
}

// @Summary     List
// @Description List system
// @ID          List
// @Tags  	    List
// @Accept      json
// @Produce     json
// @Success     200 {object} appResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /user/list [post]
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

	c.JSON(http.StatusOK, appResponse{ErrCode: errcode, Message: message, Data: appListData{MetaNode: metaNode}})
}

type doDetailRequest struct {
	MetaID string `form:"meta_id" binding:"required" example:"ksjdflksdjflksdjf"`
}

// @Summary     Detail
// @Description Detail system
// @ID          Detail
// @Tags  	    Detail
// @Accept      form
// @Produce     form
// @Param       request body doDetailRequest true "Detail System"
// @Success     200 {object} appResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /user/detail [post]
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
	MetaIDList []string `json:"meta_id_list" binding:"required"`
}

// @Summary     Delete
// @Description Delete system
// @ID          Delete
// @Tags  	    Delete
// @Accept      json
// @Produce     json
// @Param       request body doDeleteRequest true "Delete System"
// @Success     200 {object} appResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /user/delete [post]
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

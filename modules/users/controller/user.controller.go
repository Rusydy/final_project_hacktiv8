package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"final_project_hacktiv8/helpers"
	"final_project_hacktiv8/modules/users/dto"
	"final_project_hacktiv8/modules/users/service"
)

type ControllerUser interface {
	Create(ctx *gin.Context)
	Login(ctx *gin.Context)
	Update(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type controller struct {
	srv service.ServiceUser
}

func New(srv service.ServiceUser) ControllerUser {
	return &controller{srv}
}

// Create new user
// @Tags         users
// @Summary      Create new user
// @Description  Create new user
// @Accept       json
// @Produce      json
// @Param        data  body      dto.Request                                                true  "data"
// @Success      201   {object}  helpers.BaseResponse{data=dto.Response}                    "CREATED"
// @Failure      400   {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      409   {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "data conflict, like email already exist"
// @Router       /users/register [POST]
func (c *controller) Create(ctx *gin.Context) {
	data := new(dto.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helpers.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	response, err := c.srv.Create(*data)

	if err != nil {
		ctx.JSON(helpers.GetStatusCode(err), helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusCreated, response, nil))
}

// Login user
// @Tags         users
// @Summary      Login user
// @Description  Login user
// @Accept       json
// @Produce      json
// @Param        data  body      dto.RequestLogin                                           true  "data"
// @Success      200   {object}  helpers.BaseResponse{data=dto.ResponseLogin}               "OK"
// @Failure      400   {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      404   {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Record not found"
// @Router       /users/login [POST]
func (c *controller) Login(ctx *gin.Context) {
	data := new(dto.RequestLogin)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helpers.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	response, err := c.srv.Login(*data)

	if err != nil {
		ctx.JSON(helpers.GetStatusCode(err), helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, nil))
}

// Update user
// @Tags         users
// @Summary      Update user
// @Description  Update user
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Param        data           body      dto.ExampleRequestUpdate                                   true  "data"
// @Success      200            {object}  helpers.BaseResponse{data=dto.Response}                    "OK"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      401            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Unauthorization"
// @Router       /users [PUT]
func (c *controller) Update(ctx *gin.Context) {
	data := new(dto.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helpers.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	id := ctx.MustGet("user_id")

	data.ID = id.(uint)

	response, err := c.srv.Update(*data)

	if err != nil {
		ctx.JSON(helpers.GetStatusCode(err), helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, nil))
}

// DeleteByID user
// @Tags         users
// @Summary      Delete user
// @Description  Delete user
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                      true  "Bearer + user token"
// @Success      200            {object}  helpers.BaseResponse{data=dto.ExampleResponseDelete}  "OK"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}   "Bad Request"
// @Failure      404            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}   "Not Found"
// @Failure      401            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}   "Unauthorization"
// @Router       /users [DELETE]
func (c *controller) DeleteByID(ctx *gin.Context) {
	id := ctx.MustGet("user_id")

	err := c.srv.DeleteByID(id.(uint))

	if err != nil {
		ctx.JSON(helpers.GetStatusCode(err), helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, map[string]interface{}{"message": "your account has been successfully deleted"}, nil))
}

package handler

import (
	"fmt"
	"net/http"
	"strconv"

	errors "github.com/ayyaa/todo-services/lib/customerrors"
	utils "github.com/ayyaa/todo-services/lib/validator"
	"github.com/ayyaa/todo-services/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (c *Server) GetListByID(ctx echo.Context) error {
	var resp models.ErrorResponse

	// convert string to uint
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	// get service list by id
	list, err := c.Service.GetListByID(ctx.Request().Context(), uint(id))
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	return ctx.JSON(http.StatusOK, models.SuccessResponse{
		Message: fmt.Sprintf("succes get list by id = %d", id),
		Data:    list,
	})
}

func (c *Server) CreateList(ctx echo.Context) error {

	var resp models.ErrorResponse

	params := models.ListRequest{}

	// get file from form
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["file"]
	params.File = files

	// binding param
	err = ctx.Bind(&params)
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	// validate struct
	err = utils.Validate().Struct(params)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			resp.Message = utils.GetValidatorErrMsg(err.(validator.ValidationErrors))
			mapTest := utils.ValidateErrToMapString(ve)
			resp.Data = &mapTest
			return ctx.JSON(http.StatusBadRequest, resp)
		}
	}

	// iniatil new param
	paramsList := models.AdditionalRequest{
		ListRequest: params,
	}

	//  create list
	createdList, err := c.Service.CreateList(ctx.Request().Context(), &paramsList)
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	return ctx.JSON(http.StatusCreated, models.SuccessResponse{
		Message: fmt.Sprintf("%s %d", "succes create list by id = ", createdList.ID),
		Data:    createdList,
	})

}

func (c *Server) DeleteList(ctx echo.Context) error {
	var resp models.ErrorResponse

	// convert string to uint
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	// get service list by id
	err = c.Service.DeleteList(ctx.Request().Context(), uint(id))
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	return ctx.JSON(http.StatusOK, models.SuccessResponse{
		Message: fmt.Sprintf("succes deleted list by id = %d", id),
	})
}

func (c *Server) EditList(ctx echo.Context) error {
	var resp models.ErrorResponse
	params := models.ListRequestEdit{}

	// convert string to uint
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	// get file from form
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["file"]
	params.File = files

	// binding param
	err = ctx.Bind(&params)
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	// validate struct
	err = utils.Validate().Struct(params)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			resp.Message = utils.GetValidatorErrMsg(err.(validator.ValidationErrors))
			mapTest := utils.ValidateErrToMapString(ve)
			resp.Data = &mapTest
			return ctx.JSON(http.StatusBadRequest, resp)
		}
	}

	// iniatil new param
	paramsList := models.AdditionalRequestEdit{
		ListRequestEdit: params,
	}

	//  create list
	updatedList, err := c.Service.EditList(ctx.Request().Context(), &paramsList, uint(id))
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	return ctx.JSON(http.StatusCreated, models.SuccessResponse{
		Message: fmt.Sprintf("%s %d", "succes edit list by id = ", updatedList.ID),
		Data:    updatedList,
	})

}

func (c *Server) GetLists(ctx echo.Context) error {
	var resp models.ErrorResponse

	queryParams := models.ParamRequest{}

	// binding param
	err := ctx.Bind(&queryParams)
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	// validate struct
	err = utils.Validate().Struct(queryParams)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			resp.Message = utils.GetValidatorErrMsg(err.(validator.ValidationErrors))
			mapTest := utils.ValidateErrToMapString(ve)
			resp.Data = &mapTest
			return ctx.JSON(http.StatusBadRequest, resp)
		}
	}

	if queryParams.OrderBy == "" {
		queryParams.OrderBy = "title"
	}

	if queryParams.Page == 0 {
		queryParams.Page = 1
	}

	if queryParams.Size == 0 {
		queryParams.Size = 10
	}

	// if bool(queryParams.Preload) {
	queryParams.Preload = bool(queryParams.Preload)
	// }

	// get service all list
	list, err := c.Service.GetLists(ctx.Request().Context(), queryParams)
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	return ctx.JSON(http.StatusOK, models.SuccessResponse{
		Message: "succes get all list",
		Data:    list,
	})
}

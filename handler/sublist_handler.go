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

func (c *Server) GetSubListByID(ctx echo.Context) error {
	var resp models.ErrorResponse
	fmt.Print(ctx.Param("id"))
	// convert string to uint
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	// get service sub list by id
	list, err := c.Service.GetSubListByID(ctx.Request().Context(), uint(id))
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	return ctx.JSON(http.StatusOK, models.SuccessResponse{
		Message: fmt.Sprintf("succes get sublist by id = %d", id),
		Data:    list,
	})
}

func (c *Server) CreateSubList(ctx echo.Context) error {
	var resp models.ErrorResponse

	params := models.AdditionalRequest{}

	// get file from form
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["file"]
	params.File = files

	// binding struct
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

	// create sublist
	createdSubList, err := c.Service.CreateList(ctx.Request().Context(), &params)
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	return ctx.JSON(http.StatusCreated, models.SuccessResponse{
		Message: fmt.Sprintf("succes create sub list by id =  %d", createdSubList.ID),
		Data:    createdSubList,
	})

}

func (c *Server) DeleteSubList(ctx echo.Context) error {
	var resp models.ErrorResponse

	// convert string to uint
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	// get service list by id
	err = c.Service.DeleteSubList(ctx.Request().Context(), uint(id))
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	return ctx.JSON(http.StatusOK, models.SuccessResponse{
		Message: fmt.Sprintf("succes deleted sublist by id = %d", id),
	})
}

func (c *Server) EditSubList(ctx echo.Context) error {
	var resp models.ErrorResponse
	params := models.AdditionalRequestEdit{}

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

	// binding struct
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

	// create sublist
	createdSubList, err := c.Service.EditSubList(ctx.Request().Context(), &params, uint(id))
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	return ctx.JSON(http.StatusCreated, models.SuccessResponse{
		Message: fmt.Sprintf("succes edit sub list by id =  %d", createdSubList.ID),
		Data:    createdSubList,
	})

}

func (c *Server) GetSubLists(ctx echo.Context) error {
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
	list, err := c.Service.GetSubLists(ctx.Request().Context(), queryParams)
	if err != nil {
		resp.Message = err.Error()
		return ctx.JSON(errors.GetStatusCode(err), resp)
	}

	return ctx.JSON(http.StatusOK, models.SuccessResponse{
		Message: "succes get all list",
		Data:    list,
	})
}

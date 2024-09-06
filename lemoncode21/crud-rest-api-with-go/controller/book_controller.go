package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-golang/data/request"
	"rest-api-golang/data/response"
	"rest-api-golang/helper"
	"rest-api-golang/service"
	"strconv"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{BookService: bookService}
}

func (controller *BookController) Create(write http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestBody(requests, &bookCreateRequest)
	controller.BookService.Create(requests.Context(), bookCreateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "success",
		Data:   nil,
	}

	helper.WriteResponseBody(write, webResponse)
}

func (controller *BookController) Update(write http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestBody(requests, &bookUpdateRequest)

	bookId := params.ByName("id")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id
	controller.BookService.Update(requests.Context(), bookUpdateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "success",
		Data:   nil,
	}

	helper.WriteResponseBody(write, webResponse)
}

func (controller *BookController) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("id")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(request.Context(), id)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "success",
		Data:   nil,
	}

	helper.WriteResponseBody(write, webResponse)
}

func (controller *BookController) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	result := controller.BookService.FindAll(request.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "success",
		Data:   result,
	}

	helper.WriteResponseBody(write, webResponse)
}

func (controller *BookController) FindById(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("id")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	result := controller.BookService.FindById(request.Context(), id)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "success",
		Data:   result,
	}

	helper.WriteResponseBody(write, webResponse)
}

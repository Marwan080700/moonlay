package routes

import (
	"moonlay/handlers"
	"moonlay/pkg/middleware"
	"moonlay/pkg/mysql"
	"moonlay/repository"

	"github.com/labstack/echo/v4"
)

func TodoRoutes(e *echo.Group){
	r := repository.RepositoryTodo(mysql.DB)
	h := handlers.HandlerTodo(r)

	// get all todo
	e.GET("/todos", h.FindTodos)

	// Todolist (datalist)------------------
	e.GET("/list/:id", h.GetDataList)
	e.POST("/list", middleware.UploadFiles(h.CreateDataList))
	e.PUT("/list/:id",middleware.UploadFiles(h.UpdateDataList))
	e.DELETE("/list/:id", h.DeleteDataList)

	//Todolist (subdatalist)------------------------
	e.GET("/list/:id/subdatalist", h.GetSubdatalistByDataListID)
	e.GET("/list/:dataListID/subdatalist/:id", h.GetSubDataListByID)
	e.POST("/list/:id/subdatalist", middleware.UploadFiles(h.CreateSubDataList))
	e.PUT("/subdatalist/:id",middleware.UploadFiles(h.UpdateSubDataListByID))
	e.DELETE("/subdatalist/:id", h.DeleteSubDataListByID)
}
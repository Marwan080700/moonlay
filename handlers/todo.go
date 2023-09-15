package handlers

import (
	"fmt"
	sublist "moonlay/dto/Sublist"
	datalist "moonlay/dto/dataList"
	"moonlay/dto/result"
	"moonlay/models"
	"moonlay/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5000/uploads/"

type handlerTodo struct {
	TodoRepository repository.TodoRepository
}

func HandlerTodo(TodoRepository repository.TodoRepository) *handlerTodo{
	return &handlerTodo{TodoRepository}
}

func (h *handlerTodo) FindTodos(c echo.Context) error{
	todos, err := h.TodoRepository.FindTodos()

	if err != nil{
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Status: "Error",
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{
		Status: "Success",
		Data: todos,
	})
}


// Todo datalist-----------------------------------------------------------------------------------

//get data datalist todo

func (h *handlerTodo) GetDataList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	datalist,err := h.TodoRepository.GetDataList(id)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	datalist.File = path_file + datalist.File

	return c.JSON(http.StatusOK, result.SuccessResult{
		Status: "Success",
		Data: datalist})
}

// create datalist todo
func (h *handlerTodo) CreateDataList(c echo.Context) error {
    title := c.FormValue("title")
    description := c.FormValue("description")
    dataFiles := c.Get("dataFiles").([]string)

    // Define dataList outside of the loop
    var dataList models.DataList

    for _, dataFile := range dataFiles {
        fmt.Println("this is data file", dataFile)

        request := datalist.DataListRequest{
            Title:       title,
            Description: description,
            File:        dataFile,
        }

        validation := validator.New()
        err := validation.Struct(request)

        if err != nil {
            return c.JSON(http.StatusInternalServerError, result.ErrorResult{
                Status:  "Failed",
                Message: err.Error(),
            })
        }

        dataList = models.DataList{
            Title:       request.Title,
            Description: request.Description,
            File:        request.File,
        }

        dataList, err = h.TodoRepository.CreateDataList(dataList)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, result.ErrorResult{
                Status:  "Failed",
                Message: err.Error(),
            })
        }

        dataList, err = h.TodoRepository.GetDataList(dataList.ID)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, result.ErrorResult{
                Status:  "Failed",
                Message: err.Error(),
            })
        }
    }

    response := result.SuccessResult{Status: "Success", Data: dataList}
    return c.JSON(http.StatusOK, response)
}

// update datalist todo
func (h *handlerTodo) UpdateDataList(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, result.ErrorResult{
            Status:  "Failed",
            Message: "Invalid datalist ID",
        })
    }

    existingDataList, err := h.TodoRepository.GetDataList(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, result.ErrorResult{
            Status:  "Failed",
            Message: "Failed to fetch existing datalist",
        })
    }

    var updatedData struct {
        Title       string `json:"title" form:"title"`
        Description string `json:"description" form:"description"`
        File        string `json:"file" form:"file"`
    }

    if err := c.Bind(&updatedData); err != nil {
        return c.JSON(http.StatusBadRequest, result.ErrorResult{
            Status:  "Failed",
            Message: "Invalid request data",
        })
    }

    if updatedData.Title != "" {
        existingDataList.Title = updatedData.Title
    }

    if updatedData.Description != "" {
        existingDataList.Description = updatedData.Description
    }

    if updatedData.File != "" {
        existingDataList.File = updatedData.File
    }


    updatedDataList, err := h.TodoRepository.UpdateDataList(existingDataList)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, result.ErrorResult{
            Status:  "Failed",
            Message: "Failed to update datalist",
        })
    }

    return c.JSON(http.StatusOK, result.SuccessResult{
        Status: "Success",
        Data:   updatedDataList,
    })
}

// delete datalist todo
func (h *handlerTodo) DeleteDataList(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, result.ErrorResult{
            Status:  "Failed",
            Message: "Invalid datalist ID",
        })
    }

	err = h.TodoRepository.DeleteAllSubDataListsByDataListID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Status:  "Failed",
			Message: "Failed to delete associated subdatalists",
		})
	}
	
	// Now you can safely delete the datalist
	err = h.TodoRepository.DeleteDataListByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{
			Status:  "Failed",
			Message: "Failed to delete datalist",
		})
	}

    return c.JSON(http.StatusOK, result.SuccessResult{
        Status: "Success",
        Message: "DataList and associated SubDataLists deleted successfully",
    })
}




// Todo SubDataList---------------------------------------------------------------------------------

// Get all subdatalist items by datalist ID
func (h *handlerTodo) GetSubdatalistByDataListID(c echo.Context) error {
    dataListID, err := strconv.Atoi(c.Param("id")) // Use "id" as the parameter name
    if err != nil {
        return c.JSON(http.StatusBadRequest, result.ErrorResult{
            Status:  "Failed",
            Message: "Invalid datalist ID",
        })
    }

    // Use the repository to fetch all subdatalist items by datalist ID
    subdatalist, err := h.TodoRepository.GetAllSubDataListByDataListID(dataListID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, result.ErrorResult{
            Status:  "Failed",
            Message: err.Error(),
        })
    }

    // Append path_file to the File field of each subdatalist item
    for i := range subdatalist {
        subdatalist[i].File = path_file + subdatalist[i].File
    }

    return c.JSON(http.StatusOK, result.SuccessResult{
        Status: "Success",
        Data:   subdatalist,
    })
}


// get subdatalist 

// Get a single subdatalist item by its ID
func (h *handlerTodo) GetSubDataListByID(c echo.Context) error {
    subDataListID, err := strconv.Atoi(c.Param("id")) // Use "id" as the parameter name
    if err != nil {
        return c.JSON(http.StatusBadRequest, result.ErrorResult{
            Status:  "Failed",
            Message: "Invalid subdatalist ID",
        })
    }

    // Use the repository to fetch the subdatalist item by its ID
    subDataList, err := h.TodoRepository.GetSubDataListByID(subDataListID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, result.ErrorResult{
            Status:  "Failed",
            Message: err.Error(),
        })
    }

	subDataList.File = path_file + subDataList.File

    return c.JSON(http.StatusOK, result.SuccessResult{
        Status: "Success",
        Data:   subDataList,
    })
}


// create subdatalist todo
func (h *handlerTodo) CreateSubDataList(c echo.Context) error {
    title := c.FormValue("title")
    description := c.FormValue("description")
    dataFiles := c.Get("dataFiles").([]string)
    dataListID, err := strconv.Atoi(c.Param("id")) 

    if err != nil {
        return c.JSON(http.StatusBadRequest, result.ErrorResult{
            Status:  "Failed",
            Message: "Invalid datalist ID",
        })
    }


    var subDataList models.DataSublist

    for _, dataFile := range dataFiles {
        fmt.Println("this is data file", dataFile)

        request := sublist.SubListRequest{
            Title:       title,
            Description: description,
            File:        dataFile,
        }

        validation := validator.New()
        err := validation.Struct(request)

        if err != nil {
            return c.JSON(http.StatusBadRequest, result.ErrorResult{
                Status:  "Failed",
                Message: err.Error(),
            })
        }

        subDataList = models.DataSublist{
            Title:       request.Title,
            Description: request.Description,
            File:        request.File,
            DataListID:  dataListID, 
        }

        subDataList, err = h.TodoRepository.CreateSubDataList(subDataList)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, result.ErrorResult{
                Status:  "Failed",
                Message: err.Error(),
            })
        }

        subDataList, err = h.TodoRepository.GetSubDataListByID(subDataList.ID)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, result.ErrorResult{
                Status:  "Failed",
                Message: err.Error(),
            })
        }
    }

    response := result.SuccessResult{Status: "Success", Data: subDataList}
    return c.JSON(http.StatusOK, response)
}

// update subdatalist todo
func (h *handlerTodo) UpdateSubDataListByID(c echo.Context) error {
    subDataListID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, result.ErrorResult{
            Status:  "Failed",
            Message: "Invalid subdatalist ID",
        })
    }

    existingSubDataList, err := h.TodoRepository.GetSubDataListByID(subDataListID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, result.ErrorResult{
            Status:  "Failed",
            Message: "Failed to fetch existing subdatalist",
        })
    }

    var updatedData struct {
        Title       string `json:"title" form:"title"`
        Description string `json:"description" form:"description"`
        File        string `json:"file" form:"file"`
    }

    if err := c.Bind(&updatedData); err != nil {
        return c.JSON(http.StatusBadRequest, result.ErrorResult{
            Status:  "Failed",
            Message: "Invalid request data",
        })
    }

    if updatedData.Title != "" {
        existingSubDataList.Title = updatedData.Title
    }

    if updatedData.Description != "" {
        existingSubDataList.Description = updatedData.Description
    }

    if updatedData.File != "" {
        existingSubDataList.File = updatedData.File
    }

    // Update the existingSubDataList in the database
    updatedSubDataList, err := h.TodoRepository.UpdateSubDataList(existingSubDataList)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, result.ErrorResult{
            Status:  "Failed",
            Message: "Failed to update subdatalist",
        })
    }

    return c.JSON(http.StatusOK, result.SuccessResult{
        Status: "Success",
        Data:   updatedSubDataList,
    })
}

//detele subdatalist todo
// Delete a subdatalist by ID
func (h *handlerTodo) DeleteSubDataListByID(c echo.Context) error {
    subDataListID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, result.ErrorResult{
            Status:  "Failed",
            Message: "Invalid subdatalist ID",
        })
    }

    // Use the repository to delete the subdatalist by ID
    err = h.TodoRepository.DeleteSubDataListByID(subDataListID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, result.ErrorResult{
            Status:  "Failed",
            Message: "Failed to delete subdatalist",
        })
    }

    return c.JSON(http.StatusOK, result.SuccessResult{
        Status:  "Success",
        Message: "SubDataList deleted successfully",
    })
}

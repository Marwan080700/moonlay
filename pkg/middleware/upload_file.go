package middleware

import (
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

func UploadFiles(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		files := form.File["file"]
		if len(files) == 0 {
			return c.JSON(http.StatusBadRequest, "No files uploaded.")
		}

		var fileNames []string

		for _, file := range files {
			ext := strings.ToLower(filepath.Ext(file.Filename))
			if ext != ".txt" && ext != ".pdf" {
				return c.JSON(http.StatusBadRequest, "Invalid file format. Only .txt and .pdf files are allowed.")
			}

			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer src.Close()

			tempFile, err := ioutil.TempFile("uploads", "file-*.pdf")
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer tempFile.Close()

			if _, err = io.Copy(tempFile, src); err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			data := tempFile.Name()
			filename := data[8:] 
			fileNames = append(fileNames, filename)
		}

		c.Set("dataFiles", fileNames)
		return next(c)
	}
}

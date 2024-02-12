package service

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/labstack/echo/v4"
)

func Filehandler(c echo.Context) error {
    folderPath := "/home/mohammedyameen/Downloads"

    files, err := ioutil.ReadDir(folderPath)
    if err != nil {
        return err
    }

    var pdfFiles []os.FileInfo
    for _, file := range files {
        if !file.IsDir() && filepath.Ext(file.Name()) == ".pdf" {
            pdfFiles = append(pdfFiles, file)
        }
    }

    sort.Slice(pdfFiles, func(i, j int) bool {
        return pdfFiles[i].ModTime().After(pdfFiles[j].ModTime())
    })

    latestPDF := pdfFiles[0]

    file, err := os.Open(filepath.Join(folderPath, latestPDF.Name()))
    if err != nil {
        return err
    }
    defer file.Close()

	// c.Response().Header().Set("Content-Disposition", "attachment; filename=abc.pdf")
    //c.Response().Header().Set("Content-Disposition", "attachment; filename="+latestPDF.Name())

    c.Response().Header().Set("Content-Type", "application/pdf")

    c.Response().Header().Set("Content-Disposition", "inline; filename="+latestPDF.Name())

    
    if _, err := io.Copy(c.Response().Writer, file); err != nil {
        return err
    }
    return nil
}

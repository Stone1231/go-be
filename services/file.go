package services

import (
	"io"
	"log"
	"mime/multipart"
	"os"
)

type FileService struct {
}

var File = new(FileService)

func (self *FileService) Upload(file *multipart.File, filename string, folder string) (fname string, err error) {

	os.MkdirAll("static/"+folder, os.ModePerm)

	out, err := os.Create("static/" + folder + "/" + filename)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, *file)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return filename, err
}

func (self *FileService) Clear(folder string) error {
	return os.RemoveAll("static/" + folder)
}

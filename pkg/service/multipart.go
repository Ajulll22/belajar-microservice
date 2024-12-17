package service

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

func ForwardFilesToService(url string, files []*multipart.FileHeader) (*http.Response, error) {
	// Buat buffer untuk body request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Tambahkan setiap file ke form-data
	for _, fileHeader := range files {
		// Buka file
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Tambahkan file ke form-data
		part, err := writer.CreateFormFile("files", fileHeader.Filename)
		if err != nil {
			return nil, err
		}

		// Salin isi file ke form-data
		_, err = io.Copy(part, file)
		if err != nil {
			return nil, err
		}
	}

	// Tutup writer untuk menyelesaikan multipart form
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	// Buat request POST ke service media
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	// Tambahkan header Content-Type
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Kirim request menggunakan HTTP client
	client := &http.Client{}
	return client.Do(req)
}

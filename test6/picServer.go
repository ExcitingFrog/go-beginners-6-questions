package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"bytes"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"strings"
)

const (
	UPLOAD_DIR = "/home/vagrant/go/src/train1/test6"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		str := `
			<html>
			<head>
			<meta charset="utf-8">
			<title>Upload</title>
			</head>
			<body>
			<form method="POST" action="/upload" enctype="multipart/form-data">
			Choose an image to upload: <input name="image" type="file" />
			<input type="submit" value="Upload" />
			</form>
			</body>
			</html>`

		io.WriteString(w, str)
	}

	// 处理图片上传
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer t.Close()

		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}

}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	srcImage := UPLOAD_DIR + "/" + imageId
	if exists := isExists(srcImage); !exists {
		http.NotFound(w, r)
		return
	}
	dstImage, err := greyImage(srcImage)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, dstImage)
}

func greyImage(filepath string) (newPath string, err error) {
	file, _ := ioutil.ReadFile(filepath)
	buf := bytes.NewBuffer(file)
	img, _, _ := image.Decode(buf)
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := img.At(i, j)
			_, g, _, a := colorRgb.RGBA()
			g_uint8 := uint8(g >> 8)
			a_uint8 := uint8(a >> 8)
			newRgba.SetRGBA(i, j, color.RGBA{g_uint8, g_uint8, g_uint8, a_uint8})
		}
	}

	fullName := path.Base(filepath)
	fileSuffix := path.Ext(fullName)
	filePrefix := fullName[0 : len(fullName)-len(fileSuffix)]
	newPath = UPLOAD_DIR + "/" + filePrefix + "_grey" + fileSuffix
	newFile, err := os.Create(newPath)
	if err != nil {
		return newPath, err
	}
	defer newFile.Close()
	encode(fullName, newFile, newRgba)
	return newPath, nil
}

func encode(inputName string, file *os.File, rgba *image.RGBA) {
	if strings.HasSuffix(inputName, "jpg") || strings.HasSuffix(inputName, "jpeg") {
		jpeg.Encode(file, rgba, nil)
	} else if strings.HasSuffix(inputName, "png") {
		png.Encode(file, rgba)
	} else if strings.HasSuffix(inputName, "gif") {
		gif.Encode(file, rgba, nil)
	} else {
		fmt.Errorf("不支持的图片格式")
	}
}

func main() {
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

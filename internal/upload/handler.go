package upload

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 5<<20)
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		http.Error(w, "Too large", http.StatusBadRequest)
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}
	ext := filepath.Ext(header.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		http.Error(w, "only jpg and png allowed", http.StatusBadRequest)
		return
	}
	buffer := make([]byte, 512)
	file.Read(buffer)
	file.Seek(0, 0)
	fileType := http.DetectContentType(buffer)
	if fileType != "image/jpeg" && fileType != "image/png" {
		http.Error(w, "invalid image type", http.StatusBadRequest)
		return
	}
	defer file.Close()
	uniqueName := getUniqueFilename(header.Filename)
	savePath := filepath.Join("./uploads", uniqueName)
	dst, err := os.Create(savePath)
	if err != nil {
		http.Error(w, "cannot save the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error while saving file", http.StatusInternalServerError)
		return
	}

	fileOndisk, err := os.Open(savePath)
	if err != nil {
		http.Error(w, "cannot open saved file", http.StatusInternalServerError)
		return
	}
	defer fileOndisk.Close()
	var img image.Image
	var format string
	if ext == ".jpg" || ext == ".jpeg" {
		img, err = jpeg.Decode(fileOndisk)
		format = "jpeg"
	} else if ext == ".png" {
		img, err = png.Decode(fileOndisk)
		format = "png"
	}
	if err != nil {
		http.Error(w, "cannot decode image", http.StatusInternalServerError)
		return
	}
	err = createThumbnail(img, format, header.Filename)
	if err != nil {
		http.Error(w, "cannot create thumbnail", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Upload successful")
	fmt.Fprintln(w, "File name:", header.Filename)
	fmt.Fprintln(w, "File size:", header.Size)
}
func createThumbnail(src image.Image, format, filename string) error {
	const width = 150
	const height = 150
	thumbnail := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			srcX := x * src.Bounds().Dx() / width
			srcY := y * src.Bounds().Dy() / height
			thumbnail.Set(x, y, src.At(srcX, srcY))
		}
	}
	thumbPath := filepath.Join("./thumbnails", "thumb_"+filename)
	outFile, err := os.Create(thumbPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	if format == "jpeg" {
		err = jpeg.Encode(outFile, thumbnail, nil)
	} else if format == "png" {
		err = png.Encode(outFile, thumbnail)
	}
	return err
}
func getUniqueFilename(filename string) string {
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]
	newName := filename
	i := 1
	for {
		if _, err := os.Stat(filepath.Join("./uploads", newName)); os.IsNotExist(err) {
			return newName
		}
		newName = fmt.Sprintf("%s_%d%s", name, i, ext)
		i++
	}
}

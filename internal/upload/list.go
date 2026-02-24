package upload

import (
	"encoding/json"
	"net/http"
	"os"
)

type ImageInfo struct {
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("uploads")
	if err != nil {
		http.Error(w, "cannot read uploads folder", http.StatusInternalServerError)
		return
	}
	var images []ImageInfo
	for _, file := range files {
		images = append(images, ImageInfo{
			Name:      file.Name(),
			Thumbnail: "thumbnails/thumb_" + file.Name(),
		})
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(images)
}

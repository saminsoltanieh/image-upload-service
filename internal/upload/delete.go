package upload

import (
	"net/http"
	"os"
	"path/filepath"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "use DELETE method", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "file name required", http.StatusBadRequest)
		return
	}
	filePath := filepath.Join("uploads", name)
	thumbPath := filepath.Join("thumbnails", "thumb_"+name)

	os.Remove(filePath)
	os.Remove(thumbPath)

	w.Write([]byte("File deleted"))
}

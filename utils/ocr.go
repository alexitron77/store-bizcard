package utils

import (
	"bytes"
	"mime/multipart"
	"os/exec"
	"path/filepath"
)

func Ocr(file *multipart.FileHeader, c chan string) {
	path, _ := filepath.Abs("../../ocr-python/ocr.py")

	cmd := exec.Command("python3", path, "uploaded/"+file.Filename)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Run()
	c <- string(stdout.String())
}

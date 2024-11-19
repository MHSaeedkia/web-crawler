package Tests

import (
	"fmt"
	"project-root/modules/export/Lib"
	"project-root/modules/post/DB/Models"
	"project-root/modules/post/DB/Seeders"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func postGenerator() []Models.Post {
	return Seeders.GetFakePosts()
}

func TestZipExport(t *testing.T) {
	var (
		post  []Models.Post
		paths []string
		path  string
		err   error
	)
	post = postGenerator()
	for i := range 5 {
		path, err = export.FinalExport(post, i%2)
		assert.Nil(t, err)
		paths = append(paths, path)
	}

	fileName, _ := export.ZipExport(paths, "/tmp")
	fmt.Println("fileName : ", fileName)
	assert.Nil(t, err)
}

func TestDeleteExport(t *testing.T) {
	var (
		post  []Models.Post
		paths []string
		path  string
		err   error
	)
	post = postGenerator()
	for i := range 5 {
		path, err = export.FinalExport(post, i%2)
		assert.Nil(t, err)
		paths = append(paths, path)
	}

	time.Sleep(3 * time.Second)

	err = export.DeleteExport(paths)
	assert.Nil(t, err)
}

func TestEmailExport(t *testing.T) {
	var (
		email    = "mohammadhasansaeedkia@gmail.com"
		post     []Models.Post
		paths    []string
		path     string
		fileName string
		err      error
	)

	post = postGenerator()
	for i := range 5 {
		path, err = export.FinalExport(post, i%2)
		assert.Nil(t, err)
		paths = append(paths, path)
	}

	fileName, _ = export.ZipExport(paths, "/tmp")
	err = export.EmailExport(email, fileName)
	assert.Nil(t, err)
}

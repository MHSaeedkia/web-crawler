package export

import (
	"fmt"
	"testing"
	"time"

	"github.com/MHSaeedkia/web-crawler/internal/models"
	"github.com/stretchr/testify/assert"
)

func postGenerator() models.Posts {
	priceHistory := make(map[string]interface{})
	priceHistory["2024-01-02"] = 400
	priceHistory["2024-05-02"] = 440

	galleryImg := make(map[string]interface{})
	galleryImg["vanak 200 meteri"] = []string{"/tmp/image1.jpg", "/tmp/image2.jpg"}
	post := models.Posts{
		ID:             1,
		SrcSitesID:     1,
		CitiesID:       1,
		UsersID:        1,
		Status:         1,
		ExternalSiteID: "1",
		Title:          "vanak 200 meteri",
		Description:    "an apartment in junction of talegani St. and motahari St.",
		Price:          40000000000,
		PriceHistory:   priceHistory,
		MainIMG:        "/tmp/image.jpg",
		GalleryIMGs:    galleryImg,
		SellerName:     "mahmodi",
		LandArea:       1.2,
		BuiltYear:      4,
		Rooms:          2,
		IsApartment:    false,
		DealType:       1,
		Floors:         6,
		Elevator:       false,
		Storage:        false,
		Location:       "vanak Sq satartkhan St",
		PostDate:       time.Now(),
		DeletedAt:      time.Now(),
		CreatedAt:      time.Now(),
		UpdateAt:       time.Now(),
	}
	return post
}

func TestZipExport(t *testing.T) {
	var (
		post  models.Posts
		paths []string
		path  string
		err   error
	)
	post = postGenerator()
	for i := range 5 {
		path, err = Export(post, i%2)
		assert.Nil(t, err)
		paths = append(paths, path)
	}

	fileName, _ := ZipExport(paths, "/tmp")
	fmt.Println("fileName : ", fileName)
	assert.Nil(t, err)
}

func TestDeleteExport(t *testing.T) {
	var (
		post  models.Posts
		paths []string
		path  string
		err   error
	)
	post = postGenerator()
	for i := range 5 {
		path, err = Export(post, i%2)
		assert.Nil(t, err)
		paths = append(paths, path)
	}

	time.Sleep(3 * time.Second)

	err = DeleteExport(paths)
	assert.Nil(t, err)
}

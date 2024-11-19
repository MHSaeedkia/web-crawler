package export

import (
	"fmt"
	"testing"
	"time"

	"github.com/MHSaeedkia/web-crawler/internal/models"
	"github.com/stretchr/testify/assert"
)

func postGenerator() []models.Posts {
	var post []models.Posts
	priceHistory := make(map[string]interface{})
	priceHistory["2024-01-02"] = 400
	priceHistory["2024-05-02"] = 440

	galleryImg := make(map[string]interface{})
	galleryImg["vanak 200 meteri"] = []string{"/tmp/image1.jpg", "/tmp/image2.jpg"}
	post1 := models.Posts{
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

	post2 := models.Posts{
		ID:             2,
		SrcSitesID:     2,
		CitiesID:       2,
		UsersID:        2,
		Status:         2,
		ExternalSiteID: "1",
		Title:          "lavasan 1200 meteri",
		Description:    "an apartment in junction of talegani St. and motahari St.",
		Price:          140000000000,
		PriceHistory:   priceHistory,
		MainIMG:        "/tmp/image.jpg",
		GalleryIMGs:    galleryImg,
		SellerName:     "mahmodi",
		LandArea:       2.2,
		BuiltYear:      3,
		Rooms:          1,
		IsApartment:    false,
		DealType:       1,
		Floors:         1,
		Elevator:       false,
		Storage:        false,
		Location:       "vanak Sq satartkhan St",
		PostDate:       time.Now(),
		DeletedAt:      time.Now(),
		CreatedAt:      time.Now(),
		UpdateAt:       time.Now(),
	}

	post3 := models.Posts{
		ID:             3,
		SrcSitesID:     3,
		CitiesID:       3,
		UsersID:        3,
		Status:         3,
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
	post = append(post, post1, post2, post3)
	return post
}

func TestZipExport(t *testing.T) {
	var (
		post  []models.Posts
		paths []string
		path  string
		err   error
	)
	post = postGenerator()
	for i := range 5 {
		path, err = FinalExport(post, i%2)
		assert.Nil(t, err)
		paths = append(paths, path)
	}

	fileName, _ := ZipExport(paths, "/tmp")
	fmt.Println("fileName : ", fileName)
	assert.Nil(t, err)
}

func TestDeleteExport(t *testing.T) {
	var (
		post  []models.Posts
		paths []string
		path  string
		err   error
	)
	post = postGenerator()
	for i := range 5 {
		path, err = FinalExport(post, i%2)
		assert.Nil(t, err)
		paths = append(paths, path)
	}

	time.Sleep(3 * time.Second)

	err = DeleteExport(paths)
	assert.Nil(t, err)
}

func TestEmailExport(t *testing.T) {
	var (
		email    = "mohammadhasansaeedkia@gmail.com"
		post     []models.Posts
		paths    []string
		path     string
		fileName string
		err      error
	)

	post = postGenerator()
	for i := range 5 {
		path, err = FinalExport(post, i%2)
		assert.Nil(t, err)
		paths = append(paths, path)
	}

	fileName, _ = ZipExport(paths, "/tmp")
	err = EmailExport(email, fileName)
	assert.Nil(t, err)
}

package Repositories

import (
	"gorm.io/gorm"
	"project-root/modules/post/DB/Models"
)

type PostRepository struct {
	Db *gorm.DB
}

func (repo *PostRepository) FindByID(id int) (*Models.Post, error) {
	var post Models.Post
	if err := repo.Db.Preload("SrcSite").Preload("User").First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (repo *PostRepository) List() ([]Models.Post, error) {
	var posts []Models.Post
	if err := repo.Db.Preload("SrcSite").Preload("User").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (repo *PostRepository) Create(post *Models.Post) error {
	return repo.Db.Create(post).Error
}

func (repo *PostRepository) Update(post *Models.Post) error {
	return repo.Db.Save(post).Error
}

func (repo *PostRepository) Delete(id int) error {
	return repo.Db.Delete(&Models.Post{}, id).Error
}

func (repo *PostRepository) Truncate() error {
	return repo.Db.Where("1 = 1").Delete(&Models.Post{}).Error
}

func (repo *PostRepository) FindBySourceSiteID(sourceSiteID int) ([]Models.Post, error) {
	var posts []Models.Post
	if err := repo.Db.Where("source_sites_id = ?", sourceSiteID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (repo *PostRepository) FindByUserID(userID int) ([]Models.Post, error) {
	var posts []Models.Post
	if err := repo.Db.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (repo *PostRepository) FindByStatus(status int) ([]Models.Post, error) {
	var posts []Models.Post
	if err := repo.Db.Where("status = ?", status).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

var _ PostRepositoryInterface = &PostRepository{}

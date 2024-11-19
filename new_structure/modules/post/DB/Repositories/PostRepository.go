package Repositories

import (
	"gorm.io/gorm"
	"project-root/modules/post/DB/Models"
	ReportModels "project-root/modules/report/DB/Models"
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

func (repo *PostRepository) GetPostsForFilter(filter *ReportModels.ReportFilter, perPage, pageNum int) (*[]Models.Post, int, error) {
	var posts []Models.Post
	var totalRecords int64

	query := repo.Db.Model(&Models.Post{})

	if filter.BuiltStart != nil {
		query = query.Where("built_year >= ?", *filter.BuiltStart)
	}
	if filter.BuiltEnd != nil {
		query = query.Where("built_year <= ?", *filter.BuiltEnd)
	}
	if filter.AreaMin != nil {
		query = query.Where("land_area >= ?", *filter.AreaMin)
	}
	if filter.AreaMax != nil {
		query = query.Where("land_area <= ?", *filter.AreaMax)
	}
	if filter.PriceMin != nil {
		query = query.Where("price >= ?", *filter.PriceMin)
	}
	if filter.PriceMax != nil {
		query = query.Where("price <= ?", *filter.PriceMax)
	}
	if filter.RoomCountMin != nil {
		query = query.Where("rooms >= ?", *filter.RoomCountMin)
	}
	if filter.RoomCountMax != nil {
		query = query.Where("rooms <= ?", *filter.RoomCountMax)
	}
	if filter.DealType != nil {
		query = query.Where("deal_type = ?", *filter.DealType)
	}
	if filter.CityName != nil {
		query = query.Where("city_name = ?", *filter.CityName)
	}
	if filter.NeighborhoodName != nil {
		query = query.Where("neighborhood_name = ?", *filter.NeighborhoodName)
	}
	if filter.PostStartDate != nil {
		query = query.Where("post_date >= ?", *filter.PostStartDate)
	}
	if filter.PostEndDate != nil {
		query = query.Where("post_date <= ?", *filter.PostEndDate)
	}
	if filter.Elevator != nil {
		query = query.Where("has_elevator = ?", *filter.Elevator)
	}
	if filter.Storage != nil {
		query = query.Where("has_storage = ?", *filter.Storage)
	}
	if filter.Parking != nil {
		query = query.Where("parking = ?", *filter.Parking)
	}
	if filter.IsApartment != nil {
		query = query.Where("is_apartment = ?", *filter.IsApartment)
	}
	if filter.Location != nil && filter.LocationRadius != nil {
		query = query.Where("ST_Distance_Sphere(point(location), point(?, ?)) <= ?", *filter.Location, *filter.LocationRadius)
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	totalPages := int((totalRecords + int64(perPage) - 1) / int64(perPage))

	offset := (pageNum - 1) * perPage
	if err := query.Limit(perPage).Offset(offset).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return &posts, totalPages, nil
}

var _ PostRepositoryInterface = &PostRepository{}

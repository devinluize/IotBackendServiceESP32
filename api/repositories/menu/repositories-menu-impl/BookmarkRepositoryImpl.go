package MenuImplRepositories

import (
	"IotBackend/api/entities"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	menuRepository "IotBackend/api/repositories/menu"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
	"net/http"
)

type BookmarkRepositoryImpl struct {
}

func NewBookmarkRepositoryImpl() menuRepository.BookmarkRepository {
	return &BookmarkRepositoryImpl{}
}
func (repository *BookmarkRepositoryImpl) AddBookmark(db *gorm.DB, userId int, menuId int) (entities.Bookmark, *responses.ErrorResponses) {
	BookmarkEntities := entities.Bookmark{
		//BookmarkTypeId: 1,
		ArticleId: menuId,
		UserId:    userId,
	}
	//validate bookmark and user id cant be same
	isExist := 0
	errExist := db.Model(&BookmarkEntities).Where(entities.Bookmark{ArticleId: menuId, UserId: userId}).
		Select("1").Scan(&isExist).Error
	if errExist != nil {
		return BookmarkEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        errExist,
			Message:    "failed to check duplicate bookmark",
		}
	}
	if isExist == 1 {
		return BookmarkEntities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    "cannot insert duplicate bookmark",
			Err:        errors.New("cannot insert duplicate bookmark"),
			Success:    false,
			Data:       nil,
		}
	}

	err := db.Create(&BookmarkEntities).First(&BookmarkEntities).Error
	if err != nil {
		return BookmarkEntities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed To Add Bookmark ",
			Err:        err,
			Success:    false,
		}
	}
	return BookmarkEntities, nil
}

func (repository *BookmarkRepositoryImpl) RemoveBookmark(db *gorm.DB, userId int, menuId int) (bool, *responses.ErrorResponses) {
	//err := db.Where(entities.Bookmark{ArticleId: menuId, UserId: userId}).Delete(&entities.Bookmark{}).Error
	err := db.Delete(&entities.Bookmark{}, entities.Bookmark{ArticleId: menuId, UserId: userId}).Error
	if err != nil {
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed To Remove Bookmark ",
		}
	}
	return true, nil
}

func (repository *BookmarkRepositoryImpl) GetBookmarks(db *gorm.DB, userId int, cld *cloudinary.Cloudinary) ([]MenuPayloads.GetAllBookmarkResponse, *responses.ErrorResponses) {
	var InfoResponses []MenuPayloads.GetAllBookmarkResponse

	err := db.Table("trx_bookmark A").
		Joins("INNER JOIN mtr_article B ON A.article_id = B.article_id").
		Where("A.user_id = ?", userId).
		Select("B.*,A.*").Scan(&InfoResponses).Error
	for i, v := range InfoResponses {
		urls, _ := cld.Image(v.ArticleHeaderPathContent)
		//res.SortOf = url
		InfoResponses[i].ArticleHeaderPathContent = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
			"dlrd9z1mk",          // Replace with your Cloudinary cloud name
			urls.AssetType,       // e.g., "image"
			urls.DeliveryType,    // e.g., "upload"
			urls.PublicID+".jpg", // Add appropriate file extension
		)
	}

	if err != nil {
		return InfoResponses, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed fetching Bookmarks ",
			Err:        err,
			Success:    false,
		}
	}
	return InfoResponses, nil
}

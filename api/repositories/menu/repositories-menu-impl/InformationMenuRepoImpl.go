package MenuImplRepositories

import (
	"IotBackend/api/entities"
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	menuRepository "IotBackend/api/repositories/menu"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type ArticleMenu struct {
}

func NewArticleMenu() menuRepository.ArticleMenu {
	return &ArticleMenu{}
}
func (i *ArticleMenu) UpdateArticle(tx *gorm.DB, payloads MenuPayloads.ArticleUpdatePayloads) (entities.ArticleEntities, *responses.ErrorResponses) {
	var ArticleEntities entities.ArticleEntities
	var ArticleBodyEntities []entities.ArticleBodyEntities

	err := tx.Model(&ArticleEntities).Where(entities.ArticleEntities{ArticleId: payloads.ArticleId}).First(&ArticleEntities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ArticleEntities, &responses.ErrorResponses{
				Message:    err.Error(),
				StatusCode: http.StatusNotFound,
				Err:        err,
			}
		}
		return ArticleEntities,
			&responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
				Err:     err,
				Message: err.Error()}
	}
	for _, i := range payloads.ArticleBodyContent {
		ArticleBodyEntitiesData := entities.ArticleBodyEntities{
			ArticleBodyParagraph:    i.ArticleBodyParagraph,
			ArticleImageContentPath: i.ArticleImageContentPath,
			ArticleId:               payloads.ArticleId,
		}
		ArticleBodyEntities = append(ArticleBodyEntities, ArticleBodyEntitiesData)
	}
	//Article Body Inserting
	ArticleEntities.ArticleBody = ArticleBodyEntities
	err = tx.Delete(entities.ArticleBodyEntities{}, entities.ArticleBodyEntities{ArticleId: payloads.ArticleId}).Error
	err = tx.Save(&ArticleEntities).Error
	if err != nil {
		return ArticleEntities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			//Err:        nil,
			Success: false,
		}
	}
	return ArticleEntities, nil
}
func (i *ArticleMenu) DeleteArticleById(db *gorm.DB, id int) (bool, *responses.ErrorResponses) {
	err := db.Model(&entities.ArticleEntities{}).Where(entities.ArticleEntities{ArticleId: id}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &responses.ErrorResponses{
				StatusCode: http.StatusBadRequest,
				Message:    "Delete Failed Id Not Found",
				Success:    false,
			}
		}
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	err = db.Delete(entities.ArticleBodyEntities{}, entities.ArticleBodyEntities{ArticleId: id}).Error
	if err != nil {
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}

	err = db.Delete(entities.ArticleEntities{}, entities.ArticleEntities{ArticleId: id}).Error
	if err != nil {
		return false, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}
	return true, nil
}
func (i *ArticleMenu) InsertArticle(tx *gorm.DB, payloads MenuPayloads.ArticleInsertPayloads) (entities.ArticleEntities, *responses.ErrorResponses) {
	var Entities entities.ArticleEntities
	Entities = entities.ArticleEntities{
		//ArticleId:                0,
		ArticleHeader:            payloads.ArticleHeader,
		ArticleHeaderPathContent: payloads.ArticleHeaderPathContent,
		//ArticleImageContentPath1: payloads.ArticleImageContentPath1,
		//ArticleImageContentPath2: payloads.ArticleImageContentPath2,
		//ArticleImageContentPath3: payloads.ArticleImageContentPath3,
		//ArticleImageContentPath4: payloads.ArticleImageContentPath4,
		//ArticleImageContentPath5: payloads.ArticleImageContentPath5,
		//ArticleTypeId:      payloads.ArticleTypeId,
		ArticleDateCreated: time.Now(),
	}
	//for _, detail := range payloads.ArticleBodyParagraph {
	//	Entities.ArticleBody = append(Entities.ArticleBody, entities.ArticleBodyEntities{
	//		ArticleBodyParagraph:    detail.ArticleBodyParagraph,
	//		ArticleImageContentPath: detail.ArticleImageContentPath,
	//	})
	//}
	var EntitiesDetail []entities.ArticleBodyEntities

	//Entities.ArticleId = Entities.ArticleId
	err := tx.Create(&Entities).First(&Entities).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}

	//err = tx.Model(&Entities).Where(entities.ArticleEntities{ArticleId: Entities.ArticleId}).First(&Entities).Error
	for _, paragraph := range payloads.ArticleBodyParagraph {
		EntitiesDetail = append(EntitiesDetail, entities.ArticleBodyEntities{
			ArticleImageContentPath: paragraph.ArticleImageContentPath,
			ArticleBodyParagraph:    paragraph.ArticleBodyParagraph,
			ArticleId:               Entities.ArticleId,
		})
	}
	err = tx.Create(&EntitiesDetail).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}
	return Entities, nil
}
func (i *ArticleMenu) GetArticleById(db *gorm.DB, id int, userId int, cld *cloudinary.Cloudinary) (MenuPayloads.ArticleSelectResponses, *responses.ErrorResponses) {
	EntitiesInfo := entities.ArticleEntities{}
	err := db.Model(&entities.ArticleEntities{}).Where(entities.ArticleEntities{ArticleId: id}).First(&EntitiesInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return MenuPayloads.ArticleSelectResponses{}, &responses.ErrorResponses{
				StatusCode: http.StatusBadRequest,
				Message:    "Delete Failed Id Not Found",
			}
		}
	}
	ResDetail := []MenuPayloads.ArticleBodyDetail{}
	err = db.Model(&entities.ArticleBodyEntities{}).Where(entities.ArticleBodyEntities{ArticleId: id}).Scan(&ResDetail).Error
	if err != nil {
		return MenuPayloads.ArticleSelectResponses{}, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}
	SelectPayload := []MenuPayloads.ArticleBodyDetail{}
	for _, detail := range ResDetail {
		if detail.ArticleImageContentPath != "" {
			urls, _ := cld.Image(detail.ArticleImageContentPath)
			//res.SortOf = url
			detail.ArticleImageContentPath = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
				"dlrd9z1mk",          // Replace with your Cloudinary cloud name
				urls.AssetType,       // e.g., "image"
				urls.DeliveryType,    // e.g., "upload"
				urls.PublicID+".jpg", // Add appropriate file extension
			)
		}
		SelectPayloadData := MenuPayloads.ArticleBodyDetail{
			ArticleBodyParagraph:    detail.ArticleBodyParagraph,
			ArticleImageContentPath: detail.ArticleImageContentPath,
		}
		SelectPayload = append(SelectPayload, SelectPayloadData)
	}
	isBookmarkExist := false
	urls, _ := cld.Image(EntitiesInfo.ArticleHeaderPathContent)
	//res.SortOf = url
	EntitiesInfo.ArticleHeaderPathContent = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
		"dlrd9z1mk",          // Replace with your Cloudinary cloud name
		urls.AssetType,       // e.g., "image"
		urls.DeliveryType,    // e.g., "upload"
		urls.PublicID+".jpg", // Add appropriate file extension
	)

	//EntitiesInfo.ArticleHeaderPathContent
	err = db.Model(&entities.Bookmark{}).Where(entities.Bookmark{ArticleId: id, UserId: userId}).Select("1").
		Scan(&isBookmarkExist).Error
	result := MenuPayloads.ArticleSelectResponses{
		ArticleHeader:            EntitiesInfo.ArticleHeader,
		ArticleDateCreated:       EntitiesInfo.ArticleDateCreated,
		ArticleCreatedByUserId:   EntitiesInfo.ArticleCreatedByUserId,
		ArticleId:                id,
		IsBookmark:               isBookmarkExist,
		ArticleBodyContent:       SelectPayload,
		ArticleHeaderPathContent: EntitiesInfo.ArticleHeaderPathContent,
		//ArticleTypeId:          EntitiesInfo.ArticleTypeId,
	}
	return result, nil
}
func (i *ArticleMenu) GetAllArticleWithPagination(db *gorm.DB, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses) {

	var Entities []entities.ArticleEntities
	//me := db.Model(&entities.ArticleEntities{}) -> table joinan
	//cara 1
	//myJoinTable := db.Model(&entities.ArticleEntities{}).
	//err := db.Model(&entities.ArticleEntities{}).Scopes(database.Paginate(&Entities, &paginationResponses, me)).Order("article_id").Where("article_id <> 0").Scan(&Entities).Error
	//cara 2 langsung assign ke database nanti pilih aja apakah perlu buat join table atau ga kalau misalkan selectan itu merupakan hasil join table pake yang atas

	err := db.Model(&entities.ArticleEntities{}).Scopes(helper.Paginate(&Entities, &paginationResponses, db)).Order("article_id").Where("article_id <> 0").Scan(&Entities).Error
	if err != nil {
		return paginationResponses, &responses.ErrorResponses{}
	}
	paginationResponses.Rows = Entities
	fmt.Println(paginationResponses.Rows)
	return paginationResponses, nil
}
func (i *ArticleMenu) GetAllArticleWithFilter(db *gorm.DB, paginationResponses helper.Pagination, Key string, userId int, cloudinary *cloudinary.Cloudinary) (helper.Pagination, *responses.ErrorResponses) {
	//create history logging
	//ctx := context.Background()
	if Key != "" {

		historyLogging := entities.SearchHistoryEntities{
			UserId:     userId,
			SearchKey:  Key,
			DateSearch: time.Now(),
		}

		// Insert the new history record
		err := db.Create(&historyLogging).Error
		if err != nil {
			return paginationResponses, &responses.ErrorResponses{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
				Message:    "failed to log search history",
			}
		}

		// Check the total count of search history records for the user
		//historyCount := 0
		var historyCount int64
		errCount := db.Model(&entities.SearchHistoryEntities{}).
			Where("user_id = ?", userId).
			Count(&historyCount).Error
		if errCount != nil {
			return paginationResponses, &responses.ErrorResponses{
				StatusCode: http.StatusInternalServerError,
				Err:        errCount,
				Message:    "failed to check search history count",
			}
		}

		if historyCount > 10 {
			var excessCount int
			excessCount = int(historyCount) - 10
			// Delete oldest records if the count exceeds the limit
			var recordsToDelete []entities.SearchHistoryEntities
			errSelect := db.Where("user_id = ?", userId).
				Order("date_search ASC").
				Limit(excessCount).
				Find(&recordsToDelete).Error
			if errSelect != nil {
				return paginationResponses, &responses.ErrorResponses{
					StatusCode: http.StatusInternalServerError,
					Err:        errSelect,
					Message:    "failed to fetch old search history records",
				}
			}

			// Step 2: Delete the fetched records
			if len(recordsToDelete) > 0 {
				errDelete := db.Delete(&recordsToDelete).Error
				if errDelete != nil {
					return paginationResponses, &responses.ErrorResponses{
						StatusCode: http.StatusInternalServerError,
						Err:        errDelete,
						Message:    "failed to clean up old search history records",
					}
				}
			}
		}

		//	// Calculate how many records need to be deleted

		//	fmt.Println("excess : ")
		//	fmt.Println(excessCount)
		//	// Delete the oldest records
		//	errDelete := db.Where("user_id = ?", userId).
		//		Order("date_search ASC"). // Assuming `date_search` is used to track record age
		//		Limit(excessCount).
		//		Delete(&entities.SearchHistoryEntities{}).Error
		//	if errDelete != nil {
		//		return paginationResponses, &responses.ErrorResponses{
		//			StatusCode: http.StatusInternalServerError,
		//			Err:        errDelete,
		//			Message:    "failed to clean up old search history records",
		//		}
		//	}
	}

	var Entities []entities.ArticleEntities
	//me := db.Model(&entities.ArticleEntities{}) -> table joinan
	//cara 1
	joinTable := db.Model(&entities.ArticleEntities{}).Where("article_id <> 0 AND article_header LIKE ? ", "%"+Key+"%")
	//err := db.Model(&entities.ArticleEntities{}).Scopes(database.Paginate(&Entities, &paginationResponses, me)).Order("article_id").Where("article_id <> 0").Scan(&Entities).Error
	//cara 2 langsung assign ke database nanti pilih aja apakah perlu buat join table atau ga kalau misalkan selectan itu merupakan hasil join table pake yang atas
	err := joinTable.Scopes(helper.Paginate(&Entities, &paginationResponses, joinTable)).Order("article_id").Where("article_id <> 0 AND article_header LIKE ? ", "%"+Key+"%").Scan(&Entities).Error
	if err != nil {
		return paginationResponses, &responses.ErrorResponses{}
	}

	for i2, entity := range Entities {
		urls, _ := cloudinary.Image(entity.ArticleHeaderPathContent)
		//res.SortOf = url
		Entities[i2].ArticleHeaderPathContent = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
			"dlrd9z1mk",          // Replace with your Cloudinary cloud name
			urls.AssetType,       // e.g., "image"
			urls.DeliveryType,    // e.g., "upload"
			urls.PublicID+".jpg", // Add appropriate file extension
		)
	}

	paginationResponses.Rows = Entities
	fmt.Println(paginationResponses.Rows)
	return paginationResponses, nil
}
func (i *ArticleMenu) GetArticleHistory(db *gorm.DB, userId int) ([]entities.SearchHistoryEntities, *responses.ErrorResponses) {
	var entitiesData []entities.SearchHistoryEntities
	err := db.Model(&entities.SearchHistoryEntities{}).
		Where(entities.SearchHistoryEntities{UserId: userId}).
		Order("date_search DESC").
		Limit(10).
		Scan(&entitiesData).Error
	if err != nil {
		return entitiesData, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get article history search data",
		}
	}
	return entitiesData, nil
}

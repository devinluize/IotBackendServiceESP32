package menuserviceimpl

import (
	"IotBackend/api/entities"
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	menuRepository "IotBackend/api/repositories/menu"
	"IotBackend/api/service/menu"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

type BookmarkServiceImpl struct {
	db   *gorm.DB
	repo menuRepository.BookmarkRepository
	cld  *cloudinary.Cloudinary
}

func NewBookmarkServiceImpl(db *gorm.DB, repo menuRepository.BookmarkRepository, cld *cloudinary.Cloudinary) menu.BookmarkService {

	return &BookmarkServiceImpl{db: db, repo: repo, cld: cld}
}
func (s *BookmarkServiceImpl) AddBookmark(userId int, menuId int) (entities.Bookmark, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repo.AddBookmark(trans, userId, menuId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *BookmarkServiceImpl) RemoveBookmark(userId int, menuId int) (bool, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repo.RemoveBookmark(trans, userId, menuId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *BookmarkServiceImpl) GetBookmarks(userId int) ([]MenuPayloads.GetAllBookmarkResponse, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repo.GetBookmarks(trans, userId, s.cld)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

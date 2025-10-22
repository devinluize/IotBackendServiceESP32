package menu

import (
	"IotBackend/api/entities"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
)

type BookmarkService interface {
	AddBookmark(userId int, menuId int) (entities.Bookmark, *responses.ErrorResponses)
	RemoveBookmark(userId int, menuId int) (bool, *responses.ErrorResponses)
	GetBookmarks(userId int) ([]MenuPayloads.GetAllBookmarkResponse, *responses.ErrorResponses)
}

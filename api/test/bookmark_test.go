package test

import (
	configenv "IotBackend/api/config"
	"IotBackend/api/entities"
	MenuImplRepositories "IotBackend/api/repositories/menu/repositories-menu-impl"
	"IotBackend/api/service/menu"
	menuserviceimpl "IotBackend/api/service/menu/menu-service-impl"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

var (
	db           *gorm.DB
	service      menu.BookmarkService
	bookmarkTest entities.Bookmark
)

// TestMain runs once for the entire package
func TestMain(m *testing.M) {
	// Initialize environment configs and the test database
	configenv.InitEnvConfigs(true, "")
	db = configenv.InitTestDB()

	// Initialize the Bookmark service
	BookmarkRepo := MenuImplRepositories.NewBookmarkRepositoryImpl()
	service = menuserviceimpl.NewBookmarkServiceImpl(db, BookmarkRepo)

	// Run the tests
	code := m.Run()

	// Cleanup database connection after tests
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.Close()
	}

	// Exit with the test code
	os.Exit(code)
}

func TestInsertBookmark(t *testing.T) {
	userId := 1
	ArticleId := 3

	res, err := service.AddBookmark(userId, ArticleId)
	if err != nil {
		t.Errorf("Failed On: %v", err)
	}
	t.Log("Bookmark Inserted Successfully")
	assert.Nil(t, err)
	bookmarkTest = res // Save the result for other tests
}

func TestRemoveBookmark(t *testing.T) {
	_, err := service.RemoveBookmark(bookmarkTest.UserId, bookmarkTest.ArticleId)
	if err != nil {
		t.Errorf("Failed to remove bookmark: %v", err)
	}
	fmt.Println("Bookmark Removed Successfully")
	assert.Nil(t, err)
}

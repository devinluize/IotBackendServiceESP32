package menucontroller

import (
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/service/menu"
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type ArticleController interface {
	InsertArticle(writer http.ResponseWriter, request *http.Request)
	DeleteArticleById(writer http.ResponseWriter, request *http.Request)
	UpdateArticle(writer http.ResponseWriter, request *http.Request)
	GeById(writer http.ResponseWriter, request *http.Request)
	GetAllByPagination(writer http.ResponseWriter, request *http.Request)
	GetAllArticleByFilter(writer http.ResponseWriter, request *http.Request)
	GetArticleHistory(writer http.ResponseWriter, request *http.Request)
}

type ArticleControllerImpl struct {
	ArticleService menu.ArticleService
}

func NewInformatioControllerImpl(ArticleService menu.ArticleService) ArticleController {
	return &ArticleControllerImpl{ArticleService: ArticleService}
}

// InsertArticle List Via Header
//
//	@Security		BearerAuth
//	@Summary		Create New Article
//	@Description	Create New Article
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MenuPayloads.ArticleInsertPayloads	true	"Insert Request"
//	@Success		200		{object}	responses.StandarAPIResponses
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/article [post]
func (i *ArticleControllerImpl) InsertArticle(writer http.ResponseWriter, request *http.Request) {
	var ArticlePayloads MenuPayloads.ArticleInsertPayloads
	helper.ReadFromRequestBody(request, &ArticlePayloads)

	res, err := i.ArticleService.InsertArticle(ArticlePayloads)
	if err != nil {
		helper.ReturnError(writer, err)
		return

	}
	helper.HandleSuccess(writer, res, "Insert Successfull", http.StatusCreated)
}

// DeleteArticleById List Via Header
//
//	@Security		BearerAuth
//	@Summary		Delete Article
//	@Description	Delete Article
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			article_id	path int	true	"article_id"
//	@Success		200		{object}	 responses.StandarAPIResponses
//	@Router			/api/article/delete/{article_id} [delete]
func (i *ArticleControllerImpl) DeleteArticleById(writer http.ResponseWriter, request *http.Request) {
	ArticleId := chi.URLParam(request, "article_id")
	ArticleIds, err := strconv.Atoi(ArticleId)
	if err != nil {
		return
	}
	res, errs := i.ArticleService.DeleteArticleById(ArticleIds)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return

	}
	helper.HandleSuccess(writer, res, "Delete Successfull", http.StatusOK)
}

// UpdateArticle List Via Header
//
//	@Security		BearerAuth
//	@Summary		Update Article
//	@Description	Update Article
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MenuPayloads.ArticleUpdatePayloads	true	"Update Request"
//	@Success		200		{object}	 responses.ErrorResponses
//	@Router			/api/article [patch]
func (i *ArticleControllerImpl) UpdateArticle(writer http.ResponseWriter, request *http.Request) {
	var ArticlePayloads MenuPayloads.ArticleUpdatePayloads
	helper.ReadFromRequestBody(request, &ArticlePayloads)

	res, err := i.ArticleService.UpdateArticle(ArticlePayloads)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Update Successfully", http.StatusOK)
}

// GeById List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get Article By Article
//	@Description	Get Article By Article
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			article_id	path int			true		"article_id"
//	@Success		200									{object}	entities.ArticleEntities
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/article/by-id/{article_id} [get]
func (i *ArticleControllerImpl) GeById(writer http.ResponseWriter, request *http.Request) {
	ArticleId := chi.URLParam(request, "article_id")
	ArticleIds, err := strconv.Atoi(ArticleId)
	if err != nil {
		return
	}
	user := helper.GetRequestCredentialFromHeaderToken(request)
	res, errs := i.ArticleService.GetArticleById(ArticleIds, user.UserId)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return

	}
	helper.HandleSuccess(writer, res, "Get Err Successfuull", http.StatusOK)
}

// GetAllByPagination List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get All Article By Pagination
//	@Description	Get All Article By Pagination
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			sort_by								query		string	false	"sort_by"
//	@Param			sort_of								query		string	false	"sort_of"
//	@Param			page								query		string	true	"page"
//	@Param			limit								query		string	true	"limit"
//	@Success		200									{object}	[]entities.ArticleEntities
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/article [get]
func (i *ArticleControllerImpl) GetAllByPagination(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()

	pagination := helper.Pagination{
		Limit:  helper.NewGetQueryInt(queryValues, "limit"),
		Page:   helper.NewGetQueryInt(queryValues, "page"),
		SortOf: queryValues.Get("sort_of"),
		SortBy: queryValues.Get("sort_by"),
	}
	res, err := i.ArticleService.GetAllArticleWithPagination(pagination)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Get Successfully", http.StatusOK)
}

// GetAllArticleByFilter List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get All Article By Pagination With Filter
//	@Description	Get All Article By Pagination With Filter
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			key_filter							query		string	false	"key_filter"
//	@Param			sort_by								query		string	false	"sort_by"
//	@Param			sort_of								query		string	false	"sort_of"
//	@Param			page								query		string	true	"page"
//	@Param			limit								query		string	true	"limit"
//	@Success		200									{object}	[]entities.ArticleEntities
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/article/search [get]
func (i *ArticleControllerImpl) GetAllArticleByFilter(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	pagination := helper.Pagination{
		Limit:  helper.NewGetQueryInt(queryValues, "limit"),
		Page:   helper.NewGetQueryInt(queryValues, "page"),
		SortOf: queryValues.Get("sort_of"),
		SortBy: queryValues.Get("sort_by"),
	}
	user := helper.GetRequestCredentialFromHeaderToken(request)
	Key := queryValues.Get("key_filter")
	res, err := i.ArticleService.GetAllArticleWithFilter(pagination, Key, user.UserId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	//localFilePath := `C:\Users\devin\Documents\Github\IotBackend\api\controller\menu\asdasd.png`
	//
	//// Open the local file
	//file, errOpen := os.Open(localFilePath)
	//if errOpen != nil {
	//	panic("Failed to open file: %v")
	//}
	//defer file.Close()

	//CLOUDINARY_URL:=cloudinary://API_KEY:API_SECRET@CLOUD_NAME
	cld, errr := cloudinary.NewFromURL("cloudinary://695971277991789:jXnWGXSCY230XQ_5QUtMGcb9T18@dlrd9z1mk")
	fmt.Println(errr)
	cld.Config.URL.Secure = true
	ctx := context.Background()
	//resp, errcld := cld.Upload.Upload(ctx, file, uploader.UploadParams{
	//	PublicID:       "folder/12345asd",
	//	UniqueFilename: api.Bool(false),
	//	Overwrite:      api.Bool(true)})

	//cld, errr := cloudinary.NewFromURL("cloudinary://225934859532926:CJUjvSBn-bvavAE8UipYgtSaDjw@dmgpda5o7")
	//fmt.Println(errr)
	//cld.Config.URL.Secure = true
	//ctx := context.Background()
	resp, errcld := cld.Upload.Upload(ctx, "https://res.cloudinary.com/dlrd9z1mk/image/upload/v1735399592/folder/wrbcua7zbzxaqrokjia6.png", uploader.UploadParams{
		PublicID:       "folder/wrbcua7zbzxaqrokjia6DEVIN",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	if errcld != nil {
		fmt.Println("error")
		return
	}

	res.SortBy = resp.URL
	urls, _ := cld.Image(resp.PublicID)
	//res.SortOf = url
	res.SortOf = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
		"dlrd9z1mk",          // Replace with your Cloudinary cloud name
		urls.AssetType,       // e.g., "image"
		urls.DeliveryType,    // e.g., "upload"
		urls.PublicID+".jpg", // Add appropriate file extension
	)

	helper.HandleSuccess(writer, res, "Get Successfully", http.StatusOK)
}
func (i *ArticleControllerImpl) GetArticleHistory(writer http.ResponseWriter, request *http.Request) {
	user := helper.GetRequestCredentialFromHeaderToken(request)

	res, err := i.ArticleService.GetArticleHistory(user.UserId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}

	helper.HandleSuccess(writer, res, "Get Successfully", http.StatusOK)
}

package helper

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"strconv"
)

func Paniciferror(err error) {
	if err != nil {
		panic(err)
	} else {
		return
	}
}

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		tx.Rollback()
		logrus.Info(err)
	} else {
		tx.Commit()
	}
}
func NewGetQueryInt(queryValues url.Values, param string) int {
	value, _ := strconv.Atoi(queryValues.Get(param))
	return value
}
func NewGetParamInt(r *http.Request, param string) int {
	value, _ := strconv.Atoi(chi.URLParam(r, param))
	return value
}

type UserContext struct {
	UserName string
	UserId   int
}

func GetRequestCredentialFromHeaderToken(r *http.Request) UserContext {
	//return UserContext{}
	return r.Context().Value("user_credential").(UserContext)
}

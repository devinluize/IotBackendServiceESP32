package menucontroller

import (
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/service/menu"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type TimerController interface {
	InsertTimer(writer http.ResponseWriter, request *http.Request)
	InsertQueueTimer(writer http.ResponseWriter, request *http.Request)
	UpdateQueueTimer(writer http.ResponseWriter, request *http.Request)
	DeleteTimerQueueTimer(writer http.ResponseWriter, request *http.Request)
	GetTimerByUserId(writer http.ResponseWriter, request *http.Request)
	GetAllQueueTimer(writer http.ResponseWriter, request *http.Request)
	DeleteTimer(writer http.ResponseWriter, request *http.Request)
}

type TimerControllerImpl struct {
	TimerServices menu.TimerService
}

func NewTimerControllerImpl(TimerService menu.TimerService) TimerController {
	return &TimerControllerImpl{TimerServices: TimerService}
}
func (t *TimerControllerImpl) InsertTimer(writer http.ResponseWriter, request *http.Request) {
	var TimerInsert MenuPayloads.TimerInsertPayload
	helper.ReadFromRequestBody(request, &TimerInsert)
	user := helper.GetRequestCredentialFromHeaderToken(request)
	res, err := t.TimerServices.InsertTimer(TimerInsert, user.UserId)
	if err != nil {
		helper.ReturnError(writer, err)
		return

	}
	helper.HandleSuccess(writer, res, "Insert Timer Sucessfull", http.StatusCreated)

}

func (t *TimerControllerImpl) InsertQueueTimer(writer http.ResponseWriter, request *http.Request) {

	var TimerInsert MenuPayloads.TimerQueueInsertResponse
	helper.ReadFromRequestBody(request, &TimerInsert)

	res, err := t.TimerServices.InsertQueueTimer(TimerInsert)
	if err != nil {
		helper.ReturnError(writer, err)
		return

	}
	helper.HandleSuccess(writer, res, "Insert Timer queue sucesfull", http.StatusCreated)

}

func (t *TimerControllerImpl) UpdateQueueTimer(writer http.ResponseWriter, request *http.Request) {

	var TimerUpdate MenuPayloads.TimerQueueUpdatePayload
	helper.ReadFromRequestBody(request, &TimerUpdate)

	res, err := t.TimerServices.UpdateQueueTimer(TimerUpdate)
	if err != nil {
		helper.ReturnError(writer, err)
		return

	}
	helper.HandleSuccess(writer, res, "Update Timer queue sucesfull", http.StatusOK)

}

func (t *TimerControllerImpl) DeleteTimerQueueTimer(writer http.ResponseWriter, request *http.Request) {
	TimerQueueId := chi.URLParam(request, "timer_queue_id")
	TimerQueueIds, err := strconv.Atoi(TimerQueueId)
	if err != nil {
		return
	}
	res, errs := t.TimerServices.DeleteTimerQueueTimer(TimerQueueIds)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return

	}
	helper.HandleSuccess(writer, res, "Delete Timer Queue SucessFull", http.StatusOK)
}

func (t *TimerControllerImpl) GetTimerByUserId(writer http.ResponseWriter, request *http.Request) {
	//UserId := request.Context().Value("user_id").(int)
	User := helper.GetRequestCredentialFromHeaderToken(request)

	res, errs := t.TimerServices.GetTimerByUserId(User.UserId)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	helper.HandleSuccess(writer, res, "GetAll Timer By User Id", http.StatusOK)
}

func (t *TimerControllerImpl) GetAllQueueTimer(writer http.ResponseWriter, request *http.Request) {

	TimerId := chi.URLParam(request, "timer_id")
	TimerIds, err := strconv.Atoi(TimerId)
	if err != nil {
		return
	}
	res, errs := t.TimerServices.GetAllQueueTimer(TimerIds)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	helper.HandleSuccess(writer, res, "GetAll Queue By Timer Id", http.StatusOK)
}
func (t *TimerControllerImpl) DeleteTimer(writer http.ResponseWriter, request *http.Request) {

	TimerId := chi.URLParam(request, "timer_id")
	TimerIds, err := strconv.Atoi(TimerId)
	if err != nil {
		return
	}
	res, errs := t.TimerServices.DeleteTimer(TimerIds)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	helper.HandleSuccess(writer, res, "Delete By Timer Id", http.StatusOK)
}

package menucontroller

import (
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	"IotBackend/api/service/menu"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type CalendarController interface {
	InsertCalendar(writer http.ResponseWriter, request *http.Request)
	GetCalendarByUserId(writer http.ResponseWriter, request *http.Request)
	UpdateCalendar(writer http.ResponseWriter, request *http.Request)
	DeleteCalendarById(writer http.ResponseWriter, request *http.Request)
	GetCalendarByDate(writer http.ResponseWriter, request *http.Request)
}

type CalendarControllerImpl struct {
	CalendarService menu.CalendarService
}

func NewCalendarController(calendarService menu.CalendarService) CalendarController {
	return &CalendarControllerImpl{CalendarService: calendarService}
}

// InsertCalendar List Via Header
//
//	@Security		BearerAuth
//	@Summary		Create New Calendar
//	@Description	Create New Calendar
//	@Tags			Calendar
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MenuPayloads.CalendarInsertPayload	true	"Insert Request"
//	@Success		200		{object}	responses.StandarAPIResponses
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/calendar [post]
func (controller *CalendarControllerImpl) InsertCalendar(writer http.ResponseWriter, request *http.Request) {
	var insertCalendar MenuPayloads.CalendarInsertPayload
	helper.ReadFromRequestBody(request, &insertCalendar)
	user := helper.GetRequestCredentialFromHeaderToken(request)
	insertCalendar.UserId = user.UserId
	res, err := controller.CalendarService.InsertCalendar(insertCalendar)
	if err != nil {
		helper.ReturnError(writer, err)
	}
	helper.HandleSuccess(writer, res, "success insert calendar", http.StatusCreated)
}

// GetCalendarByUserId List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get Calendar by Id
//	@Description	Get Calendar by Id
//	@Tags			Calendar
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path int	true	"user_id"
//	@Success		200		{object}	 responses.StandarAPIResponses
//	@Router			/api/calendar/by-user-id/ [get]
func (controller *CalendarControllerImpl) GetCalendarByUserId(writer http.ResponseWriter, request *http.Request) {
	//calendarId := chi.URLParam(request, "user_id")
	User := helper.GetRequestCredentialFromHeaderToken(request)
	//ArticleIds, err := strconv.Atoi(calendarId)
	//if err != nil {
	//	return
	//}
	res, errs := controller.CalendarService.GetCalendarByUserId(User.UserId)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	helper.HandleSuccess(writer, res, "Get Successful", http.StatusOK)
}

// UpdateCalendar List Via Header
//
//	@Security		BearerAuth
//	@Summary		Update Calendar
//	@Description	Update Calendar
//	@Tags			Calendar
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MenuPayloads.CalendarUpdatePayload	true	"Update Request"
//	@Success		200		{object}	 responses.ErrorResponses
//	@Router			/api/calendar [patch]
func (controller *CalendarControllerImpl) UpdateCalendar(writer http.ResponseWriter, request *http.Request) {

	var UpdateCalendar MenuPayloads.CalendarUpdatePayload
	helper.ReadFromRequestBody(request, &UpdateCalendar)
	calendarId := chi.URLParam(request, "event_id")
	calendarIdInt, errConvert := strconv.Atoi(calendarId)
	if errConvert != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        errConvert,
			Message:    errConvert.Error(),
		})
	}
	user := helper.GetRequestCredentialFromHeaderToken(request)
	UpdateCalendar.EventId = calendarIdInt
	UpdateCalendar.UserId = user.UserId
	res, err := controller.CalendarService.UpdateCalendar(UpdateCalendar)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Update Successfully", http.StatusOK)
}

// DeleteCalendarById List Via Header
//
//	@Security		BearerAuth
//	@Summary		Delete Calendar by Id
//	@Description	Delete Calendar by Id
//	@Tags			Calendar
//	@Accept			json
//	@Produce		json
//	@Param			calendar_id	path int	true	"calendar_id"
//	@Success		200		{object}	 responses.StandarAPIResponses
//	@Router			/api/calendar/delete/{calendar_id} [delete]
func (controller *CalendarControllerImpl) DeleteCalendarById(writer http.ResponseWriter, request *http.Request) {
	calendarId := chi.URLParam(request, "event_id")

	ArticleIds, err := strconv.Atoi(calendarId)
	if err != nil {
		return
	}
	res, errs := controller.CalendarService.DeleteCalendarById(ArticleIds)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	helper.HandleSuccess(writer, res, "Delete Successful", http.StatusOK)
}
func (controller *CalendarControllerImpl) GetCalendarByDate(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	date := queryValues.Get("event_date")
	user := helper.GetRequestCredentialFromHeaderToken(request)
	res, err := controller.CalendarService.GetCalendarByDate(date, user.UserId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Get Successful", http.StatusOK)
}

package EquipmentController

import (
	"IotBackend/api/helper"
	"IotBackend/api/payloads/responses"
	"IotBackend/api/service/EquipmentService"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type EquipmentBookmarkController interface {
	AddEquipmentBookmark(writer http.ResponseWriter, request *http.Request)
	RemoveEquipmentBookmark(writer http.ResponseWriter, request *http.Request)
	GetEquipmentBookmarkByUserId(writer http.ResponseWriter, request *http.Request)
}

type EquipmentBookmarkControllerImpl struct {
	service EquipmentService.EquipmentBookmarkService
}

func (e *EquipmentBookmarkControllerImpl) AddEquipmentBookmark(writer http.ResponseWriter, request *http.Request) {
	user := helper.GetRequestCredentialFromHeaderToken(request)
	//get equipment course id
	equipmentCourse := chi.URLParam(request, "equipment_course_id")
	equipmentCourseId, err := strconv.Atoi(equipmentCourse)
	if err != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to fetch equipment course id",
		})
		return
	}
	res, errRes := e.service.AddEquipmentBookmark(user.UserId, equipmentCourseId)
	if errRes != nil {
		helper.ReturnError(writer, errRes)
		return
	}

	helper.HandleSuccess(writer, res, "success to insert bookmark", http.StatusOK)
}

func (e *EquipmentBookmarkControllerImpl) RemoveEquipmentBookmark(writer http.ResponseWriter, request *http.Request) {
	user := helper.GetRequestCredentialFromHeaderToken(request)
	//get equipment course id
	equipmentCourse := chi.URLParam(request, "equipment_course_id")
	equipmentCourseId, err := strconv.Atoi(equipmentCourse)
	if err != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to fetch equipment course id",
		})
		return
	}
	res, errRes := e.service.RemoveEquipmentBookmark(user.UserId, equipmentCourseId)
	if errRes != nil {
		helper.ReturnError(writer, errRes)
		return
	}
	helper.HandleSuccess(writer, res, "success to delete  bookmark", http.StatusOK)
}
func (e *EquipmentBookmarkControllerImpl) GetEquipmentBookmarkByUserId(writer http.ResponseWriter, request *http.Request) {
	user := helper.GetRequestCredentialFromHeaderToken(request)
	res, err := e.service.GetEquipmentBookmarkByUserId(user.UserId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	if len(res) == 0 {
		helper.HandleSuccess(writer, []string{}, "success to get bookmark", http.StatusOK)
		return

	}
	helper.HandleSuccess(writer, res, "success to get bookmark", http.StatusOK)
}

func NewEquipmentBookmarkControllerImpl(service EquipmentService.EquipmentBookmarkService) EquipmentBookmarkController {
	return &EquipmentBookmarkControllerImpl{service: service}
}

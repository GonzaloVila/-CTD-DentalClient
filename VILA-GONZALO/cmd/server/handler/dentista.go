package handler

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/GonzaloVila/clinica/internal/dentistas"

	"github.com/GonzaloVila/clinica/pkg/web"

	"github.com/gin-gonic/gin"
)

type dentistaHandler struct {
	dService dentistas.Service
}

func NewDentistaHandler(d dentistas.Service) *dentistaHandler {
	return &dentistaHandler{
		dService: d,
	}
}

type requestDentista struct {
	ID        int     `json:"id"`
	Apellido  *string `json:"apellido"`
	Nombre    *string `json:"nombre"`
	Matricula *int    `json:"matricula"`
}

func (h *dentistaHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		dentista, err := h.dService.GetByID(ctx, id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("dentist not found"))
			return
		}
		web.Success(ctx, 200, dentista)
	}
}

func (h *dentistaHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestDentista
		if err := ctx.ShouldBindJSON(&req); err!=nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		dentistaRef := reflect.ValueOf(req)
		var valuesNil []string
		for i := 0; i < dentistaRef.NumField(); i++ {
			if e := dentistaRef.Field(i); e.IsNil() &&
				dentistaRef.Type().Field(i).Name != "ID" {
				valuesNil = append(valuesNil, dentistaRef.Type().Field(i).Name)
			}
		}
		if len(valuesNil) > 0 {
			web.Error(ctx, 422, "required fields: %s", strings.Join(valuesNil, ", "))
			return
		}
		d, err := h.dService.Create(ctx, *req.Matricula, *req.Apellido, *req.Nombre)
		if err != nil {
			web.Failure(d, 400, err)
			return
		}
		web.Success(ctx, 201, d)
	}
}

func (h *dentistaHandler) Delete() gin.HandlerFunc { //probar
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		err = h.dService.Delete(ctx, int(id))
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 204, "delete from dentista")
	}
}

func (h *dentistaHandler) Put() gin.HandlerFunc { //probar
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		var req requestDentista
		if err := ctx.ShouldBindJSON(&req); err!=nil {
			web.Failure(ctx, 422, errors.New("data cannot be processed"))
			return
		}
		myMap := make(map[string]interface{})
		if req.Nombre != nil {
			myMap["nombre"] = *req.Nombre
		}
		if req.Apellido != nil {
			myMap["apellido"] = *req.Apellido
		}

		dentistaRef := reflect.ValueOf(req)
		var valuesNil []string
		
		for i := 0; i < dentistaRef.NumField(); i++ {
			if e := dentistaRef.Field(i); e.IsNil() &&
				dentistaRef.Type().Field(i).Name != "ID" {
				valuesNil = append(valuesNil, dentistaRef.Type().Field(i).Name)
			}
		}
		if len(valuesNil) > 0 {
			web.Error(ctx, 422, "required fields: %s", strings.Join(valuesNil, ", "))
			return
		}
		d, err := h.dService.Update(ctx, id, myMap)
		if err != nil {
			web.Failure(d, 400, err)
			return
		}

		web.Success(ctx, 200, d)
	}
}


func (h *dentistaHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestDentista

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Failure(ctx, 422, errors.New("invalid json"))
			return
		}

		myMap := make(map[string]interface{})
		if req.Nombre != nil {
			myMap["nombre"] = *req.Nombre
		}
		if req.Apellido != nil {
			myMap["apellido"] = *req.Apellido
		}

		d, err := h.dService.Update(ctx, int(id), myMap)
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		web.Success(ctx, 200, d)
	}
}
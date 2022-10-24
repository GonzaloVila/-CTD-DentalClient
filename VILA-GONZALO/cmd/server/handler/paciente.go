package handler

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/GonzaloVila/clinica/internal/pacientes"

	"github.com/GonzaloVila/clinica/pkg/web"

	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	pService pacientes.Service
}

func NewPacienteHandler(p pacientes.Service) *pacienteHandler {
	return &pacienteHandler{
		pService: p,
	}
}

type requestPaciente struct {
	ID        int     `json:"id"`
	Apellido  *string `json:"apellido"`
	Nombre    *string `json:"nombre"`
	Domicilio *string `json:"domicilio"`
	DNI       *int    `json:"dni"`
	FechaAlta *string `json:"fecha_de_alta"`
}

func (h *pacienteHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		paciente, err := h.pService.GetByID(ctx, id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("pacient not found"))
			return
		}
		web.Success(ctx, 200, paciente)
	}
}

func (h *pacienteHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestPaciente
		if err := ctx.ShouldBindJSON(&req); err!=nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		pacienteRef := reflect.ValueOf(req)
		var valuesNil []string
		for i := 0; i < pacienteRef.NumField(); i++ {
			if e := pacienteRef.Field(i); e.IsNil() &&
				pacienteRef.Type().Field(i).Name != "ID" {
				valuesNil = append(valuesNil, pacienteRef.Type().Field(i).Name)
			}
		}
		if len(valuesNil) > 0 {
			web.Error(ctx, 422, "required fields: %s", strings.Join(valuesNil, ", "))
			return
		}
		p, err := h.pService.Create(ctx, *req.Nombre, *req.Apellido, *req.Domicilio, *req.DNI, *req.FechaAlta)
		if err != nil {
			web.Failure(p, 400, err)
			return
		}
		web.Success(ctx, 201, p)
	}
}

func (h *pacienteHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		var req requestPaciente
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
		if req.Domicilio != nil {
			myMap["domicilio"] = *req.Domicilio
		}
		if req.FechaAlta != nil {
			myMap["fecha_de_alta"] = *req.FechaAlta
		}
		pacienteRef := reflect.ValueOf(req)
		var valuesNil []string
		for i := 0; i < pacienteRef.NumField(); i++ {
			if e := pacienteRef.Field(i); e.IsNil() &&
				pacienteRef.Type().Field(i).Name != "ID" {
				valuesNil = append(valuesNil, pacienteRef.Type().Field(i).Name)
			}
		}
		if len(valuesNil) > 0 {
			web.Error(ctx, 422, "required fields: %s", strings.Join(valuesNil, ", "))
			return
		}
		p, err := h.pService.Update(ctx, id, myMap)
		if err != nil {
			web.Failure(p, 400, err)
			return
		}
		web.Success(ctx, 200, p)
	}
}

func (h *pacienteHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		err = h.pService.Delete(ctx, int(id))
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 204, "delete from paciente")
	}
}

func (h *pacienteHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestPaciente

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
		if req.Domicilio != nil {
			myMap["domicilio"] = *req.Domicilio
		}
		if req.FechaAlta != nil {
			myMap["fecha_de_alta"] = *req.FechaAlta
		}
		p, err := h.pService.Update(ctx, int(id), myMap)
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		web.Success(ctx, 200, p)
	}
}
package handler

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/GonzaloVila/clinica/internal/domain"

	"github.com/GonzaloVila/clinica/internal/turnos"

	"github.com/GonzaloVila/clinica/pkg/web"

	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	tService turnos.Service
}

func NewTurnoHandler(t turnos.Service) *turnoHandler {
	return &turnoHandler{
		tService: t,
	}
}

func (h *turnoHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		turnos, err := h.tService.GetByID(ctx, id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("turno not found"))
			return
		}
		web.Success(ctx, 200, turnos)
	}
}

func (h *turnoHandler) GetByPacienteDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Error(c, 400, "invalid dni")
			return
		}

		turnos, err := h.tService.Get(c, dni)

		if err != nil {
			if errors.Is(err, turnos.ErrNotFound) {
				web.Error(c, 404, err.Error())
				return
			}
		}
		web.Success(c, 200, turnos)
	}
}

func (h *turnoHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.Turno
		if err := ctx.ShouldBindJSON(&req); err!=nil {
			web.Failure(ctx, 422, errors.New("data cannot be processed"))
			return
		}

		turnoRef := reflect.ValueOf(req)
		var valuesNil []string
		
		for i := 0; i < turnoRef.NumField(); i++ {
			if e := turnoRef.Field(i); e.IsNil() &&
				turnoRef.Type().Field(i).Name != "ID" {
				valuesNil = append(valuesNil, turnoRef.Type().Field(i).Name)
			}
		}
		if len(valuesNil) > 0 {
			web.Error(ctx, 422, "required fields: %s", strings.Join(valuesNil, ", "))
			return
		}
		t, err := h.tService.Save(ctx, req)
		if err != nil {
			web.Failure(t, 400, err)
			return
		}
		web.Success(ctx, 201, t)
	}
}



func (h *turnoHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		var req domain.Turno
		if err := ctx.ShouldBindJSON(&req); err!=nil {
			web.Failure(ctx, 422, errors.New("data cannot be processed"))
			return
		}

		turnoRef := reflect.ValueOf(req)
		var valuesNil []string
		
		for i := 0; i < turnoRef.NumField(); i++ {
			if e := turnoRef.Field(i); e.IsNil() &&
				turnoRef.Type().Field(i).Name != "ID" {
				valuesNil = append(valuesNil, turnoRef.Type().Field(i).Name)
			}
		}
		if len(valuesNil) > 0 {
			web.Error(ctx, 422, "required fields: %s", strings.Join(valuesNil, ", "))
			return
		}
		t, err := h.tService.Update(ctx, id, req)
		if err != nil {
			web.Failure(t, 400, err)
			return
		}
		web.Success(ctx, 200, t)
	}
}

func (h *turnoHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		err = h.tService.Delete(ctx, int(id))
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 204, "delete from turno")
	}
}

func (h *turnoHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		var req domain.Turno
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Failure(ctx, 422, errors.New("data cannot be processed"))
			return
		}
		t, err := h.tService.Update(ctx, int(id), req)
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 200, t)
	}
}
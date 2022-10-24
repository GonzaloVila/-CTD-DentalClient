package routes

import (
	"database/sql"

	"github.com/GonzaloVila/clinica/cmd/server/handler"

	"github.com/GonzaloVila/clinica/internal/dentistas"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.buildDentistaRoutes()
}

func (r *router) buildDentistaRoutes() {

	repository := dentistas.NewRepository((r.db))
	service := dentistas.NewService(repository)
	h := handler.NewDentista(service)

	odonto := r.rg.Group("/dentista")
	{
		odonto.GET("/dentista/:id", h.GetDentista())
		odonto.POST("/dentista/", h.CreateDentista())
		odonto.PATCH("/dentista/:id", h.PatchDentista())
		odonto.PUT("/dentista/:id", h.PutDentista())
		odonto.DELETE("/dentista/:id", h.DeleteDentista())
	}	
}

func (r *router) buildPacienteRoutes() {

	repo := pacientes.NewRepository(r.db)
	service := pacientes.NewService(repo)
	h := handler.NewPaciente(service)

	pac := r.rg.Group("pacientes")

	pac.POST("/", h.Create())
	pac.GET("/:id", h.Get())
	pac.DELETE("/:id", h.Delete())
	pac.PUT("/:id", h.Put())
	pac.PATCH("/:id", h.Patch())
}

func (r *router) buildTurnoRoutes() {

	repo := turnos.NewRepository(r.db)
	service := turnos.NewService(repo)
	h := handler.NewTurno(service)

	tur := r.rg.Group("turnos")

	tur.POST("/", h.Create())
	tur.GET("/:id", h.Get())
	tur.GET("/dni/:dni", h.GetByPacienteDNI())
	tur.DELETE("/:id", h.Delete())
	tur.PUT("/:id", h.Put())
	tur.PATCH("/:id", h.Patch())
}

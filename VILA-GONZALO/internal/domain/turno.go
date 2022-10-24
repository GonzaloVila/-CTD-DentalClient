package domain

type Turno struct {
	ID          int     `json:"id"`
	Dentista    Dentista `json:"dentista" binding:"required"`
	Paciente    Paciente `json:"paciente"binding:"required"`
	Fecha       string   `json:"fecha" binding:"required"`
	Hora        string   `json:"hora" binding:"required"`
	Descripcion string   `json:"descripcion"`
}

type TurnoDTO struct {
	ID                int    `json:"id"`
	DentistaMatricula int    `json:"dentista_matricula"`
	PacienteDNI       int    `json:"paciente_dni"`
	Fecha             string `json:"fecha"`
	Hora              string `json:"hora"`
	Descripcion       string `json:"descripcion"`
}
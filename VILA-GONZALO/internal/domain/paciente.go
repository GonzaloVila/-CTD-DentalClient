package domain

type Paciente struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre" binding:"required"`
	Apellido    string  `json:"apellido" binding:"required"`
	Domicilio   string  `json:"domicilio" binding:"required"`
	DNI   		int    	`json:"dni" binding:"required"`
	FechaAlta	string 	`json:"fecha_de_alta"`
}
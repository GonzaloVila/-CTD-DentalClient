package turnos

import (
	"errors"
	"context"

	"github.com/GonzaloVila/clinica/internal/domain"
)

type Service interface {
	GetByID(ctx context.Context, id int) (domain.turno, error)
	Create(ctx context.Context, id int, nombre, apellido string) (domain.turno, error)
	Update(ctx context.Context, id int, d domain.turno,data map[string]interface{}) (domain.turno, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{
		r: r,
	}
}

//Trae un turno por id
func (s *service) GetByID(ctx context.Context, id int) (domain.turno, error) {
	d, err := s.r.GetByID(ctx, id)
	if err != nil {
		return domain.turno{}, errors.New("Turno not found")
	}
	return d, nil
}

//Crea un turno 
func (s *service) Create(ctx context.Context, id int, nombre, apellido, domicilio, fechaAlta string, dni int) (domain.turno, error) {
	if s.r.Exists(ctx, id) {
		return domain.turno{}, errors.New("Turno already exists")
	}
	paciente := domain.Paciente{ID: id, Nombre: nombre, Apellido: apellido, Domicilio: domicilio, DNI: dni, FechaAlta: fechaAlta}
	p, err := s.r.Create(ctx, paciente)
	if err != nil {
		return domain.Paciente{}, err
	}
	paciente.ID = p
	return paciente, nil
}

//Edita un turno
func (s *service) Update(ctx context.Context, id int, p domain.turno,data map[string]interface{}) (domain.turno, error) {
	paciente, err := s.r.GetByID(ctx, id)
	if err != nil {
		return domain.Paciente{}, errors.New("Turno not found")
	}
	if nombre, ok := data["nombre"].(string); ok && &nombre != nil {
		paciente.Nombre = nombre
	}
	if apellido, ok := data["apellido"].(string); ok && &apellido != nil {
		paciente.Apellido = apellido
	}
	if domicilio, ok := data["domicilio"].(string); ok && &domicilio != nil {
		paciente.Domicilio = domicilio
	}
	if fechaAlta, ok := data["fecha_de_alta"].(string); ok && &fechaAlta != nil {
		paciente.FechaAlta = fechaAlta
	}
	paciente, err = s.r.Update(ctx, p)
	if err != nil {
		return domain.Paciente{}, errors.New("Internal error")
	}
	return paciente, nil
}

//Elimina un turno
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.r.Delete(ctx, id)
	if err != nil {
		return errors.New("Turno not found")
	}
	return nil
}


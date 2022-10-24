package dentista

import (
	"errors"
	"context"

	"github.com/GonzaloVila/clinica/internal/domain"
)

type Service interface {
	GetByID(ctx context.Context, id int) (domain.Dentista, error)
	Create(ctx context.Context, id int, nombre, apellido string) (domain.Dentista, error)
	Update(ctx context.Context, id int, d domain.Dentista,data map[string]interface{}) (domain.Dentista, error)
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

//Trae dentista por id
func (s *service) GetByID(ctx context.Context, id int) (domain.Dentista, error) {
	d, err := s.r.GetByID(ctx, id)
	if err != nil {
		return domain.Dentista{}, errors.New("Dentist not found")
	}
	return d, nil
}


//Crea dentista
func (s *service) Create(ctx context.Context, id int, nombre, apellido string) (domain.Dentista, error) {
	if s.r.Exists(ctx, id) {
		return domain.Dentista{}, errors.New("Dentist already exists")
	}
	dentista := domain.Dentista{ID: id, Nombre: nombre, Apellido: apellido}
	d, err := s.r.Save(ctx, dentista)
	if err != nil {
		return domain.Dentista{}, err
	}
	dentista.ID = d
	return dentista, nil
}

//Edita dentista
func (s *service) Update(ctx context.Context, id int, d domain.Dentista,data map[string]interface{}) (domain.Dentista, error) {
	dentista, err := s.r.GetByID(ctx, id)
	if err != nil {
		return domain.Dentista{}, errors.New("Dentist not found")
	}
	if nombre, ok := data["nombre"].(string); ok && &nombre != nil {
		dentista.Nombre = nombre
	}
	if apellido, ok := data["apellido"].(string); ok && &apellido != nil {
		dentista.Apellido = apellido
	}

	dentista, err = s.r.Update(ctx, d)
	if err != nil {
		return domain.Dentista{}, errors.New("Internal error")
	}
	return dentista, nil
}

//Elimina dentista
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return errors.New("Dentist not found")
	}
	return nil
}


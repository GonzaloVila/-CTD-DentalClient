package dentistas

import (
	"context"
	"database/sql"
	"errors"

	"github.com/GonzaloVila/clinica/internal/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (domain.Dentista, error)
	Create(ctx context.Context, d domain.Dentista) (int, error)
	Update(ctx context.Context, d domain.Dentista) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, matricula string) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

//Busca dentista por id
func (r *repository) GetByID(ctx context.Context, id int) (domain.Dentista, error) {
	query := "SELECT * FROM dentistas WHERE id=?;"
	row := r.db.QueryRow(query, id)
	dentista := domain.Dentista{}
	err := row.Scan(&dentista.ID, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
	if err != nil {
		return domain.Dentista{}, errors.New("dentist not found")
	}
	return dentista, nil

}

//Crea un dentista 
func (r *repository) Create(ctx context.Context, d domain.Dentista) (domain.Dentista, error) {
	query := "INSERT INTO dentistas(nombre, apellido, matricula) VALUES(?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(&d.ID, &d.Nombre, &d.Apellido, &d.Matricula)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

//Edita un dentista
func (r *repository) Update(ctx context.Context, d domain.Dentista) error {
	query := "UPDATE dentistas SET nombre=?, apellido=? WHERE matricula=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(&d.Nombre, &d.Apellido, &d.Matricula)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

//Borra un dentista
func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM dentistas WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("Dentist not found")
	}

	return nil
}

//Verifica si existe
func (r *repository) Exists(ctx context.Context, matricula int) bool {
	query := "SELECT matricula FROM odontologos WHERE matricula=?"
	row := r.db.QueryRow(query, matricula)
	err := row.Scan(&matricula)
	return err == nil
}

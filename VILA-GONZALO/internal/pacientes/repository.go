package pacientes

import (
	"context"
	"database/sql"
	"errors"

	"github.com/GonzaloVila/clinica/internal/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (domain.Paciente, error)
	Create(ctx context.Context, p domain.Paciente) (int, error)
	Update(ctx context.Context, p domain.Paciente) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, id int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

//Busca un paciente por id
func (r *repository) GetByID(ctx context.Context, id int) (domain.Paciente, error) {
	query := "SELECT * FROM pacientes WHERE id=?"
	row := r.db.QueryRow(query, id)
	p := domain.Paciente{}
	err := row.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.Domicilio, &p.DNI, &p.FechaAlta)
	if err != nil {
		return domain.Paciente{}, errors.New("pacient not found")
	}
	return p, nil

}

//Crea un paciente
func (r *repository) Create(ctx context.Context, p domain.Paciente) (domain.Paciente, error) {
	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_de_alta) values (?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(&p.Nombre, &p.Apellido, &p.Domicilio, &p.DNI, &p.FechaAlta)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

//Edita un paciente
func (r *repository) Update(ctx context.Context, p domain.Paciente) error {
	query := "UPDATE pacientes SET nombre=?, apellido=?, domicilio=?, fecha_de_alta=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(&p.Nombre, &p.Apellido, &p.Domicilio, &p.FechaAlta)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

//Borra un paciente
func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM pacientes WHERE id=?"
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
		return errors.New("Pacient not found")
	}

	return nil
}

//Verificacion si existe
func (r *repository) Exists(ctx context.Context, id int) bool {
	query := "SELECT id FROM pacientes WHERE id=?"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&id)
	return err == nil
}

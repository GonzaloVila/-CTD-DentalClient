package turnos

import (
	"context"
	"database/sql"
	"errors"

	"github.com/GonzaloVila/clinica/internal/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (domain.TurnoDTO, error)
	Create(ctx context.Context, t domain.Turno) (int, error)
	Update(ctx context.Context, t domain.TurnoDTO) error
	Delete(ctx context.Context, id int) error
	GetByPacienteDNI(ctx context.Context, dni int) ([]domain.Turno, error)
	Exists(ctx context.Context, matricula, id int, fecha, hora string) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

//Trae un turno por id
func (r *repository) GetByID(ctx context.Context, id int) (domain.TurnoDTO, error) {
	query := "SELECT * FROM turnos WHERE id=?"
	row := r.db.QueryRow(query, id)
	t := domain.Turno{}
	err := row.Scan(&t.ID,&t.Fecha, &t.Hora, &t.PacienteDNI, &t.DentistaMatricula, &t.Descripcion)
	if err != nil {
		return domain.Turno{}, errors.New("Turno not found")
	}
	return t, nil

}

//Crea un turno 
func (r *repository) Create(ctx context.Context, t domain.Turno) (domain.Turno, error) {
	query := "INSERT INTO turnos (fecha, hora, paciente_dni, dentista_matricula, descripcion) values (?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(&t.Fecha, &t.Hora, &t.Paciente.DNI, &t.Dentista.Matricula, &t.Descripcion)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

//Edita un turno
func (r *repository) Update(ctx context.Context, t domain.TurnoDTO) error {
	query := "UPDATE turnos SET fecha=?, hora=?, paciente_dni=?, dentista_matricula=?, descripcion=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(&t.Fecha, &t.Hora, &t.PacienteDNI, &t.DentistaMatricula, &t.Descripcion)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

//Elimina un turno
func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM turnos WHERE id=?"
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
		return errors.New("Turnos not found")
	}
	return nil
}


//Trae un turno por id (del paciente)
func (r *repository) GetByPacienteDNI(ctx context.Context, dni int) ([]domain.Turno, error) {
	query := "SELECT t.id, t.fecha, t.hora, t.descripcion, d.id, d.nombre, d.apellido, d.matricula, p.id," + " p.nombre, p.apellido, p.domicilio, p.dni, p.fecha_de_alta FROM turnos t JOIN dentistas d ON d.matricula = t.dentista_matricula" +" JOIN pacientes p ON p.dni = t.paciente_dni WHERE t.paciente_dni=?;"
	row := r.db.QueryRow(query, dni)
	turno := domain.Turno{}
	err := row.Scan(&turno.ID, &turno.Fecha, &turno.Hora, &turno.Descripcion,
		&turno.Dentista.ID, &turno.Dentista.Nombre, &turno.Dentista.Apellido, &turno.Dentista.Matricula,
		&turno.Paciente.ID, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.DNI, &turno.Paciente.FechaAlta)
	if err != nil {
		return []domain.Turno{}, err
	}

	return []domain.Turno{turno}, nil
}

//Verifica si existe
func (r *repository) Exists(ctx context.Context, matricula, id int, fecha, hora string) bool {
	query := "SELECT dentista_matricula FROM turnos WHERE dentista_matricula=? AND paciente_dni=? AND fecha=? AND hora=?;"
	row := r.db.QueryRow(query, matricula, id, fecha, hora)
	err := row.Scan(&matricula)
	return err == nil
}

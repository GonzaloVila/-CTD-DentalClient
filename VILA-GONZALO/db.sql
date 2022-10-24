use my_db

drop table turnos if exists;
drop table dentistas if exists;
drop table pacientes if exists;

create table `odontologos`(
    `id` int not null primary key auto_increment,
    `nombre` text not null,
    `apellido` text not null,
    `matricula` int not null
);
create table `pacientes`(
    `id` int not null primary key auto_increment,
    `nombre` text not null,
    `apellido` text not null,
    `domicilio` text not null,
    `dni` int not null,
    `fecha_de_alta` text not null
);
create table `turnos`(
    `id` int not null primary key auto_increment,
    `paciente_dni` int not null, 
    `dentista_matricula` int not null,
    `fecha` text not null,
    `hora` text not null,
    `descripcion` text 
);

CREATE INDEX `dni`  ON `my_db`.`pacientes` (dni);
 ALTER TABLE turnos
 ADD FOREIGN KEY (paciente_dni) REFERENCES pacientes(dni);
CREATE INDEX `matricula`  ON `my_db`.`dentistas` (matricula); -- probar si crea 
ALTER TABLE turnos
ADD FOREIGN KEY (dentista_matricula) REFERENCES dentistas(matricula);
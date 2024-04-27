CREATE DATABASE sala;

CREATE TABLE persona (
id int(10) unsigned not null auto_increment,
nombre varchar(20) NOT NULL,
apellido varchar(20) DEFAULT NULL,
telefono varchar(40) DEFAULT NULL,
PRIMARY KEY (id)
);

CREATE TABLE sala (
id int(10) unsigned not null auto_increment,
nombre varchar(20) NOT NULL,
precio float NOT NULL,
color varchar(20) NOT NULL,
PRIMARY KEY (id)
);

CREATE TABLE grupo (
id int(10) unsigned not null auto_increment,
nombre varchar(20) NOT NULL,
PRIMARY KEY (id)
);

CREATE TABLE persona_grupo (
id int(10) unsigned not null auto_increment,
grupo_id int(10) unsigned NOT NULL,
persona_id int(10) unsigned NOT NULL,
PRIMARY KEY (id),
KEY idx_persona_grupo_grupo (grupo_id),
KEY idx_persona_grupo_persona (persona_id),
CONSTRAINT fk_persona_grupo_grupo FOREIGN KEY (grupo_id) REFERENCES grupo (id),
CONSTRAINT fk_persona_grupo_persona FOREIGN KEY (persona_id) REFERENCES persona (id)
);

CREATE TABLE sala_grupo (
id int(10) unsigned not null auto_increment,
grupo_id int(10) unsigned NOT NULL,
sala_id int(10) unsigned NOT NULL,
hora_desde datetime NOT NULL,
hora_hasta datetime NOT NULL,
PRIMARY KEY (id),
KEY idx_sala_grupo_grupo (grupo_id),
KEY idx_sala_grupo_sala (sala_id),
CONSTRAINT fk_sala_grupo_grupo FOREIGN KEY (grupo_id) REFERENCES grupo (id),
CONSTRAINT fk_sala_grupo_sala FOREIGN KEY (sala_id) REFERENCES sala (id)
);
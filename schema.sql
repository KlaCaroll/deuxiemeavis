CREATE DATABASE IF NOT EXISTS deuxiemeavis;

USE deuxiemeavis;

DROP TABLE IF EXISTS patients;
DROP TABLE IF EXISTS hospitals;
DROP TABLE IF EXISTS doctors;
DROP TABLE IF EXISTS diseases;
DROP TABLE IF EXISTS doctors_diseases;
DROP TABLE IF EXISTS requests;

-- Patients
CREATE TABLE patients (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `creation_date` datetime DEFAULT current_timestamp(),
  `modification_date` datetime DEFAULT NULL ON UPDATE current_timestamp(),
  `first_name` varchar(100) DEFAULT NULL,
  `last_name` varchar(100) DEFAULT NULL,
  `city` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO patients (first_name, last_name) VALUES ('Gaspard', 'Foucault');
INSERT INTO patients (first_name, last_name) VALUES ('Zoé', 'Belisle');
INSERT INTO patients (first_name, last_name) VALUES ('Dave', 'Boulanger');
INSERT INTO patients (first_name, last_name) VALUES ('Emilie', 'Lacasse');
INSERT INTO patients (first_name, last_name) VALUES ('Claire', 'Fournier');
INSERT INTO patients (first_name, last_name) VALUES ('Jacques', 'Gaulin');

-- Hospitals
CREATE TABLE hospitals (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `creation_date` datetime DEFAULT current_timestamp(),
  `modification_date` datetime DEFAULT NULL ON UPDATE current_timestamp(),
  `name` varchar(100) DEFAULT NULL,
  `city` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO hospitals (name, city) VALUES ('Institut Curie', 'Paris');
INSERT INTO hospitals (name, city) VALUES ('Centre Léon Bérard', 'Lyon');
INSERT INTO hospitals (name, city) VALUES ('Hopital Cochin', 'Paris');
INSERT INTO hospitals (name, city) VALUES ('Centre Antoine Lacassagne', 'Nice');
INSERT INTO hospitals (name, city) VALUES ('Institut Pasteur', 'Paris');
INSERT INTO hospitals (name, city) VALUES ('Institut Gustave Roussy', 'Créteil');

-- Doctors
CREATE TABLE doctors (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `creation_date` datetime DEFAULT current_timestamp(),
  `modification_date` datetime DEFAULT NULL ON UPDATE current_timestamp(),
  `first_name` varchar(100) DEFAULT NULL,
  `last_name` varchar(100) DEFAULT NULL,
  `hospital_id` smallint(5) unsigned DEFAULT NULL,  
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO doctors (first_name, last_name, hospital_id) VALUES ('Marie', 'Curie', 1);
INSERT INTO doctors (first_name, last_name, hospital_id) VALUES ('Léon', 'Bérard', 2);
INSERT INTO doctors (first_name, last_name, hospital_id) VALUES ('Madeleine', 'Brès', 3);
INSERT INTO doctors (first_name, last_name, hospital_id) VALUES ('Antoine', 'Lacassagne', 4);
INSERT INTO doctors (first_name, last_name, hospital_id) VALUES ('Françoise', 'Barré-Sinoussi', 5);
INSERT INTO doctors (first_name, last_name, hospital_id) VALUES ('Gustave', 'Roussy', 6);

-- Diseases
CREATE TABLE diseases (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `creation_date` datetime DEFAULT current_timestamp(),
  `modification_date` datetime DEFAULT NULL ON UPDATE current_timestamp(),
  `name` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO diseases (name) VALUES ('Fracture du poignet');
INSERT INTO diseases (name) VALUES ('Acouphènes');
INSERT INTO diseases (name) VALUES ('Cancer du thorax');

-- Doctors diseases
CREATE TABLE doctors_diseases (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `creation_date` datetime DEFAULT current_timestamp(),
  `modification_date` datetime DEFAULT NULL ON UPDATE current_timestamp(),
  `doctor_id` smallint(5) DEFAULT NULL,
  `disease_id` smallint(5) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO doctors_diseases (doctor_id, disease_id) VALUES (1, 1);
INSERT INTO doctors_diseases (doctor_id, disease_id) VALUES (1, 2);
INSERT INTO doctors_diseases (doctor_id, disease_id) VALUES (2, 2);
INSERT INTO doctors_diseases (doctor_id, disease_id) VALUES (3, 3);
INSERT INTO doctors_diseases (doctor_id, disease_id) VALUES (5, 1);
INSERT INTO doctors_diseases (doctor_id, disease_id) VALUES (6, 1);
INSERT INTO doctors_diseases (doctor_id, disease_id) VALUES (6, 2);
INSERT INTO doctors_diseases (doctor_id, disease_id) VALUES (6, 3);

-- Requests
CREATE TABLE requests (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `creation_date` datetime DEFAULT current_timestamp(),
  `modification_date` datetime DEFAULT NULL ON UPDATE current_timestamp(),
  `patient_id` smallint(5) NOT NULL,
  `doctor_id` smallint(5) NOT NULL,
  `disease_id` smallint(5) NOT NULL,
  `status` varchar(50) NOT NULL,
  `diagnosis` varchar(255) NOT NULL,
  `second_opinion` varchar(255) DEFAULT NULL,
  `second_opinion_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO requests (patient_id, doctor_id, disease_id, status, diagnosis) VALUES (1, 6, 1, 'new', 'Lorem ipsum');
INSERT INTO requests (patient_id, doctor_id, disease_id, status, diagnosis) VALUES (2, 1, 2, 'new', 'Lorem ipsum');
INSERT INTO requests (patient_id, doctor_id, disease_id, status, diagnosis) VALUES (3, 2, 2, 'archived', 'Lorem ipsum');


-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
USE `empresa_internet` ;

-- -----------------------------------------------------
-- Table `mydb`.`packs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `empresa_internet`.`packs` (
  `idpacks` INT NOT NULL,
  `velocidad` INT NULL,
  `precio` DECIMAL(2) NULL,
  `descuento` DECIMAL(2) NULL,
  PRIMARY KEY (`idpacks`))
ENGINE = InnoDB;


CREATE TABLE IF NOT EXISTS `empresa_internet`.`clientes` (
  `idclientes` INT NOT NULL,
  `DNI` VARCHAR(45) NULL,
  `Nombre` VARCHAR(45) NULL,
  `apellido` VARCHAR(45) NULL,
  `fecha_nacimiento` DATETIME NULL,
  `provincia` VARCHAR(45) NULL,
  `ciudad` VARCHAR(45) NULL,
    `pack_id` INT NOT NULL,

  PRIMARY KEY (`idclientes`),
  CONSTRAINT 
    FOREIGN KEY fk_pack_id (`pack_id`)
    REFERENCES `empresa_internet`.`packs` (`idpacks`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;


insert into packs (idpacks, velocidad, precio, descuento)
values
(1,10, 10.5, 5),
(2,100, 100, 5),
(3,1000, 1000.2, 5),
(4,300, 300.2, 5),
(5,500, 500.2, 5);

insert into clientes (idclientes,DNI,Nombre,apellido,fecha_nacimiento,provincia,ciudad,pack_id)
values
(1,"231321", "Pedro1","sanchez", now(), "salta", "salta",1),
(2,"231321", "Pedro2","sanchez", now(), "salta", "salta",2),
(3,"231321", "Pedro3","sanchez", now(), "salta", "salta",3),
(4,"231321", "Pedro4","sanchez", now(), "salta", "salta",4),
(5,"231321", "Pedro5","sanchez", now(), "salta", "salta",5),
(6,"231321", "Pedro6","sanchez", now(), "salta", "salta",4),
(7,"231321", "Pedro7","sanchez", now(), "salta", "salta",3),
(8,"231321", "Pedro8","sanchez", now(), "salta", "salta",2),
(9,"231321", "Pedro9","sanchez", now(), "salta", "salta",1),
(10,"231321", "Pedro10","sanchez", now(), "salta", "salta",5);


-- Plantear 10 consultas SQL que se podrían realizar a la base de datos. Expresar las sentencias.

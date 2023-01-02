-- MySQL Script generated by MySQL Workbench
-- Mon Jan  2 09:03:48 2023
-- Model: New Model    Version: 1.0
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
CREATE SCHEMA IF NOT EXISTS `mydb` DEFAULT CHARACTER SET utf8 ;
USE `mydb` ;

-- -----------------------------------------------------
-- Table `mydb`.`table1`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`table1` (
  `ss` INT NOT NULL,
  `table1col` VARCHAR(45) NULL,
  `df` VARCHAR(45) NULL,
  PRIMARY KEY (`ss`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`table2`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`table2` (
  `dadda` INT NOT NULL,
  PRIMARY KEY (`dadda`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`User`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`User` (
  `id` INT NOT NULL,
  `name` VARCHAR(255) NULL,
  `email` VARCHAR(255) NULL,
  `password` VARCHAR(45) NULL,
  `profile` LONGTEXT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`MentorRequirement`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`MentorRequirement` (
  `id` INT NOT NULL,
  `title` VARCHAR(255) NULL,
  `category` ENUM('programming', 'marketing', 'design', 'writing', 'movie', 'business', 'language', 'lifestyle') NULL,
  `contact_type` ENUM('one_off', 'continuous') NULL,
  `consultation_method` ENUM('chat', 'video') NULL,
  `description` VARCHAR(2000) NULL,
  `budget` INT NULL,
  `application_period` INT NULL,
  `status` ENUM('publish', 'cancel') NULL,
  `user_id` INT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `user_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`User` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`MentoringProposal`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`MentoringProposal` (
  `id` INT NOT NULL,
  `content` VARCHAR(255) NULL,
  `user_id` INT NULL,
  `mentor_requirement_id` INT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `user_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`User` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `mentor_requirement_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`MentorRequirement` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`MentoringPlan`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`MentoringPlan` (
  `id` INT NOT NULL,
  `title` VARCHAR(255) NULL,
  `category` ENUM('programming', 'marketing', 'design', 'writing', 'movie', 'business', 'language', 'lifestyle') NULL,
  `content` VARCHAR(255) NULL,
  `status` ENUM('publish', 'cancel') NULL,
  `pricing` VARCHAR(45) NULL,
  `consultation_method` ENUM('chat', 'video') NULL,
  `user_id` INT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `user_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`User` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`Contract`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`Contract` (
  `id` INT NOT NULL,
  `mentor_id` INT NULL,
  `mentee_id` INT NULL,
  `mentoring_plan_id` INT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `mentor_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`User` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `mentee_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`User` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `mentoring_plan_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`MentoringPlan` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`Career`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`Career` (
  `id` INT NOT NULL,
  `detail` VARCHAR(1000) NULL,
  `start_year` INT NULL,
  `end_year` INT NULL,
  `user_id` INT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `user_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`User` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`Skill`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`Skill` (
  `id` INT NOT NULL,
  `evaluation` INT NULL,
  `years` INT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`tag`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`tag` (
  `id` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`SkillTagging`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`SkillTagging` (
  `id` INT NOT NULL,
  `skill_id` INT NULL,
  `tag_id` INT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `skill_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`Skill` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `tag_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`tag` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`MonitoringPlanTagging`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`MonitoringPlanTagging` (
  `id` INT NOT NULL,
  `monitaring_plan_id` INT NULL,
  `tag_id` INT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `tag_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`tag` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `monitoring_plan_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`MentoringPlan` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`MentorRequirementTagging`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`MentorRequirementTagging` (
  `id` INT NOT NULL,
  `tag_id` INT NULL,
  `mentor_requirement_id` INT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `tag_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`tag` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `mentor_requirement_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`MentorRequirement` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`UserSkillBelonging`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`UserSkillBelonging` (
  `id` INT NOT NULL,
  `user_id` INT NULL,
  `skill_id` INT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `user_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`User` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `skill_id`
    FOREIGN KEY ()
    REFERENCES `mydb`.`Skill` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

USE `mydb` ;

-- -----------------------------------------------------
-- Placeholder table for view `mydb`.`view1`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`view1` (`id` INT);

-- -----------------------------------------------------
-- View `mydb`.`view1`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `mydb`.`view1`;
USE `mydb`;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
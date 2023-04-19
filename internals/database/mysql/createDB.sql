CREATE SCHEMA IF NOT EXISTS NaturalYSalvaje;

USE NaturalYSalvaje;

CREATE TABLE boat(
  id INT AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE stateRoom(
  id INT NOT NULL,
  boatId INT NOT NULL,
  PRIMARY KEY(boatId, id),
  FOREGIN KEY(boatId) REFERENCES boat(id)
)

CREATE TABLE reservation(
  id INT AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  phone VARCHAR(255) NOT NULL,
  firstDay DATETIME NOT NULL,
  lastDay DATETIME NOT NULL
  boatId INT NOT NULL,
  stateRoomId INT NOT NULL,
  
  PRIMARY KEY(id),
  FOREGIN KEY(boatId, stateRoomId) REFERENCES stateRoom(boatId,id)
)



CREATE SCHEMA IF NOT EXISTS naturalYSalvaje;

USE naturalYSalvaje;

CREATE TABLE user(
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  PRIMARY KEY(email)
);

CREATE TABLE boat(
  id INT AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  owner VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
  FOREIGN KEY(owner) REFERENCES user(email)
);

CREATE TABLE stateRoom(
  id INT NOT NULL,
  boatId INT NOT NULL,
  PRIMARY KEY(boatId, id),
  FOREIGN KEY(boatId) REFERENCES boat(id)
);

CREATE TABLE reservation(
  id INT AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  phone VARCHAR(255) NOT NULL,
  firstDay DATE NOT NULL,
  lastDay DATE NOT NULL,
  boatId INT NOT NULL,
  stateRoomId INT NOT NULL,
  
  PRIMARY KEY(id),
  FOREIGN KEY(boatId, stateRoomId) REFERENCES stateRoom(boatId,id)
);



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
  
  PRIMARY KEY(id),
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
  firstDay DATE NOT NULL,
  lastDay DATE NOT NULL,
  passengers INT,
  isOpen BOOLEAN,
  boatId INT NOT NULL,
  
  PRIMARY KEY(id),
  FOREIGN KEY(boatId) REFERENCES boat(id)
);

CREATE TABLE client(
  id INT AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  phone VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE client_reservation(
  client_id INT NOT NULL,
  reservation_id INT NOT NULL,


  PRIMARY KEY(client_id,reservation_id)
  FOREIGN KEY(client_id) REFERENCES client(id),
  FOREIGN KEY(reservation_id) REFERENCES reservation(id)
);


CREATE TABLE stateRoom_reservation(
  reservation_id INT NOT NULL,
  stateroom_id INT NOT NULL,
  boat_id INT NOT NULL,

  PRIMARY KEY(reservation_id,stateroom_id,boat_id)

  FOREIGN KEY(reservation_id) REFERENCES reservation(id),
  FOREIGN KEY(boat_id, stateroom_id) REFERENCES stateRoom(boatId,id)
);
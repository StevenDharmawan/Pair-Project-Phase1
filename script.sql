CREATE DATABASE pair_project;

CREATE TABLE User (
  UserID INTEGER PRIMARY KEY AUTO_INCREMENT,
  Email VARCHAR(255) NOT NULL,
  Password VARCHAR(255) NOT NULL,
  Nama VARCHAR(255) NOT NULL,
  Address VARCHAR(255) NOT NULL,
  PhoneNumber INT NOT NULL,
  Balance DECIMAL(10,2) NOT NULL,
  Role VARCHAR(255) NOT NULL
);

CREATE TABLE Size (
  SizeID INTEGER PRIMARY KEY AUTO_INCREMENT,
  SizeName VARCHAR(10) NOT NULL
);

CREATE TABLE Products (
  ProductsID INTEGER PRIMARY KEY AUTO_INCREMENT,
  Name VARCHAR(255) NOT NULL,
  Description VARCHAR(255) NOT NULL
);

CREATE TABLE ProductsDetail (
  ProductsDetailID INTEGER PRIMARY KEY AUTO_INCREMENT,
  ProductID INT NOT NULL,
  SizeID INT NOT NULL,
  Price DECIMAL(10,2) NOT NULL,
  Stock INT,
  FOREIGN KEY (ProductID) REFERENCES Products(ProductsID),
  FOREIGN KEY (SizeID) REFERENCES Size (SizeID)
);


-- Insert Table

INSERT INTO Products (Name, Description)
VALUES 
('T-Shirt', 'Cotton T-Shirt'), 
('Graphic T-Shirt', 'T-shirts that have some sort of image or other graphic design on them'), 
('Denim Jacket', 'a jacket that made from denim');

INSERT INTO Size (SizeName)
VALUES ('S'), ('M'), ('L'), ('XL');

INSERT INTO ProductsDetail (ProductID, SizeID, Price, Stock)
VALUES 
(1, 1, 200000, 50),
(2, 2, 350000, 40),
(3, 1, 500000, 30);


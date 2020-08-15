CREATE TABLE ProductType(
    TypeCode CHAR(8),
    TypeName VARCHAR(50) not null,
    Warehouse int,
    
    PRIMARY KEY (TypeCode)
)

CREATE TABLE Employee (
    EmployeeNumber INT,
    EmployeeName VARCHAR(50) not null,
    Male bit not NULL,
    EmployeeAddress ntext,
    ManageProductType char(8),
    ManageWarehouse INT,

    primary key (EmployeeNumber)
)
GO

create table WareHouse(
    WarehouseNumber INT,
    WarehouseAddress text,
     
    primary key (WarehouseNumber)
)
GO

CREATE table ProductLocation (
    Floor INT,
    Stall CHAR(3),
    ContainProduct CHAR(8),
    ContainProductType CHAR(8),

    PRIMARY key (Floor,Stall)
)
GO

CREATE TABLE TimeSheet (
    Employee INT,
    WorkOnFloor INT,
    WorkAtStall CHAR(3),
    NightShift BIT,

    PRIMARY KEY (Employee, WorkOnFloor, WorkAtStall),
    FOREIGN KEY (WorkOnFloor,WorkAtStall) REFERENCES ProductLocation (Floor ,Stall),
)
go  
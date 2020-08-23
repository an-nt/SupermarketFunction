CREATE DATABASE Supermarket
go

use Supermarket
go

create table Product(
	Code char(3) not null,
	Name nvarchar(50) not null,
	Type char(3),
	Warehouse int,

	primary key (Code),
	foreign key (Type) references dbo.ProductType(Code),
	foreign key (WareHouse) references dbo.Warehouse(Number),
)
go

create table ProductType(
	Code char(3),
	Name varchar(20) not null,

	Primary key (Code),
)
go

create table Employee(
	ID int,
	FullName nvarchar(40) not null,
	Male bit,
	Nationality text,
	ManageWarehouse int,
	ManageProductType char(3),
	DirectManager int,

	primary key (ID),
	foreign key (ManageWareHouse) references dbo.WareHouse(Number),
	foreign key (ManageProductType) references dbo.ProductType(Code),
	foreign key (DirectManager) references dbo.Employee(ID),
)
go

create table WareHouse(
	Number int,
	District nvarchar(10),
	City nvarchar(10) not null,

	primary key(Number),
)
go

create table ProductLocation(
	Floor int,
	Stall char(3),
	ContainProduct char(3),

	Primary key (Floor, Stall),
	Foreign key (ContainProduct) references dbo.Product(Code),
)
go

create table Timesheet(
	Employee int,
	WorkOnFloor int,
	WorkAtStall char(3),
	NightShift bit,

	Primary key (Employee, WorkOnFloor, WorkAtStall),
	Foreign key (Employee) references dbo.Employee(ID),
	Foreign key (WorkOnFloor, WorkAtStall) references dbo.ProductLocation (Floor, Stall)

)
go

create table Customer(
	Number int unique not null,
	Name nvarchar(50),
	Year int,
	Nationality	nvarchar(20),
	Male bit,
)
go

create table StockHistory (
	ProductCode char(3) not null,
	InWarehouse int not null,
	Day date,
	Time time,
	Change int,
	DoneBy int,

	primary key (ProductCode,InWarehouse,Day, Time),
	constraint FK_StockHistory_ProductCode_Product foreign key (ProductCode) references dbo.Product(Code),
	constraint FK_StockHistory_InWarehouse_Warehouse foreign key (InWarehouse) references dbo.Warehouse(Number),	constraint FK_StockHistory_InWarehouse_Warehouse foreign key (InWarehouse) references dbo.Warehouse(Number),
	constraint FK_StockHistory_DoneBy_Employee foreign key (DoneBy) references dbo.Employee(ID),
)
go

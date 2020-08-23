use Supermarket
go

-- Add data to dbo.WareHouse
insert dbo.WareHouse (Number,District,City) 
values (1, N'Tân Bình', N'Sài Gòn')

insert dbo.WareHouse (Number,District,City) 
values (2, N'7', N'Sài Gòn')

insert dbo.WareHouse (Number,District,City) 
values (3, N'Dĩ An', N'Bình Dương')

--Add data to dbo.ProductType
insert dbo.ProductType (Code, Name)
values ('T01','Electrical Device')

insert dbo.ProductType (Code, Name)
values ('T02','Food')

insert dbo.ProductType (Code, Name)
values ('T03','Health Product')

insert dbo.ProductType (Code, Name)
values ('T04','Book')

insert dbo.ProductType (Code, Name)
values ('T05','Clothe')


--Add data to dbo.Product
insert dbo.Product (Code, Name, Type, Warehouse)
values (SUBSTRING(REPLACE(CONVERT(varchar(36), NEWID()), '-', ''), 1, 3), N'Iphone 11', 'T01',1)


insert dbo.Product (Code, Name, Type, Warehouse)
values ('P02', N'Galaxy Note 10', 'T01',2)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P03', N'Macbook Pro 13inch 2020', 'T01',3)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P04', N'Ipad Air 2020', 'T01',1)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P05', N'AirPod Pro', 'T01',2)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P06', N'Pizza', 'T02',3)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P07', N'Hot Dog', 'T02',1)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P08', N'Sushi', 'T02',2)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P09', N'Spagetti', 'T02',3)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P10', N'Instance Noodle', 'T02',1)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P11', N'Clear men', 'T03',2)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P12', N'Nivea', 'T03',3)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P13', N'Kotex', 'T03',1)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P14', N'Dove', 'T03',2)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P15', N'Face Mask', 'T03',3)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P16', N'Quần Việt Tiến Nam', 'T04',1)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P17', N'Quần An Phước Nam', 'T04',2)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P18', N'Áo khoác Uniqlo', 'T04',3)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P19', N'Giày Nike', 'T04',2)

insert dbo.Product (Code, Name, Type, Warehouse)
values ('P20', N'Túi Gucci', 'T04',3)

--Add data to dbo.Employee
insert dbo.Employee (ID, FullName, Male, Nationality, ManageWarehouse, ManageProductType,DirectManager)
values (1, N'' ,'' ,'' ,null, null, null)

--Add data to dbo.StockHistory
INSERT INTO [dbo].[StockHistory] ([ProductCode],[InWarehouse],[Day],[Time],[Change])
VALUES ('P01',1,getdate(),Getdate(),-10)
GO



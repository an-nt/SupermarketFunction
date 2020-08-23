USE [Supermarket]
GO
-- Delete Foreign key and column ManageWarehouse on table Employee
ALTER TABLE [dbo].[Employee] DROP CONSTRAINT [FK__Employee__Manage__6477ECF3]
GO

alter table dbo.Employee drop column ManageWarehouse
go
--Add Foreign key on table Warehouse
alter table dbo.Warehouse add constraint FK_ManagedBy_Employee foreign key (ManagedBy) references dbo.Employee(ID)

--Add new column on table ProductType and foreign key
alter table dbo.ProductType add ManagedBy int 
alter table dbo.ProductType add constraint FK_ProductType_ManagedBy_Employee foreign key (Managedby) references dbo.Employee(ID)


use Supermarket
go

--Create function to return the present stock of a product in a warehouse
create function UF_GetBalance(@Product char(3),@Warehouse int) 
returns int
begin
	declare @Balance int
	select @Balance = sum(Change) from dbo.StockHistory
		where ProductCode = @Product 
		and InWarehouse = @Warehouse
	return @Balance
end
go

--Demo function
select dbo.UF_GetBalance('P01',1)

--Create function to show stock history within a given period of time
create function UF_GetStockHistory(@StartDay date,@EndDay date,@Product char(3),@Warehouse int)
returns table
as return
	select * from dbo.StockHistory
		where Day >= @StartDay and Day <= @EndDay
		and ProductCode = @Product
		and InWarehouse = @Warehouse
go

--Demo function
select * from dbo.UF_GetStockHistory('2020-08-22','2020-08-23','P01',2)

--drop function UF_GetBalance
--drop function UF_GetStockHistory
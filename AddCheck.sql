declare @LastBalance int 
select @LastBalance = sum(Change) from dbo.StockHistory
print @LastBalance
go
alter table dbo.StockHistory add constraint CK_Balance check (
	sum(Change) >= 0
)
--Show Balance
select StockHistory.InWarehouse, StockHistory.ProductCode, sum(StockHistory.Change) from dbo.StockHistory
group by StockHistory.InWarehouse, StockHistory.ProductCode

select StockHistory.InWarehouse, Product.Name, sum(StockHistory.Change) from dbo.StockHistory, dbo.Product
where StockHistory.ProductCode = Product.Code
group by StockHistory.InWarehouse, StockHistory.ProductCode, Product.Name


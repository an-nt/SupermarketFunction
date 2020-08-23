alter trigger UTG_Change_StockHistory
on dbo.StockHistory
for insert
as
begin
	declare @balance int

	select @balance = sum(StockHistory.Change) from dbo.StockHistory, inserted
	where dbo.StockHistory.ProductCode = inserted.ProductCode
		and dbo.StockHistory.InWarehouse = inserted.InWarehouse

	if @balance < 0
	begin
		print N'Không được xuất quá số lượng hiện có'
		rollback tran
	end
end
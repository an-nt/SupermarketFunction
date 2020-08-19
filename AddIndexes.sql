USE [Supermarket]

GO

CREATE NONCLUSTERED INDEX [NonClusteredIndex-Type] ON [dbo].[Product]
(
	[Type]
)
INCLUDE([Name])

GO


-- Index Seek (NonClustered)
select Type, Name from dbo.Product
where Type = 'T01'

-- Clustered Index Scan
select * from dbo.Product
where Type = 'T02'
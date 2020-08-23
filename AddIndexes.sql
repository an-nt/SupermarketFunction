USE [Supermarket]

GO

CREATE NONCLUSTERED INDEX [NonClusteredIndex-Type] ON [dbo].[Product]
(
	[Type]
)
INCLUDE([Name])

GO
select count(*) from dbo.Product

-- Nonclustered Index Scan
select Type from dbo.Product
where Name = 'Pizza'

-- Index Seek (NonClustered)
select Name from dbo.Product
where Type = 'T03'

-- Index Seek (NonClustered)
select Type, Name from dbo.Product
where Type = 'T01'

-- Index Seek (NonClustered)
select Code, Name, Type from dbo.Product
where Type = 'T04'

-- Clustered Index Scan
select Code, Name, Type, Warehouse from dbo.Product
where Type = 'T04'

-- NonClustered Index Seek --> Key Lookup (Clustered)
select Code, Name, Type, Warehouse
from dbo.Product
with (index([NonClusteredIndex-Type]))
where Type = 'T04'

-- Clustered Index Scan
select * from dbo.Product
where Type = 'T02'


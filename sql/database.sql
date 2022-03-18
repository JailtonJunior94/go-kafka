CREATE DATABASE KafkaPoC
GO

USE KafkaPoC
GO

EXEC sys.sp_cdc_enable_db
GO

CREATE TABLE dbo.Customers
(
    [Id] INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
    [Name] VARCHAR (50) NOT NULL,
    [Email] VARCHAR (100) NOT NULL,
    [CreatedAt] DATETIME NOT NULL,
    [UpdatedAt] DATETIME NULL,
    [Active] BIT NOT NULL
);

EXEC sys.sp_cdc_enable_table
@source_schema = N'dbo',
@source_name   = N'Customers',
@role_name     = N'Admin',
@supports_net_changes = 1
GO

EXEC sys.sp_cdc_help_change_data_capture 
GO

INSERT INTO dbo.Customers
VALUES ('Antony', 'antony.teixeira@outlook.com', GETDATE(), NULL, 1)

SELECT * FROM dbo.Customers (NOLOCK)
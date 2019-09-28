-- Create a new database called 'account'
-- Connect to the 'master' database to run this snippet
USE master
GO
-- Create the new database if it does not exist already
IF NOT EXISTS (
  SELECT name
    FROM sys.databases
    WHERE name = N'account'
)
CREATE DATABASE account
GO

-- Create a new table called 'app' in schema 'SchemaName'
-- Drop the table if it already exists
IF OBJECT_ID('SchemaName.app', 'U') IS NOT NULL
DROP TABLE SchemaName.app
GO
-- Create the table in the specified schema
CREATE TABLE SchemaName.app
(
  AppId INT NOT NULL PRIMARY KEY, -- primary key column
  UserDomain [NVARCHAR](50) NOT NULL,
  AppSecret [NVARCHAR](50) NOT NULL
  -- specify more columns here
);
GO
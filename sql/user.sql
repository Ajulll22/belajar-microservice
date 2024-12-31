-- USE [master]
-- GO

-- DROP DATABASE [service-user]
-- GO
CREATE DATABASE [service-user]
GO

USE [service-user]
GO

CREATE TABLE [user] (
	id INT IDENTITY(1,1) PRIMARY KEY,
	username VARCHAR(100) NOT NULL UNIQUE,
	email VARCHAR(100) NOT NULL UNIQUE,
	phone_number VARCHAR(100) NOT NULL UNIQUE,
	dob VARCHAR(30) NOT NULL,
	password VARCHAR(255) NOT NULL,
	created_at DATETIME DEFAULT GETDATE(),
	updated_at DATETIME DEFAULT GETDATE()
)
GO

CREATE TABLE user_address (
	id INT IDENTITY(1,1) PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	address VARCHAR(255) NOT NULL,
	note VARCHAR(100),
	user_id INT NOT NULL FOREIGN KEY REFERENCES [user](id),
	created_at DATETIME DEFAULT GETDATE(),
	updated_at DATETIME DEFAULT GETDATE()
)
GO

INSERT INTO [user] ( username, email, phone_number, dob, password )
VALUES ( 'administrator', 'ajulrizki@gmail.com', '089503941064', '2000-07-22', '$2a$10$Lv7PbUZL9SsVWelin1HtcuJf4cOlTxOB6Rb65h01i/pGOoNAmXI4e' )
GO

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Get user data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_user_data]
	-- Add the parameters for the stored procedure here
	@id INT
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	SELECT 
		*
	FROM [user]
	WHERE (
			@id = 0
			OR
			(
				@id <> 0
				AND
				id = @id
			)
		)

END
GO


SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Get user data by username
-- =============================================
CREATE PROCEDURE [dbo].[spMS_user_data_by_username]
	-- Add the parameters for the stored procedure here
	@username VARCHAR(100)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	SELECT 
		*
	FROM [user]
	WHERE username = @username

END
GO
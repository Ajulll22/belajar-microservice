-- USE [master]
-- GO

-- DROP DATABASE [service-product]
-- GO
CREATE DATABASE [service-product]
GO

USE [service-product]
GO

CREATE TABLE category (
	id INT IDENTITY(1,1) PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	created_at DATETIME DEFAULT GETDATE(),
	updated_at DATETIME DEFAULT GETDATE(),
	deleted_at DATETIME
)
GO

CREATE TABLE product (
	id INT IDENTITY(1,1) PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	price INT NOT NULL DEFAULT 0,
	stock INT NOT NULL DEFAULT 0,
	description TEXT,
	created_at DATETIME DEFAULT GETDATE(),
	updated_at DATETIME DEFAULT GETDATE(),
	deleted_at DATETIME
)
GO

CREATE TABLE product_picture (
	id INT IDENTITY(1,1) PRIMARY KEY,
	url VARCHAR(100),
	product_id INT NOT NULL FOREIGN KEY REFERENCES product(id)
)
GO

CREATE TABLE tran_product_category (
	product_id INT NOT NULL FOREIGN KEY REFERENCES product(id),
	category_id INT NOT NULL FOREIGN KEY REFERENCES category(id)
)
GO

INSERT INTO category(name)
VALUES ('Alas Kaki'),
	('Baju'),
	('Celana'),
	('Sneakers')
GO

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Get product data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_product_data]
	-- Add the parameters for the stored procedure here
	@id INT
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	SELECT 
		a.id AS product_id, a.name AS product_name, a.price AS product_price, a.stock AS product_stock, a.description AS product_description, a.created_at, a.updated_at,
		b.id AS picture_id, b.url AS picture_url,
		d.id AS category_id, d.name AS category_name
	FROM product a
	JOIN product_picture b ON a.id = b.product_id
	JOIN tran_product_category c ON a.id = c.product_id
	JOIN category d ON c.category_id = d.id
	WHERE a.deleted_at IS NULL AND d.deleted_at IS NULL
		AND
		(
			@id = 0
			OR
			(
				@id <> 0
				AND
				a.id = @id
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
-- Description:	Insert product data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_product_data_insert]
	-- Add the parameters for the stored procedure here
	@name VARCHAR(255), 
	@price INT, 
	@stock INT, 
	@description TEXT
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	INSERT INTO product( name, price, stock, description )
	VALUES ( @name, @price, @stock, @description )

	SELECT 
		* 
	FROM product
	WHERE id = SCOPE_IDENTITY()

END
GO


SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Update product data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_product_data_update]
	-- Add the parameters for the stored procedure here
	@id INT,
	@name VARCHAR(255), 
	@price INT, 
	@stock INT, 
	@description TEXT
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	UPDATE product
	SET name = @name,
		price = @price,
		stock = @stock,
		description = @description,
		updated_at = GETDATE()
	WHERE id = @id

END
GO


SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Delete product data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_product_data_delete]
	-- Add the parameters for the stored procedure here
	@id INT
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	UPDATE product
	SET deleted_at = GETDATE()
	WHERE id = @id

END
GO

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Get category data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_category_data]
	-- Add the parameters for the stored procedure here
	@id INT
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	SELECT 
		id, name, created_at, updated_at
	FROM category
	WHERE deleted_at IS NULL
		AND
		(
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
-- Description:	Insert category data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_category_data_insert]
	-- Add the parameters for the stored procedure here
	@name VARCHAR(100)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	INSERT INTO category (name)
	VALUES (@name)

	SELECT 
		* 
	FROM category
	WHERE id = SCOPE_IDENTITY()

END
GO

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Update category data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_category_data_update]
	-- Add the parameters for the stored procedure here
	@id INT,
	@name VARCHAR(100)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	UPDATE category
	SET name = @name,
		updated_at = GETDATE()
	WHERE id = @id

END
GO


SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Delete category data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_category_data_delete]
	-- Add the parameters for the stored procedure here
	@id INT
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	UPDATE category
	SET deleted_at = GETDATE()
	WHERE id = @id

END
GO

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Insert product category data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_product_category_data_insert]
	-- Add the parameters for the stored procedure here
	@product_id INT,
	@array_category_id VARCHAR(MAX)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	INSERT INTO tran_product_category (product_id, category_id)
	SELECT 
		@product_id, value 
	FROM OPENJSON(@array_category_id)

END
GO

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Update product category data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_product_category_data_update]
	-- Add the parameters for the stored procedure here
	@product_id INT,
	@array_category_id VARCHAR(MAX)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	DECLARE @tmp_tran_category TABLE (
		product_id INT,
		category_id INT
	)

	INSERT INTO @tmp_tran_category
	SELECT @product_id, value 
	FROM OPENJSON(@array_category_id)

	MERGE tran_product_category AS target
	USING @tmp_tran_category AS source
	ON target.category_id = source.category_id AND target.product_id = source.product_id

	WHEN NOT MATCHED BY target AND @product_id = source.product_id THEN
	INSERT (product_id, category_id)
	VALUES (@product_id, source.category_id)
	
	WHEN NOT MATCHED BY source AND target.product_id = @product_id THEN
	DELETE;

END
GO

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Insert product picture data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_product_picture_data_insert]
	-- Add the parameters for the stored procedure here
	@product_id INT,
	@array_url VARCHAR(MAX)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	INSERT INTO product_picture(product_id, url)
	SELECT 
		@product_id, value 
	FROM OPENJSON(@array_url)

	SELECT * FROM product_picture
	WHERE product_id = @product_id

END
GO

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Update product picture data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_product_picture_data_update]
	-- Add the parameters for the stored procedure here
	@product_id INT,
	@array_url VARCHAR(MAX)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	DECLARE @tmp_picture TABLE (
		product_id INT,
		url VARCHAR(100)
	)
	INSERT INTO @tmp_picture
	SELECT @product_id, value FROM OPENJSON(@array_url)

	MERGE product_picture AS target
	USING @tmp_picture AS source
	ON target.url = source.url AND target.product_id = source.product_id

	WHEN NOT MATCHED BY target AND @product_id = source.product_id THEN
	INSERT (product_id, url)
	VALUES (@product_id, source.url)
	
	WHEN NOT MATCHED BY source AND target.product_id = @product_id THEN
	DELETE;

END
GO

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Ajulll
-- Create date: 20241002
-- Description:	Insert product picture data
-- =============================================
CREATE PROCEDURE [dbo].[spMS_product_picture_data_delete]
	-- Add the parameters for the stored procedure here
	@array_id VARCHAR(MAX)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	DELETE FROM product_picture
	WHERE id IN (
		SELECT value FROM OPENJSON(@array_id)
	)

END
GO
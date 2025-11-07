ALTER TABLE facilities
ALTER COLUMN icon TYPE varchar(255) USING icon::text;

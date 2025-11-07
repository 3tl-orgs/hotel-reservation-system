ALTER TABLE facilities
ALTER COLUMN icon TYPE jsonb USING to_jsonb(icon);

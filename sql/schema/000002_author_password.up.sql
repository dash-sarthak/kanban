-- Add password field to author
ALTER TABLE authors ADD COLUMN
password TEXT UNIQUE NOT NULL
DEFAULT md5(random()::text);

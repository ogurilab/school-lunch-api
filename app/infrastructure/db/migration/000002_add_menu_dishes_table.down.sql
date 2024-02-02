DROP TABLE IF EXISTS menu_dishes;

ALTER TABLE dishes
ADD COLUMN menu_id varchar(255) NOT NULL;
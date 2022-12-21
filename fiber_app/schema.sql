DROP TABLE IF EXISTS MENUS;
DROP TABLE IF EXISTS RESTAURANTS;
DROP TABLE IF EXISTS RESTAURANT_TYPES;
DROP TABLE IF EXISTS USERS;

CREATE TABLE USERS (
    id serial primary key,
    login_id varchar(20) unique not null,
    login_password bytea not null,
    first_name varchar(100),
    last_name varchar(100),
    email varchar(50) not null
);

CREATE TABLE RESTAURANT_TYPES (
    restaurant_type varchar(20) primary key
);

CREATE TABLE RESTAURANTS (
    id serial primary key,
    user_id int not null references USERS(id) on delete cascade,
    name varchar(30) unique,
    restaurant_type varchar(20) references RESTAURANT_TYPES(restaurant_type) on delete set null,
    origin varchar(1000),
    menu_types varchar(20)[]
);

CREATE TABLE MENUS (
    id serial primary key,
    restaurant_id int references RESTAURANTS(id) on delete cascade,
    name varchar(30) unique,
    menu_type varchar(20),
    images varchar[],
    contents varchar(1000),
    allergies varchar(500)
);

DROP INDEX IF EXISTS idx_login_id;
DROP INDEX IF EXISTS idx_user_id;

CREATE INDEX idx_login_id on USERS using hash (login_id);
CREATE INDEX idx_user_id on RESTAURANTS using hash (user_id);

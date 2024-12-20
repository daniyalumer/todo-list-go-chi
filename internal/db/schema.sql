CREATE DATABASE todolist

CREATE TABLE todo{
    id int primary key,
    descr varchar(255) not null,
    completed boolean default=false,
    user_id int not null foreign key references user(id),
    created_at datetime,
    updated_at datetime,
    completed_at datetime,
    deleted_at datetime,
}

CREATE TABLE user{
    id int primary key,
    username varchar(255) not null,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
}
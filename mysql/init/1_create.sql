USE sample;


create table users
  (id int, name varchar(64), email varchar(64), created_at timestamp not null default current_timestamp, updated_at timestamp not null default current_timestamp on update current_timestamp, primary key(id));

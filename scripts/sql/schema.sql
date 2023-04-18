create table if not exists schools (
  id serial primary key, 
  name varchar not null, 
  type varchar not null, 
  logo varchar not null default '',
  createdAt timestamp default now(),
  updatedAt timestamp default now()
);
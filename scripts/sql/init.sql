create database gouniversity;

create user user_gouniversity;

alter user user_gouniversity with encrypted password '1234';

grant all privileges on database gouniversity to user_gouniversity;
grant all privileges on all tables in schema public to user_gouniversity;
grant all privileges on all sequences in schema public to user_gouniversity;
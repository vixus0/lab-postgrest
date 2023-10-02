create schema api;

create table api.models (
    id serial primary key,
    name text not null,
    description text not null,
    owner text not null default current_setting('request.jwt.claims', true)::json->>'username'
);

create table api.tags (
    id serial primary key,
    name text not null
);

create table api.model_tags (
    model_id int not null references api.models(id),
    tag_id int not null references api.tags(id),
    primary key (model_id, tag_id),
    tagger text not null default current_setting('request.jwt.claims', true)::json->>'username'
);

insert into api.tags (name) values ('car'), ('toy'), ('tool');

create role api_user with login password 'api_user';
grant usage on schema api to api_user;

grant 
  select
, insert
, update
, delete 
on all tables in schema api to api_user;

create role api_anon with nologin;
grant usage on schema api to api_anon;
grant select on all tables in schema api to api_anon;

create role authenticator with noinherit nocreatedb nocreaterole nosuperuser login password 'authenticator';
grant api_anon to authenticator;

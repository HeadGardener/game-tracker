CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE games
(
    id       serial       not null unique,
    title    varchar(255) not null,
    platform varchar(50)  not null,
    status   varchar(50)  not null,
    notes    varchar(50)
);

CREATE TABLE users_games
(
    id      serial                                      not null unique,
    user_id int references users (id) on delete cascade not null,
    game_id int references games (id) on delete cascade not null
);
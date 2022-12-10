CREATE TABLE IF NOT EXISTS monster (
    id BIGSERIAL PRIMARY KEY,
    name varchar(100) not null,
    sub_name varchar(155) not null,
    description text not null,
    height double precision default 0.0 not null,
    weight integer not null,
    image varchar(155) not null,
    type int[] not null,
    hit_point_stat int not null,
    attach_stat int not null,
    defence_stat int not null,
    speed_stat int not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE IF NOT EXISTS user (
    id BIGSERIAL PRIMARY KEY,
    name varchar(100) not null,
    email varchar(100) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE IF NOT EXISTS user_monster_link (
    id BIGSERIAL PRIMARY KEY,
    moster_id BIGINT not null,
    user_id BIGINT not null,
    capture_status boolean not null,
    created_at timestamp not null
);
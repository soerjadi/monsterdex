CREATE TABLE IF NOT EXISTS monster (
    id BIGSERIAL PRIMARY KEY,
    name varchar(100) not null,
    tag_name varchar(155) not null,
    description text not null,
    height double precision default 0.0 not null,
    weight integer not null,
    image varchar(155) not null,
    type int[] not null,
    hit_point_stat int not null,
    attack_stat int not null,
    defence_stat int not null,
    speed_stat int not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name varchar(100) not null,
    email varchar(100) not null,
    password varchar(100) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE IF NOT EXISTS user_monster_link (
    id BIGSERIAL PRIMARY KEY,
    monster_id BIGINT not null,
    user_id BIGINT not null,
    capture_status boolean not null,
    created_at timestamp not null,
    FOREIGN KEY (monster_id) REFERENCES monster(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS access_token (
    token text PRIMARY KEY not null,
    user_id BIGINT not null,
    created_at timestamp not null,
    valid_thru timestamp not null,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX user_id_index ON user_monster_link USING btree(user_id);
CREATE INDEX user_id_monster_id_index ON user_monster_link USING btree(user_id, monster_id);

CREATE INDEX name_index ON users USING btree(name);

CREATE INDEX name_index ON monster USING btree(name);
CREATE INDEX type_index ON monster USING btree(type);

CREATE INDEX token_access_token ON access_token USING btree(token);
CREATE INDEX user_id_access_token ON access_token USING btree(user_id);
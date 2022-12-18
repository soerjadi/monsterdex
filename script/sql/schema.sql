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
    role int not null default(0),
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE IF NOT EXISTS user_monster_link (
    monster_id BIGINT not null,
    user_id BIGINT not null,
    capture_status boolean not null,
    created_at timestamp not null,
    FOREIGN KEY (monster_id) REFERENCES monster(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (monster_id, user_id)
);

CREATE TABLE IF NOT EXISTS access_token (
    token text PRIMARY KEY not null,
    user_id BIGINT not null,
    created_at timestamp not null,
    valid_thru timestamp not null,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX name_index ON users USING btree(name);

CREATE INDEX name_monster_index ON monster USING btree(name);
CREATE INDEX type_index ON monster USING btree(type);

CREATE INDEX token_access_token ON access_token USING btree(token);
CREATE INDEX user_id_access_token ON access_token USING btree(user_id);

INSERT INTO monster(id, name, tag_name, description, height, weight, image, type, hit_point_stat, attack_stat, defence_stat, speed_stat, created_at, updated_at)
    VALUES(1, 'Lugia', 'Diving Monster', 'In eleifend nec dui vitae condimentum. Maecenas varius euismod orci, id facilisis nisi tincidunt id. Lorem ipsum dolor sit amet, consectetur adipiscing elit.', 5.1, 216, 'https://assets.pokemon.com/assets/cms2/img/pokedex/full/249.png', '{2,3}', 350, 420, 280, 480, now(), now());

INSERT INTO users(id, name, email, password, created_at, updated_at) VALUES(1, 'pokemon', 'pokemon@monsterdex.com', '$2a$10$9Lfiyg7MEpz32Aa/mgEZd.Kw0aDu0qqHImllNtOOKl88jO0EsmsXm', now(), now());

INSERT INTO access_token(token, user_id, created_at, valid_thru) VALUES('fe3d771d45fe33c33de003af028149ece72b14f386da5f2f2a7ccc942e2f5cb090702091dbb1488aae45a3076e8ed5fa06b6aa21beaa488ddaa33c0d036fb09a', 1, now(), now() + interval '7' day);

INSERT INTO user_monster_link(monster_id, user_id, capture_status, created_at) VALUES(1, 1, true, now());


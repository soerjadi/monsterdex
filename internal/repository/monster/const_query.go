package monster

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	getListMonster   *sqlx.Stmt
	getDetailMonster *sqlx.Stmt
	captureMonster   *sqlx.Stmt

	insertMonster *sqlx.Stmt
	updateMonster *sqlx.Stmt
	deleteMonster *sqlx.Stmt
}

const (
	getListMonster = `
		SELECT 
			id,
			name,
			tag_name,
			description,
			height,
			weight,
			image,
			type,
			hit_point_stat,
			attack_stat,
			defence_stat,
			speed_stat,
			COALESCE((
				SELECT 
					capture_status 
				FROM user_monster_link uml 
				WHERE
					user_id = $1
			), false) as captured,
			created_at,
			updated_at
		FROM 
			monster m
	`

	getDetailMonster = `
	SELECT 
		m.id,
		m.name,
		m.tag_name,
		m.description,
		m.height,
		m.weight,
		m.image,
		m.type,
		m.hit_point_stat,
		m.attack_stat,
		m.defence_stat,
		m.speed_stat,
		COALESCE((
			SELECT 
				capture_status 
			FROM user_monster_link uml 
			WHERE
				user_id = $1
		), false) as captured,
		m.created_at,
		m.updated_at
	FROM 
		monster m
	WHERE
		id = $2
	`

	captureMonster = `
	INSERT INTO user_monster_link (
		monster_id,
		user_id,
		capture_status,
		created_at
	) VALUES (
		$1,
		$2,
		$3,
		NOW()
	) ON CONFLICT (monster_id, user_id) DO 
	UPDATE SET
		capture_status=$3
	RETURNING monster_id, user_id, capture_status
	`

	insertMonster = `
	INSERT INTO monster(
		name,
		tag_name,
		description,
		height,
		weight,
		image,
		type,
		hit_point_stat,
		attack_stat,
		defence_stat,
		speed_stat,
		created_at,
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		NOW(),
		NOW()
	) RETURNING 
		id, 
		name,
		tag_name,
		description,
		height,
		weight,
		image,
		type,
		hit_point_stat,
		attack_stat,
		defence_stat,
		speed_stat,
		created_at,
		updated_at
	`

	updateMonster = `
	UPDATE monster 
	SET
		name=$1,
		tag_name=$2,
		description=$3,
		height=$4,
		weight=$5,
		image=$6,
		type=$7,
		hit_point_stat=$8,
		attack_stat=$9,
		defence_stat=$10,
		speed_stat=$11,
		updated_at=NOW()
	WHERE
		id=$12
	`

	deleteMonster = `
	DELETE FROM monster WHERE id=$1
	`
)

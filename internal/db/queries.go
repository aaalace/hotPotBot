package db

// User

const UpdCorrectUsername = `UPDATE users
	SET telegram_username = :telegram_username
	WHERE telegram_id = :telegram_id`

const SelectUserQuery = `SELECT * FROM users WHERE telegram_id = $1`

const AddUserQuery = `INSERT INTO users (telegram_id, telegram_username)
	VALUES (:telegram_id, :telegram_username)`

const SelectUserByUsernameQuery = `SELECT * FROM users WHERE telegram_username = $1`

const CountUserWeight = `SELECT COALESCE(SUM(c.weight * uc.quantity), 0) AS weight
	FROM 
		user_cards uc
	JOIN 
		cards c ON uc.card_id = c.id
	WHERE 
		uc.user_id = $1`

// Cooldown

const GetCooldown = `SELECT * FROM cooldowns WHERE user_id = $1`

const UpdateCooldown = `UPDATE cooldowns SET next_accept = $1 WHERE user_id = $2`

const AddCooldown = `INSERT INTO cooldowns (user_id, next_accept)
	VALUES (:user_id, :next_accept)`

// User + Card

const GiveUserRandomCard = `INSERT INTO user_cards (user_id, card_id, quantity)
	VALUES (:user_id, :card_id, 1) 
	ON CONFLICT (user_id, card_id) 
	DO UPDATE SET quantity = user_cards.quantity + 1`

// Card

const SelectCardById = `SELECT * FROM cards WHERE id = $1`

const SelectAllCardsIds = `SELECT id FROM cards`

// CardType

const SelectTypeNameById = `SELECT name FROM card_types WHERE id = $1`

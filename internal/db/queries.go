package db

const SelectUserQuery = `SELECT * FROM users WHERE telegram_id = $1`

const AddUserQuery = `INSERT INTO users (telegram_id)
	VALUES (:telegram_id)`

const GiveUserRandomCard = `INSERT INTO user_cards (user_id, card_id, quantity)
	VALUES (:user_id, :card_id, 1) 
	ON CONFLICT (user_id, card_id) 
	DO UPDATE SET quantity = user_cards.quantity + 1`

const SelectCardById = `SELECT * FROM cards WHERE id = $1`

const SelectAllCardsIds = `SELECT id FROM cards`

const SelectTypeNameById = `SELECT name FROM card_types WHERE id = $1`

const CountUserWeight = `SELECT COALESCE(SUM(c.weight * uc.quantity), 0) AS weight
	FROM 
		user_cards uc
	JOIN 
		cards c ON uc.card_id = c.id
	WHERE 
		uc.user_id = $1`

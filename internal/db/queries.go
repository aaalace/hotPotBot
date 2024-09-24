package db

// -----USERS-----

const SelectUserByLocalIdQuery = `SELECT * FROM users WHERE id = $1`

const SelectUserQuery = `SELECT * FROM users WHERE telegram_id = $1`

const SelectUserByUsernameQuery = `SELECT * FROM users WHERE telegram_username = $1`

const AddUserQuery = `INSERT INTO users (telegram_id, telegram_username)
	VALUES (:telegram_id, :telegram_username)`

const UpdCorrectUsername = `UPDATE users SET telegram_username = :telegram_username
    WHERE telegram_id = :telegram_id`

// -----CARDS-----

const SelectAllCardsIds = `SELECT id FROM cards`

const SelectCardById = `SELECT * FROM cards WHERE id = $1`

const SelectCardsByType = `SELECT c.*
	FROM cards c
	WHERE c.type_id = $1 OR $1 = 0`

const SelectAllCardsIdsByType = `SELECT id FROM cards WHERE type_id = $1`

// -----USER_CARDS-----

const SelectUserCardsByType = `SELECT c.*
	FROM cards c
	JOIN user_cards uc ON c.id = uc.card_id
	JOIN users u ON uc.user_id = u.id
	WHERE u.id = $1
  		AND (c.type_id = $2 OR $2 = 0)`

const SelectUserDuplicates = `SELECT c.*
	FROM cards c
	JOIN user_cards uc ON c.id = uc.card_id
	JOIN users u ON uc.user_id = u.id
	WHERE u.id = $1
  		AND uc.quantity > 1`

const SelectUserDuplicatesByType = `SELECT c.*
	FROM cards c
	JOIN user_cards uc ON c.id = uc.card_id
	JOIN users u ON uc.user_id = u.id
	WHERE u.id = $1
  		AND uc.quantity > 1 AND c.type_id = $2`

const GiveUserCard = `INSERT INTO user_cards (user_id, card_id, quantity)
	VALUES (:user_id, :card_id, 1) 
	ON CONFLICT (user_id, card_id) 
	DO UPDATE SET quantity = user_cards.quantity + 1`

const CountUserWeight = `SELECT COALESCE(SUM(c.weight * uc.quantity), 0) AS weight
	FROM 
		user_cards uc
	JOIN 
		cards c ON uc.card_id = c.id
	WHERE 
		uc.user_id = $1`

const SelectUserCardQuantity = `SELECT quantity FROM user_cards WHERE user_id = $1 AND card_id = $2`

const MinusUserCardQuantity = `UPDATE user_cards SET quantity = quantity - :to_remove
    WHERE user_id = :user_id AND card_id = :card_id`

const DeleteUserCard = `DELETE FROM user_cards
	WHERE (user_id = $1 AND card_id = $2)`

// -----CARD_TYPES-----

const SelectTypeNameById = `SELECT name FROM card_types WHERE id = $1`

// -----COOLDOWNS-----

const GetCooldown = `SELECT * FROM cooldowns WHERE user_id = $1`

const AddCooldown = `INSERT INTO cooldowns (user_id, next_accept) VALUES (:user_id, :next_accept)`

const UpdateCooldown = `UPDATE cooldowns SET next_accept = $1 WHERE user_id = $2`

// EXCHANGE

const CheckExchangeInitialized = `SELECT EXISTS (
	SELECT 1
	FROM exchanges
	WHERE (user_init_id = $1 AND user_continue_id = $2))`

const InitExchange = `INSERT INTO exchanges
	(user_init_id, user_continue_id, card_init_id) VALUES ($1, $2, $3)
	RETURNING
	    id, user_init_id, card_init_id, user_init_accept, user_continue_id, card_continue_id, user_continue_accept`

const ContinueExchange = `UPDATE exchanges
	SET card_continue_id = $3 WHERE user_init_id = $2 AND user_continue_id = $1
	RETURNING
	    id, user_init_id, card_init_id, user_init_accept, user_continue_id, card_continue_id, user_continue_accept`

const AcceptExchange = `UPDATE exchanges
	SET 
		user_init_accept = CASE WHEN user_init_id = $1 THEN TRUE ELSE user_init_accept END,
		user_continue_accept = CASE WHEN user_continue_id = $1 THEN TRUE ELSE user_continue_accept END
	WHERE 
		(user_init_id = $1 AND user_continue_id = $2) OR
		(user_init_id = $2 AND user_continue_id = $1)
	RETURNING
	    id, user_init_id, card_init_id, user_init_accept, user_continue_id, card_continue_id, user_continue_accept`

const DeleteExchange = `DELETE FROM exchanges
	WHERE (user_init_id = $1 AND user_continue_id = $2
		OR user_init_id = $2 AND user_continue_id = $1)`

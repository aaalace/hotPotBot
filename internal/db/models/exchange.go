package models

import "database/sql"

type Exchange struct {
	Id                 int           `db:"id"`
	UserInitId         sql.NullInt32 `db:"user_init_id"`
	CardInitId         sql.NullInt32 `db:"card_init_id"`
	UserInitAccept     bool          `db:"user_init_accept"`
	UserContinueId     sql.NullInt32 `db:"user_continue_id"`
	CardContinueId     sql.NullInt32 `db:"card_continue_id"`
	UserContinueAccept bool          `db:"user_continue_accept"`
}

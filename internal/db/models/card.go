package models

type Card struct {
	Id       int    `db:"id"`
	ImageUrl string `db:"image_url"`
	Name     string `db:"name"`
	Price    int    `db:"price"`
	Weight   int    `db:"weight"`
	TypeId   int    `db:"type_id"`
}

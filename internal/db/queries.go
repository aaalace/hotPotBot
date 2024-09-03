package db

const SelectUserQuery = "SELECT * FROM users WHERE telegram_id = $1"
const AddUserQuery = "INSERT INTO users (telegram_id, weight) VALUES (:telegram_id, :weight)"

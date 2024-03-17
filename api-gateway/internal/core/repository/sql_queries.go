package repository

const(

	getUserQuery = "SELECT user_id, password_hash, display_name FROM users WHERE username = $1"
)

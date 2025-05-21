package domain

import "time"

// User représente un utilisateur de l'application.
type User struct {
	ID        string    `db:"id" json:"id"` // PostgreSQL ID (souvent un UUID ou un int auto-incrémenté)
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password_hash" json:"-"` // Ne jamais exposer le hash du mot de passe en JSON
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	// Ajoutez d'autres champs (ex: RoleID, IsActive)
}
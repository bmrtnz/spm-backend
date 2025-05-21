package domain

import "time"

// Portfolio représente un portefeuille de projets/initiatives.
type Portfolio struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"` // MongoDB ID
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt" json:"updatedAt"`
	// Ajoutez d'autres champs pertinents pour un portefeuille (ex: OwnerID, Status)
}
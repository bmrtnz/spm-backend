package driven

import (
	"context"

	"github.com/bmrtnz/spm-backend/internal/domain" // Assurez-vous que le chemin du module est correct
)

// UserRepository définit les opérations de persistance pour les Users.
// Ceci est un port "driven" (piloté par l'application).
type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	// Ajoutez d'autres méthodes (ex: Update, Delete)
}
package driven

import (
	"context"

	"github.com/bmrtnz/spm-backend/internal/domain" // Assurez-vous que le chemin du module est correct
)

// PortfolioRepository définit les opérations de persistance pour les Portfolios.
// Ceci est un port "driven" (piloté par l'application).
type PortfolioRepository interface {
	Save(ctx context.Context, portfolio *domain.Portfolio) error
	GetByID(ctx context.Context, id string) (*domain.Portfolio, error)
	// Ajoutez d'autres méthodes nécessaires (ex: Update, Delete, ListByOwner)
}
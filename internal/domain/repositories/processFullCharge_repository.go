package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/thiagohmm/processCargaFull/internal/domain/entities"
)

type ProcessFullChargeRepository struct {
	db *sql.DB
}

func (r *ProcessFullChargeRepository) GetProcessFullCharge(ctx context.Context, idProcessFullCharge int) (*entities.ProcessFullCharge, error) {
	query := `SELECT IdProcessoCargaFull, CodigoIbm, DataInicioProcesso, Processando, DataFimProcesso, Sucesso, Cancelado 
              FROM ProcessFullCharge 
              WHERE IdProcessoCargaFull = ?`

	row := r.db.QueryRowContext(ctx, query, idProcessFullCharge)

	var processFullCharge entities.ProcessFullCharge
	err := row.Scan(
		&processFullCharge.IdProcessoCargaFull,
		&processFullCharge.CodigoIbm,
		&processFullCharge.DataInicioProcesso,
		&processFullCharge.Processando,
		&processFullCharge.DataFimProcesso,
		&processFullCharge.Sucesso,
		&processFullCharge.Cancelado,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &processFullCharge, nil
}

package repositories

import (
	"context"
	"database/sql"
	"fmt"
)

type SuccessRepository struct {
	db *sql.DB
}

func (r *SuccessRepository) GenerateMarketingStruct(ctx context.Context, idDealer int) error {
	query := `BEGIN SP_MOVERSTAGINGESTRUTURAMERCADOLOGICAREVENDEDOR(:1); END;`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.ExecContext(ctx, query, idDealer)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to execute query: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *SuccessRepository) MoverMarketingStruct(ctx context.Context, idDealer int) error {
	query := `BEGIN SP_GRAVARINTEGRACAOESTRUTURAMERCADOLOGICASTAGING(:1); END;`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.ExecContext(ctx, query, idDealer)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to execute query: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *SuccessRepository) ClearIntegrationMarketingStructureStagingDealer(ctx context.Context, idDealer int) error {
	query := `BEGIN sp_LimparIntegracaoEstruturaMercadologicaStagingRevendedor(:1); END;`
	_, err := r.db.ExecContext(ctx, query, idDealer)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *SuccessRepository) ClearIntegrationMarketingStructureDealer(ctx context.Context, idDealer int) error {
	query := `BEGIN sp_LimparIntegracaoEstruturaMercadologicaRevendedor(:1); END;`
	_, err := r.db.ExecContext(ctx, query, idDealer)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *SuccessRepository) ClearIntegrationMarketingStructureTransaction(ctx context.Context, transaction string) error {
	query := `BEGIN sp_LimparIntegracaoEstruturaMercadologicaTransacao(:1); END;`
	_, err := r.db.ExecContext(ctx, query, transaction)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *SuccessRepository) MoveIntegrationMarketingStructure(ctx context.Context, dataCorte string) error {
	query := `BEGIN sp_MoverStagingEstruturaMercadologica(:1); END;`
	_, err := r.db.ExecContext(ctx, query, dataCorte)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *SuccessRepository) DoMoveIntegrationMarketingStructure(ctx context.Context, idDealer int, json string) error {
	query := `BEGIN sp_IntegrarEstruturaMercadologica(:1, :2); END;`
	_, err := r.db.ExecContext(ctx, query, idDealer, json)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (r *SuccessRepository) MoveIntegrationMarketingStructureDealer(ctx context.Context, idDealer int) error {
	query := `BEGIN sp_MoverStagingEstruturaMercadologicaRevendedor(:1); END;`
	_, err := r.db.ExecContext(ctx, query, idDealer)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

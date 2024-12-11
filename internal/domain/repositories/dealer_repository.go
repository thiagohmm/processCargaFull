package repositories

import (
	"context" // Add the missing import
	"database/sql"

	"github.com/thiagohmm/processCargaFull/internal/domain/entities"
)

type DealerRepositoryDB struct {
	DB *sql.DB
}

func (d *DealerRepositoryDB) GetDealerByIBM(ctx context.Context, ibm string) (*entities.Dealer, error) { // Add context.Context parameter
	var dealer entities.Dealer
	dealer.CodigoIBM = ibm
	query := `
		SELECT *
		FROM Revendedor
		WHERE CodigoIBM = ?
	`
	err := d.DB.QueryRowContext(ctx, query, dealer.CodigoIBM).Scan(&dealer)
	if err != nil {
		return nil, err
	}

	return &dealer, nil
}

package paymentmethod

import (
	"database/sql"
	"log"

	"github.com/avivbintangaringga/ayobayar/db/jet/model"
	. "github.com/avivbintangaringga/ayobayar/db/jet/table"
	"github.com/avivbintangaringga/ayobayar/types"
	. "github.com/go-jet/jet/v2/postgres"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List() ([]types.PaymentMethod, error) {
	statement := SELECT(PaymentMethods.AllColumns).FROM(PaymentMethods)

	var queryResult []model.PaymentMethods

	err := statement.Query(r.db, &queryResult)
	if err != nil {
		return nil, err
	}

	log.Printf("result: %v", queryResult)

	var result []types.PaymentMethod

	for _, q := range queryResult {
		result = append(result, modelToType(q))
	}

	return result, nil
}

func (r *Repository) FindById(id string) (*types.PaymentMethod, error) {
	statement := SELECT(PaymentMethods.AllColumns).
		FROM(PaymentMethods).
		WHERE(AND(
			PaymentMethods.IsAvailable.EQ(Bool(true)),
			PaymentMethods.ID.EQ(VarChar(2)(id)),
		)).LIMIT(1)

	var queryResult model.PaymentMethods

	err := statement.Query(r.db, &queryResult)
	if err != nil {
		return nil, err
	}

	result := modelToType(queryResult)
	return &result, nil
}

func modelToType(m model.PaymentMethods) types.PaymentMethod {
	return types.PaymentMethod{
		Id:            m.ID,
		Name:          m.Name,
		SmallImageUrl: m.SmallImageURL,
		BigImageUrl:   m.BigImageURL,
		Category:      m.Category,
		IsAvailable:   m.IsAvailable,
	}
}

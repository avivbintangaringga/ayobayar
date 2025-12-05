package payment

import (
	"database/sql"
	"time"

	"github.com/avivbintangaringga/ayobayar/db/jet/ayobayar/public/model"
	. "github.com/avivbintangaringga/ayobayar/db/jet/ayobayar/public/table"
	"github.com/avivbintangaringga/ayobayar/types"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/teris-io/shortid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(data types.Payment) (*types.Payment, error) {
	m := typeToModel(data)
	m.ID = shortid.MustGenerate()
	m.Status = "PENDING"
	m.CreatedAt = time.Now()
	s := Payments.INSERT().MODEL(m).RETURNING(Payments.AllColumns)

	var qr model.Payments
	err := s.Query(r.db, &qr)
	if err != nil {
		return nil, err
	}

	result := modelToType(qr)

	return &result, nil
}

func (r *Repository) FindById(id string) (*types.Payment, error) {
	s := SELECT(Payments.AllColumns).FROM(Payments).
		WHERE(Payments.ID.EQ(Text(id))).
		LIMIT(1)

	var qr model.Payments
	err := s.Query(r.db, &qr)
	if err != nil {
		return nil, err
	}

	result := modelToType(qr)
	return &result, nil
}

func (r *Repository) Update(id string, data types.Payment) (*types.Payment, error) {
	return nil, nil
}

func (r *Repository) Delete(id string) error {
	s := Payments.DELETE().WHERE(Payments.ID.EQ(Text(id))).RETURNING(Payments.AllColumns)

	var qr model.Payments
	err := s.Query(r.db, &qr)
	if err != nil {
		return err
	}
	//TODO: CHECK
	return nil
}

func (r *Repository) List() ([]types.Payment, error) {
	s := SELECT(Payments.AllColumns).FROM(Payments)
	var qr []model.Payments
	err := s.Query(r.db, &qr)
	if err != nil {
		return nil, err
	}

	result := make([]types.Payment, len(qr))
	for i, m := range qr {
		result[i] = modelToType(m)
	}
	return result, nil
}

func modelToType(m model.Payments) types.Payment {
	return types.Payment{
		Id:              m.ID,
		PaymentMethodId: m.PaymentMethodID,
		Amount:          m.Amount,
		Status:          m.Status,
		ExpiryTime:      m.ExpiryTime,
		CallbackUrl:     m.CallbackURL,
		RedirectUrl:     m.RedirectURL,
		MerchantId:      m.MerchantID,
		MerchantOrderId: m.MerchantOrderID,
		CustomerEmail:   m.CustomerEmail,
		CustomerName:    m.CustomerName,
		CustomerPhone:   m.CustomerPhone,
		ProductDetails:  m.ProductDetails,
		CreatedAt:       m.CreatedAt,
	}
}

func typeToModel(t types.Payment) model.Payments {
	return model.Payments{
		ID:              t.Id,
		PaymentMethodID: t.PaymentMethodId,
		Amount:          t.Amount,
		Status:          t.Status,
		ExpiryTime:      t.ExpiryTime,
		CallbackURL:     t.CallbackUrl,
		RedirectURL:     t.RedirectUrl,
		MerchantID:      t.MerchantId,
		MerchantOrderID: t.MerchantOrderId,
		CustomerEmail:   t.CustomerEmail,
		CustomerName:    t.CustomerName,
		CustomerPhone:   t.CustomerPhone,
		ProductDetails:  t.ProductDetails,
		CreatedAt:       t.CreatedAt,
	}
}

package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*Queries
	db *pgxpool.Pool
}

// constructor
func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

// private function
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// start transaction
	tx, err := store.db.Begin(ctx)
	if err != nil {
		return err
	}

	// execute function
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	// commit transaction
	return tx.Commit(ctx)
}

// public functions
func (store *Store) HealthCheck() error {
	return store.db.Ping(context.Background())
}

type CreateProductTxParams struct {
	LangCode       LanguageCode  `json:"code"`
	Name           string        `json:"name"`
	IsHot          bool          `json:"is_hot"`
	Status         ProductStatus `json:"status"`
	Category       string        `json:"category"`
	ImageURLs      []string      `json:"image_urls"`
	Price          int32         `json:"price"`
	Quantity       int32         `json:"quantity"`
	Introduction   string        `json:"introduction"`
	Prize          string        `json:"prize"`
	ItemInfo       string        `json:"item_info"`
	Recommendation string        `json:"recommendation"`
}

type ProductTxResult struct {
	Product      Product
	ProductTrans ProductTranslation
	ProductDesc  ProductDescription
}

func (store *Store) CreateProductTx(ctx context.Context, args CreateProductTxParams) (*ProductTxResult, error) {
	var result ProductTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// create product
		result.Product, err = q.CreateProduct(ctx, CreateProductParams{
			Price:     args.Price,
			ImageURLs: args.ImageURLs,
			Status:    args.Status,
			Quantity:  args.Quantity,
			IsHot:     args.IsHot,
		})
		if err != nil {
			return err
		}

		// create product translation
		result.ProductTrans, err = q.CreateProductTranslation(ctx, CreateProductTranslationParams{
			ProductPk: result.Product.Pk,
			Language:  args.LangCode,
			Name:      args.Name,
			Category:  args.Category,
		})

		// create product description
		result.ProductDesc, err = q.CreateProductDescription(ctx, CreateProductDescriptionParams{
			ProductTranslationPk: result.ProductTrans.Pk,
			Introduction:         args.Introduction,
			Prize:                args.Prize,
			ItemInfo:             args.ItemInfo,
			Recommendation:       args.Recommendation,
		})

		return nil
	})
	return &result, err
}

func (store *Store) Close() {
	store.db.Close()
}

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/RafaelRochaS/quick-commerce-ddd/internal/domain/order"
	_ "modernc.org/sqlite"
)

type SQLite struct {
	db  *sql.DB
	ctx context.Context
}

func NewSQL(ctx context.Context) (order.Repository, error) {
	log.Println("Connecting to SQLite db...")
	db, err := sql.Open("sqlite", "./order_repository.db")

	if err != nil {
		return nil, fmt.Errorf("failed to create/open db: %w", err)
	}

	setDbConfig(db)
	err = validateConnection(db, ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite db: %w", err)
	}

	if err := ensureSchema(ctx, db); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("failed to ensure schema: %w", err)
	}

	defer db.Close()

	log.Println("Connected to SQLite db")

	return &SQLite{
		db:  db,
		ctx: ctx,
	}, nil
}

func setDbConfig(db *sql.DB) {
	log.Println("Setting DB Config...")

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	log.Println("Done.")
}

func validateConnection(db *sql.DB, ctx context.Context) error {
	log.Println("Validating DB connection...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return err
	}

	log.Println("Done.")
	return nil
}

func ensureSchema(ctx context.Context, db *sql.DB) error {
	schema := `
PRAGMA journal_mode = WAL;
PRAGMA busy_timeout = 5000;
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS orders (
	id TEXT PRIMARY KEY,
	customer_id TEXT NOT NULL,
	status TEXT NOT NULL,
	total_cents INTEGER NOT NULL,
	version INTEGER NOT NULL DEFAULT 0,
	updated_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS outbox_events (
	id TEXT PRIMARY KEY,
	aggregate_type TEXT NOT NULL,
	aggregate_id TEXT NOT NULL,
	event_type TEXT NOT NULL,
	payload TEXT NOT NULL,
	occurred_at TEXT NOT NULL,
	published_at TEXT NULL
);

CREATE INDEX IF NOT EXISTS idx_outbox_unpublished
	ON outbox_events (occurred_at)
	WHERE published_at IS NULL;
`
	if _, err := db.ExecContext(ctx, schema); err != nil {
		return err
	}
	return nil
}

func (S *SQLite) PlaceOrder(order *order.Order) error {
	log.Println("PlaceOrder called:", order)
	log.Println("Order placed")

	return nil
}

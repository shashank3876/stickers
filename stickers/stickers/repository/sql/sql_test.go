package sql_test

import (
	"context"
	"testing"
	"time"

	domain "shashank-gusain-backend-onboarding/domain"
	sql "shashank-gusain-backend-onboarding/stickers/repository/sql"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	err = db.AutoMigrate(&domain.Stickers{})
	if err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}

	return db
}

func closeTestDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)

	}
	sqlDB.Close()
}

func TestGetStickers(t *testing.T) {
	db := setupTestDB(t)
	defer closeTestDB(db)

	repo := sql.NewSQLStickerRepository(db)

	stickers := []domain.Stickers{
		{ID: 1,
			Name:     "Sticker 1",
			Priority: 0,
			AddedAt:  time.Now()},

		{ID: 2,
			Name:     "Sticker 2",
			Priority: 5,
			AddedAt:  time.Now()},
	}

	resultStickers, err := repo.GetStickers(context.Background())
	if err != nil {

		t.Errorf("error getting stickers: %v", err)
	}

	if len(resultStickers) != len(stickers) {
		t.Errorf("expected %d stickers, got %d", len(stickers), len(resultStickers))
	}
}

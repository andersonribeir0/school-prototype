package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/andersonribeir0/school-prototype/internal/models"
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id string) (*models.User, error)
	InsertUser(ctx context.Context, user *models.User) (string, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}

type Database struct {
	logger *slog.Logger
	db     *memdb.MemDB
}

func NewDatabase(logger *slog.Logger) *Database {
	db, err := createMemDB()
	if err != nil {
		logger.Error("createMemDB", err)
		panic(err)
	}

	return &Database{
		logger: logger,
		db:     db,
	}
}

func (db *Database) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	txn := db.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("user", "username", username)
	if err != nil {
		return nil, err
	}

	if raw == nil {
		return nil, nil
	}

	user, ok := raw.(*models.User)
	if !ok {
		return nil, fmt.Errorf("[GetUserByUsername] invalid format for user")
	}

	return user, nil
}

func (db *Database) GetUserById(ctx context.Context, id string) (*models.User, error) {
	txn := db.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("user", "id", id)
	if err != nil {
		return nil, err
	}
	if raw == nil {
		return nil, nil
	}

	user, ok := raw.(*models.User)
	if !ok {
		return nil, fmt.Errorf("[GetUserById] invalid format for user")
	}

	return user, nil
}

func (db *Database) InsertUser(ctx context.Context, user *models.User) (string, error) {
	txn := db.db.Txn(true)
	defer txn.Abort()

	user.ID = uuid.New().String()
	if err := txn.Insert("user", user); err != nil {
		return "", err
	}

	txn.Commit()

	return user.ID, nil
}

func createMemDB() (*memdb.MemDB, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"user": {
				Name: "user",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"username": {
						Name:    "username",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Username"},
					},
					"password": {
						Name:    "password",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Password"},
					},
				},
			},
			"course": {
				Name: "course",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"roomID": {
						Name:    "roomID",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "RoomID"},
					},
				},
			},
		},
	}

	// Create a new database
	return memdb.NewMemDB(schema)
}

package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"
)

func NewMemeCoinRepository(db *sql.DB) *MemeCoinRepository {
	return &MemeCoinRepository{
		db: db,
	}
}

func (repo *MemeCoinRepository) FindOne(id int) (*MemeCoin, error) {
	const sqlStatement string = `
		SELECT id, name, description, created_at, popularity_score
		FROM meme_coins
		WHERE id = $1`

	var memeCoin MemeCoin
	row := repo.db.QueryRowContext(context.Background(), sqlStatement, id)
	err := row.Scan(&memeCoin.Id, &memeCoin.Name, &memeCoin.Description, &memeCoin.CreatedAt, &memeCoin.PopularityScore)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &memeCoin, nil
}

func (repo *MemeCoinRepository) CreateOne(name string, description string) (*MemeCoin, error) {
	const sqlStatement string = `
		INSERT INTO meme_coins (name, description) 
		VALUES ($1, $2)
	 	ON CONFLICT (name) DO NOTHING
		RETURNING id, name, description, created_at, popularity_score`

	var newMemeCoin MemeCoin
	row := repo.db.QueryRowContext(context.Background(), sqlStatement, name, description)
	err := row.Scan(&newMemeCoin.Id, &newMemeCoin.Name, &newMemeCoin.Description, &newMemeCoin.CreatedAt, &newMemeCoin.PopularityScore)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &newMemeCoin, nil
}

func (repo *MemeCoinRepository) UpdateOne(id int, description string) (*MemeCoin, error) {
	const sqlStatement string = `
		UPDATE meme_coins
		SET description = $2
		WHERE id = $1
		RETURNING id, name, description, created_at, popularity_score`

	var updatedMemeCoin MemeCoin
	row := repo.db.QueryRowContext(context.Background(), sqlStatement, id, description)
	err := row.Scan(&updatedMemeCoin.Id, &updatedMemeCoin.Name, &updatedMemeCoin.Description, &updatedMemeCoin.CreatedAt, &updatedMemeCoin.PopularityScore)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &updatedMemeCoin, nil
}

func (repo *MemeCoinRepository) DeleteOne(id int) (*MemeCoin, error) {
	const sqlStatement string = `
		DELETE FROM meme_coins
		WHERE id = $1
		RETURNING id, name, description, created_at, popularity_score`

	var deletedMemeCoin MemeCoin
	row := repo.db.QueryRowContext(context.Background(), sqlStatement, id)
	err := row.Scan(&deletedMemeCoin.Id, &deletedMemeCoin.Name, &deletedMemeCoin.Description, &deletedMemeCoin.CreatedAt, &deletedMemeCoin.PopularityScore)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &deletedMemeCoin, nil
}

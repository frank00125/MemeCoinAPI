package repositories

import (
	"context"
	"errors"
	"portto-assignment/config"

	"github.com/jackc/pgx/v5"
)

func NewMemeCoinRepository(connectionPool config.DatabaseConnectionPoolInterface) *MemeCoinRepository {
	return &MemeCoinRepository{
		pool: connectionPool,
	}
}

func (repo *MemeCoinRepository) FindOne(id int) (*MemeCoin, error) {
	const sqlStatement string = `
		SELECT (id, name, description, created_at, popularity_score)
		FROM meme_coin
		WHERE id = @id`

	queryArgs := pgx.NamedArgs{
		"id": id,
	}

	var memeCoin MemeCoin
	row := repo.pool.QueryRow(context.Background(), sqlStatement, queryArgs)
	err := row.Scan(&memeCoin)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &memeCoin, nil
}

func (repo *MemeCoinRepository) CreateOne(name string, description string) (*MemeCoin, error) {
	const sqlStatement string = `
		INSERT INTO meme_coin (name, description) 
		VALUES (@name, @description)
	 	ON CONFLICT (name) DO NOTHING
		RETURNING (id, name, description, created_at, popularity_score)`

	queryArgs := pgx.NamedArgs{
		"name":        name,
		"description": description,
	}

	var newMemeCoin MemeCoin
	err := repo.pool.QueryRow(context.Background(), sqlStatement, queryArgs).Scan(&newMemeCoin)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &newMemeCoin, nil
}

func (repo *MemeCoinRepository) UpdateOne(id int, description string) (*MemeCoin, error) {
	const sqlStatement string = `
		UPDATE meme_coin
		SET description = @description
		WHERE id = @id
		RETURNING (id, name, description, created_at, popularity_score)`

	queryArgs := pgx.NamedArgs{
		"id":          id,
		"description": description,
	}

	var updatedMemeCoin MemeCoin
	err := repo.pool.QueryRow(context.Background(), sqlStatement, queryArgs).Scan(&updatedMemeCoin)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &updatedMemeCoin, nil
}

func (repo *MemeCoinRepository) DeleteOne(id int) (*MemeCoin, error) {
	const sqlStatement string = `
		DELETE FROM meme_coin
		WHERE id = @id
		RETURNING (id, name, description, created_at, popularity_score)`

	queryArgs := pgx.NamedArgs{
		"id": id,
	}

	var deletedMemeCoin MemeCoin
	err := repo.pool.QueryRow(context.Background(), sqlStatement, queryArgs).Scan(&deletedMemeCoin)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &deletedMemeCoin, nil
}

func (repo *MemeCoinRepository) PokeOne(id int) error {
	const sqlStatement string = `
		UPDATE meme_coin
		SET popularity_score = popularity_score + 1
		WHERE id = @id`

	queryArgs := pgx.NamedArgs{
		"id": id,
	}

	tx, err := repo.pool.Begin(context.Background())
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), sqlStatement, queryArgs)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	return nil
}

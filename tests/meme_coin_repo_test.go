package tests

import (
	"math/rand"
	"testing"
	"time"

	"portto-assignment/internal/repositories"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
)

type MemeCoinRepositoryTest struct {
	mockConnectionPool pgxmock.PgxPoolIface
	memeCoinRepository *repositories.MemeCoinRepository
}

func TestMemeCoinRepository(t *testing.T) {
	// Mocking the database connection
	mockConnectionPool, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal()
	}
	defer mockConnectionPool.Close()

	// Get the repository
	memeCoinRepository := repositories.NewMemeCoinRepository(mockConnectionPool)

	// Run the tests
	memeCoinRepositoryTest := MemeCoinRepositoryTest{
		mockConnectionPool: mockConnectionPool,
		memeCoinRepository: memeCoinRepository,
	}
	t.Run("FindOne", memeCoinRepositoryTest.testFindOne)
	t.Run("CreateOne", memeCoinRepositoryTest.testCreateOne)
	t.Run("UpdateOne", memeCoinRepositoryTest.testUpdateOne)
	t.Run("DeleteOne", memeCoinRepositoryTest.testDeleteOne)
	t.Run("PokeOne", memeCoinRepositoryTest.testPokeOne)
}

func (repo *MemeCoinRepositoryTest) testFindOne(t *testing.T) {
	fakeMemeCoin := repositories.MemeCoin{
		Id:              rand.Intn(100),
		Name:            "Test MemeCoin",
		Description:     "Test MemeCoin Description",
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}

	// Mocking the database connection
	returnRows := pgxmock.NewRows([]string{"row"}).AddRow(fakeMemeCoin)
	repo.mockConnectionPool.ExpectQuery(`
		SELECT (.+)
		FROM meme_coin
		WHERE id = @id`).WithArgs(pgx.NamedArgs{
		"id": fakeMemeCoin.Id,
	}).WillReturnRows(returnRows)
	memeCoin, err := repo.memeCoinRepository.FindOne(fakeMemeCoin.Id)
	if err != nil {
		t.Errorf("FindOne() failed, got error: %v", err)
	}

	assert.Equal(t, fakeMemeCoin.Id, memeCoin.Id)
	assert.Equal(t, fakeMemeCoin.Name, memeCoin.Name)
	assert.Equal(t, fakeMemeCoin.Description, memeCoin.Description)
	assert.Equal(t, fakeMemeCoin.CreatedAt, memeCoin.CreatedAt)
	assert.Equal(t, fakeMemeCoin.PopularityScore, memeCoin.PopularityScore)
}

func (repo *MemeCoinRepositoryTest) testCreateOne(t *testing.T) {
	fakeMemeCoin := repositories.MemeCoin{
		Id:              rand.Intn(100),
		Name:            "Test MemeCoin",
		Description:     "Test MemeCoin Description",
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}

	// Mocking the database connection
	returnRows := pgxmock.NewRows([]string{"row"}).AddRow(fakeMemeCoin)
	repo.mockConnectionPool.ExpectQuery(`
		INSERT INTO meme_coin \(name, description\)
	 	VALUES \(@name, @description\)
		ON CONFLICT \(name\) DO NOTHING
		RETURNING (.+)`).WithArgs(pgx.NamedArgs{
		"name":        fakeMemeCoin.Name,
		"description": fakeMemeCoin.Description,
	}).WillReturnRows(returnRows)
	memeCoin, err := repo.memeCoinRepository.CreateOne(fakeMemeCoin.Name, fakeMemeCoin.Description)
	if err != nil {
		t.Errorf("CreateOne() failed, got error: %v", err)
	}

	assert.Equal(t, fakeMemeCoin.Id, memeCoin.Id)
	assert.Equal(t, fakeMemeCoin.Name, memeCoin.Name)
	assert.Equal(t, fakeMemeCoin.Description, memeCoin.Description)
	assert.Equal(t, fakeMemeCoin.CreatedAt, memeCoin.CreatedAt)
	assert.Equal(t, fakeMemeCoin.PopularityScore, memeCoin.PopularityScore)
}

func (repo *MemeCoinRepositoryTest) testUpdateOne(t *testing.T) {
	fakeMemeCoin := repositories.MemeCoin{
		Id:              rand.Intn(100),
		Name:            "Test MemeCoin",
		Description:     "Test MemeCoin Description",
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}

	// Mocking the database connection
	returnRows := pgxmock.NewRows([]string{"row"}).AddRow(fakeMemeCoin)
	repo.mockConnectionPool.ExpectQuery(`
		UPDATE meme_coin
		SET description = @description
		WHERE id = @id
		RETURNING (.+)`).WithArgs(pgx.NamedArgs{
		"id":          fakeMemeCoin.Id,
		"description": fakeMemeCoin.Description,
	}).WillReturnRows(returnRows)
	memeCoin, err := repo.memeCoinRepository.UpdateOne(fakeMemeCoin.Id, fakeMemeCoin.Description)
	if err != nil {
		t.Errorf("UpdateOne() failed, got error: %v", err)
	}

	assert.Equal(t, fakeMemeCoin.Id, memeCoin.Id)
	assert.Equal(t, fakeMemeCoin.Name, memeCoin.Name)
	assert.Equal(t, fakeMemeCoin.Description, memeCoin.Description)
	assert.Equal(t, fakeMemeCoin.CreatedAt, memeCoin.CreatedAt)
	assert.Equal(t, fakeMemeCoin.PopularityScore, memeCoin.PopularityScore)
}

func (repo *MemeCoinRepositoryTest) testDeleteOne(t *testing.T) {
	fakeMemeCoin := repositories.MemeCoin{
		Id:              rand.Intn(100),
		Name:            "Test MemeCoin",
		Description:     "Test MemeCoin Description",
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}

	// Mocking the database connection
	returnRows := pgxmock.NewRows([]string{"row"}).AddRow(fakeMemeCoin)
	repo.mockConnectionPool.ExpectQuery(`
		DELETE FROM meme_coin
		WHERE id = @id
		RETURNING (.+)`).WithArgs(pgx.NamedArgs{
		"id": fakeMemeCoin.Id,
	}).WillReturnRows(returnRows)
	memeCoin, err := repo.memeCoinRepository.DeleteOne(fakeMemeCoin.Id)
	if err != nil {
		t.Errorf("DeleteOne() failed, got error: %v", err)
	}

	assert.Equal(t, fakeMemeCoin.Id, memeCoin.Id)
	assert.Equal(t, fakeMemeCoin.Name, memeCoin.Name)
	assert.Equal(t, fakeMemeCoin.Description, memeCoin.Description)
	assert.Equal(t, fakeMemeCoin.CreatedAt, memeCoin.CreatedAt)
	assert.Equal(t, fakeMemeCoin.PopularityScore, memeCoin.PopularityScore)
}

func (repo *MemeCoinRepositoryTest) testPokeOne(t *testing.T) {
	fakeMemeCoin := repositories.MemeCoin{
		Id:              rand.Intn(100),
		Name:            "Test MemeCoin",
		Description:     "Test MemeCoin Description",
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}

	// Mocking the database actions
	repo.mockConnectionPool.ExpectBegin()
	repo.mockConnectionPool.ExpectExec(`
		UPDATE meme_coin
		SET popularity_score = popularity_score \+ 1
		WHERE id = @id`).WithArgs(pgx.NamedArgs{
		"id": fakeMemeCoin.Id,
	}).WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	repo.mockConnectionPool.ExpectCommit()

	err := repo.memeCoinRepository.PokeOne(fakeMemeCoin.Id)
	if err != nil {
		t.Errorf("PokeOne() failed, got error: %v", err)
	}
}

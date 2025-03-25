package tests

import (
	"math/rand"
	"portto-assignment/internal/repositories"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type MemeCoinRepositoryTest struct {
	mockConnectionPool sqlmock.Sqlmock
	memeCoinRepository *repositories.MemeCoinRepository
}

func TestMemeCoinRepository(t *testing.T) {
	// Mocking the database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal()
	}
	defer mockDB.Close()

	// Get the repository
	memeCoinRepository := repositories.NewMemeCoinRepository(mockDB)

	// Run the tests
	memeCoinRepositoryTest := MemeCoinRepositoryTest{
		mockConnectionPool: mock,
		memeCoinRepository: memeCoinRepository,
	}
	t.Run("FindOne", memeCoinRepositoryTest.testFindOne)
	t.Run("CreateOne", memeCoinRepositoryTest.testCreateOne)
	t.Run("UpdateOne", memeCoinRepositoryTest.testUpdateOne)
	t.Run("DeleteOne", memeCoinRepositoryTest.testDeleteOne)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
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
	sqlStatement := "SELECT id, name, description, created_at, popularity_score FROM meme_coins WHERE id = $1"
	repo.mockConnectionPool.ExpectQuery(regexp.QuoteMeta(sqlStatement)).
		WithArgs(fakeMemeCoin.Id).
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "name", "description", "created_at", "popularity_score"}).
			AddRow(fakeMemeCoin.Id, fakeMemeCoin.Name, fakeMemeCoin.Description, fakeMemeCoin.CreatedAt, fakeMemeCoin.PopularityScore))
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
	sqlStatement := "INSERT INTO meme_coins (name, description) VALUES ($1, $2) ON CONFLICT (name) DO NOTHING RETURNING id, name, description, created_at, popularity_score"
	repo.mockConnectionPool.ExpectQuery(regexp.QuoteMeta(sqlStatement)).
		WithArgs(fakeMemeCoin.Name, fakeMemeCoin.Description).
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "name", "description", "created_at", "popularity_score"}).
			AddRow(fakeMemeCoin.Id, fakeMemeCoin.Name, fakeMemeCoin.Description, fakeMemeCoin.CreatedAt, fakeMemeCoin.PopularityScore))
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
	sqlStatement := "UPDATE meme_coins SET description = $2 WHERE id = $1 RETURNING id, name, description, created_at, popularity_score"
	repo.mockConnectionPool.ExpectQuery(regexp.QuoteMeta(sqlStatement)).
		WithArgs(fakeMemeCoin.Id, fakeMemeCoin.Description).
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "name", "description", "created_at", "popularity_score"}).
			AddRow(fakeMemeCoin.Id, fakeMemeCoin.Name, fakeMemeCoin.Description, fakeMemeCoin.CreatedAt, fakeMemeCoin.PopularityScore))
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
	sqlStatement := "DELETE FROM meme_coins WHERE id = $1 RETURNING id, name, description, created_at, popularity_score"
	repo.mockConnectionPool.ExpectQuery(regexp.QuoteMeta(sqlStatement)).
		WithArgs(fakeMemeCoin.Id).
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "name", "description", "created_at", "popularity_score"}).
			AddRow(fakeMemeCoin.Id, fakeMemeCoin.Name, fakeMemeCoin.Description, fakeMemeCoin.CreatedAt, fakeMemeCoin.PopularityScore))
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

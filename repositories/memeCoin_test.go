package repositories_test

import (
	"math/rand"
	"testing"
	"time"

	"portto-assignment/repositories"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
)

var mockConnectionPool pgxmock.PgxPoolIface
var memeCoinRepository *repositories.MemeCoinRepository

func TestMemeCoinRepository(t *testing.T) {
	// Mocking the database connection
	var err error
	mockConnectionPool, err = pgxmock.NewPool()
	if err != nil {
		t.Fatal()
	}
	defer mockConnectionPool.Close()

	// Get the repository
	repositories.Init(mockConnectionPool)
	memeCoinRepository = repositories.GetMemeCoinRepository()

	// Run the tests
	t.Run("FindOne", testFindOne)
	t.Run("CreateOne", testCreateOne)
	t.Run("UpdateOne", testUpdateOne)
	t.Run("DeleteOne", testDeleteOne)
	t.Run("PokeOne", testPokeOne)
}

func testFindOne(t *testing.T) {
	fakeMemeCoin := repositories.MemeCoin{
		Id:              rand.Intn(100),
		Name:            "Test MemeCoin",
		Description:     "Test MemeCoin Description",
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}

	// Mocking the database connection
	returnRows := pgxmock.NewRows([]string{"row"}).AddRow(fakeMemeCoin)
	mockConnectionPool.ExpectQuery(`
		SELECT (.+)
		FROM meme_coin
		WHERE id = @id`).WithArgs(pgx.NamedArgs{
		"id": fakeMemeCoin.Id,
	}).WillReturnRows(returnRows)
	memeCoin, err := memeCoinRepository.FindOne(fakeMemeCoin.Id)
	if err != nil {
		t.Errorf("FindOne() failed, got error: %v", err)
	}

	assert.Equal(t, fakeMemeCoin.Id, memeCoin.Id)
	assert.Equal(t, fakeMemeCoin.Name, memeCoin.Name)
	assert.Equal(t, fakeMemeCoin.Description, memeCoin.Description)
	assert.Equal(t, fakeMemeCoin.CreatedAt, memeCoin.CreatedAt)
	assert.Equal(t, fakeMemeCoin.PopularityScore, memeCoin.PopularityScore)
}

func testCreateOne(t *testing.T) {
	fakeMemeCoin := repositories.MemeCoin{
		Id:              rand.Intn(100),
		Name:            "Test MemeCoin",
		Description:     "Test MemeCoin Description",
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}

	// Mocking the database connection
	returnRows := pgxmock.NewRows([]string{"row"}).AddRow(fakeMemeCoin)
	mockConnectionPool.ExpectQuery(`
		INSERT INTO meme_coin \(name, description\)
	 	VALUES \(@name, @description\)
		ON CONFLICT \(name\) DO NOTHING
		RETURNING (.+)`).WithArgs(pgx.NamedArgs{
		"name":        fakeMemeCoin.Name,
		"description": fakeMemeCoin.Description,
	}).WillReturnRows(returnRows)
	memeCoin, err := memeCoinRepository.CreateOne(fakeMemeCoin.Name, fakeMemeCoin.Description)
	if err != nil {
		t.Errorf("CreateOne() failed, got error: %v", err)
	}

	assert.Equal(t, fakeMemeCoin.Id, memeCoin.Id)
	assert.Equal(t, fakeMemeCoin.Name, memeCoin.Name)
	assert.Equal(t, fakeMemeCoin.Description, memeCoin.Description)
	assert.Equal(t, fakeMemeCoin.CreatedAt, memeCoin.CreatedAt)
	assert.Equal(t, fakeMemeCoin.PopularityScore, memeCoin.PopularityScore)
}

func testUpdateOne(t *testing.T) {
	fakeMemeCoin := repositories.MemeCoin{
		Id:              rand.Intn(100),
		Name:            "Test MemeCoin",
		Description:     "Test MemeCoin Description",
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}

	// Mocking the database connection
	returnRows := pgxmock.NewRows([]string{"row"}).AddRow(fakeMemeCoin)
	mockConnectionPool.ExpectQuery(`
		UPDATE meme_coin
		SET description = @description
		WHERE id = @id
		RETURNING (.+)`).WithArgs(pgx.NamedArgs{
		"id":          fakeMemeCoin.Id,
		"description": fakeMemeCoin.Description,
	}).WillReturnRows(returnRows)
	memeCoin, err := memeCoinRepository.UpdateOne(fakeMemeCoin.Id, fakeMemeCoin.Description)
	if err != nil {
		t.Errorf("UpdateOne() failed, got error: %v", err)
	}

	assert.Equal(t, fakeMemeCoin.Id, memeCoin.Id)
	assert.Equal(t, fakeMemeCoin.Name, memeCoin.Name)
	assert.Equal(t, fakeMemeCoin.Description, memeCoin.Description)
	assert.Equal(t, fakeMemeCoin.CreatedAt, memeCoin.CreatedAt)
	assert.Equal(t, fakeMemeCoin.PopularityScore, memeCoin.PopularityScore)
}

func testDeleteOne(t *testing.T) {
	fakeMemeCoin := repositories.MemeCoin{
		Id:              rand.Intn(100),
		Name:            "Test MemeCoin",
		Description:     "Test MemeCoin Description",
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}

	// Mocking the database connection
	returnRows := pgxmock.NewRows([]string{"row"}).AddRow(fakeMemeCoin)
	mockConnectionPool.ExpectQuery(`
		DELETE FROM meme_coin
		WHERE id = @id
		RETURNING (.+)`).WithArgs(pgx.NamedArgs{
		"id": fakeMemeCoin.Id,
	}).WillReturnRows(returnRows)
	memeCoin, err := memeCoinRepository.DeleteOne(fakeMemeCoin.Id)
	if err != nil {
		t.Errorf("DeleteOne() failed, got error: %v", err)
	}

	assert.Equal(t, fakeMemeCoin.Id, memeCoin.Id)
	assert.Equal(t, fakeMemeCoin.Name, memeCoin.Name)
	assert.Equal(t, fakeMemeCoin.Description, memeCoin.Description)
	assert.Equal(t, fakeMemeCoin.CreatedAt, memeCoin.CreatedAt)
	assert.Equal(t, fakeMemeCoin.PopularityScore, memeCoin.PopularityScore)
}

func testPokeOne(t *testing.T) {
	fakeMemeCoin := repositories.MemeCoin{
		Id:              rand.Intn(100),
		Name:            "Test MemeCoin",
		Description:     "Test MemeCoin Description",
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}

	// Mocking the database actions
	mockConnectionPool.ExpectBegin()
	mockConnectionPool.ExpectExec(`
		UPDATE meme_coin
		SET popularity_score = popularity_score \+ 1
		WHERE id = @id`).WithArgs(pgx.NamedArgs{
		"id": fakeMemeCoin.Id,
	}).WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	mockConnectionPool.ExpectCommit()

	err := memeCoinRepository.PokeOne(fakeMemeCoin.Id)
	if err != nil {
		t.Errorf("PokeOne() failed, got error: %v", err)
	}
}

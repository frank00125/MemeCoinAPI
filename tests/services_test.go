package tests

import (
	"testing"
	"time"

	"portto-assignment/internal/services"
	"portto-assignment/tests/mocks"

	"github.com/stretchr/testify/assert"
)

var memeCoinService *services.MemeCoinService

func TestMemeCoinService(t *testing.T) {
	// Mock the repository
	mockMemeCoinRepository := &mocks.MockMemeCoinRepository{}
	mockRedisCachedRepository := &mocks.MockRedisCachedRepository{}

	memeCoinService = services.NewMemeCoinService(mockMemeCoinRepository, mockRedisCachedRepository)

	t.Run("CreateMemeCoin", testCreateMemeCoin)
	t.Run("GetMemeCoin", testGetMemeCoin)
	t.Run("UpdateMemeCoin", testUpdateMemeCoin)
	t.Run("DeleteMemeCoin", testDeleteMemeCoin)
	t.Run("PokeMemeCoin", testPokeMemeCoin)
}

func testCreateMemeCoin(t *testing.T) {
	// Test case 1: name is not empty
	timeBeforeExecute := time.Now()
	memeCoin, err := memeCoinService.CreateMemeCoin(services.CreateMemeCoinInput{
		Name:        "name",
		Description: "description",
	})
	timeAfterExecute := time.Now()

	assert.NoError(t, err)
	assert.NotNil(t, memeCoin)
	assert.Greater(t, memeCoin.Id, 0)
	assert.Less(t, memeCoin.PopularityScore, 100)
	assert.Equal(t, "name", memeCoin.Name)
	assert.Equal(t, "description", memeCoin.Description)
	assert.GreaterOrEqual(t, memeCoin.CreatedAt.UnixNano(), timeBeforeExecute.UnixNano())
	assert.LessOrEqual(t, memeCoin.CreatedAt.UnixNano(), timeAfterExecute.UnixNano())
	assert.Greater(t, memeCoin.PopularityScore, 0)
	assert.Less(t, memeCoin.PopularityScore, 100)
}

func testGetMemeCoin(t *testing.T) {
	// Test case 1: id is invalid (id = 0 => invalid)
	memeCoin, err := memeCoinService.GetMemeCoin(0)
	assert.Error(t, err)
	assert.Nil(t, memeCoin)

	// Test case 2: id is valid
	timeBeforeExecute := time.Now()
	memeCoin, err = memeCoinService.GetMemeCoin(1)
	timeAfterExecute := time.Now()

	assert.NoError(t, err)
	assert.NotNil(t, memeCoin)
	assert.Greater(t, memeCoin.Id, 0)
	assert.Less(t, memeCoin.PopularityScore, 100)
	assert.Equal(t, "FakeCoin", memeCoin.Name)
	assert.Equal(t, "A fake meme coin", memeCoin.Description)
	assert.GreaterOrEqual(t, memeCoin.CreatedAt.UnixNano(), timeBeforeExecute.UnixNano())
	assert.LessOrEqual(t, memeCoin.CreatedAt.UnixNano(), timeAfterExecute.UnixNano())
	assert.Greater(t, memeCoin.PopularityScore, 0)
	assert.Less(t, memeCoin.PopularityScore, 100)
}

func testUpdateMemeCoin(t *testing.T) {
	// Test case 1: id is invalid (id = 0 => invalid)
	memeCoin, err := memeCoinService.UpdateMemeCoin(0, "new description")
	assert.Error(t, err)
	assert.Nil(t, memeCoin)

	// Test case 2: id is valid
	timeBeforeExecute := time.Now()
	memeCoin, err = memeCoinService.UpdateMemeCoin(1, "new description")
	timeAfterExecute := time.Now()

	assert.NoError(t, err)
	assert.NotNil(t, memeCoin)
	assert.Greater(t, memeCoin.Id, 0)
	assert.Less(t, memeCoin.PopularityScore, 100)
	assert.Equal(t, "FakeCoin", memeCoin.Name)
	assert.Equal(t, "new description", memeCoin.Description)
	assert.GreaterOrEqual(t, memeCoin.CreatedAt.UnixNano(), timeBeforeExecute.UnixNano())
	assert.LessOrEqual(t, memeCoin.CreatedAt.UnixNano(), timeAfterExecute.UnixNano())
	assert.Greater(t, memeCoin.PopularityScore, 0)
	assert.Less(t, memeCoin.PopularityScore, 100)
}

func testDeleteMemeCoin(t *testing.T) {
	// Test case 1: id is invalid (id = 0 => invalid)
	memeCoin, err := memeCoinService.DeleteMemeCoin(0)
	assert.Error(t, err)
	assert.Nil(t, memeCoin)

	// Test case 2: id is valid
	timeBeforeExecute := time.Now()
	memeCoin, err = memeCoinService.DeleteMemeCoin(1)
	timeAfterExecute := time.Now()

	assert.NoError(t, err)
	assert.NotNil(t, memeCoin)
	assert.Greater(t, memeCoin.Id, 0)
	assert.Less(t, memeCoin.PopularityScore, 100)
	assert.Equal(t, "FakeCoin", memeCoin.Name)
	assert.Equal(t, "A fake meme coin", memeCoin.Description)
	assert.GreaterOrEqual(t, memeCoin.CreatedAt.UnixNano(), timeBeforeExecute.UnixNano())
	assert.LessOrEqual(t, memeCoin.CreatedAt.UnixNano(), timeAfterExecute.UnixNano())
	assert.Greater(t, memeCoin.PopularityScore, 0)
	assert.Less(t, memeCoin.PopularityScore, 100)
}

func testPokeMemeCoin(t *testing.T) {
	// Test case 1: id is invalid (id = 0 => invalid)
	err := memeCoinService.PokeMemeCoin(0)
	assert.Error(t, err)

	// Test case 2: id is valid
	err = memeCoinService.PokeMemeCoin(1)
	assert.NoError(t, err)
}

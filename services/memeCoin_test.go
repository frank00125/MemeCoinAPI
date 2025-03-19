package services_test

import (
	"testing"
	"time"

	"portto-assignment/mocks"
	"portto-assignment/services"

	"github.com/stretchr/testify/assert"
)

func TestMemeCoinService(t *testing.T) {
	mockMemeCoinRepository := &mocks.MockMemeCoinRepository{}
	services.Init(mockMemeCoinRepository)
	memeCoinService := services.GetMemeCoinService()

	t.Log("Running tests for function 'CreateMemeCoin'")
	testCreateMemeCoin(t, memeCoinService)

	t.Log("Running tests for function 'GetMemeCoin'")
	testGetMemeCoin(t, memeCoinService)

	t.Log("Running tests for function 'UpdateMemeCoin'")
	testUpdateMemeCoin(t, memeCoinService)

	t.Log("Running tests for function 'DeleteMemeCoin'")
	testDeleteMemeCoin(t, memeCoinService)

	t.Log("Running tests for function 'PokeMemeCoin'")
	testPokeMemeCoin(t, memeCoinService)
}

func testCreateMemeCoin(t *testing.T, memeCoinService *services.MemeCoinService) {
	// Test case 1: name is empty
	_, err := memeCoinService.CreateMemeCoin(services.CreateMemeCoinInput{
		Name:        "",
		Description: "description",
	})
	assert.NotNil(t, err)

	// Test case 2: name is not empty
	timeBeforeExecute := time.Now()
	memeCoin, err := memeCoinService.CreateMemeCoin(services.CreateMemeCoinInput{
		Name:        "name",
		Description: "description",
	})
	timeAfterExecute := time.Now()

	assert.Nil(t, err)
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

func testGetMemeCoin(t *testing.T, memeCoinService *services.MemeCoinService) {
	// Test case 1: id is invalid (id = 0 => invalid)
	memeCoin, err := memeCoinService.GetMemeCoin(0)
	assert.NotNil(t, err)
	assert.Nil(t, memeCoin)

	// Test case 2: id is valid
	timeBeforeExecute := time.Now()
	memeCoin, err = memeCoinService.GetMemeCoin(1)
	timeAfterExecute := time.Now()

	assert.Nil(t, err)
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

func testUpdateMemeCoin(t *testing.T, memeCoinService *services.MemeCoinService) {
	// Test case 1: id is invalid (id = 0 => invalid)
	memeCoin, err := memeCoinService.UpdateMemeCoin(0, "new description")
	assert.NotNil(t, err)
	assert.Nil(t, memeCoin)

	// Test case 2: id is valid
	timeBeforeExecute := time.Now()
	memeCoin, err = memeCoinService.UpdateMemeCoin(1, "new description")
	timeAfterExecute := time.Now()

	assert.Nil(t, err)
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

func testDeleteMemeCoin(t *testing.T, memeCoinService *services.MemeCoinService) {
	// Test case 1: id is invalid (id = 0 => invalid)
	memeCoin, err := memeCoinService.DeleteMemeCoin(0)
	assert.NotNil(t, err)
	assert.Nil(t, memeCoin)

	// Test case 2: id is valid
	timeBeforeExecute := time.Now()
	memeCoin, err = memeCoinService.DeleteMemeCoin(1)
	timeAfterExecute := time.Now()

	assert.Nil(t, err)
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

func testPokeMemeCoin(t *testing.T, memeCoinService *services.MemeCoinService) {
	// Test case 1: id is invalid (id = 0 => invalid)
	memeCoin, err := memeCoinService.PokeMemeCoin(0)
	assert.NotNil(t, err)
	assert.Nil(t, memeCoin)

	// Test case 2: id is valid
	timeBeforeExecute := time.Now()
	memeCoin, err = memeCoinService.PokeMemeCoin(1)
	timeAfterExecute := time.Now()

	assert.Nil(t, err)
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

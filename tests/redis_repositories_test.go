package tests

import (
	"portto-assignment/internal/repositories"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

type RedisCachedRepositoryTest struct {
	dbmock                sqlmock.Sqlmock
	redismock             redismock.ClientMock
	redisCachedRepository *repositories.RedisCachedRepository
}

func TestRedisCachedRepository(t *testing.T) {
	mockDB, dbmock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()

	mockRedisClient, redismock := redismock.NewClientMock()
	defer mockRedisClient.Close()

	redisCachedRepositoryTest := RedisCachedRepositoryTest{
		dbmock:    dbmock,
		redismock: redismock,
		redisCachedRepository: repositories.NewRedisCachedRepository(mockDB, mockRedisClient, repositories.RepositoryConfig{
			SyncBatchSize: repositories.DefaultSyncBatchSize,
			SyncInterval:  repositories.DefaultSyncInterval,
			NeedToSync:    false,
		}),
	}

	t.Run("IncrementMemeCoinPopularityScore", redisCachedRepositoryTest.testIncrementMemeCoinPopularityScore)
	t.Run("RemoveMemeCoinPopularityScore", redisCachedRepositoryTest.testRemoveMemeCoinPopularityScore)
}

func (repo *RedisCachedRepositoryTest) testIncrementMemeCoinPopularityScore(t *testing.T) {
	memeCoinId := 10
	key := "meme:popularity_score:" + strconv.Itoa(memeCoinId)

	repo.redismock.ExpectIncrBy(key, 1).SetVal(1)

	err := repo.redisCachedRepository.IncrementPopularityScore(memeCoinId)
	if err != nil {
		t.Errorf("IncrementMemeCoinPopularityScore() failed, got error: %v", err)
	}

	err = repo.redismock.ExpectationsWereMet()
	assert.NoError(t, err)

	err = repo.dbmock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func (repo *RedisCachedRepositoryTest) testRemoveMemeCoinPopularityScore(t *testing.T) {
	memeCoinId := 10
	key := "meme:popularity_score:" + strconv.Itoa(memeCoinId)

	repo.redismock.ExpectDel(key).SetVal(1)

	err := repo.redisCachedRepository.RemovePopularityScore(memeCoinId)
	if err != nil {
		t.Errorf("RemoveMemeCoinPopularityScore() failed, got error: %v", err)
	}

	err = repo.redismock.ExpectationsWereMet()
	assert.NoError(t, err)

	err = repo.dbmock.ExpectationsWereMet()
	assert.NoError(t, err)
}

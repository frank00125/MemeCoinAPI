package tests

import (
	"portto-assignment/internal/repositories"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
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

	t.Run("TestSet", redisCachedRepositoryTest.testSet)
	t.Run("TestIncr", redisCachedRepositoryTest.testIncrBy)
	t.Run("TestDelete", redisCachedRepositoryTest.testDelete)
	t.Run("TestExists", redisCachedRepositoryTest.testExists)
}

func (r *RedisCachedRepositoryTest) testIncrBy(t *testing.T) {
	key := "test_key"
	r.redismock.ExpectIncrBy(key, 1).SetVal(1)

	err := r.redisCachedRepository.IncrBy("test_key", 1)
	if err != nil {
		t.Fatal(err)
	}
}

func (r *RedisCachedRepositoryTest) testSet(t *testing.T) {
	key := "test_key"
	r.redismock.ExpectSet(key, 0, 0).SetVal("OK")

	err := r.redisCachedRepository.Set("test_key", 0)
	if err != nil {
		t.Fatal(err)
	}
}

func (r *RedisCachedRepositoryTest) testDelete(t *testing.T) {
	key := "test_key"
	r.redismock.ExpectDel(key).SetVal(1)

	err := r.redisCachedRepository.Delete("test_key")
	if err != nil {
		t.Fatal(err)
	}
}

func (r *RedisCachedRepositoryTest) testExists(t *testing.T) {
	key := "test_key"
	r.redismock.ExpectExists(key).SetVal(1)

	exists, err := r.redisCachedRepository.Exists(key)
	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		t.Fatal("Key should exist")
	}
}

package mocks

import (
	"errors"
	"fmt"
	"math/rand"
	"portto-assignment/internal/repositories"
	"time"
)

type MockMemeCoinRepository struct {
}

type MockRedisCachedRepository struct {
}

func (m *MockMemeCoinRepository) FindOne(id int) (*repositories.MemeCoin, error) {
	if id == 0 {
		return nil, errors.New("invalid ID")
	}

	fakeMemeCoin := m.getFakeMemeCoin()
	fakeMemeCoin.Id = id

	return &fakeMemeCoin, nil
}

func (m *MockMemeCoinRepository) CreateOne(name string, description string) (*repositories.MemeCoin, error) {
	fakeMemeCoin := m.getFakeMemeCoin()
	fakeMemeCoin.Name = name
	fakeMemeCoin.Description = description

	return &fakeMemeCoin, nil
}

func (m *MockMemeCoinRepository) UpdateOne(id int, description string) (*repositories.MemeCoin, error) {
	if id == 0 {
		return nil, errors.New("invalid ID")
	}

	fakeMemeCoin := m.getFakeMemeCoin()
	fakeMemeCoin.Id = id
	fakeMemeCoin.Description = description

	return &fakeMemeCoin, nil
}

func (m *MockMemeCoinRepository) DeleteOne(id int) (*repositories.MemeCoin, error) {
	if id == 0 {
		return nil, errors.New("invalid ID")
	}

	fakeMemeCoin := m.getFakeMemeCoin()
	fakeMemeCoin.Id = id

	return &fakeMemeCoin, nil
}

func (m *MockMemeCoinRepository) PokeOne(id int) error {
	if id == 0 {
		return errors.New("invalid ID")
	}

	return nil
}

func (m *MockMemeCoinRepository) getFakeMemeCoin() repositories.MemeCoin {
	return repositories.MemeCoin{
		Id:              rand.Intn(9999) + 1,
		Name:            "FakeCoin",
		Description:     "A fake meme coin",
		CreatedAt:       time.Now(),
		PopularityScore: rand.Intn(99) + 1,
	}
}

func (m *MockRedisCachedRepository) IncrBy(key string, increment int) error {
	if key == fmt.Sprintf("meme:popularity_score:%d", 0) {
		return fmt.Errorf("key %s does not exist", key)
	}
	return nil
}

func (m *MockRedisCachedRepository) Set(key string, value int) error {
	if key == fmt.Sprintf("meme:popularity_score:%d", 0) {
		return fmt.Errorf("key %s does not exist", key)
	}
	return nil
}

func (m *MockRedisCachedRepository) Delete(key string) error {
	if key == fmt.Sprintf("meme:popularity_score:%d", 0) {
		return fmt.Errorf("key %s does not exist", key)
	}
	return nil
}

func (m *MockRedisCachedRepository) Exists(key string) (bool, error) {
	if key == fmt.Sprintf("meme:popularity_score:%d", 0) {
		return false, nil
	}

	return true, nil
}

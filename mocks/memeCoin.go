package mocks

import (
	"errors"
	"math/rand"
	"portto-assignment/repositories"
	"time"
)

type MockMemeCoinRepository struct {
}

func (m *MockMemeCoinRepository) FindOne(id int) (*repositories.MemeCoin, error) {
	if id == 0 {
		return nil, errors.New("invalid ID")
	}

	fakeMemeCoin := getFakeMemeCoin()
	fakeMemeCoin.Id = id

	return &fakeMemeCoin, nil
}

func (m *MockMemeCoinRepository) CreateOne(name string, description string) (*repositories.MemeCoin, error) {
	fakeMemeCoin := getFakeMemeCoin()
	fakeMemeCoin.Name = name
	fakeMemeCoin.Description = description

	return &fakeMemeCoin, nil
}

func (m *MockMemeCoinRepository) UpdateOne(id int, description string) (*repositories.MemeCoin, error) {
	if id == 0 {
		return nil, errors.New("invalid ID")
	}

	fakeMemeCoin := getFakeMemeCoin()
	fakeMemeCoin.Id = id
	fakeMemeCoin.Description = description

	return &fakeMemeCoin, nil
}

func (m *MockMemeCoinRepository) DeleteOne(id int) (*repositories.MemeCoin, error) {
	if id == 0 {
		return nil, errors.New("invalid ID")
	}

	fakeMemeCoin := getFakeMemeCoin()
	fakeMemeCoin.Id = id

	return &fakeMemeCoin, nil
}

func (m *MockMemeCoinRepository) PokeOne(id int) error {
	if id == 0 {
		return errors.New("invalid ID")
	}

	return nil
}

func getFakeMemeCoin() repositories.MemeCoin {
	return repositories.MemeCoin{
		Id:              rand.Intn(9999) + 1,
		Name:            "FakeCoin",
		Description:     "A fake meme coin",
		CreatedAt:       time.Now(),
		PopularityScore: rand.Intn(99) + 1,
	}
}

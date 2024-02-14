package sqlite

import (
	"database/sql"
	"services/pkg/agent"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storagePath string) (*Storage, error) {

	database, err := sql.Open("sqlite3",storagePath)
	if err != nil {
		return nil, err
	}

	return &Storage{db: database}, nil
}

func (s *Storage) SaveUserTask(userTask agent.UserTask) error {
	return nil
}

func (s *Storage) SaveUserAnswer(userAnswer agent.UserAnswer) error {
	return nil
}

func (s *Storage) GetUserTask(id int) (agent.UserTask, error) {
	var result agent.UserTask

	return result, nil
}

func (s *Storage) GetUserAnswer(id int) (agent.UserAnswer, error) {
	var result agent.UserAnswer

	return result, nil
}


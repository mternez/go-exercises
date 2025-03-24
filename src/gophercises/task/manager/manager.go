package manager

import (
	"encoding/json"
	"task/data"

	"github.com/boltdb/bolt"
)

type Task struct {
	Title       string
	Description string
	Done        bool
}

func (t *Task) GetKey() string {
	return t.Title
}

var _ data.Value = (*Task)(nil)

type TaskManager struct {
	dbFileName string
	db         *bolt.DB
	repository *data.Repository[*Task]
}

func NewTaskManager(dbFileName string) *TaskManager {
	repo := &data.Repository[*Task]{
		Name: "tasks",
		Serialize: func(val *Task) ([]byte, error) {
			return json.Marshal(val)
		},
		Deserialize: func(valueBytes []byte) (*Task, error) {
			value := &Task{}
			err := json.Unmarshal(valueBytes, value)
			return value, err
		},
	}
	return &TaskManager{dbFileName: dbFileName, repository: repo}
}

func (m *TaskManager) openConnection() error {
	db, err := bolt.Open(m.dbFileName, 0600, nil)
	if err != nil {
		return err
	}
	m.db = db
	return nil
}

func (m *TaskManager) closeConnection() error {
	return m.db.Close()
}

func (m *TaskManager) Save(task *Task) error {
	err := m.openConnection()
	if err != nil {
		return err
	}
	defer m.closeConnection()
	err = m.repository.Put(m.db, task)
	if err != nil {
		return err
	}
	return nil
}

func (m *TaskManager) FindByTitle(title string) (*Task, error) {
	err := m.openConnection()
	if err != nil {
		return nil, err
	}
	defer m.closeConnection()
	return m.repository.Get(m.db, title)
}

func (m *TaskManager) FindAll() ([]*Task, error) {
	err := m.openConnection()
	if err != nil {
		return nil, err
	}
	defer m.closeConnection()
	return m.repository.Values(m.db)
}

func (m *TaskManager) Remove(title string) error {
	err := m.openConnection()
	if err != nil {
		return err
	}
	defer m.closeConnection()
	return m.repository.Delete(m.db, title)
}
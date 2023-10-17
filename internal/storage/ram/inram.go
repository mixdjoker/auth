package ram

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/mixdjoker/auth/internal/model"
)

type UserStore struct {
	counter int64
	users   map[int64]model.User
	mails   map[string]int64
	mu      sync.RWMutex
}

func NewUserStore() *UserStore {
	return &UserStore{
		counter: 0,
		users:   make(map[int64]model.User),
		mails:   make(map[string]int64),
		mu:      sync.RWMutex{},
	}
}

func (s *UserStore) Create(ctx context.Context, u model.User) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if exID, ok := s.mails[u.Email]; ok {
		exUser := s.users[exID]
		errStr := fmt.Sprintf("email already exists: user %v with id %v has this email", exUser.Name, exUser.ID)
		return 0, errors.New(errStr)
	}
	u.ID = s.counter
	t := time.Now()
	u.CreatedAt = int64(t.Unix())
	s.users[u.ID] = u
	s.mails[u.Email] = u.ID

	s.counter++

	return u.ID, nil
}

func (s *UserStore) Get(ctx context.Context, id int64) (model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, ok := s.users[id]; ok {
		return s.users[id], nil
	}

	return model.User{}, errors.New("user not found")
}

func (s *UserStore) Update(ctx context.Context, u model.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[u.ID]; !ok {
		return errors.New("user not found")
	}

	curUser := s.users[u.ID]

	if exID, ok := s.mails[u.Email]; ok && exID != u.ID {
		exUser := s.users[exID]
		errStr := fmt.Sprintf("email already exists: user %v with id %v has this email", exUser.Name, exUser.ID)
		return errors.New(errStr)
	}

	delete(s.mails, curUser.Email)
	delete(s.users, curUser.ID)

	if u.Name != "" {
		curUser.Name = u.Name
	}

	if u.Password != "" {
		curUser.Password = u.Password
	}

	if u.Role != 0 {
		curUser.Role = u.Role
	}

	if u.Email != "" {
		curUser.Email = u.Email
	}

	curUser.UpdatedAt = int64(time.Now().Unix())
	s.users[curUser.ID] = curUser
	s.mails[curUser.Email] = curUser.ID

	return nil
}

func (s *UserStore) Delete(ctx context.Context, id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[id]; !ok {
		return errors.New("user not found")
	}

	delete(s.mails, s.users[id].Email)
	delete(s.users, id)

	return nil
}

package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/collection_manager_memory"
)

type Manager struct {
	collection *collection_manager_memory.Manager[*User]
	users      []*User
}

func NewManager(path string) (*Manager, error) {

	manager := &Manager{}

	var err error
	manager.collection, err = collection_manager_memory.New[*User](path, "users")
	if err != nil {
		return nil, err
	}

	err = manager.load()
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) load() error {

	items, err := m.collection.ReadAll()
	if err != nil {
		return err
	}

	m.users = []*User{}

	fmt.Println("--------------------------------------------------")
	for _, item := range items {
		fmt.Println(item.ID, item.Username)
		m.users = append(m.users, item)
	}
	fmt.Println("--------------------------------------------------")

	return nil
}

func (m *Manager) Create(u *User) (*User, error) {
	err := u.Validate()
	if err != nil {
		return nil, err
	}

	users, err := m.collection.ReadAll()
	for _, user := range users {
		if user.Username == u.Username {
			return nil, fmt.Errorf("user with username %s already exists", u.Username)
		}
		if user.Email == u.Email {
			return nil, fmt.Errorf("user with email %s already exists", u.Email)
		}
		if user.PhoneNumber == u.PhoneNumber {
			return nil, fmt.Errorf("user with phone number %s already exists", u.PhoneNumber)
		}
	}

	if u.DisplayName == "" {
		u.DisplayName = u.FirstName + " " + u.LastName
	}

	//avatarID, err := uuid.NewV7()
	//if err != nil {
	//	return nil, fmt.Errorf("failed to generate avatar id: %w", err)
	//}
	//u.ThumbnailURL = config.RootDir + "/assets/" + avatarID.String() + ".jpg"

	u.Version = "1"

	item, err := m.collection.Create(u)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *Manager) Read(id uuid.UUID) (*User, error) {
	item, err := m.collection.Read(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *Manager) ReadAll(with *SearchOptions) ([]*User, error) {
	items, err := m.collection.ReadAll()
	if err != nil {
		return nil, err
	}

	//filterItems := Search(items, with)

	return items, nil
}

func (m *Manager) Update(with UpdateOptions) (*User, error) {

	if with.ID == uuid.Nil {
		return nil, fmt.Errorf("cannot update user without an ID")
	}

	item, err := m.collection.Read(with.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to read u1 %s: %w", with.ID, err)
	}

	err = with.ValidateUpdate()
	if err != nil {
		return nil, err
	}

	if with.Username != nil {
		if item.Username == *with.Username {
			return nil, fmt.Errorf("username %s already exists", *with.Username)
		}
	}
	if with.Email != nil {
		if item.Email == *with.Email {
			return nil, fmt.Errorf("email %s already exists", *with.Email)
		}
	}
	if with.PhoneNumber != nil {
		if item.PhoneNumber == *with.PhoneNumber {
			return nil, fmt.Errorf("phone number %s already exists", *with.PhoneNumber)
		}
	}

	Update(item, with)

	create, err := m.collection.Update(item)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (m *Manager) Delete(id uuid.UUID) error {
	err := m.collection.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) IsExist(id uuid.UUID) error {
	_, err := m.collection.Read(id)
	if err != nil {
		return fmt.Errorf("album not found: %s", id)
	}

	return nil
}

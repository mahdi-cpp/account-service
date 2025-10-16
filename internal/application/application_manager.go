package application

import (
	"sync"
	"time"

	"github.com/mahdi-cpp/account-service/internal/collections/user"
)

const (
	CommandChannel      = "application/command"
	userChannel         = "application/user"
	ListChannel         = "application/list"
	userAddChannel      = "application/user/add"
	userDeleteChannel   = "application/user/delete"
	userUpdateChannel   = "application/user/update"
	publishTimeout      = 2 * time.Second
	subscriptionTimeout = 3 * time.Second
)

type AppManager struct {
	mu          sync.RWMutex
	UserManager *user.Manager

	//rdb        *redis.Client
	//ctx      context.Context
	//cancel   context.CancelFunc
	//wg       sync.WaitGroup
	//subReady chan struct{}
}

func New() (*AppManager, error) {

	//ctx, cancel := context.WithCancel(context.Background())
	manager := &AppManager{
		//ctx:      ctx,
		//cancel:   cancel,
		//subReady: make(chan struct{}),
	}

	var err error
	manager.UserManager, err = user.NewManager("/app/iris/services/accounts")
	if err != nil {
		return nil, err
	}

	// Initialize Redis client
	//manager.rdb = redis.NewClient(&redis.Options{
	//	Addr: "localhost:6389",
	//	DB:   0,
	//})

	// Verify Redis connection
	//ctxPing, cancel := context.WithTimeout(ctx, 2*time.Second)
	//defer cancel()
	//if _, err := manager.rdb.Ping(ctxPing).Result(); err != nil {
	//	return nil, fmt.Errorf("redis connection failed: %w", err)
	//}

	// Initialize user collection
	//var err error
	//manager.collection, err = collection_manager_memory.New[*user.User](config.RootDir, "users")
	//if err != nil {
	//	return nil, fmt.Errorf("failed to initialize user collection: %w", err)
	//}
	//
	//manager.collectionJson, err = collection_manager_json.New[*user.User](config.RootDir + "/metadata")
	//if err != nil {
	//	return nil, fmt.Errorf("failed to initialize user collection: %w", err)
	//}

	// Start subscription handler
	//manager.wg.Add(1)
	//go manager.runSubscription()

	// Wait for subscription to be ready
	//select {
	//case <-manager.subReady:
	//	log.Println("Redis subscription established")
	//case <-time.After(subscriptionTimeout):
	//	log.Println("Warning: Subscription setup timed out")
	//case <-ctx.Done():
	//	return nil, context.Canceled
	//}

	return manager, nil
}

//func (m *AppManager) Close() error {
//	m.cancel()  // Signal shutdown
//	m.wg.Wait() // Wait for goroutines
//
//	if err := m.rdb.Close(); err != nil {
//		return fmt.Errorf("redis close error: %w", err)
//	}
//	return nil
//}
//
//func (m *AppManager) runSubscription() {
//	defer m.wg.Done()
//
//	channels := []string{
//		commandChannel,
//		userChannel,
//		listChannel,
//		userAddChannel,
//		userDeleteChannel,
//		userUpdateChannel,
//	}
//
//	pubsub := m.rdb.Subscribe(m.ctx, channels...)
//	defer pubsub.Close()
//
//	// Confirm subscription
//	if _, err := pubsub.ReceiveTimeout(m.ctx, 500*time.Millisecond); err != nil {
//		log.Printf("Subscription confirmation failed: %v", err)
//		return
//	}
//	close(m.subReady)
//
//	ch := pubsub.Channel()
//	for {
//		select {
//		case msg, ok := <-ch:
//			if !ok {
//				log.Println("Subscription channel closed")
//				return
//			}
//			m.handleMessage(msg)
//		case <-m.ctx.Done():
//			log.Println("Subscription exiting due to shutdown")
//			return
//		}
//	}
//}

//func (m *AppManager) handleMessage(msg *redis.Message) {
//	switch msg.Channel {
//	case commandChannel:
//		switch msg.Payload {
//		case "list":
//			if err := m.Publish(); err != nil {
//				log.Printf("Publish failed: %v", err)
//			}
//		case "user":
//			// Handle user command
//		}
//	case userAddChannel:
//		// Handle user addition
//	case userDeleteChannel:
//		// Handle user deletion
//	case userUpdateChannel:
//		// Handle user update
//	}
//}

//func (m *AppManager) Publish() error {
//
//	users, err := m.collection.ReadAll()
//	if err != nil {
//		return fmt.Errorf("get users failed: %w", err)
//	}
//
//	toJSON, err := help.ToStringJson(users)
//	if err != nil {
//		return fmt.Errorf("JSON conversion failed: %w", err)
//	}
//
//	ctx, cancel := context.WithTimeout(m.ctx, publishTimeout)
//	defer cancel()
//
//	if err := m.rdb.Publish(ctx, listChannel, toJSON).Err(); err != nil {
//		return fmt.Errorf("redis publish failed: %w", err)
//	}
//
//	return nil
//}

//func (m *AppManager) CreateUser(u *user.User) (*user.User, error) {
//
//	err := u.Validate()
//	if err != nil {
//		return nil, err
//	}
//
//	users, err := m.collection.ReadAll()
//	for _, user := range users {
//		if user.Username == u.Username {
//			return nil, fmt.Errorf("user with username %s already exists", u.Username)
//		}
//		if user.Email == u.Email {
//			return nil, fmt.Errorf("user with email %s already exists", u.Email)
//		}
//		if user.PhoneNumber == u.PhoneNumber {
//			return nil, fmt.Errorf("user with phone number %s already exists", u.PhoneNumber)
//		}
//	}
//
//	if u.DisplayName == "" {
//		u.DisplayName = u.FirstName + " " + u.LastName
//	}
//
//	avatarID, err := uuid.NewV7()
//	if err != nil {
//		return nil, fmt.Errorf("failed to generate avatar id: %w", err)
//	}
//	u.AvatarURL = config.RootDir + "/assets/" + avatarID.String() + ".jpg"
//
//	u.Version = "1"
//
//	// Step 3: create the user in the database
//	_, err = m.collection.Create(u)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create user in database: %w", err)
//	}
//
//	_, err = m.collectionJson.Create(u)
//	if err != nil {
//		return nil, err
//	}
//
//	return u, nil
//}
//
//func (m *AppManager) ReadUser(userID uuid.UUID) (*user.User, error) {
//
//	user1, err := m.collection.Read(userID)
//	if err != nil {
//		return nil, err
//	}
//
//	return user1, nil
//}

//func (m *AppManager) ReadUsers(with *user.SearchOptions) ([]*user.User, error) {
//
//	users, err := m.collection.ReadAll()
//	if err != nil {
//		return nil, err
//	}
//
//	filterUsers := user.Search(users, with)
//
//	return filterUsers, nil
//}
//
//func (m *AppManager) UpdateUsers(with *user.UpdateOptions) (*user.User, error) {
//
//	if with.ID == uuid.Nil {
//		return nil, fmt.Errorf("cannot update user without an ID")
//	}
//
//	selectUser, err := m.collection.Read(with.ID)
//	if err != nil {
//		return nil, fmt.Errorf("failed to read u1 %s: %w", with.ID, err)
//	}
//
//	err = with.ValidateUpdate()
//	if err != nil {
//		return nil, err
//	}
//
//	users, err := m.collection.ReadAll()
//	for _, u1 := range users {
//		if u1.ID != with.ID {
//			if with.Username != nil {
//				if u1.Username == *with.Username {
//					return nil, fmt.Errorf("username %s already exists", *with.Username)
//				}
//			}
//			if with.Email != nil {
//				if u1.Email == *with.Email {
//					return nil, fmt.Errorf("email %s already exists", *with.Email)
//				}
//			}
//			if with.PhoneNumber != nil {
//				if u1.PhoneNumber == *with.PhoneNumber {
//					return nil, fmt.Errorf("phone number %s already exists", *with.PhoneNumber)
//				}
//			}
//		}
//	}
//
//	user.Update(selectUser, with)
//
//	updateUser, err := m.collection.Update(selectUser)
//	if err != nil {
//		return nil, fmt.Errorf("failed to update u1 %s: %w", with.ID, err)
//	}
//
//	_, err = m.collectionJson.Update(selectUser)
//	if err != nil {
//		return nil, err
//	}
//
//	return updateUser, nil
//}

//func (m *AppManager) DeleteUser(userID uuid.UUID) error {
//
//	err := m.collection.Delete(userID)
//	if err != nil {
//		fmt.Println("error deleting user")
//		return err
//	}
//	return nil
//}

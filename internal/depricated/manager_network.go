package depricated

import (
	"log"

	"github.com/mahdi-cpp/account-service/internal/user"
	"github.com/mahdi-cpp/iris-tools/network"
)

type NetworkManager struct {
	networkUser     *network.Control[user.User]
	networkUserList *network.Control[[]user.User]
}

type requestBody struct {
	UserID string `json:"userID"`
}

func NewNetworkAccountManager() *NetworkManager {
	manager := &NetworkManager{
		networkUser:     network.NewNetworkManager[user.User]("http://localhost:8080/api/v1/user/get_user"),
		networkUserList: network.NewNetworkManager[[]user.User]("http://localhost:8080/api/v1/user/list"),
	}

	return manager
}

func (m *NetworkManager) GetUser(id string) (*user.User, error) {

	user, err := m.networkUser.Read("", requestBody{UserID: id})
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}
	return user, nil
}

func (m *NetworkManager) GetAll() (*[]user.User, error) {

	users, err := m.networkUserList.Read("", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}
	return users, nil
}

func (m *NetworkManager) GetByFilterOptions(userIDs []string) (*[]user.User, error) {

	users, err := m.networkUserList.Read("", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}
	return users, nil
}

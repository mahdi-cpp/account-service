package user

import (
	"fmt"
	"testing"

	"github.com/mahdi-cpp/account-service/internal/config"
	"github.com/mahdi-cpp/account-service/internal/help"
)

const workDir = "/app/iris/services/accounts"

func TestManager_Create(t *testing.T) {
	manager, err := NewManager(workDir)
	if err != nil {
		return
	}

	//u1 := &User{
	//	Username:     "parsa_1368",
	//	FirstName:    "Parsa",
	//	LastName:     "Nasiri",
	//	Email:        "parasa@gmail.com",
	//	PhoneNumber:  "09125640233",
	//	OriginalURL:  workDir + "/assets/00/chat_41.jpg",
	//	ThumbnailURL: workDir + "/assets/thumbnails/chat_41",
	//}

	//u := &User{
	//	Username:     "maryam",
	//	FirstName:    "Maryam",
	//	LastName:     "Farahmand",
	//	Email:        "maryam_1369@gmail.com",
	//	PhoneNumber:  "09354442388",
	//	OriginalURL:  "",
	//	ThumbnailURL: workDir + "/assets/thumbnails/maryam",
	//}

	//u := &User{
	//	Username:     "ali_safari",
	//	FirstName:    "Ali",
	//	LastName:     "Safari",
	//	Email:        "ali_safari_1375@gmail.com",
	//	PhoneNumber:  "09124456978",
	//	OriginalURL:  "",
	//	ThumbnailURL: workDir + "/assets/thumbnails/chat_26",
	//}

	u := &User{
		Username:     "parastoo_1375",
		FirstName:    "Parastoo",
		LastName:     "Aslani",
		DisplayName:  "Parastoo Aslani",
		Email:        "parastoo_1375@gmail.com",
		PhoneNumber:  "09122456978",
		OriginalURL:  "",
		ThumbnailURL: workDir + "/assets/thumbnails/Parastoo_Aslani",
	}

	create, err := manager.Create(u)
	if err != nil {
		t.Fatalf("%v", err)
	}

	fmt.Println(create.ID)
}

func TestManager_Update(t *testing.T) {

	config.Init()

	manager, err := NewManager(workDir)
	if err != nil {
		return
	}

	u := UpdateOptions{
		ID:           config.Maryam,
		ThumbnailURL: help.StrPtr(workDir + "/assets/thumbnails/maryam2"),
	}

	user, err := manager.Update(u)
	if err != nil {
		t.Fatalf("%v", err)
	}

	fmt.Println(user.ThumbnailURL)
}

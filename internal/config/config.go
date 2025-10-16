package config

import (
	"log"

	"github.com/google/uuid"
)

const RootDir = "/app/iris/services/accounts"

var (
	Mahdi    uuid.UUID
	Parsa    uuid.UUID
	Ali      uuid.UUID
	Maryam   uuid.UUID
	Parastoo uuid.UUID
)

func Init() {
	initUsers()
}

func initUsers() {
	var err error

	Mahdi, err = uuid.Parse("0199b306-d156-7c6b-b122-1b309599fb82")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}

	Parsa, err = uuid.Parse("0199edc1-469b-7de7-bf36-c6cbf299b874")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}

	Ali, err = uuid.Parse("0199ee86-f4ef-722a-8355-f860ce9513a3")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}

	Maryam, err = uuid.Parse("0199ee39-47f9-7673-88aa-e311f3ad0575")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}

	Parastoo, err = uuid.Parse("0199ee92-2bed-7a9b-921a-e897e3cbc42c")
	if err != nil {
		log.Fatalf("failed to parse Mahdi: %v", err)
	}

}

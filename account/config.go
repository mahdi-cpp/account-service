package account

import (
	"path/filepath"
)

const rootDir = "/app/iris/"
const serviceDir = "services"

func GetServicesPath(file string) string {
	return filepath.Join(rootDir, serviceDir, file)
}

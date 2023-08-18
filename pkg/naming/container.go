package naming

import (
	"os"
)

const (
	LOCAL_ENV              = "local"
	LOCAL_CONTAINER_PREFIX = "dev.local/ktwin/"
	CONTAINER_REGISTRY     = "ghcr.io/ktwins/"
)

func GetContainerRegistry(container string) string {
	if os.Getenv("ENV") == LOCAL_ENV {
		return LOCAL_CONTAINER_PREFIX + container
	}
	return CONTAINER_REGISTRY + container
}
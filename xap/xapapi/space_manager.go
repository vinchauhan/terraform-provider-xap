package xapapi

import (
	"log"
)

type Space struct {
	name              string
	partitions        int
	backups           bool
	requiresIsolation bool
}

type SpaceManager struct {
	log         *log.Logger
	apiEndpoint string
}

func NewSpaceManager(logger *log.Logger, apiEndpoint string, session *Session) *SpaceManager {
	dm := &SpaceManager{
		log:         logger,
		apiEndpoint: apiEndpoint,
	}
	return dm
}

func (sm *SpaceManager) CreateSpace() {

}

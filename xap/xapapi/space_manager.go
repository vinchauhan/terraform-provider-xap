package xapapi

import (
	"log"
)

type SpaceManager struct {
	log         *log.Logger
	apiEndpoint string
	session     *Session
}

func newSpaceManager(logger *log.Logger, apiEndpoint string, session *Session) *SpaceManager {
	dm := &SpaceManager{
		log:         logger,
		apiEndpoint: apiEndpoint,
		session:     session,
	}
	return dm
}
func (sm *SpaceManager) CreateSpace() {

}

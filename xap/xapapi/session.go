package xapapi

import (
	"crypto/tls"
	"log"
	"net/http"
)

type Session struct {
	httpClient   *http.Client
	log          *log.Logger
	spaceManager *SpaceManager
}

func NewSession(endpoint string, user string, password string, skipSslValidation bool) (*Session, error) {
	s = &Session{
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: skipSslValidation},
			},
		},
		log: nil,
	}
	//err = s.initSession(endpoint, user, password, skipSslValidation)
	//if err != nil {
	//	return nil, err
	//}

	return s, nil
}

//func (s *Session) initSession(endpoint string, user string, password string)  {
//		s.spaceManager = newSpaceManager(s.log, endpoint,s. )
//}

package storage

import (
	"errors"
	"log"

	"github.com/segmentio/ksuid"
)

type StorageRepo struct {
}

type AuthData struct {
	Success   bool
	MovieData []string
	ID        string
	FullName  string
	UserName  string
	Password  string
	SessionID string
}

var authData AuthData

var mockData map[string]AuthData
var mockDataPurchases map[string][]string

func (s *StorageRepo) AuthenticateRoute(id string, sessionID string) error {
	if val, ok := mockData[id]; ok {
		if val.SessionID != sessionID {
			return errors.New("Unauthorized")
		}
	}
	return nil
}

func (s *StorageRepo) GetAppDataById(id string) (AuthData, error) {
	return mockData[id], nil
}

func (s *StorageRepo) Login(userName string, password string) (AuthData, error) {
	return mockData["1"], nil
}

func (s *StorageRepo) SetMockData() {
	mockData = make(map[string]AuthData)
	mockDataPurchases = make(map[string][]string)
	mockData["1"] = AuthData{true, []string{"Star Wars, 1 day rental - $3.99", "Die Hard, 1 day rental - $4.99", "Joker, 2 day rental - $9.99", "Interstellar, 1 day rental - $4.99", "Caddyshack, 1 day rental - $5.99"}, "1", "Carl Bellitti", "c@b.com", "pw", ksuid.New().String()}
	mockData["2"] = AuthData{true, []string{"XXX", "YYY", "ZZZ"}, "2", "Joe Schmoe", "j@s.com", "pw", ksuid.New().String()}
	mockData["3"] = AuthData{true, []string{"AAA", "BBB", "CCC"}, "3", "John Doe", "j@d.com", "pw", ksuid.New().String()}
	mockDataPurchases["1"] = []string{"Star Wars, $3.99", "Caddyshack, $5.99", "Joker, $9.99"}
}

func (s *StorageRepo) CheckLogin(username string, password string) (AuthData, error) {
	var a AuthData
	log.Println("MOCK1:" + mockData["1"].FullName)
	for k, v := range mockData {
		if v.UserName == username && v.Password == password {
			newSessionID := ksuid.New().String()
			v.SessionID = newSessionID
			mockData[k] = v
			a = mockData[k]
			return a, nil
		}
	}
	return AuthData{}, errors.New("USER NOT FOUND")
}

func (s *StorageRepo) SetNewSessionID(id string) {
	if _, ok := mockData[id]; ok {
		newSessionID := ksuid.New().String()
		mockDataForID := mockData[id]
		mockDataForID.SessionID = newSessionID
		mockData[id] = mockDataForID
	}
}

func (s *StorageRepo) GetPurchases(id string) []string {
	if _, ok := mockDataPurchases[id]; ok {
		return mockDataPurchases[id]
	}
	return []string{}
}

func (s *StorageRepo) Logout(id string) bool {
	s.SetNewSessionID(id)
	return true
}

package code 

import (
	"sort"
	//"math"
	//"sync"
	//"time"
	//"strings"
	//"encoding/json"

	"fmt"
)

type AnybaseService struct {
	bin BinStorage
}

var _ Server = new(AnybaseService)

func (self *AnybaseService) findUserBin(userName string) (Storage, error) {
	return nil, nil
}
func (self *AnybaseService) userExists(user string) (bool, error) {
	return false, nil 
}
func (self *AnybaseService) addUser(userName string) error {
	return nil
}
// Creates a user
func (self *AnybaseService) SignUp(user string) error {
	return nil
}
// List 20 registered users.  When there are less than 20 users that
// signed up the service, all of them needs to be listed.  When there
// are more than 20 users that signed up the service, an arbitrary set
// of at least 20 of them needs to be listed.
// The result should be sorted in alphabetical order.
func (self *AnybaseService) ListUsers() ([]string, error) {
	var secretBin Storage
	var twentyUsersList List
	var err error

	secretBin = self.bin.Bin(SECRET_BIN_KEY)
	err = secretBin.ListGet(USER_LIST, &twentyUsersList)
	if err != nil{
		return nil, fmt.Errorf("error in getting secret bin", USER_LIST)
	}

	//twentyUsersList = RemoveDuplicates(twentyUsersList)
	//sort the list
	sort.Strings(twentyUsersList.L)
	if len(twentyUsersList.L) > MIN_DOCS {
		twentyUsersList.L = twentyUsersList.L[:MIN_DOCS]
	}
	//fmt.Println("[LOG][TWENTY]", twentyUsersList.L)

	return twentyUsersList.L, nil
}
func (self *AnybaseService) ListDocs() ([]string, error) {

	var secretBin Storage
	var docList List
	var err error

	secretBin = self.bin.Bin(SECRET_BIN_KEY)
	err = secretBin.ListGet(DOC_LIST, &docList)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving doc_list from secret bin", DOC_LIST)
	}

	sort.Strings(docList.L)
	if len(docList.L) > MIN_DOCS {
		docList.L = docList.L[:MIN_DOCS]
	}
	//docList.L = append(docList.L, "Sample Document", "Test Doc")
	return docList.L, nil	
}
// Post a tribble.  The clock is the maximum clock value this user has
// seen so far by reading tribbles or clock sync.
func (self *AnybaseService) Post(who, post string, clock uint64) error {
	//fmt.Println("[AnybaseService][Post]")
	return nil
}
// List the tribs that a particular user posted.
func (self *AnybaseService) Tribs(user string) ([]*Trib, error){
	var randomV []*Trib
return randomV, nil
}
// Follow someone's timeline. Returns error when who == whom.
func (self *AnybaseService) Follow(who, whom string) error {
	return nil
}
// Unfollow someone's timeline. Returns error when who == whom.
func (self *AnybaseService) Unfollow(who, whom string) error {
	return nil
}
// Returns true when who following whom. Returns error when who == whom.
func (self *AnybaseService) IsFollowing(who, whom string) (bool, error) {
	return false, nil 
}
// Returns the list of following users.
func (self *AnybaseService) Following(who string) ([]string, error){
	var randomV []string
	return randomV, nil
}
// List the tribs of someone's following users (including himself).
func (self *AnybaseService) Home(user string) ([]*Trib, error){
	var randomV []*Trib
	return randomV, nil
}

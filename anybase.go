package code 

import (
	//"sort"
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
	fmt.Println("[AnybaseService][ListUsers]")
	tempList := []string{"user", "dummy", "singh"}
	return tempList, nil
}
func (self *AnybaseService) ListDocs() ([]string, error) {
	fmt.Println("[AnybaseService][ListDocs]")
	tempList := []string{"Sample_document_1", "dummy_document_1", "moar doc"}
	return tempList, nil	
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

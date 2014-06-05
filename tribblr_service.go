package code 

import (
	//"sort"
	//"math"
	//"sync"
	//"time"
	//"strings"
	//"encoding/json"

	//"fmt"
)

type TribblrService struct {
	bin BinStorage
}

var _ Server = new(TribblrService)

func (self *TribblrService) findUserBin(userName string) (Storage, error) {
	return nil, nil
}
func (self *TribblrService) userExists(user string) (bool, error) {
	return false, nil 
}
func (self *TribblrService) addUser(userName string) error {
	return nil
}
// Creates a user
func (self *TribblrService) SignUp(user string) error {
	return nil
}
// List 20 registered users.  When there are less than 20 users that
// signed up the service, all of them needs to be listed.  When there
// are more than 20 users that signed up the service, an arbitrary set
// of at least 20 of them needs to be listed.
// The result should be sorted in alphabetical order.
func (self *TribblrService) ListUsers() ([]string, error) {
	var tempList []string 
	return tempList, nil
}
// Post a tribble.  The clock is the maximum clock value this user has
// seen so far by reading tribbles or clock sync.
func (self *TribblrService) Post(who, post string, clock uint64) error {
	return nil
}
// List the tribs that a particular user posted.
func (self *TribblrService) Tribs(user string) ([]*Trib, error){
	var randomV []*Trib
return randomV, nil
}
// Follow someone's timeline. Returns error when who == whom.
func (self *TribblrService) Follow(who, whom string) error {
	return nil
}
// Unfollow someone's timeline. Returns error when who == whom.
func (self *TribblrService) Unfollow(who, whom string) error {
	return nil
}
// Returns true when who following whom. Returns error when who == whom.
func (self *TribblrService) IsFollowing(who, whom string) (bool, error) {
	return false, nil 
}
// Returns the list of following users.
func (self *TribblrService) Following(who string) ([]string, error){
	var randomV []string
	return randomV, nil
}
// List the tribs of someone's following users (including himself).
func (self *TribblrService) Home(user string) ([]*Trib, error){
	var randomV []*Trib
	return randomV, nil
}

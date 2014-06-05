package code 

import (
	//"net"

	"fmt"
)

type DSClient struct {
	userName string
	backs []string
	storageList []Storage
	storageStatus []bool
}

var _ Storage = new(DSClient)
//returns a Storage object which will contain three Storages
//this class will be responsible for simultaneously updating all three etc.
func NewDSClient(backs []string, binName string) Storage {
	storageStatus := []bool{true, true, true}
	return &DSClient{ 
		backs 		: backs,
		storageList : nil,
		userName 	: binName,
		storageStatus : storageStatus,
	}
}
// Gets a value. Empty string by default.
func (self *DSClient) Get(key string, value *string) error {
		fmt.Println("[DSClient][Get]")
		fmt.Println(self.backs, self.storageList, self.userName, self.storageStatus)
	    return nil
}

// Set kv.Key to kv.Value. Set succ to true when no error.
func (self *DSClient) Set(kv *KeyValue, succ *bool) error {
		fmt.Println("[DSClient][Set]")
    return nil 
}

// List all the keys of non-empty pairs where the key matches
// the given pattern.
func (self *DSClient) Keys(p *Pattern, list *List) error {
		fmt.Println("[DSClient][Keys]")
    return nil
}
// Get the list.
func (self *DSClient) ListGet(key string, list *List) error {
		fmt.Println("[DSClient][ListGet]")
    return nil
}

// Append a string to the list. Set succ to true when no error.
func (self *DSClient) ListAppend(kv *KeyValue, succ *bool) error {
		fmt.Println("[DSClient][ListAppend]")
    return nil 
}

// Removes all elements that equals to kv.Value in list kv.Key
// n is set to the number of elements removed.
func (self *DSClient) ListRemove(kv *KeyValue, n *int) error {
		fmt.Println("[DSClient][ListRemove]")
    return nil 
}

// List all the keys of non-empty lists, where the key matches
// the given pattern.
func (self *DSClient) ListKeys(p *Pattern, list *List) error {
		fmt.Println("[DSClient][ListKeys]")
    return nil
}

// Returns an auto-incrementing clock. The returned value of each call will
// be unique, no smaller than atLeast, and strictly larger than the value
// returned last time, unless it was math.MaxUint64.
func (self *DSClient) Clock(atLeast uint64, ret *uint64) error {
		fmt.Println("[DSClient][Clock]")
    return nil
}

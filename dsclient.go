package code 

import (

	"fmt"
)

type DSClient struct {
	userName string
	backs []string
	storageList []Storage
	storageStatus []bool
}

var _ Storage = new(DSClient)
//returns a trib.Storage object which will contain three Storages
//this class will be responsible for simultaneously updating all three etc.
func NewDSClient(backs []string, binName string) Storage {
	//id := BinNumber(binName)
	storageList := make([]Storage, 3)
	//primary
	//slide to find next alive bin
	nBin1 := HashNSlide(backs, BinNumber(len(backs), binName))
	storageList[0] = NewStorageClient(backs[nBin1], binName)
	//replica 1
	//slide to find next alive bin
	nBin2 := HashNSlide(backs, nBin1 + 1)
	storageList[1] = NewStorageClient(backs[nBin2], binName)
	//replica 2
	//slide to find next alive bin
	nBin3 := HashNSlide(backs, nBin2 + 1)
	storageList[2] = NewStorageClient(backs[nBin3], binName)
	fmt.Println("[LOG][NewDSClient]", binName, storageList[0], storageList[1], storageList[2])

	storageStatus := []bool{true, true, true} 
	return &DSClient{ 
		backs 		: backs,
		storageList : storageList,
		userName 	: binName,
		storageStatus : storageStatus,
	}
}
// Gets a value. Empty string by default.
func (self *DSClient) Get(key string, value *string) error {
	//retrieve from primary bin
	//first check if its online
	//if not then refresh list
	//try again
	//get from primary, replica 1, replica 2
	//see if all agree
	//if not then make them agree
	//return the damn value

	//current approach
	//return from anywhere possible!
	//marshall key
	e := self.storageList[0].Get(key, value)
	if e != nil {
		//set refresh flag true
		//#TODO refreshPlaceHolder
		e = self.storageList[1].Get(key, value)
		if e != nil {
			//set refresh flag true
			//#TODO refreshPlaceHolder
			e = self.storageList[2].Get(key, value)
			if e != nil {
				//set refresh flag true
				//#TODO refreshPlaceHolder
				return e
			}
		}
	}
    return nil
}

// Set kv.Key to kv.Value. Set succ to true when no error.
func (self *DSClient) Set(kv *KeyValue, succ *bool) error {
	//set to all 3 storages
	//on failure
	//refreshes
	//marshall key

	//special case when addUser is called
	if kv.Key == EXISTS_KEY {
		err := self.storageList[0].ListAppend(KV(PRIMARY_USERS_KEY, self.userName), succ)
		if err != nil {
			return err
		}
		err = self.storageList[1].ListAppend(KV(REPLICA_ONE_USERS_KEY, self.userName), succ)
		if err != nil {
			return err
		}
		err = self.storageList[2].ListAppend(KV(REPLICA_TWO_USERS_KEY, self.userName), succ)
		if err != nil {
			return err
		}
	}
	e := self.storageList[0].Set(kv, succ)
	if e != nil {
		self.storageStatus[0] = false
	}
	e = self.storageList[1].Set(kv, succ)
	if e != nil {
		self.storageStatus[1] = false
	}
	e = self.storageList[2].Set(kv, succ)
	if e != nil {
		self.storageStatus[2] = false
	}
	if self.allStoragesDown(){
		return fmt.Errorf("All Storages down for %q", self.userName)
	}
    return nil 
}

// List all the keys of non-empty pairs where the key matches
// the given pattern.
func (self *DSClient) Keys(p *Pattern, list *List) error {
	e := self.storageList[0].Keys(p, list)
	if e != nil {
		//set refresh flag true
		//#TODO refreshPlaceHolder
		e = self.storageList[1].Keys(p, list)
		if e != nil {
			//set refresh flag true
			//#TODO refreshPlaceHolder
			e = self.storageList[2].Keys(p, list)
			if e != nil {
				//set refresh flag true
				//#TODO refreshPlaceHolder
				return e
			}
		}
	}
    return nil
}
// Get the list.
func (self *DSClient) ListGet(key string, list *List) error {
	fmt.Println("[DSClient][ListGet]", key)
	e := self.storageList[0].ListGet(key, list)
	if e != nil {
		//set refresh flag true
		//#TODO refreshPlaceHolder
		e = self.storageList[1].ListGet(key, list)
		if e != nil {
			//set refresh flag true
			//#TODO refreshPlaceHolder
			e = self.storageList[2].ListGet(key, list)
			if e != nil {
				//set refresh flag true
				//#TODO refreshPlaceHolder
				return e
			}
		}
	}
    return nil
}

// Append a string to the list. Set succ to true when no error.
func (self *DSClient) ListAppend(kv *KeyValue, succ *bool) error {
	e := self.storageList[0].ListAppend(kv, succ)
	if e != nil {
		self.storageStatus[0] = false
	}
	//to handle the case <3 backends
	if len(self.backs) > 1{
		e = self.storageList[1].ListAppend(kv, succ)
		if e != nil {
			self.storageStatus[1] = false
		}

	}
	//to handle the case <3 backends
	if len(self.backs) > 2{
		e = self.storageList[2].ListAppend(kv, succ)
		if e != nil {
			self.storageStatus[2] = false
		}
	}
	if self.allStoragesDown(){
		return fmt.Errorf("All Storages down for %q", self.userName)
	}
    return nil 
}

// Removes all elements that equals to kv.Value in list kv.Key
// n is set to the number of elements removed.
func (self *DSClient) ListRemove(kv *KeyValue, n *int) error {
	e := self.storageList[0].ListRemove(kv, n)
	if e != nil {
		self.storageStatus[0] = false
	}
	//to handle the case <3 backends
	if len(self.backs) > 1{
		e = self.storageList[1].ListRemove(kv, n)
		if e != nil {
			self.storageStatus[1] = false
		}
	}
	//to handle the case <3 backends
	if len(self.backs) > 2{
		e = self.storageList[2].ListRemove(kv, n)
		if e != nil {
			self.storageStatus[2] = false
		}
	}
	if self.allStoragesDown(){
		return fmt.Errorf("All Storages down for %q", self.userName)
	}
    return nil 
}

// List all the keys of non-empty lists, where the key matches
// the given pattern.
func (self *DSClient) ListKeys(p *Pattern, list *List) error {
	e := self.storageList[0].ListKeys(p, list)
	if e != nil {
		//set refresh flag true
		//#TODO refreshPlaceHolder
		e = self.storageList[1].ListKeys(p, list)
		if e != nil {
			//set refresh flag true
			//#TODO refreshPlaceHolder
			e = self.storageList[2].ListKeys(p, list)
			if e != nil {
				//set refresh flag true
				//#TODO refreshPlaceHolder
				return e
			}
		}
	}
    return nil
}

// Returns an auto-incrementing clock. The returned value of each call will
// be unique, no smaller than atLeast, and strictly larger than the value
// returned last time, unless it was math.MaxUint64.
func (self *DSClient) Clock(atLeast uint64, ret *uint64) error {
	clocks := make([]uint64, 3)
	var e error
	e = self.storageList[0].Clock(atLeast, &clocks[0])
	e = self.storageList[1].Clock(atLeast, &clocks[1])
	e = self.storageList[2].Clock(atLeast, &clocks[2])

	for _, clock := range clocks {
		if *ret < clock {
			*ret = clock
		}
	}

	if e != nil {
		return e
	}
	fmt.Println("[LOG][dsclient][Clock]", *ret)
    return nil
}
// added for lab3
// will be called when it appears that the 1st of the three storages is down 
func (self *DSClient) RefreshStorage() error {
	return nil
}
func (self *DSClient) performSet(key string, succ *bool) error {
	return nil
}
//check if all storages are down
func (self *DSClient) allStoragesDown() bool {
	return !self.storageStatus[0] && 
		!self.storageStatus[1] &&
		!self.storageStatus[2]
}
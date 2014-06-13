package code

import(
	"net/rpc"
	"strings"

	//"fmt"
)

type StorageClient struct {
	addr string
	binName string
}

var _ Storage = new(StorageClient)

func (self *StorageClient) Clock(atLeast uint64, ret *uint64) error {
    //connect to the server
    conn, e := rpc.DialHTTP("tcp", self.addr)
    if e != nil {
        return e
    }

    //perform the call
    e = conn.Call("Storage.Clock", atLeast, ret)
    if e != nil {
        conn.Close()
        return e
    }

    //close the connection
    return conn.Close()
}
func (self *StorageClient) Get(key string, value *string) error {
	//marshall key
	key = self.Marshall(key)

	// connect to the server
    conn, e := rpc.DialHTTP("tcp", self.addr)
    if e != nil {
        return e
    }

    // perform the call
    e = conn.Call("Storage.Get", key, value)
    if e != nil {
        conn.Close()
        return e
    }
    //fmt.Println("[GET] ", key, *value)
    // close the connection
    return conn.Close()
}
func (self *StorageClient) Set(kv *KeyValue, succ *bool) error {
	//marshall key
	kv.Key = self.Marshall(kv.Key)
    e := self.SetRaw(kv, succ)
    if e != nil {
        return e
    }
    return nil
}
func (self *StorageClient) SetRaw(kv *KeyValue, succ *bool) error {
    //connect to the server
    conn, e := rpc.DialHTTP("tcp", self.addr)
    if e != nil {
        return e
    }

    //perform the call
    e = conn.Call("Storage.Set", kv, succ)
    if e != nil {
        conn.Close()
        return e
    }

    //fmt.Println("[SET] ", kv.Key, *succ)
    //close the connection
    return conn.Close()
}
func (self *StorageClient) Keys(p *Pattern, r *List) error {
	//marshall key
	p.Prefix = self.Marshall(p.Prefix)

    //connect to the server
    conn, e := rpc.DialHTTP("tcp", self.addr)
    if e != nil {
        return e
    }

    r.L = nil
    //perform the call
    e = conn.Call("Storage.Keys", p, r)
    if e != nil {
        conn.Close()
        return e 
    }

    if r.L == nil {
       r.L = []string{}
    } else {
    	var tempList = []string{}
    	for _, key := range r.L {
    		tempList = append(tempList, strings.SplitAfter(key, "::")[1])
    	}	
    	r.L = tempList
        //fmt.Println("[LOG][keys] removing dups")
        //*r = RemoveDuplicates(*r)
    	//fmt.Println("[KEYS]", r.L)
    }

    //close the connection
    return conn.Close()
}
func (self *StorageClient) ListKeys(p *Pattern, r *List) error {
	//marshall key
	p.Prefix = self.Marshall(p.Prefix)

    //connect to the server
    conn, e := rpc.DialHTTP("tcp", self.addr)
    if e != nil {
        return e
    }

    r.L = nil
    //perform the call
    e = conn.Call("Storage.ListKeys", p, r)
    if e != nil {
        conn.Close()
        return e
    }

    if r.L == nil {
        r.L = []string{}
    } else {
    	var tempList = []string{}
    	for _, key := range r.L {
    		tempList = append(tempList, strings.SplitAfter(key, "::")[1])
    	}	
    	r.L = tempList
        //fmt.Println("[LOG][ListKeys] removing dups")
        //*r = RemoveDuplicates(*r)
    	//fmt.Println("[LIST_KEYS]", r.L)
    }

    //close the connection
    return conn.Close()
}
func (self *StorageClient) ListGet(key string, ret *List) error {
    //marshall the key
    key = self.Marshall(key)

    //connect to the server
    conn, e := rpc.DialHTTP("tcp", self.addr)
    if e != nil {
        return e
    }

    ret.L = nil
    //perform the call
    e = conn.Call("Storage.ListGet", key, ret)
    if e != nil {
        conn.Close()
        return e
    }
    if ret.L == nil {
        ret.L = make([]string, 0)
    } 
    ret.L = append(ret.L, "Sample Document", "Lorem Ipsum", "Magna Carta")
    //ret.L = []string{"user", "santa", "banta"}
    //fmt.Println("[LOG][listget] Removing duplicates")
    //*ret = RemoveDuplicates(*ret)
    //close the connection 
    return conn.Close() 
} 
func (self *StorageClient) ListAppend(kv *KeyValue, succ *bool) error {
    //marshall the key
    kv.Key = self.Marshall(kv.Key)

    //connect to the server 
    conn, e := rpc.DialHTTP("tcp", self.addr) 

    if e != nil {
        return e 
    } 
    //perform the call 
    e = conn.Call("Storage.ListAppend", kv, succ) 
    if e != nil {
        conn.Close() 
        return e
    }
    //close the connection
    return conn.Close()
}
// Removes all elements that equals to kv.Value in list kv.Key
// n is set to the number of elements removed.
func (self *StorageClient) ListRemove(kv *KeyValue, n *int) error {
    //marshall the key
    kv.Key = self.Marshall(kv.Key)

    //connect to the server
    conn, e := rpc.DialHTTP("tcp", self.addr)
    if e != nil {
        return e
    }

    //perform the call
    e = conn.Call("Storage.ListRemove", kv, n)
    if e != nil {
        conn.Close()
        return e
    }

    //close the connection
    return conn.Close()
}
func (self *StorageClient) Marshall(key string) string{
	if len(key) < 0 || IsSpecialKey(key) || IsAlreadyMarshalled(key) {
		return key
	}
	return self.binName + "::" + key
}


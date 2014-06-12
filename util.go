package code 

import (
	"net"
	"hash/fnv"
	"strings"

)

func CheckAlive(backs []string) ([]bool, error) {
	aliveArr := make([]bool, len(backs)) 
	for i := 0; i < len(backs); i++ {
		conn, err := net.Dial("tcp", backs[i])
		if err != nil {
			aliveArr[i] = false
			//fmt.Println("[LOG][isAlive] Storage down!" , backs[i])
		} else {
			aliveArr[i] = true 
			conn.Close()
		}
	}
	//fmt.Println(aliveMap)
	return aliveArr, nil 
}

func BinNumber(nBacks int, binName string) int {
	hashFunction := fnv.New64a()
	hashFunction.Write([]byte(binName))
	binNumber := hashFunction.Sum64() % uint64(nBacks)
	return int(binNumber)
}

func IsSpecialKey(key string) bool {
    for _, a := range SPECIAL_KEYS { if a == key { return true } }
    return false
}

func IsAlreadyMarshalled(key string) bool {
	return strings.Contains(key, "::") 
}

//finds the next Alive Storage index based on provided index
func HashNSlide(backs []string, index int) int {
	//will find the next alive backend
	//#TODO check for ready flag of that backend
	i := 0
	for ; ; i++ {
		conn, err := net.Dial("tcp", backs[(index + i) % len(backs)])
		if err != nil {
			//fmt.Println("[LOG][HashNSlide] Storage down!" , backs[(index + i) % len(backs)])
		} else {
			conn.Close()
			break	
		}
	}
	return (index + i) % len(backs)
}
//finds the previous Alive Storage index based on provided index
func HashNReverseSlide(backs []string, index int) int {
	//will find the previous alive backend
	//#TODO check for ready flag of that backend
	i := 0
	for ; ; i++ {
		conn, err := net.Dial("tcp", backs[(index - i) % len(backs)])
		if err != nil {
			//fmt.Println("[LOG][HashNSlide] Storage down!" , backs[(index + i) % len(backs)])
		} else {
			conn.Close()
			break	
		}
	}
	return (index - i) % len(backs)
}
//pattern creator
func Pat(pre, suf string) *Pattern {
	return &Pattern{pre, suf}
}

func NewStorageClient(addr string, binName string) Storage {
	return &StorageClient{
		addr : addr,
		binName : binName,
	}
}

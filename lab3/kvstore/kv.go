// +build !solution

// Leave an empty line above this comment.
package main

import (
	"fmt"
	"sync"
)

type Pair struct {
	Key, Value string
}

type KVStore struct {
	lock  *sync.Mutex
	store map[string]string
}

func (kv *KVStore) Lookup(key string, val *string) error {
   defer kv.lock.Unlock()
   // kv.lock.Lock()

	if v, ok := kv.store[key]; ok {
		*val = v
	} else {
		*val = fmt.Sprintf("key '%s' not found", key)
	}
	return nil
}
func (kv *KVStore) Insert(input Pair, reply *bool) error {
    defer kv.lock.Unlock()
    kv.lock.Lock()

	kv.store[input.Key] = input.Value
	*reply = true
	return nil

}
//  keys
func (kv *KVStore) Keys(dummy bool, keys *[]string) error {
    //defer kv.lock.Unlock()
    //kv.lock.Lock()
  
    i:=0
	ks := make([]string,len(kv.store),len(kv.store))
	for key, _ := range kv.store {
        
		ks[i] = key
		i++
	} 
   fmt.Printf("Num of keys: %d",i);
	*keys = ks
	return nil
}



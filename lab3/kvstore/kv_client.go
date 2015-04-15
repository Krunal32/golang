// +build !solution

// Leave an empty line above this comment.
package main

import (
	"fmt"
	"net/rpc"
	"strconv"
)

const NUM_KEYS int = 10

var sourcemap map[string]string

func clientLoop(srvAddr string) {
	sourcemap = make(map[string]string)
	for k := 0; k < NUM_KEYS; k++ {
		key := "keynr" + strconv.Itoa(k) + ""
		value := "valuenr" + strconv.Itoa(k) + ""
		sourcemap[key] = value
		fmt.Printf("\nKey: %s  |  Value: %s created ", key, value)
	}
	fmt.Printf("\n-------------------------------------------\n")
	client, err := rpc.Dial("tcp", srvAddr)
	checkError(err)

	for key, value := range sourcemap {
		var isReturn bool
		kvPair := Pair{key, value}
		go func(pr Pair, isRet bool) {
			err = client.Call("KVStore.Insert", &pr, &isReturn)
			checkError(err)

		}(kvPair, isReturn)

	}
	retrieved := make([]string, NUM_KEYS)               // store received keys
	err = client.Call("KVStore.Keys", true, &retrieved) //ask server for all key values//
	checkError(err)
	isEqualKeys := CheckEqualKeys(retrieved)
	fmt.Printf("\n-------------------------------------------------------\n")
	fmt.Printf("All retrieved keys matches the ones sent: %t \n", isEqualKeys)

}

// check if all retrieved keys have a match in  .//
// the keys sent to the server -> return true  //
func CheckEqualKeys(retrieved []string) bool {

	allExist := true
	if len(retrieved) != NUM_KEYS {
		fmt.Printf("len retrieved != len sent\n")
		allExist = false
	}
	for i := 0; i < len(retrieved); i++ {
		key := retrieved[i]
		_, exists := sourcemap[key]
		if exists {
			fmt.Printf("Key retrieved: %s .Exists in source map: %t \n", key, exists)
		} else {
			fmt.Printf("non match (Key: %s )\n %t", key, exists)
			allExist = false
		}
	}
	return allExist
}

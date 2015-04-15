package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const filepath string = "/dev/simp_rw"

func main() {
	r := bufio.NewReader(os.Stdin) // read input from terminal
	fmt.Printf("Echo/Cat simulator. Type \"echo=<your message here>\" to write, \"cat\" to read, \"exit\" to quit\n")
	msg := ""
	for msg != "exit" {
		if msg == "" { 
			fmt.Printf("Type command:")
			l, _ := r.ReadString('\n')               //read command + input
			msg = strings.TrimRight(l, "\n")         // trim msg for newlines
		}
		if strings.HasPrefix(msg, "echo=") {  // if echo command is used
			msg = strings.TrimPrefix(msg, "echo=") //remove cammand phrase to contain msg only
			err := ioutil.WriteFile(filepath, []byte(msg), 0644) //write msg to driver 

			errCheck(err)
			msg = ""
		} else if strings.HasPrefix(msg, "cat") { //if cat command is used
			msg = ""
			br, err := ioutil.ReadFile(filepath)  //read echo message
			errCheck(err)
			str := string(br[:])                 //convert bytearray to string
			fmt.Printf("Reads: %s\n", str)

		} else if !strings.Contains(msg, "exit") && msg != "" {
			fmt.Printf("Wrong syntax. Use \"echo=\"<your message here>. or \"cat\"  \n")  //else: Wrong Syntax
		}
	}
}
//
func errCheck(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

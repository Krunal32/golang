package main
import (
"os"
"fmt"
"bufio")
func main(){
a,_:=os.Open("/dev/simp_read") //open file
r:= bufio.NewReader(a) // create a reader
b,_,_ := r.ReadLine() //read a line
fmt.Printf("This is what I read from the driver:  %s \n",string(b)) //print
}

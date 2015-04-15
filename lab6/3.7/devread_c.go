package main 
import (
"os" 
"fmt"
"log"
)
const filepath string ="/dev/simp_read" //path
func main(){

dfile,err:=os.Open(filepath)  //open file
errCheck(err)                // check for errors
defer dfile.Close()         // close at last
b:=make([]byte,1024)   
_,err1:=dfile.Read(b)     // read from file into byte array
str:=string(b[:])
errCheck(err1)
fmt.Printf("Message from kernel module: %s \n",str)      
}
func errCheck(err error ){
 
  if err !=nil {
   log.Fatal(err)
     
  }
}

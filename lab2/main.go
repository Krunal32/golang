// +build !solution

// Leave an empty line above this comment.
package main

import (
	"fmt"
	"os"
    "flag"
	"github.com/uis-dat320-fall2014/labs/lab2/config"
     
    
)
const txtpath,gobpath string ="config.txt","config.gob"

func main() {
    //flags fom commandline
   
    mnum_ptr:=flag.Int("Number",0,"Error. Check syntax")
    mnam_ptr:=flag.String("Name","anonymous","Error. Check syntax")
    flag.Parse()
    cfg := config.Configuration{*mnum_ptr, *mnam_ptr}
	
    fmt.Printf("From input or default: %v\n",cfg)
    //saves
    err:=cfg.Save()
    err=cfg.SaveGob()
    Errcheck(err)
    //test config.txt
    loadconfig,err1:=config.LoadConfig(txtpath)
    Errcheck(err1)
   //
    TestEqual("Load Config.txt",cfg,*loadconfig)
    //test config.gob
  
    loadGOBconfig,err2:=config.LoadGobConfig(gobpath)
   
    Errcheck(err2)

   TestEqual("Load Config.gob",cfg,*loadGOBconfig)
    

}
func Errcheck( err error){
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func TestEqual(str string,runTConf config.Configuration, loadConf config.Configuration) {

	b := loadConf.Name == runTConf.Name && loadConf.Number == runTConf.Number
    working:="NO"
    if b{ working="YES!"}      
	fmt.Println("Loaded "+str+" ",loadConf)
    fmt.Println("\nIs it working?:"+working)
    

}

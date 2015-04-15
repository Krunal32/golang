package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    
    yslice:=make([][] uint8,dy)
    for y:= range yslice {
    
        xslice:=make([] uint8,dx)
        for x:=range xslice {
            xslice[x]=(uint8)(x^y*x^y/2)
        }
        yslice[y]=xslice
    
    }
  
   return yslice

}

func main() {
    pic.Show(Pic)
}

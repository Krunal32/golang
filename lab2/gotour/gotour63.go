	package main
	 
	import (
	"io"
	"os"
	"strings"
	)
 
	type rot13Reader struct {
	r io.Reader
	}
 
	func (rot13 rot13Reader) Read(p []byte) (n int, err error) {
    n, err = rot13.r.Read(p)
    for i,ch := range p {
        switch {
        case ch > 'A' && ch <='Z':
            p[i] = (ch - 'A' + 13) % 26 + 'A'
        case ch > 'a' && ch < 'z':
            p[i] = (ch - 'a' + 13) % 26 + 'a'
        }
    }
    return
}
     

    
	func main() {
	s := strings.NewReader(
	"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	}

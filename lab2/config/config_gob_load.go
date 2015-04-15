// +build !solution

// Leave an empty line above this comment.

package config
import ("os"
        "encoding/gob")
func LoadGobConfig(file string) (conf *Configuration, err error) {
	gobfile, err := os.Open(file)
	decoder := gob.NewDecoder(gobfile)
	err = decoder.Decode(&conf)
	gobfile.Close()
	return
}

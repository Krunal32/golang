package config
import ("fmt")
func TestEqual(str string,con1 Configuration, con2 Configuration) {

	b := con1.Name == con2.Name && con1.Number == con2.Number
	fmt.Printf("%s is Working:  %t \n", str, b)
	fmt.Printf("loaded%s: %+v\n%18sVS\nstored config:  %+v\n", str, con1, " ", con2)

}

package main
import "fmt"
import "math"
type ErrNegativeSqrt float64
func main(){

fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))

}
func (e ErrNegativeSqrt) Error() string{
return fmt.Sprintf("cannot return negative number: %g",float64(e))
}

func Sqrt(f float64)( float64, error){

if f<0{return 0, ErrNegativeSqrt(f)}
z,z1:=float64(1),float64(0)
for  math.Abs(z-z1)>0.001 && z!=z1{
z1=z
z=z-(z*z-f)/2*z
}
return z, nil
}

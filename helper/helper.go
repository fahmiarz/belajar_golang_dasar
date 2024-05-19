package helper

var version = "1.0.0"      //tidak bisa diakses dari luar package
var Application = "golang" // bisa diakses

// tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return "Hello " + name
}
func SayHello(name string) string {
	return "Hello " + name

}

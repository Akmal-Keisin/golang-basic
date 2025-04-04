package helper

var version = "1.0.0" // unexported variable
var AppName = "MyApp" // exported variable

func sayGoodBye(name string) string {
	return "Goodbye " + name // unexported function
}

func SayHello(name string) string {
	return "Hello " + name
}

func AccessingSayGoodbye(name string) string {
	return sayGoodBye(name) // accessing unexported function within the same package
}

func AccessingVersion() string {
	return version // accessing unexported variable within the same package
}
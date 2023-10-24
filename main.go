package main

func main() {
	runInterface := NewCLI(NewPatientQueue())
	runInterface.Run()
}

package main

func main() {
	if err := SetupApplication().Run(); err != nil {
		panic(err)
	}
}

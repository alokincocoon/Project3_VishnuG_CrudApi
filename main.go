package main

func main() {
	a := App{}
	a.Initialize()
	a.Run(":8080")
	defer a.DBClient.Close()
}

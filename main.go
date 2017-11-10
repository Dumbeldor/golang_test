package main

// our main function
func main() {
	a := App{}
	/*
		a.Initialize(
			os.Getenv("APP_DB_USERNAME"),
			os.Getenv("APP_DB_PASSWORD"),
			os.Getenv("APP_DB_NAME"))
	*/
	a.Initialize("test", "", "test")
	a.Run(":8080")
}

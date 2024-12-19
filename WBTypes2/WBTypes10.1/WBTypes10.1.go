package main

import "fmt"

// Regular structure
type WebBrowserUser struct {
	name string
	age  int
}

func main() {
	// Anonymous structure for a session
	var session = struct {
		userName     string
		visitedSites []string
	}{
		userName:     "Andrew Argentina",
		visitedSites: []string{"example.com", "golang.org", "github.com"},
	}

	// Regular type for a web browser user
	var user = WebBrowserUser{
		name: "Vice Lead Nickolay",
		age:  177,
	}

	fmt.Println("Session User:", session.userName)
	fmt.Println("Visited Sites:", session.visitedSites)
	fmt.Println("Web Browser User:", user.name, user.age)
}

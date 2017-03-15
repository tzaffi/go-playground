package main

import (
	"fmt"
	"os"
	"syscall"
	
	"golang.org/x/crypto/ssh/terminal"
	"github.com/ktrysmt/go-bitbucket"
)

func main() {
	var username string
	fmt.Print("Bitbucket Username: ")
	fmt.Scanln(&username)
	
	fmt.Print("Bitbucket Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	/*
	if err == nil {
		fmt.Println("\nPassword typed: " + string(bytePassword))
	}
*/
	password := string(bytePassword)
	fmt.Print("Thanks [" + username + "] !!!!\n")
	os.Exit(0)
	
	c := bitbucket.NewBasicAuth(username, password)

	opt := &bitbucket.PullRequestsOptions{
		Owner:      "your-team",
		Repo_slug:  "awesome-project",
		Source_branch: "develop",
		Destination_branch: "master",
		Title: "fix bug. #9999",
		Close_source_branch: true,
	}
	res := c.Repositories.PullRequests.Create(opt)

	fmt.Println(res) // receive the data as json format
}

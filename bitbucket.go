package main

import (
	"encoding/json"
	"fmt"
	//	"os"
	"github.com/ktrysmt/go-bitbucket"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

func getMyRepos(client *bitbucket.Client, owner string, team string) interface{} {
	opt := &bitbucket.RepositoriesOptions{
		Owner: owner,
		Team:  team,
	}
	res := client.Repositories.ListForTeam(opt)
	return res

	//res := c.Repositories.ListForAccount(opt)
	//var result interface{}
	//return result
}

func printPretty(res *interface{}){
	resJson, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(resJson))
}
	
func main() {
	var username string
	fmt.Print("Bitbucket Email: ")
	fmt.Scanln(&username)

	fmt.Print("Bitbucket Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)
	fmt.Print("Thanks [" + username + "] !!!!\n")

	c := bitbucket.NewBasicAuth(username, password)

	//	os.Exit(0)
	res := getMyRepos(c, "edlabtc", "edlabtc")
	printPretty(&res)
}

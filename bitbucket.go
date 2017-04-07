package main

import (
	"encoding/json"
	"fmt"
	//	"os"
	"github.com/tzaffi/go-bitbucket"
	"golang.org/x/crypto/ssh/terminal"
	"reflect"
	"sort"
	"strconv"
	"syscall"
)

func getMyRepos(client *bitbucket.Client, owner string, team string, options ...string) interface{} {
	opt := &bitbucket.RepositoriesOptions{
		Owner: owner,
		Team:  team,
	}
	/*
	if options != nil {
		fmt.Println("something:")
	} else {
		fmt.Println("nada:")
	}
  */
	fmt.Printf("options = %v\tTtype = %T\n", options, options)
	getAllPages := options != nil && options[0] == "ALL_PAGES"
	fmt.Println("getting all pages ?", getAllPages)
	var pages []uint;
	if(!getAllPages) {
		pages = []uint{1}
	} else {
		pages = []uint{1, 11}
	}
	
	res := client.Repositories.ListForTeam(opt, pages...)

	return res

	//res := c.Repositories.ListForAccount(opt)
	//var result interface{}
	//return result
}


func getPretty(res *interface{}) string {
	resJson, _ := json.MarshalIndent(res, "", "  ")
	return string(resJson)
}

func printPretty(res *interface{}) {
	fmt.Println(getPretty(res))
}

func reflectionLength(res *interface{}) int {
	resVal := *res
	fmt.Printf("reflect.TypeOf(resVal) = %v\nreflect.TypeOf(resVal).Kind() = %v\n",
		reflect.TypeOf(resVal), reflect.TypeOf(resVal).Kind())
	switch reflect.TypeOf(resVal).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(resVal)
		return s.Len()
	default:
		return -1
	}
}

// cf. https://blog.golang.org/json-and-go#TOC_5.
func reflectionParse(res *interface{}) {
	resVal := *res
	switch t0 := resVal.(type) {
	case []interface{}:
		fmt.Println("array")
	case map[string]interface{}:
		fmt.Println("map")
	default:
		fmt.Printf("Surprise, surprise. Is %v\n", t0)
	}
}

//find all values that have the given key and a string value
func grepByKey(res *interface{}, key string) []string {
  var result []string
	resVal := *res
	switch t0 := resVal.(type) {
	case []interface{}:
		for _, v := range resVal.([]interface{}) {
			result = append(result, grepByKey(&v, key)...)
		}
	case map[string]interface{}:
		for k, v := range resVal.(map[string]interface{}) {
			if k == key && reflect.TypeOf(v).Kind() == reflect.String {
				result = append(result, v.(string))
			} else {
				result = append(result, grepByKey(&v, key)...)
			}
		}
	default:
		fmt.Printf("Surprise, surprise. %v is type %T\n", t0, t0)
		return result
	}
	return result
}

//aggregate the values arrays present in bitbucket JSON
func aggregateValues(res *interface{}) []interface{} {
	var result []interface{}
	respVal := *res
	switch t0 := respVal.(type) {
	case []interface{}:
		for _, v := range respVal.([]interface{}) {
			result = append(result, aggregateValues(&v)...)
		}
		return result
	case map[string]interface{}:
		respMap := respVal.(map[string]interface{})
		if val, ok := respMap["values"]; ok {
			return val.([]interface{})
			//TODO: NOT SAFE!!!!
		}
		return result
	default:
		fmt.Printf("Surprise, surprise. %v is type %T\n", t0, t0)
		return result
	}
}

//return the first object which has a key with given value
//otherwise, return an empty struct
func findByKeyVal(res *interface{}, key string, val string) map[string]interface{} {
	var result map[string]interface{}
	respVal := *res
	switch t0 := respVal.(type) {
	case []interface{}:
		respSlice := respVal.([]interface{})
		fmt.Printf("case []interface{} of length %d\n",len(respSlice))
		for _, v := range respSlice {
			subResult := findByKeyVal(&v, key, val)
			if len(subResult) > 0 {
				return subResult
			}
		}
		return result
	case map[string]interface{}:
		respMap := respVal.(map[string]interface{})
		fmt.Printf("case map[string]interface{} of length %d\n",len(respMap))
		if valTest, ok := respMap[key]; ok && val == valTest {
			return respMap
		}
		return result
	default:
		fmt.Printf("Surprise, surprise. %v is type %T\n", t0, t0)
		return result
	}
}

//Given slice return slice with all 
func filterByPredicate(res *[]interface{}, pred func(*interface{}) bool) []interface{} {
	var result []interface{}
	respSlice := *res
	for _, v := range respSlice {
		if pred(&v) {
			result = append(result, v)
		}
	}
	return result
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
	res := getMyRepos(c, "edlabtc", "edlabtc", "ALL_PAGES")
	fmt.Println("reflectionLength(&res) == ", reflectionLength(&res))	
	fmt.Println("len(getPretty(&res)) == ", len(getPretty(&res)))
	//reflectionParse(&res)
	repos := grepByKey(&res, "full_name")
	sort.Strings(repos)
	reposM, _ := json.MarshalIndent(repos, "", " ")
	fmt.Println("repos:", string(reposM))
	aggs := aggregateValues(&res)
	fmt.Printf("len(aggs) = %d\n", len(aggs))
	aggsM, _ := json.MarshalIndent(aggs, "", " ")
	fmt.Println("aggs:", string(aggsM))
	var aggsIntf interface{} = aggs
	repo := findByKeyVal(&aggsIntf, "full_name", "edlabtc/library-pocketknowledge")
	fmt.Printf("repo:\n%v\n", repo)

	isSmallerThan10Megs := func(repo *interface{}) bool{
		if val, ok := (*repo).(map[string] interface{} )["size"]; ok  {
			if f, err := strconv.ParseFloat(val); err == nil && f < 10000000 {
				return true
			}
		}
		return false
	}
	//printPretty(&res)	
}

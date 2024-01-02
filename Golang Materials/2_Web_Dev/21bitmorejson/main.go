package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"Course_Name"` // Alias name
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` // dont give space after, :-  tags, omit
}

func main() {
	fmt.Println("Json in Golang")
	// encoding the json
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "lco.dev", "abc123", []string{"web-dev", "JS"}},
		{"MERN Bootcamp", 199, "lco.dev", "bcd123", []string{"full-stack", "JS"}},
		{"Angular Bootcamp", 299, "lco.dev", "hit123", nil},
	}

	// package this data as json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	CheckError(err)
	fmt.Printf("Json : %s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"Course_Name": "ReactJs Bootcamp",
		"Price": 299,
		"website": "lco.dev",
		"tags": [ "web-dev", "JS" ]
	}`)

	var lcoCourse course

	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("JSON was valid!")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n",lcoCourse) // special format verb in string for json 
	} else {
		fmt.Println("Json in invalid!")
	}

	// some cases where we want to add data to key value pair 

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for k,v := range(myOnlineData){
		fmt.Printf("Key : %v and Value : %v and Type is %T\n",k,v,v)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

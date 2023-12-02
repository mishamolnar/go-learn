package encoding

import (
	"encoding/json"
	"fmt"
	"os"
)

type MyUser struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Fired    bool   `json:"fired,omitempty"`
	Position string
	Salary   int `json:"salary"`
}

func displayUsers() {
	users := make([]MyUser, 0)
	users = append(users, MyUser{"John", 123, false, "dev", 123123})
	data, err := json.Marshal(users[0])
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
	var use MyUser
	if err := json.Unmarshal(data, &use); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(use)
}

func JsonStart() {
	displayUsers()
}

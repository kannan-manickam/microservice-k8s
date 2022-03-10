package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func main() {
	http.HandleFunc("/friends", get_friends_json)
	fmt.Printf("API endpoint -> http://localhost:8080/friends")
	http.ListenAndServe(":8080", nil)
}

type Friend struct {
	Id       string `json:"id"`
	First      string `json:"first"`
	Last        string `json:"last"`
	Location  string `json:"location ,omitempty"`
}

var friends = map[string]Friend{
	"0000000000": Friend{Id: "1", First: "Karthikeyan", Last: "Varadharajan", Location: "Texas"},
	"0000000001": Friend{Id: "2", First: "Saravanakumar", Last: "Sugumaran", Location: "Texas"},
	"0000000002": Friend{Id: "3", First: "Sheik", Last: "Sajith", Location: "Pensylvania"},
	"0000000003": Friend{Id: "4", First: "Milton", Last: "Devasagayam", Location: "Pensylvania"},
	"0000000004": Friend{Id: "5", First: "Mandeep", Last: "Sah", Location: "Ohio"},
	"0000000005": Friend{Id: "6", First: "Karthikeyan", Last: "Duraisamy", Location: "Ohio"}
	"0000000003": Friend{Id: "7", First: "Ashak", Last: "Prabhu", Location: "Ohio"}
}

func get_friends() []Friend {
	values := make([]Friend, len(friends))
	i := 0
	for _, emp := range friends {
		values[i] = emp
		i++
	}
	return values
}

func get_friend_json(w http.ResponseWriter, r *http.Request) {
	frnds := get_friends()
	data, err := json.Marshal(frnds)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}

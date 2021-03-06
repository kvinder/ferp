package main

import (
	"encoding/json"
	"ferp/pkg/app"
	"ferp/pkg/model"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	port, databaseName, databaseURL, setUpDatabase := readFileConfig()
	model.InitConnection(databaseName, databaseURL)
	if setUpDatabase == "create" {
		model.CreateDatabase()
	}
	mux := http.NewServeMux()
	app.Router(mux)
	fmt.Println("localhost:" + port + " runing...")
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println(err)
	}
}

func readFileConfig() (portOut, databaseNameOut, databaseURLOut, setUpDatabaseOut string) {
	config, err := ioutil.ReadFile("./config.json")
	checkErr(err)
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(config, &objmap)
	checkErr(err)
	var port, databaseName, databaseURL, setUpDatabase string
	err = json.Unmarshal(*objmap["port"], &port)
	checkErr(err)
	err = json.Unmarshal(*objmap["databaseName"], &databaseName)
	checkErr(err)
	err = json.Unmarshal(*objmap["databaseURL"], &databaseURL)
	checkErr(err)
	err = json.Unmarshal(*objmap["setUpDatabase"], &setUpDatabase)
	checkErr(err)
	return port, databaseName, databaseURL, setUpDatabase
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

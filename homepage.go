package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	crypto "crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	PK       []byte `json:"publickey"`
}

func post(http.ResponseWriter, *http.Request, httprouter.Params) {
	fmt.Println("2. Performing Http Post...")
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	key := crypto.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	user := User{
		Username: "maria",
		PK:       key,
	}
	jsonReq, err := json.Marshal(user)
	resp, err := http.Post("http://localhost:4000/create-user", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	fmt.Printf("Status Code: %d\n", resp.StatusCode)

	// Convert response body to Todo struct
	//var todoStruct User
	//json.Unmarshal(bodyBytes, &todoStruct)
	//	fmt.Printf("%+v\n", todoStruct)
	return
}

func authenticate(pk rsa.PublicKey) {
	fmt.Println("2. Performing Authenticate...")
	key := crypto.MarshalPKCS1PublicKey(&pk)
	user := User{
		Username: "maria",
		PK:       key,
	}
	jsonReq, err := json.Marshal(user)
	resp, err := http.Post("http://localhost:4000/authenticate-user", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	fmt.Printf("Status Code: %d\n", resp.StatusCode)

	// Convert response body to Todo struct
	//var todoStruct User
	//json.Unmarshal(bodyBytes, &todoStruct)
	//	fmt.Printf("%+v\n", todoStruct)
}

func getPing(http.ResponseWriter, *http.Request, httprouter.Params) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:4000/ping", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	//req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	//defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	//var responseObject Response
	//json.Unmarshal(bodyBytes, &responseObject)
	//fmt.Printf("API Response as struct %+v\n", responseObject)
	s := string(bodyBytes)
	fmt.Printf("msg: %s", s)
}

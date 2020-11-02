package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

var baseURL = "http://localhost:8000"

type Student struct {
	ID string
	Name string
	NIM string
	Email string
	PhoneNumber string
	Class string
	Major string
	Faculty string
	University string
}

func fetchUsers() ([]Student, error) {
	var err error
	var client = &http.Client{}
	var data []Student

	request, err := http.NewRequest("GET", baseURL + "/students", nil)
	if err != nil {
		return nil,err
	}

	response,err := client.Do(request)
	if err != nil {
		return nil,err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil,err
	}
	return data,nil
}


func fetchUser(ID string) (Student, error) {
	var err error
	var client = &http.Client{}
	var data Student

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/student", payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}


func getUsers() {
	users, err := fetchUsers()
	if err != nil {
		fmt.Println("ERROR!!!", err.Error())
		return
	}

	for _, each := range users {
		fmt.Println("" +
			" ID : " + each.ID + ", "+
			" Name : " + each.Name + ", "+
			" NIM : " + each.NIM + ", "+
			" Email : " + each.Email + ", "+
			" PhoneNumber : " + each.PhoneNumber + ", "+
			" Major : " + each.Major + ", "+
			" Faculty : " + each.Faculty + ", "+
			" University : " + each.University)
	}
	start()
}

func getUser(id string) {
	user1, err := fetchUser(id)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	fmt.Println(
		"{ \n" +
			"    ID : " + user1.ID + "\n"+
			"    Name : " + user1.Name + "\n"+
			"    NIM : " + user1.NIM + "\n"+
			"    Email : " + user1.Email + "\n"+
			"    PhoneNumber : " + user1.PhoneNumber + "\n"+
			"    Major : " + user1.Major + "\n"+
			"    Faculty : " + user1.Faculty + "\n"+
			"    University : " + user1.University + "\n}")

	start()
}

func getUserWithManualID()  {
	var input string
	fmt.Print("Masukkan ID : ")
	fmt.Scan(&input)
	getUser(input)
}

func menu() int {
	var input int

	fmt.Println("1. GetAll")
	fmt.Println("2. GetWithID")
	fmt.Println("3. exit")

	fmt.Print("Pilihan anda : ")
	fmt.Scanln(&input)
	return input
}

func start()  {
	a := menu()
	status := true

	if status {
		if a == 1 {
			getUsers()
		}else if a == 2 {
			getUserWithManualID()
		}else if a == 3	 {
			os.Exit(0)
		}else {
			start()
		}

		status = isNext()
	}
}

func main() {
	start()
}

func isNext() bool {
	var input string

	fmt.Print("Lanjut (yes/no) : ")
	fmt.Scanln(&input)
	if input != "yes" {
		return false
	}
	return true
}
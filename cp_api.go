package main

import (
	api "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"fmt"
	"os"
)

func ShowHosts() {

	// var apiServer string
	// var username string
	// var password string

	// fmt.Printf("Enter server IP address or hostname: ")
	// fmt.Scanln(&apiServer)

	// fmt.Printf("Enter username: ")
	// fmt.Scanln(&username)

	// fmt.Printf("Enter password: ")
	// fmt.Scanln(&password)

	apiServer := "192.168.168.140"
	domain := "Cti"
	apiKey := os.Getenv("CHECKPOINT_API_KEY")

	args := api.APIClientArgs(api.DefaultPort, "", "", apiServer, "", -1, "", true, true, "deb.txt", api.WebContext, api.TimeOut, api.SleepTime, "", "", -1)

	client := api.APIClient(args)


	if x, _ := client.CheckFingerprint(); !x {
		print("Could not get the server's fingerprint - Check connectivity with the server.\n")
		os.Exit(1)
	}

	// loginRes, err := client.Login(username, password, false, "", false, "")
	loginRes, err := client.ApiLoginWithApiKey(apiKey, false, domain, false, map[string]interface{}{})
	if err != nil {
		print("Login error.\n")
		os.Exit(1)
	}

	if !loginRes.Success {
		print("Login failed:\n" + loginRes.ErrorMsg)
		os.Exit(1)
	}

	showHostsRes,err2 := client.ApiQuery("show-hosts", "full", "objects", false, map[string]interface{}{})

	if err2 != nil {
		print("Failed to retrieve the hosts\n")
		return
	}

	//fmt.Println(show_sessions_res.GetData())
	for _,sessionObj := range showHostsRes.GetData(){
		fmt.Println("-------------")
		fmt.Println(sessionObj.(map[string]interface{})["name"].(string))
		fmt.Println(sessionObj.(map[string]interface{})["ipv4-address"].(string))
	}

}

func main() {
	fmt.Println("Testing Check Point API")
	ShowHosts()
}

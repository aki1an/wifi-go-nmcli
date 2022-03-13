package main

import (
	"fmt"
	"strings"
	"os/exec"
)
/**
author : akil4n
**/
func main() {
	out, err := exec.Command("nmcli", "device", "wifi", "list").Output()
	output := strings.Replace(string(out), "*", "", -1)
	available_wifi := strings.Trim(output,"\n");

	devices_details := strings.Split(available_wifi, "\n");
	device_list := make([][]string, len(devices_details))
	
	for i:= 0; i<len(devices_details); i++ {
		inter := strings.Fields(devices_details[i])
		device_list = append(device_list, inter);
	}
	if(err != nil){
		fmt.Println("ERROR OCCURRED MAKE SURE nmcli IS INSTALLED")
	}
	fmt.Println("\n")
	fmt.Println("---|---------------------|-----------------------")
	fmt.Println("id | BSSID               | SSID")
	fmt.Println("---|---------------------|-----------------------")
	for i:=1;i<len(device_list);i++ {
		if len(device_list[i]) > 1 && len(device_list[i][0]) > 6 {
			fmt.Println(i, " | ",device_list[i][0], " | ", device_list[i][1], device_list[i][2])
		}
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("\n")
	fmt.Println("Enter the id : ")
	// reader := bufio.NewReader(os.Stdin)
	var id int
	fmt.Scanf("%d", &id)
	fmt.Println("Password : ")
	var Password string
	fmt.Scanf("%s", &Password)
	result, err := exec.Command("nmcli", "device", "wifi", "connect", device_list[id][0], "password", Password).Output()
	for (!strings.Contains(string(result), "successfully activated")) {
		fmt.Println("wrong password..\nretry : ")
		fmt.Scanf("%s", &Password)
		result, err = exec.Command("nmcli", "device", "wifi", "connect", device_list[id][0], "password", Password).Output()
	}
	fmt.Println("connected to wifi successfully")
}

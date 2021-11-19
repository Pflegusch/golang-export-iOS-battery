package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/melbahja/goph"
)

func logData(data string) {
	f, err := os.OpenFile("../../data.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	f.WriteString(data)
}

func getData(ip string, sshkey string) string {
	// Start new ssh connection with private key.
	auth, err := goph.Key(sshkey, "")
	if err != nil {
		log.Fatal(err)
	}

	client, err := goph.New("root", ip, auth)
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing the network connection.
	defer client.Close()

	// Execute your command.
	out, err := client.Run("bstat")

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func main() {
	var ip, sshkey string

	if len(os.Args) != 3 {
		fmt.Println("Usage: ./binary <ipAddr> <path_to_sshkey>")
		return
	} else {
		ip = os.Args[1]
		sshkey = os.Args[2]
	}

	for {
		fmt.Println(getData(ip, sshkey))
		logData(getData(ip, sshkey))
		time.Sleep(5 * time.Second)
	}
}

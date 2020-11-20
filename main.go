package main
 
import (
	"io/ioutil"
	"net"
	"net/http"
	"fmt"
	"log"
	"errors"
	"time"
)

func getIpAddr() (string, error) {
	url := "https://api.ipify.org?format=text"
    resp, err := http.Get(url)
    if err != nil {
		return "", errors.New("Could not get response from ipify")
    }
    defer resp.Body.Close()
    ip, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		return "", errors.New("Could not read response from ipify")
	}
	return string(ip), nil;
}

func sendMsg(msg string) error {
	return nil
}

func handleConnection(conn net.Conn) {
	log.Println("Connected to", conn.RemoteAddr().String())
	i := 0
	for {
		i += 1
		fmt.Printf("Sending hello world%d...\n", i)
		_, err := fmt.Fprintf(conn, "hello world %d\n", i)
		if (err != nil) {
			log.Println(err)
		}
		time.Sleep(3 * time.Second)
	}
}
 
func main() {
	log.Println("Getting IP address from ipify")
	ipAddr, err := getIpAddr()
    if (err != nil) { 
		log.Println("Could not get IP adress:", err)
	} else {
		log.Println("IP address is:", ipAddr)
	}

	port := "8080"
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Printf("Could not listen on port %s using tcp: %s\n", port, err)
	} else {
		log.Printf("Listening on port %s\n", port)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Failed to accept connection")
		}
		go handleConnection(conn)
	}
}
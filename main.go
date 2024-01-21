package main

import (
	"bufio"
	"net"
	"os"
	"fmt"
	"log"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMx, hasSPF, sprRecord, hasDMARC, dmarcRecord\n" )

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string){

}
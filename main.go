package main

import (
	"crypto/x509"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("export.cer")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	data := make([]byte, 2)
	var s string
	for{
		n, err := file.Read(data)
		if err == io.EOF{
			break
		}
		//fmt.Print(string(data[:n]))
		s = s+string(data[:n])
	}

	r := strings.NewReader(s)
	dataFile:=make([]byte,len(s))
	for{
		_, err := r.Read(dataFile)
		if err == io.EOF {
			break
		}
	}

	fmt.Println(dataFile)

	cert, err := x509.ParseCertificate(dataFile)
	fmt.Println(cert)
	if err != nil{
		fmt.Println(err)
	}
}
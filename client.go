package main
import (
	"fmt"
	"os"
	"net"
	"io"
	"strings"
)
func main() {
	file, err := os.Open("DOCX_.docx.sig")
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

	httpRequest:=dataFile
	conn, err := net.Dial("tcp", "golang.org:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write(httpRequest); err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("Done")
}
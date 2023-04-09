package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func hello(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	buf, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(buf))

	remove()

	cmd := exec.Command("hugo", "--cleanDestinationDir")
	cmd.Dir = "./prebuild"
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func remove() {
	cmd := exec.Command("rm", "-rf", "./prebuild/public")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	go func() {
		log.Println("Starting...")
		http.HandleFunc("/rebuild", hello)
		err := http.ListenAndServe(":5001", nil)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()

	cmd := exec.Command("hugo", "server", "--quiet", "--disableFastRender")
	stdOut, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdOut)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}

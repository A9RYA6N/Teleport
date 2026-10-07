package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
    "strconv"
)

func main(){
    fmt.Println("Welcome to teleport")
    file, err:=os.Open("/proc/net/tcp")
    defer file.Close()
    if err!=nil{
        log.Fatalf("Failed to open file: %s", err)
    }
    scanner:=bufio.NewScanner(file)
    headerFlag:=true
    for scanner.Scan(){
	if headerFlag{
	    headerFlag=false
	    continue
	}
	line:=scanner.Text()
	words:=strings.Fields(line)
	if words[3]!="0A"{
	    continue
	}
	hexIp:=strings.Split(words[1], ":")[0]
	decIp, err:=strconv.ParseInt(hexIp, 16, 64)
	if err!=nil{
	    log.Fatalf("Error parsing hex ip: %s", err)
	}
	fmt.Printf("%d ", decIp)
	fmt.Print(words[3]+" ")
	fmt.Println(words[11])
    }
    if err:=scanner.Err(); err!=nil{
	log.Fatalf("Error scanning file: %s", err)
    }
}

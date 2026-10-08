package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
    "encoding/hex"
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
	hexArr:=strings.Split(words[1], ":")
	hexIp:=hexArr[0]
	hexPort:=hexArr[1]
	bytes, err:=hex.DecodeString(hexIp)
	if err!=nil{
	    log.Fatalf("Error parsing hex ip: %s", err)
	}
	realIp:=fmt.Sprintf("%d.%d.%d.%d", bytes[3], bytes[2], bytes[1], bytes[0])
	port, err:=strconv.ParseInt(hexPort, 16, 64)
	if err!=nil{
	    log.Fatalf("Error parsing hex port: %s", err)
	}
	fmt.Printf("%s:%d, %s:%s ",realIp, port, hexIp, hexPort)
	fmt.Print(words[3]+" ")
	fmt.Println(words[9])
    }
    if err:=scanner.Err(); err!=nil{
	log.Fatalf("Error scanning file: %s", err)
    }
}

package main

import (
    "log"
    "os"
    "backend/ssh_server"
)

func main(){
    port := os.Getenv("SSH_PORT")
    if port = ""{
        port = "2222"
    }
    
    err := ssh_server.StartSSHServer("0.0.0.0" + port)
    if err!=nil{
        log.Fatalf("Failed to start SSH server %v", err)
    }

}


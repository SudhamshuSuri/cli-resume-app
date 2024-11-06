package handlers

import (
    "os"
    "io"
    "log"
    "golang.org/x/crypto/ssh"
)

func HandleSession(channel ssh.Channel, reqeusts <-chan *ssh.Request) {
    defer channel.Close()

    for req := range requests{
        switch req.Type {
        case "exec":
            if req.WantReply {
                req.Reply(true, nil)
            }
            command := string(req.Payload[4:])
            if command == "download_resume" {
                sendResume(channel)
            }
        }
    }
}


func sendResume(channel ssh.Channel) {
    file, err := os.Open("resume.pdf")
    if err!= nil {
        log.Printf("Could not open resume.pdf %v", err)
        return
    }

    defer file.Close()

    _, err = io.Copy(channel, file)
    if err!=nil {
        log.Printf("Failed to send file %v", err)
    }
}

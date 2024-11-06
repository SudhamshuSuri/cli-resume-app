package ssh_server

import (
    "log"
    "os"
    "golang.org/x/crypto/ssh"
    "net"
    "backend/handlers"
)


func StartSSHServer( addr string ) error {
    config := &ssh.ServerConfig{
        NoClientAuth: true,
    }

    privateBytes, err := os.ReadFile("ssh_host_key")
    if err != nil{
        return err
    }

    private, err := ssh.ParsePrivateKey(privateBytes)
    if err ! nil{
        return err
    }

    config.AddHostKey(private)

    listner, err := net.Listen("tcp", addr)
    if err!= nil{
        return err
    }

    log.Printf("SSH server listening on %s", addr)

    for{
        conn, err := listener.Accept()
        if err!= nil{
            log.Printf("Failed to accept incoming connection %v", err)
            continue
        }

        go func() {
            sshConn, chans, _, err := ssh.NewServerConn(conn, config)
            if err!=nil {
                log.Printf("Failed to handshake %v", err)
                return
            }
            defer sshConn.close()

            for newChannel := range chans{
                if newChannel.ChannelType() == "session" {
                    channel, requests, err := newChannel.Accept()
                    if err!= nil {
                        log.Printf("Could not accept channel :%v", err)
                        return
                    }
                    go handlers.HandleSession(channel, requests)
                } else {
                    newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
                }
            }


        }()
    }
}

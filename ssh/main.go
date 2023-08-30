package main

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"time"
)

func main() {
	err := handle()

	if err != nil {
		panic(err)
	}

	os.Exit(0)
}

func handle() error {
	ip := os.Getenv("SERVER_IP")

	if ip == "" {
		return errors.New("服务器IP不能为空")
	}

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		return errors.New("服务器端口不能为空")
	}

	addr := ip + ":" + port

	username := os.Getenv("SERVER_USERNAME")

	if username == "" {
		return errors.New("服务器用户名不能为空")
	}

	password := os.Getenv("SERVER_PASSWORD")

	if password == "" {
		return errors.New("服务器密码不能为空")
	}

	command := os.Getenv("SERVER_COMMAND")

	if command == "" {
		return errors.New("服务器命令不能为空")
	}

	log.Printf("服务器地址: %s", addr)
	log.Printf("服务器用户名: %s", username)

	client, err := ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 10,
	})

	if err != nil {
		return err
	}

	session, err := client.NewSession()

	if err != nil {
		return err
	}

	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	return session.Run(command)
}

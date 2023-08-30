package main

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"os"
)

func main() {
	err := handle()

	if err != nil {
		panic(err)
	}
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

	client, err := ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
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

	err = session.Run(command)

	if err != nil {
		return err
	}

	return session.Wait()
}

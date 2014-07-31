package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	tag := "main"

	if len(os.Args) >= 2 {
		tag = os.Args[1]
	}

	log.SetFlags(0)
	log.SetPrefix(tag + ": ")

	status, err := os.Open("/proc/self/status")
	if err != nil {
		log.Fatal(err)
	}
	defer status.Close()

	r := bufio.NewReader(status)

	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Print(line)
	}

	if tag != "main" {
		return
	}

	os.Stdout.Write([]byte{'\n'})

	{
		cmd := exec.Command(os.Args[0], "fork")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			log.Println("failed to exec command: ", err)
		} else {
			cmd.Wait()
		}

		os.Stdout.Write([]byte{'\n'})
	}

	{
		cmd := exec.Command(os.Args[0], "NEWNS")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWNS,
		}

		if err := cmd.Start(); err != nil {
			log.Println("failed to exec command:", err)
		} else {
			cmd.Wait()
		}

		os.Stdout.Write([]byte{'\n'})
	}

	{
		cmd := exec.Command(os.Args[0], "NEWPID")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWPID,
		}

		if err := cmd.Start(); err != nil {
			log.Println("failed to exec command:", err)
		} else {
			cmd.Wait()
		}

		os.Stdout.Write([]byte{'\n'})
	}
}

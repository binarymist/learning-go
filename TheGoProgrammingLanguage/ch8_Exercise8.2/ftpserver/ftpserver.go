package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// FTP has two channels:
// command: transmitting commands as well as replies to those commands
// data: transferring data

// cd: change directory
// ls: list a directory
// get: send contents of a file
// close: close the connection

const (
	StatusSyntaxErrorNotRecognised = 500            // RFC 959, 4.2.1
	DateFormatStatTime             = "Jan _2 15:04" // LIST date formatting with hour and minute
)

var port int

type FTPConnection struct {
	command net.Conn
	data    net.Conn
	pasv    bool
	root    string
	cwd     string
}

func fileStat(file os.FileInfo) string {
	return fmt.Sprintf(
		"%s 1 ftp ftp %12d %s %s",
		file.Mode(),
		file.Size(),
		file.ModTime().Format(DateFormatStatTime),
		file.Name(),
	)
}

func list(ftpConnection *FTPConnection, msgs []string) {
	cwd := ftpConnection.cwd
	if len(msgs) > 0 {
		cwd = path.Join(cwd, msgs[0]) // Todo: Can we use filepath.Join?
	}

	defer ftpConnection.data.Close()
	// files is []os.FileInfo
	if files, err := ioutil.ReadDir(filepath.Join(cwd, msgs[0])); err != nil {
		reply(ftpConnection, fmt.Sprintf("%d %s", StatusSyntaxErrorNotRecognised, fmt.Sprintf("Could not list: %v", err)))
	} else {
		if ftpConnection.pasv {
			reply(ftpConnection, "150 File status okay. About to open data connection.")
		} else {
			reply(ftpConnection, "125 Data connection already open. Transfer starting.")
		}

		for _, file := range files {
			fmt.Fprintf(ftpConnection.command, "%s\r\n", fileStat(file))
		}

		reply(ftpConnection, "226 Transfer complete.")
	}
}

func reply(ftpConnection *FTPConnection, msg string) {
	fmt.Fprintf(ftpConnection.command, msg+"\r\n")
}

// Fix this function to make it similar to handlePASV in the ftpserver sample project.
func pasv(ftpConnection *FTPConnection, command string) {
	if addr, ok := ftpConnection.command.LocalAddr().(*net.TCPAddr); ok {
		ip := addr.IP.String()
		for i := 40000; i < 40050; i++ {
			ln, err := net.Listen("tcp", ip+":"+strconv.Itoa(i))
			if err == nil {
				if command == "pasv" {
					ip = strings.Join(strings.Split(ip, "."), ",")
					h := strconv.Itoa(i >> 8)
					l := strconv.Itoa(i % 256)
					msg := "227 Entering passive mode (" + ip + "," + h + "," + l + ")"
					log.Print(msg + "; port " + strconv.Itoa(i))
					reply(ftpConnection, msg)
				} else {
					// I think we call reply here? ....................
				}
				go (func() {
					data, err := ln.Accept()
					if err == nil {
						ftpConnection.data = data
						ftpConnection.pasv = true
					}
				})()
			}
		}
	}
}

func handle(command string, msgs []string, ftpConnection *FTPConnection) {
	switch command {
	case "user":
		reply(ftpConnection, "331 Username ok, send password.")
	case "pass":
		reply(ftpConnection, "230 Login successful.")
	case "syst":
		reply(ftpConnection, "215 UNIX Type: L8.")
	case "feat":
		reply(ftpConnection, "211-Features:\r\n  FEAT\r\n  MDTM\r\n  PASV\r\n  SIZE\r\n  TYPE A;I\r\n211 End")
	// case "port":
	// 	port(msgs, c)
	case "list":
		list(ftpConnection, msgs)
	case "pwd":
		reply(ftpConnection, "257 \"/\" is the current directory.")
	case "type":
		reply(ftpConnection, "200 Type set to binary.")
	case "cwd":
		dir := msgs[0]
		ftpConnection.cwd = path.Join(ftpConnection.root, dir)
		reply(ftpConnection, "250 Directory changed to "+dir+".")
	case "pasv", "epsv":
		pasv(ftpConnection, command)
	// case "retr":
	// 	retr(c, msgs)
	case "quit":
		reply(ftpConnection, "221 bye")
		// Close ftpConnection.command in handleConnection.
	}
}

// func handleConnection(ftpConnection FTPConnection) {
// 	defer ftpConnection.command.Close()
// 	// What about ftpConnection.data.Close()?
// 	// Send connection success message to FTP client:

// }

func handleConnection(ftpConnection *FTPConnection) {
	defer ftpConnection.command.Close()
	reply(ftpConnection, "220 Kim's FTP server ready.")
	reader := bufio.NewReader(ftpConnection.command)
	for {
		clientLine, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
			return
		}
		log.Print("Reqeust: ", clientLine)
		command := strings.TrimSpace(strings.ToLower(clientLine[0:4]))
		msgs := strings.Split(strings.Trim(clientLine, "\r\n "), " ")[1:]
		handle(command, msgs, ftpConnection)
	}
}

func main() {
	flag.IntVar(&port, "port", 2120, "The port to listen on ('2121')")
	flag.Parse()
	netListener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		if port < 1024 {
			log.Printf("Port: %d is a privileged port. Elevated permissions might be required.", port)
		}
		log.Fatal(err)
	}
	defer netListener.Close()
	wd, _ := os.Getwd()
	fmt.Println("Listening on port:", netListener.Addr().(*net.TCPAddr).Port)
	for {
		conn, err := netListener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Connection from %v established.\n", conn.RemoteAddr())
		ftpConnection := FTPConnection{command: conn, pasv: false, cwd: wd, root: wd}

		go handleConnection(&ftpConnection) // Why pass the address?
	}
}

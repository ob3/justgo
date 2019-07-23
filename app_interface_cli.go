package justgo

import (
	"fmt"
	"net"
	"runtime/debug"
	"strings"
)

type CliCommand func(string) string

type CliInterface struct {
	listener net.Listener
	Address  string
	commands map[string]CliCommand
}

func (cli *CliInterface) AddCommand(command string, function CliCommand) {
	if cli.commands == nil {
		cli.commands = map[string]CliCommand{}
	}
	cli.commands[command] = function
}

func (cli *CliInterface) Serve() {
	log.Info("serving CLI: ", cli.Address)
	var err error
	cli.listener, err = net.Listen("tcp", cli.Address)
	if err == nil {
		listen := make(chan net.Conn, 100)
		go acceptConnection(cli.listener, listen)
		for {
			conn := <-listen
			go cli.handleClient(conn)
		}
	} else {
		log.Error(err)
	}
}

func (cli *CliInterface) handleClient(client net.Conn) {
	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}

		command, params := getCommandAndParams(string(buf))

		returnValue := ""
		if len(command) == 0 {
			returnValue = ""
		} else {
			function := cli.commands[command]
			if function == nil {
				availableCommands := ""
				for key, _ := range cli.commands {
					availableCommands += key + " "
				}

				returnValue = fmt.Sprintf("command %s not found. available commands: %s", command, availableCommands)
			} else {
				execAndReturn(function, params, client)
			}

		}

		returnValue += "\r\n"
		client.Write([]byte(returnValue))
	}

}

func execAndReturn(function CliCommand, param string, conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			stack := debug.Stack()
			log.Error(string(stack))
			conn.Write(stack)
		}
	}()
	returnValue := function(param)
	returnValue += "\r\n"
	conn.Write([]byte(returnValue))
}


func getCommandAndParams(input string) (string, string) {
	cleanedInput := cleanInput(input)
	fields := strings.Fields(cleanedInput)
	if len(fields) == 0 {
		return "", ""
	}

	return fields[0], strings.Join(fields[1:], " ")
}

func cleanInput(input string) string {
	split := strings.Split(input, "\r\n")
	return split[0]
}

func (cli *CliInterface) ShutDown() {
	log.Info("shutting down CLI")
	cli.listener.Close()
}

func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		listen <- conn
	}
}

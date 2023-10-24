package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EXIT       = "EXIT"
	ROUNDROBIN = "ROUNDROBIN"
	DEFAULT    = "DEFAULT"
	OUT        = "OUT"
	IN         = "IN"
	MODE       = "MODE"
)

type RunInterface interface {
	Run()
}

type cli struct {
	queue Queue
}

func NewCLI(q Queue) RunInterface {
	return &cli{
		queue: q,
	}
}

func (c *cli) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		// Remove newline character from the input
		trimmedInput := strings.Split(strings.TrimSpace(input), " ")
		command := trimmedInput[0]

		switch command {
		case EXIT:
			fmt.Println("Exiting the program. Goodbye!")
			return
		case OUT:
			item, err := c.queue.Dequeue()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("Send: %s %s\n", item.Number, item.GetGender())
		case IN:
			if len(trimmedInput) != 3 {
				fmt.Println("Input format incorrect, try again")
				continue
			}

			number := trimmedInput[1]
			gender := trimmedInput[2]
			if gender != "F" && gender != "M" {
				fmt.Println("Gender format incorrect, try again")
				continue
			}

			err := c.queue.Enqueue(NewPatient(number, gender))
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println("Success")
		case ROUNDROBIN:
			c.queue.ChangeMode(RoundRobin)
			fmt.Println("Changing mode to Round Robin")
		case DEFAULT:
			c.queue.ChangeMode(Default)
			fmt.Println("Changing mode to Default")
		case MODE:
			fmt.Println(c.queue.GetMode())
		default:
			fmt.Println("Unknown command. Please try again.")
		}

	}

}

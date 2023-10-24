# Hospital Queue program in Go

## Description
Simple Go program in CLI to receive patient data in queue.

## How to run
```bash
go run .
```

## Program Overview
The program in run fully in CLI by inputting various commands. List of commands:
- IN **NUMBER** **GENDER**: Input the number and gender. Only accepts "F" and "M" gender.
- OUT: Print and remove the top item in Queue FIFO style. In Round Robin mode the program will output item alternately between gender, if no alternate gender available, program will return error. If the mode is switched back to Default, then the current gender information will be resetted.
- MODE: Print current mode.
- ROUNDROBIN: Change mode to Round Robin mode.
- DEFAULT: Change mode to Default mode.
- EXIT: Exit the program.

## Queue Implementation
The queue is implemented by using Go's slice data structure using list comprehension to manage the current available items.

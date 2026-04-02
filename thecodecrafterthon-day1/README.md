# Terminal Calculator — thecodecrafterthon Day 1

**Author: Ihwe Mathias**

A terminal calculator I built in Go that runs in a loop, takes math expressions as input, and returns the result. It handles bad input cleanly and never crashes — no matter what you type.

Standard library only. No external packages needed.


## How to Run

### Run directly
```bash
cd thecodecrafterthon-day1
go run main.go
```

### Or build a binary first
```bash
go build -o calc main.go
./calc
```


## Example

✦ Terminal Calculator ✦
Type 'help' for instructions.
> add 10 4
✦ Result: 14
> mul 3 7
✦ Result: 21
> sub 100 37
✦ Result: 63
> div 9 3
✦ Result: 3
> add -10 -5
✦ Result: -15
> quit
Goodbye! ✦
```


## Supported Commands

| Command       | What it does        |
|---------------|---------------------|
| `add <a> <b>` | adds a and b        |
| `sub <a> <b>` | subtracts b from a  |
| `mul <a> <b>` | multiplies a by b   |
| `div <a> <b>` | divides a by b      |
| `help`        | shows all commands  |
| `quit`        | exits the program   |

Negative numbers work fine — just type them normally like `add -10 -5`.


## Error Handling

Nothing you type will crash the program. Every bad input gets a clear message.

| What you type    | What happens                                        |
|------------------|-----------------------------------------------------|
| `div 9 0`        | Error: division by zero is not allowed.             |
| `add cat dog`    | Error: arguments must be valid integers.            |
| `add 5`          | Error: 'add' requires exactly 2 arguments.          |
| `mul 1 2 3`      | Error: 'mul' requires exactly 2 arguments.          |
| `blorp`          | Unknown command: 'blorp'. Type 'help' for usage.    |
| *(empty line)*   | nothing happens, the prompt just comes back         |


## How I Built It

I used `bufio.NewReader` to read input line by line from the terminal. Each line gets trimmed and split into parts using `strings.Fields` — the first part is the command, the rest are the arguments.

I used a `switch` statement to handle each command. For the math operations I first check that exactly 2 arguments were given, then I parse them with `strconv.Atoi` — if either one is not a valid integer, I print an error and wait for the next input. Division by zero is checked separately before doing the calculation.

The whole thing runs inside a `for` loop so it keeps going until the user types `quit`.


## Project Structure

thecodecrafterthon-day1/
├── main.go      — all the logic
└── README.md    — this file
```
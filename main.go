package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start() position {
	board, _ := fen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBKQBNR")
	return position{
		board: board,
	}
}

func cli() {
	pos := start()
	searcher := &searcher{tp: map[position]entry{}}
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(pos.board)
		valid := false
		for !valid {
			fmt.Print("Enter move: ")
			input, _ := r.ReadString('\n')
			input = strings.TrimSpace(input)
			valid = false
			for _, m := range pos.moves() {
				if input == m.String() {
					pos = pos.move(m)
					valid = true
					break
				}
			}
		}
		fmt.Println(pos.rotate().board)
		m := searcher.search(pos, 10000)
		score := pos.value(m)
		if score <= -MateValue {
			fmt.Println("You won")
			return
		}
		if score >= MateValue {
			fmt.Println("You lost")
			return
		}
		pos = pos.move(m)
	}
}

func uci() {
	pos := start()
	searcher := &searcher{tp: map[position]entry{}}
	r := bufio.NewReader(os.Stdin)
	sqr := map[string]square{}
	for i := square(0); i < 120; i++ {
		sqr[i.String()] = i
	}
	white := true
	for {
		input, _ := r.ReadString('\n')
		input = strings.TrimSpace(input)
		switch {
		case input == "quit":
			return
		case input == "isready":
			fmt.Println("readyok")
		case input == "uci":
			fmt.Println("id name carnatus")
			fmt.Println("id author zserge")
			fmt.Println("uciok")
		case input == "ucinewgame" || input == "position startpos":
			pos = start()
			white = true
		case strings.HasPrefix(input, "position startpos moves "):
			pos = start()
			moves := strings.Split(input[24:], " ")
			for i, s := range moves {
				m := move{from: sqr[s[0:2]], to: sqr[s[2:4]]}
				if i%2 != 0 {
					m = move{from: 119 - m.from, to: 119 - m.to}
				}
				pos = pos.move(m)
			}
			white = len(moves)%2 == 0
		case strings.HasPrefix(input, "position fen "):
			b, _ := fen(input[13:])
			fmt.Println(b)
			pos = position{board: b}
		case strings.HasPrefix(input, "go"):
			m := searcher.search(pos, 10000)
			if !white {
				m = move{from: 119 - m.from, to: 119 - m.to}
			}
			fmt.Println("bestmove", m)
		}
	}
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "cli" {
		cli()
	} else {
		uci()
	}
}

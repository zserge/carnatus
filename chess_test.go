package main

import (
	"sort"
	"strings"
	"testing"
)

func TestAbs(t *testing.T) {
	if x := abs(3); x != 3 {
		t.Error(x)
	}
	if x := abs(-4); x != 4 {
		t.Error(x)
	}
	if x := abs(0); x != 0 {
		t.Error(x)
	}
	if x := abs(1e10); x != 1e10 {
		t.Error(x)
	}
}

func TestFEN(t *testing.T) {
	if b, err := fen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"); err != nil {
		t.Error(err)
	} else if b.String() != "\nrnbqkbnr\npppppppp\n........\n........\n........\n........\nPPPPPPPP\nRNBQKBNR\n" {
		t.Error(b.String())
	}
	if b, err := fen("7K/3P4/8/8/8/8/1p6/k7"); err != nil {
		t.Error(err)
	} else if b[28] != 'K' || b[34] != 'P' || b[82] != 'p' || b[91] != 'k' {
		t.Error(b)
	} else if b[21] != '.' || b[20] != ' ' {
		t.Error(b)
	}
	for _, s := range []string{
		"",
		"hello",
		"8/8/8/8/8/8/8/8/8",
		"8/8/8/8/8/8/8/9",
		"8/8/8/8/8/8/8",
		"8/1p1/8/8/8/8/8/8",
		"8/1x7/8/8/8/8/8/8",
		"8/1 7/8/8/8/8/8/8",
		"8/1.7/8/8/8/8/8/8",
	} {
		if b, err := fen(s); err == nil {
			t.Error(s, "should return an error, but got:", b)
		}
	}
}

func TestBoardSwap(t *testing.T) {
	b, _ := fen("1k6/2p5/8/8/8/8/8/K7")
	if b.Swap().Swap().String() != b.String() {
		t.Error(b, b.Swap().Swap())
	}
	if b[22] != 'k' || b[33] != 'p' || b[91] != 'K' {
		t.Error(b)
	}
	b = b.Swap()
	if b[28] != 'k' || b[86] != 'P' || b[97] != 'K' {
		t.Error(b)
	}
}

func TestSquare(t *testing.T) {
	for sq, s := range map[square]string{
		A1: "a1", H1: "h1", A1 + 1: "b1", A1 - 10: "a2", A8: "a8", H8: "h8",
	} {
		if sq.String() != s {
			t.Error(int(sq), sq.String())
		}
	}
}

func TestMoves(t *testing.T) {
	for game, expected := range map[string]string{
		"r4rk1/ppp2ppp/2n2n2/5P2/2pb4/2N2N2/PPP2PPP/RQ2K2R": "a2a3 a2a4 b1c1 b1d1 b2b3 b2b4 c3a4 c3b5 c3d1 c3d5 c3e2 c3e4 e1d1 e1d2 e1e2 e1f1 f3d2 f3d4 f3e5 f3g1 f3g5 f3h4 g2g3 g2g4 h1f1 h1g1 h2h3 h2h4",
	} {
		b, err := fen(game)
		if err != nil {
			t.Error(game, b, err)
		}
		p := position{board: b}
		moves := []string{}
		for _, m := range p.moves() {
			moves = append(moves, m.String())
		}
		sort.Strings(moves)
		if s := strings.Join(moves, " "); s != expected {
			t.Error("\n", expected, "\n", s)
		}
	}
}

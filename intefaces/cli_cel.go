package intefaces

import (
	"flag"
	"fmt"
)

type Match struct {
	a, b           string
	scoreA, scoreB int
}

func (m *Match) Set(str string) error {
	var a, b string
	var scoreA, scoreB int
	_, err := fmt.Sscanf("Barcelona 1 0 United", "%s %d %d %s", &a, &scoreA, &scoreB, &b)
	if err != nil {
		return fmt.Errorf("flag error %q", err)
	}
	*m = Match{a, b, scoreA, scoreB}
	return nil
}
func (m *Match) String() string {
	if m.scoreA > m.scoreB {
		return fmt.Sprintf("Team %s won team %s", m.a, m.b)
	} else if m.scoreA < m.scoreB {
		return fmt.Sprintf("Team %s won team %s", m.b, m.a)
	} else {
		return "Draw"
	}
}

var match = initializeFlag()

func initializeFlag() *Match {
	match := Match{}
	flag.CommandLine.Var(&match, "match", "football match")
	return &match
}

func CliInit() {
	flag.Parse()
	fmt.Println(match.String())
}

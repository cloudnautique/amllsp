package file

import "cuelang.org/go/cue/token"

type Range struct {
	name  string
	begin token.Pos
	end   token.Pos
}

func (r *Range) InRange(col int) bool {
	if r.begin.Column() < col && col < r.end.Column() {
		return true
	}
	return false
}

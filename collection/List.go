package collection

import (
	"github.com/samber/lo"
)

type List struct {
	Val []string
}

func (l *List) Distinct() {
	l.Val = lo.Uniq(l.Val)
}

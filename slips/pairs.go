package slips

import (
	"sort"

	"github.com/iloginow/esportsdifference/compare"
)

type Pairs struct {
	Elements [][]compare.Record
	Used     []int
}

func NewPairs(e [][]compare.Record) *Pairs {
	pairs := new(Pairs)
	pairs.Elements = e
	return pairs
}

func (p *Pairs) Find() [][][]compare.Record {
	var paired [][][]compare.Record
	for i := 0; i < len(p.Elements); i += 1 {
		if p.findUsed(i) {
			continue
		}
		pair := p.findPair(i)
		if pair != nil {
			paired = append(paired, pair)
		}
	}
	
	sort.Slice(paired, func(i, j int) bool {
		iDiff := getTotalPairDiff(paired[i])
		jDiff := getTotalPairDiff(paired[j])
		if iDiff != jDiff {
			return iDiff > jDiff
		}
		return true
	})
	return paired
}

func (p *Pairs) findPair(i int) [][]compare.Record {
	it := NewTeams(p.Elements[i])
	for j := 0; j < len(p.Elements); j += 1 {
		if j == i || p.findUsed(j) {
			continue
		}
		jt := NewTeams(p.Elements[j])
		if !it.match(jt) {
			p.use(i)
			p.use(j)
			return [][]compare.Record{p.Elements[i], p.Elements[j]}
		}
	}
	return nil
}

func (p *Pairs) use(i int) {
	p.Used = append(p.Used, i)
}

func (p *Pairs) findUsed(i int) bool {
	for _, u := range p.Used {
		if u == i {
			return true
		}
	}
	return false
}

func getTotalPairDiff(p [][]compare.Record) float64 {
	var total float64
	for _, e := range p {
		total += getTotalElementDiff(e)
	}
	return total
}

func getTotalElementDiff(e []compare.Record) float64 {
	var total float64
	for _, r := range e {
		total += r.Difference
	}
	return total
}

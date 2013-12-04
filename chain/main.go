package chain

type Chain struct {
	ref      string
	backref  *Chain
	hasDown  bool
	downFile string
	upFile   string
}

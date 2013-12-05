package chain

type Direction int

const (
	up   Direction = iota
	down Direction = iota
)

type Chain struct {
	ref      string
	backref  *Chain
	hasDown  bool
	downFile string
	upFile   string
}

type Meta struct {
	ref       string
	backref   string
	direction Direction
	filename  string
}

func fileList() /*[]string*/ {

}

func cwdIsSchemaDir() /*bool*/ {

}

func parseMeta(filepath string) /* *Meta*/ {

}

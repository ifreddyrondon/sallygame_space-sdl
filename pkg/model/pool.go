package model

type Pool struct {
	elements []*Element
}

func NewPool(elements []*Element) Pool {
	return Pool{elements: elements}
}

func (p Pool) Get() (*Element, bool) {
	for _, bul := range p.elements {
		if !bul.Active {
			return bul, true
		}
	}
	return nil, false
}

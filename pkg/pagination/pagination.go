package pagination

type Pagination struct {
	PageSize int
	Page     int
}

func (p *Pagination) DefaultIfNotSet() *Pagination {
	if p.PageSize == 0 {
		p.PageSize = 10
	}

	return p
}

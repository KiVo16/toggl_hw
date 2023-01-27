package db

import "base/pkg/pagination"

func getLimitOffsetFromModelPagination(p pagination.Pagination) (limit int, offset int) {
	limit = p.PageSize
	offset = p.PageSize * p.Page

	return limit, offset
}

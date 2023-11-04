package middleware

import "github.com/morkid/paginate"

var Pagination = paginate.New(&paginate.Config{
	DefaultSize: 10,
})
package main

import (
	"net/http"

	"github.com/manoj-H-C/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

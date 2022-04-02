package dummydb

import (
	tkpdSqlt "github.com/kokka-team/nakama-investor-be/lib/sqlt"
)

type dummyDB struct {
	db *tkpdSqlt.DB
}

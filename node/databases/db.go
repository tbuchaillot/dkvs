package databases

import (
	"github.com/tbuchaillot/dkvs/node/databases/bolt"
	"strings"
)

type  Database interface{
	Close() error
	SetKey(key string, value []byte) error
	GetKey(key string) ([]byte,error)
}
const (
	BOLT_TYPE = "bolt"
)

func NewDatabase(dbType string ,dbPath string, dbExtension string) (Database, error){
	switch strings.ToLower(dbType) {
	default:
		return bolt.NewDatabase(dbPath+"."+dbExtension)
	}
	return nil, nil
}

package mongo

import "github.com/rs/xid"

func ObjectId() string {
	guid := xid.New()
	return guid.String()
}

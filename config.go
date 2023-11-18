package bot

import (
	"os"

	"github.com/aiteung/atdb"
)

var DBmongoinfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRING"),
	DBName:   "waapi",
}

var Mongoconn = atdb.MongoConnect(DBmongoinfo)

const WAKeyword = "wh4t5auth0"

var WebHookSecret = os.Getenv("SECRET")

var TokenAPIWA = os.Getenv("TOKEN")

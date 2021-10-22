package postgres

import (
	"fmt"
)

var (
	PGDBRead  *TGPGDB // concurrent access ok
	PGDBWrite *TGPGDB
)

func init() {

	fmt.Println("DBWriteURL", DBWriteURL)
	fmt.Println("DBReadURL", DBReadURL)
	PGDBWrite = CreateTGPGDB(DBWriteURL)
	PGDBRead = CreateTGPGDB(DBReadURL)
	// Keep DB connected as long as process is alive
}

package bot

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/aiteung/atdb"
)

func TestGetENVToken(t *testing.T) {
	fmt.Println(os.Getenv("MONGOSTRINGAWANGGA"))
}

func TestUpdateGetData(t *testing.T) {
	result, err := atdb.GetRandomDoc[Reply](Mongoconn, "reply", 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", strings.ReplaceAll(result[0].Message, "#BOTNAME#", "aku"))
}

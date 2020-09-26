package permutations

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestPerm(t *testing.T) {
	ans := Perm("abc")
	assert.Len(t, ans, 6)
}

package base

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"regexp"
	"testing"
	"time"
)

func TestGenerate_generateRegex(t *testing.T) {
	rand.Seed(time.Now().Unix())
	g := &Generate{MaxLength: 12}


	regexs:=[]string{"\\d+\\.\\d{2,4}","[a-zA-Z0-9]{4,4}"}
	for i := range regexs {
		val:=g.Generate2(regexs[i])
		assert.True(t, regexp.MustCompile(regexs[i]).MatchString(val))
	}
}
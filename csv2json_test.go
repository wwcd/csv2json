package csv2json

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConv(t *testing.T) {
	input := `
h1,h2,h3,h4
a,b,c,d
aa,bb,cc,dd
aaa,bbb,ccc,ddd
aaaa,bbbb,cccc,dddd
aaaaa,bbbbb,ccccc,ddddd
`
	expected := `[{"h2":"bb","h3":"cc"},{"h2":"bbb","h3":"ccc"}] `

	output := &bytes.Buffer{}

	err := Conv(bytes.NewBufferString(input), output, WithCol(1, 2), WithRow(2, 3))
	assert.NoError(t, err)
	assert.JSONEq(t, expected, output.String())
}

package file

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tliron/kutil/logging"
)

const TestDataDir = "testdata"

type TC struct {
	position []int
	output   []string
}

var tcs = []TC{
	{
		position: []int{1, 1},
		output: []string{`#### Type
## #Container: {
	#ContainerBase
	labels: [string]:      string
	annotations: [string]: string
	scale?: >=0
	sidecars: [string]: #Sidecar
}
`}}, {
		position: []int{7, 7},
		output: []string{`#### Type
## #Secret: *#SecretOpaque | #SecretBasicAuth | #SecretGenerated | #SecretTemplate | #SecretToken
`},
	}, {
		position: []int{12, 13},
		output: []string{`#### Type
## #Volume: {
	labels: [string]:      string
	annotations: [string]: string
	class:       string | *""
	size:        int | *10 | string
	accessModes: [#AccessMode, ...#AccessMode] | #AccessMode | *"readWriteOnce"
}
`}},
}

func TestPosition(t *testing.T) {
	logger := logging.GetLogger("testing")

	f, err := New(filepath.Join(TestDataDir, "simple", "Acornfile"), logger)
	if err != nil {
		t.Error(err)
	}

	for _, tc := range tcs {
		md, err := f.GetDocDefinitionMarkDown(tc.position[0], tc.position[1])
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tc.output[0], md)
	}
}

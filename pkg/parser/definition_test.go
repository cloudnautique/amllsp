package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const TestSourceDir = "./testdata"

var testCases = map[string]struct {
	input  string
	output []string
}{
	"loads": {
		input: "",
		output: []string{
			`#Probes: s[99:34]|e[99:41]`,
			`#ContextDirRef: s[110:1]|e[110:15]`,
			`#Image: s[288:28]|e[288:34]`,
			`#AcornVolumeBinding: s[262:30]|e[262:49]`,
			`#Router: s[292:28]|e[292:35]`,
			`#ProbeMap: s[68:19]|e[68:28]`,
			`#PortSingle: s[261:25]|e[261:36]`,
			`#Port: s[261:44]|e[261:49]`,
			`#RouteTargetName: s[274:1]|e[274:17]`,
			`#Acorn: s[291:28]|e[291:34]`,
			`#Args: s[284:32]|e[284:37]`,
			`#PortSpec: s[121:1]|e[121:10]`,
			`#DNSName: s[292:14]|e[292:22]`,
			`#SecretBase: s[207:2]|e[207:13]`,
			`#Route: s[241:1]|e[241:7]`,
			`#RouteTarget: s[253:23]|e[253:35]`,
			`#ProbeSpec: s[68:35]|e[68:45]`,
			`#PortMap: s[261:53]|e[261:61]`,
			`#EphemeralRef: s[109:1]|e[109:14]`,
			`#RouteMap: s[252:1]|e[252:10]`,
			`#Build: s[152:20]|e[152:26]`,
			`#ContainerBase: s[87:1]|e[87:15]`,
			`#FileContent: s[88:36]|e[88:48]`,
			`#ScopedLabel: s[258:30]|e[258:42]`,
			`#AcornServiceBinding: s[264:30]|e[264:50]`,
			`#FileSecretSpec: s[81:11]|e[81:26]`,
			`#VolumeRef: s[108:1]|e[108:11]`,
			`#SecretRef: s[111:1]|e[111:11]`,
			`#RuleSpec: s[270:21]|e[270:30]`,
			`#ShortVolumeRef: s[107:1]|e[107:16]`,
			`#ScopedLabelMapKey: s[133:20]|e[133:38]`,
			`#ScopedLabelMap: s[258:46]|e[258:61]`,
			`#Volume: s[289:28]|e[289:35]`,
			`#EnvVars: s[265:25]|e[265:33]`,
			`#Sidecar: s[36:22]|e[36:30]`,
			`#FileSpec: s[85:135]|e[85:144]`,
			`#SecretTemplate: s[216:65]|e[216:80]`,
			`#Secret: s[290:28]|e[290:35]`,
			`#App: s[282:1]|e[282:5]`,
			`#SecretToken: s[216:83]|e[216:95]`,
			`#SecretBasicAuth: s[216:27]|e[216:43]`,
			`#Container: s[286:28]|e[286:38]`,
			`#Job: s[287:28]|e[287:32]`,
			`#Dir: s[115:1]|e[115:5]`,
			`#AccessMode: s[232:47]|e[232:58]`,
			`#SecretOpaque: s[216:11]|e[216:24]`,
			`#SecretGenerated: s[216:46]|e[216:62]`,
			`#PathName: s[276:1]|e[276:10]`,
			`#AcornBuild: s[260:34]|e[260:45]`,
			`#PortRegexp: s[119:1]|e[119:12]`,
			`#AcornSecretBinding: s[263:30]|e[263:49]`,
		},
	},
}

func TestDefinitionParsing(t *testing.T) {
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			defs, err := ParseAcornDefinitions()
			if err != nil {
				t.Fatal(err)
			}
			output := defs.String()
			for _, o := range tc.output {
				require.Contains(t, output, o)
			}
		})
	}
}

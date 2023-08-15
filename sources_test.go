package solgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txpull/solgo/tests"
)

func TestSources(t *testing.T) {
	// Define multiple test cases
	testCases := []struct {
		name          string
		sources       Sources
		expected      string
		expectedUnits int
	}{
		{
			name: "Test Case 1",
			sources: Sources{
				SourceUnits: []*SourceUnit{
					{
						Name:    "Source 1",
						Path:    "path/to/source1",
						Content: "Content of Source 1",
					},
					{
						Name:    "Source 2",
						Path:    "path/to/source2",
						Content: "Content of Source 2",
					},
				},
				EntrySourceUnitName: "Source",
				LocalSourcesPath:    "./sources/",
			},
			expected:      "Content of Source 1\n\nContent of Source 2",
			expectedUnits: 2,
		},
		{
			name: "Openzeppelin import",
			sources: Sources{
				SourceUnits: []*SourceUnit{
					{
						Name:    "Import",
						Path:    "Import.sol",
						Content: tests.ReadContractFileForTestFromRootPath(t, "contracts/cheelee/Import").Content,
					},
				},
				EntrySourceUnitName: "Cheelee",
				LocalSourcesPath:    "./sources/",
			},
			expected:      tests.ReadContractFileForTestFromRootPath(t, "contracts/cheelee/Combined").Content, // @TODO
			expectedUnits: 15,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.sources.Prepare()
			assert.NoError(t, err)
			//combinedSource := testCase.sources.GetCombinedSource()
			//assert.Equal(t, testCase.expected, combinedSource)
			//os.WriteFile(fmt.Sprintf("combined_%d.sol", i), []byte(combinedSource), 0755)
			assert.Equal(t, testCase.expectedUnits, len(testCase.sources.SourceUnits))
			assert.NotNil(t, testCase.sources.ToProto())
			assert.NoError(t, testCase.sources.WriteToDir("./data/tests/sources/"))
			assert.NoError(t, testCase.sources.TruncateDir("./data/tests/sources/"))
			assert.True(t, testCase.sources.ArePrepared())
			assert.True(t, testCase.sources.SourceUnitExistsIn(testCase.sources.SourceUnits[0].Name, testCase.sources.SourceUnits))
			assert.True(t, testCase.sources.SourceUnitExists(testCase.sources.SourceUnits[0].Name))
			assert.NotNil(t, testCase.sources.GetSourceUnitByName(testCase.sources.SourceUnits[0].Name))
			assert.NotNil(t, testCase.sources.GetSourceUnitByPath(testCase.sources.SourceUnits[0].Path))
		})
	}
}

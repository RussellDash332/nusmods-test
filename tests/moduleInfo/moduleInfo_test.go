package moduleInfo

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"testing"
	"reflect"
	"nusmods-test/config"

	"github.com/stretchr/testify/assert"
)

func TestModuleInformation(t *testing.T) {
	assert := assert.New(t)

	for _, scenario := range config.Scenarios {
		t.Run(fmt.Sprintf("%s", scenario.Name), func(t *testing.T) {
			resp, err := http.Get("https://api.nusmods.com/v2/" + scenario.Name + "/moduleInfo.json")
			if err != nil {
				t.Fatalf("(%s) Cannot get response: %+v\n", scenario.Name, err.Error())
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("(%s) Cannot read response body: %+v\n", scenario.Name, err.Error())
			}

			var moduleInfoList []ModuleInfo
			if err := json.Unmarshal(body, &moduleInfoList); err != nil {
				t.Fatalf("(%s) Cannot unmarshal response JSON: %+v\n", scenario.Name, err.Error())
			}

			bodyExpected, err := ioutil.ReadFile(scenario.ExpectedJsonPath)
			if err != nil {
				t.Fatalf("(%s) Cannot read test data: %+v\n", scenario.Name, err.Error())
			}

			var moduleInfoListExpected []ModuleInfo
			if err := json.Unmarshal(bodyExpected, &moduleInfoListExpected); err != nil {
				t.Fatalf("(%s) Cannot unmarshal expected JSON: %+v\n", scenario.Name, err.Error())
			}

			if !reflect.DeepEqual(moduleInfoList, moduleInfoListExpected) {
				t.Errorf("%s: test JSON differs with expected JSON", scenario.Name)
			}

			// Show diff in local
			if config.ShowDiff {
				assert.Equalf(moduleInfoList, moduleInfoListExpected, "%s: test JSON differs with expected JSON", scenario.Name)
			}
		})
	}
}

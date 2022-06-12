package moduleInfo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"nusmods-test/config"

	"github.com/stretchr/testify/assert"
)

func TestModuleInformation(t *testing.T) {
	assert := assert.New(t)

	for _, scenario := range config.Scenarios {
		resp, err := http.Get("https://api.nusmods.com/v2/" + scenario.Name + "/moduleInfo.json")
		if err != nil {
			assert.Truef(false, "(%s) Cannot get response: %+v\n", scenario.Name, err.Error())
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			assert.Truef(false, "(%s) Cannot read response body: %+v\n", scenario.Name, err.Error())
		}

		var moduleInfoList []ModuleInfo
		if err := json.Unmarshal(body, &moduleInfoList); err != nil {
			assert.Truef(false, "(%s) Cannot unmarshal response JSON: %+v\n", scenario.Name, err.Error())
		}

		bodyExpected, err := ioutil.ReadFile(scenario.ExpectedJsonPath)
		if err != nil {
			assert.Truef(false, "(%s) Cannot read test data: %+v\n", scenario.Name, err.Error())
		}

		var moduleInfoListExpected []ModuleInfo
		if err := json.Unmarshal(bodyExpected, &moduleInfoListExpected); err != nil {
			assert.Truef(false, "(%s) Cannot unmarshal expected JSON: %+v\n", scenario.Name, err.Error())
		}

		assert.Equalf(moduleInfoList, moduleInfoListExpected, "%s: test JSON differs with expected JSON", scenario.Name)
		t.Logf("%s", scenario.Name)
	}
}

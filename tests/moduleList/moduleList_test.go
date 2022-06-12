package moduleList

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"nusmods-test/config"

	"github.com/stretchr/testify/assert"
)

func TestModuleList(t *testing.T) {
	assert := assert.New(t)

	for _, scenario := range config.Scenarios {
        resp, err := http.Get("https://api.nusmods.com/v2/" + scenario.Name + "/moduleList.json")
		if err != nil {
			assert.Truef(false, "(%s) Cannot get response: %+v\n", scenario.Name, err.Error())
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
            assert.Truef(false, "(%s) Cannot read response body: %+v\n", scenario.Name, err.Error())
		}

		var modulesList []Module
		if err := json.Unmarshal(body, &modulesList); err != nil {
			assert.Truef(false, "(%s) Cannot unmarshal response JSON: %+v\n", scenario.Name, err.Error())
		}

		bodyExpected, err := ioutil.ReadFile(scenario.ExpectedJsonPath)
		if err != nil {
			assert.Truef(false, "(%s) Cannot read test data: %+v\n", scenario.Name, err.Error())
		}

		var modulesListExpected []Module
		if err := json.Unmarshal(bodyExpected, &modulesListExpected); err != nil {
			assert.Truef(false, "(%s) Cannot unmarshal expected JSON: %+v\n", scenario.Name, err.Error())
		}

		assert.Equalf(modulesList, modulesListExpected, "%s: test JSON differs with expected JSON", scenario.Name)
		t.Logf("%s", scenario.Name)
	}
}

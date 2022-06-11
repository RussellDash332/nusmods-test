package modules

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"strings"
	"nusmods-test/config"

	"github.com/stretchr/testify/assert"
)

func TestModules(t *testing.T) {
	assert := assert.New(t)

	for _, scenario := range config.Scenarios {
		for _, module := range config.Modules {
			module = strings.ReplaceAll(module, "\r", "")
			resp, err := http.Get("https://api.nusmods.com/v2/" + scenario.Name + "/modules/" + module + ".json")
			if err != nil {
				assert.Truef(false, "(%s, %s) Cannot get response: %+v\n", module, scenario.Name, err.Error())
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				assert.Truef(false, "(%s, %s) Cannot read response body: %+v\n", module, scenario.Name, err.Error())
			}

			var moduleData ModuleData
			if err := json.Unmarshal(body, &moduleData); err != nil {
				assert.Truef(false, "(%s, %s) Cannot unmarshal response JSON: %+v\n", module, scenario.Name, err.Error())
			}

			bodyExpected, err := ioutil.ReadFile("testdata/" + module + "_" + scenario.Name + ".json")
			if err != nil {
				assert.Truef(false, "(%s, %s) Cannot read test data: %+v\n", module, scenario.Name, err.Error())
			}

			var moduleDataExpected ModuleData
			if err := json.Unmarshal(bodyExpected, &moduleDataExpected); err != nil {
				assert.Truef(false, "(%s, %s) Cannot unmarshal expected JSON: %+v\n", module, scenario.Name, err.Error())
			}

			assert.Equalf(moduleData, moduleDataExpected, "%s (%s): test JSON differs with expected JSON", module, scenario.Name)
		}
	}
}

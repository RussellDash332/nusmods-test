package modules

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"testing"
	"strings"
	"reflect"
	"nusmods-test/config"

	"github.com/stretchr/testify/assert"
)

func TestModules(t *testing.T) {
	assert := assert.New(t)

	for _, scenario := range config.Scenarios {
		for _, module := range config.Modules {
			module = strings.ReplaceAll(module, "\r", "")
			t.Run(fmt.Sprintf("%s %s", module, scenario.Name), func(t *testing.T) {
				resp, err := http.Get("https://api.nusmods.com/v2/" + scenario.Name + "/modules/" + module + ".json")
				if err != nil {
					t.Fatalf("(%s, %s) Cannot get response: %+v\n", module, scenario.Name, err.Error())
				}

				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("(%s, %s) Cannot read response body: %+v\n", module, scenario.Name, err.Error())
				}

				var moduleData ModuleData
				if err := json.Unmarshal(body, &moduleData); err != nil {
					t.Fatalf("(%s, %s) Cannot unmarshal response JSON: %+v\n", module, scenario.Name, err.Error())
				}

				bodyExpected, err := ioutil.ReadFile("testdata/" + scenario.Name + "/" + module + ".json")
				if err != nil {
					t.Fatalf("(%s, %s) Cannot read test data: %+v\n", module, scenario.Name, err.Error())
				}

				var moduleDataExpected ModuleData
				if err := json.Unmarshal(bodyExpected, &moduleDataExpected); err != nil {
					t.Fatalf("(%s, %s) Cannot unmarshal expected JSON: %+v\n", module, scenario.Name, err.Error())
				}

				if !reflect.DeepEqual(moduleData, moduleDataExpected) {
					t.Errorf("%s (%s): test JSON differs with expected JSON", module, scenario.Name)
				}

				// Show diff in local
				if config.ShowDiff {
					assert.Equalf(moduleData, moduleDataExpected, "%s (%s): test JSON differs with expected JSON", module, scenario.Name)
				}
			})
		}
	}
}

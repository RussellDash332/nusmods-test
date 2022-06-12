package moduleList

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

func TestModuleList(t *testing.T) {
	assert := assert.New(t)

	for _, scenario := range config.Scenarios {
        t.Run(fmt.Sprintf("%s", scenario.Name), func(t *testing.T) {
			resp, err := http.Get("https://api.nusmods.com/v2/" + scenario.Name + "/moduleList.json")
			if err != nil {
				t.Fatalf("(%s) Cannot get response: %+v\n", scenario.Name, err.Error())
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("(%s) Cannot read response body: %+v\n", scenario.Name, err.Error())
			}

			var modulesList []Module
			if err := json.Unmarshal(body, &modulesList); err != nil {
				t.Fatalf("(%s) Cannot unmarshal response JSON: %+v\n", scenario.Name, err.Error())
			}

			bodyExpected, err := ioutil.ReadFile(scenario.ExpectedJsonPath)
			if err != nil {
				t.Fatalf("(%s) Cannot read test data: %+v\n", scenario.Name, err.Error())
			}

			var modulesListExpected []Module
			if err := json.Unmarshal(bodyExpected, &modulesListExpected); err != nil {
				t.Fatalf("(%s) Cannot unmarshal expected JSON: %+v\n", scenario.Name, err.Error())
			}

			if !reflect.DeepEqual(modulesList, modulesListExpected) {
				t.Errorf("%s: test JSON differs with expected JSON", scenario.Name)
			}

			// Show diff in local
			if config.ShowDiff {
				assert.Equalf(modulesList, modulesListExpected, "%s: test JSON differs with expected JSON", scenario.Name)
			}
		})
	}
}

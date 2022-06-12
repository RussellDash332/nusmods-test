package venueInformation

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

func TestVenueInformation(t *testing.T) {
	assert := assert.New(t)

	for _, scenario := range config.Scenarios {
		for _, semester := range config.Semesters {
			semester = strings.ReplaceAll(semester, "\r", "")
			t.Run(fmt.Sprintf("Semester %s, %s", semester, scenario.Name), func(t *testing.T) {
				resp, err := http.Get("https://api.nusmods.com/v2/" + scenario.Name + "/semesters/" + semester + "/venueInformation.json")
				if err != nil {
					t.Fatalf("(Semester %s, %s) Cannot get response: %+v\n", semester, scenario.Name, err.Error())
				}

				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("(Semester %s, %s) Cannot read response body: %+v\n", semester, scenario.Name, err.Error())
				}

				var venueInfo map[string][]VenueInfo
				if err := json.Unmarshal(body, &venueInfo); err != nil {
					t.Fatalf("(Semester %s, %s) Cannot unmarshal response JSON: %+v\n", semester, scenario.Name, err.Error())
				}

				bodyExpected, err := ioutil.ReadFile("testdata/" + scenario.Name + "/" + semester + ".json")
				if err != nil {
					t.Fatalf("(Semester %s, %s) Cannot read test data: %+v\n", semester, scenario.Name, err.Error())
				}

				var venueInfoExpected map[string][]VenueInfo
				if err := json.Unmarshal(bodyExpected, &venueInfoExpected); err != nil {
					t.Fatalf("(Semester %s, %s) Cannot unmarshal expected JSON: %+v\n", semester, scenario.Name, err.Error())
				}

				if !reflect.DeepEqual(venueInfo, venueInfoExpected) {
					t.Errorf("(Semester %s, %s): test JSON differs with expected JSON", semester, scenario.Name)
				}

				// Show diff in local
				if config.ShowDiff {
					assert.Equalf(venueInfo, venueInfoExpected, "(Semester %s, %s): test JSON differs with expected JSON", semester, scenario.Name)
				}
			})
		}
	}
}

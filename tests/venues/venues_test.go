package venues

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"testing"
	"sort"
	"strings"
	"reflect"
	"nusmods-test/config"

	"github.com/stretchr/testify/assert"
)

func TestVenues(t *testing.T) {
	assert := assert.New(t)

	for _, scenario := range config.Scenarios {
		for _, semester := range config.Semesters {
			semester = strings.ReplaceAll(semester, "\r", "")
			t.Run(fmt.Sprintf("Semester %s, %s", semester, scenario.Name), func(t *testing.T) {
				resp, err := http.Get("https://api.nusmods.com/v2/" + scenario.Name + "/semesters/" + semester + "/venues.json")
				if err != nil {
					t.Fatalf("(Semester %s, %s) Cannot get response: %+v\n", semester, scenario.Name, err.Error())
				}

				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("(Semester %s, %s) Cannot read response body: %+v\n", semester, scenario.Name, err.Error())
				}

				var venueList []string
				if err := json.Unmarshal(body, &venueList); err != nil {
					t.Fatalf("(Semester %s, %s) Cannot unmarshal response JSON: %+v\n", semester, scenario.Name, err.Error())
				}

				bodyExpected, err := ioutil.ReadFile("testdata/" + scenario.Name + "/" + semester + ".json")
				if err != nil {
					t.Fatalf("(Semester %s, %s) Cannot read test data: %+v\n", semester, scenario.Name, err.Error())
				}

				var venueListExpected []string
				if err := json.Unmarshal(bodyExpected, &venueListExpected); err != nil {
					t.Fatalf("(Semester %s, %s) Cannot unmarshal expected JSON: %+v\n", semester, scenario.Name, err.Error())
				}

				// List needs to be sorted
				sort.Strings(venueList)
				sort.Strings(venueListExpected)
				if !reflect.DeepEqual(venueList, venueListExpected) {
					t.Errorf("(Semester %s, %s): test JSON differs with expected JSON", semester, scenario.Name)
				}

				// Show diff in local
				if config.ShowDiff {
					assert.Equalf(venueList, venueListExpected, "(Semester %s, %s): test JSON differs with expected JSON", semester, scenario.Name)
				}
			})
		}
	}
}

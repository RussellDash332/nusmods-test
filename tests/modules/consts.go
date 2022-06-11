package modules

type ModuleData struct {
    AcadYear        string
    Preclusion      string  `json:"preclusion,omitempty"`
    Description     string
    Title           string
    Department      string
    Faculty      	string
    Workload     	interface{} // ad-hoc conversion
    ModuleCredit	string
    ModuleCode      string
    Prerequisite    string  `json:"prerequisite,omitempty"`
    // Might as well change below to map[string]bool
    Attributes struct {
        SU    bool `json:"su,omitempty"`    // SU exercise
        GRSU  bool `json:"grsu,omitempty"`  // SU for graduates
        ISM   bool `json:"ism,omitempty"`   // Independent study module
        FYP   bool `json:"fyp,omitempty"`   // Final year
        LAB   bool `json:"lab,omitempty"`   // Lab-based module
        SFS   bool `json:"sfs,omitempty"`   // SkillsFuture series
        SSGF  bool `json:"ssgf,omitempty"`  // SkillsFuture funded
        UROP  bool `json:"urop,omitempty"`  // UROPS program
        YEAR  bool `json:"year,omitempty"`  // Year long module
        MPES1 bool `json:"mpes1,omitempty"` // Module Planning Exercise Semester 1
        MPES2 bool `json:"mpes2,omitempty"` // Module Planning Exercise Semester 2
    }
    Aliases         []string `json:"aliases,omitempty"`
    SemesterData []struct {
        Semester    int
        Timetable []struct {
            ClassNo     string
            StartTime   string
            EndTime     string
            Weeks       interface{} // either map[string]string or []int
            Venue       string
            Day         string
            LessonType  string
            Size        int
        } `json:"timetable,omitempty"`
        ExamDate     string   `json:"examDate,omitempty"`
        ExamDuration int      `json:"examDuration,omitempty"`
        /**
        CovidZones   []string `json:"covidZones,omitempty"`
        **/
    }
    PrereqTree          interface{} `json:"prereqTree,omitempty"` // either string or map[string][string]
    FulfillRequirements []string    `json:"fulfillRequirements,omitempty"`
}

type Scenario struct {
    Name string
}

var modulesList = []string {
    "CS1010S", "GER1000", "MA2104", "DSA2102",
    "ACC1701X", "IT5002", "GEH1028", "SP3172",
    "SP1541", "CS2040", "MA2101S", "PL2131",
}

var Scenarios = []Scenario{
    {
        Name:   "2019-2020",
    },
    {
        Name:   "2020-2021",
    },
    {
        Name:   "2021-2022",
    },
}

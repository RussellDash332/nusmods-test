package moduleInfo

type ModuleInfo struct {
	ModuleCode   string
	Title        string
	Description  string
	ModuleCredit string
	Department   string
	Faculty      string
	Workload     interface{} // ad-hoc conversion
	Preclusion   string
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
	SemesterData []struct {
		Semester     int
		ExamDate     string   `json:"examDate,omitempty"`
		ExamDuration int      `json:"examDuration,omitempty"`
		/**
		CovidZones   []string `json:"covidZones,omitempty"`
		**/
	}
}

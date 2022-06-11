package venueInformation

type VenueInfo struct {
    Day             string
    Classes []struct {
        ClassNo     string
        StartTime   string
        EndTime     string
        Weeks       interface{}
        Day         string
        LessonType  string
        Size        int
        ModuleCode  string
    }
    Availability    map[string]string `json:"availability, omitempty`
}

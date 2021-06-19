package db

type DeathStatistics struct {
	Id            string `gorm:"primary_key"`
	Location      string
	Date          string
	Number        int64
	ReportingUnit string
}

func (DeathStatistics) TableName() string {
	return "death_statistics"
}

type CommDisaster struct {
	Id            string `gorm:"primary_key"`
	Date          string
	Location      string
	Type          string
	Grade         string
	Picture       []byte `gorm:"type:blob"`
	Note          string
	ReportingUnit string
}

func (CommDisaster) TableName() string {
	return "comm_disaster"
}

type BrickwoodStructure struct {
	Id                    string `gorm:"primary_key"`
	Date                  string
	Location              string
	BasicallyIntactSquare string
	DamagedSquare         string
	DestroyedSquare       string
	Note                  string
	ReportingUnit         string
}

func (BrickwoodStructure) TableName() string {
	return "brickwood_structure"
}

type CivilStructure struct {
	Id                    string `gorm:"primary_key"`
	Date                  string
	Location              string
	BasicallyIntactSquare string
	DamagedSquare         string
	DestroyedSquare       string
	Note                  string
	ReportingUnit         string
}

func (CivilStructure) TableName() string {
	return "civil_structure"
}

type CollapseRecord struct {
	Id            string `gorm:"primary_key"`
	Date          string
	Location      string
	Type          string
	Picture       []byte `gorm:"type:blob"`
	Note          string
	ReportingUnit string
	Status        string
}

func (CollapseRecord) TableName() string {
	return "collapse_record"
}

type CrackRecord struct {
	Id            string `gorm:"primary_key"`
	Date          string
	Location      string
	Type          string
	Picture       []byte `gorm:"type:blob"`
	Note          string
	ReportingUnit string
	Status        string
}

func (CrackRecord) TableName() string {
	return "crack_record"
}

type DebrisRecord struct {
	Id            string `gorm:"primary_key"`
	Date          string
	Location      string
	Type          string
	Picture       []byte `gorm:"type:blob"`
	Note          string
	ReportingUnit string
	Status        string
}

func (DebrisRecord) TableName() string {
	return "debris_record"
}

type DisasterInfo struct {
	DId           string `gorm:"primary_key"`
	Id            string `gorm:"primary_key"`
	Date          string
	Location      string
	Longitude     float64
	Latitude      float64
	Depth         float32
	Magnitude     float32
	Picture       []byte `gorm:"type:blob"`
	ReportingUnit string
}

func (DisasterInfo) TableName() string {
	return "disaster_info"
}

type DisasterPrediction struct {
	Did           string `gorm:"primary_key"`
	Date          string
	Picture       []byte `gorm:"type:blob"`
	ReportingUnit string
	Sid           string
	Grade         string
	Intensity     string
	Type          string
	Note          string
}

func (DisasterPrediction) TableName() string {
	return "disaster_prediction"
}

type DisasterRequest struct {
	Id           string `gorm:"primary_key"`
	Date         string
	DisasterType string
	Status       string
	OUrl         string
	Requestunit  string
}

func (DisasterRequest) TableName() string {
	return "disaster_request"
}

type FrameworkStructure struct {
	Id                    string `gorm:"primary_key"`
	Date                  string
	Location              string
	BasicallyIntactSquare string
	SlightDamagedSquare   string
	DestroyedSquare       string
	Note                  string
	ReportingUnit         string
	ModerateDamagedSquare string
	SeriousDamagedSquare  string
}

func (FrameworkStructure) TableName() string {
	return "framework_structure"
}

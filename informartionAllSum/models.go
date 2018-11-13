package informartionAllSum

type TASKINFORS struct {
	//Id     int    `gorm:"column:id"`
	TaskId    string `gorm:"column:task_id"`
	ProjectId string `gorm:"column:project_id"`
}

type TASKINFORSSUN struct {
	//Id int `gorm:"column:id"`
	TaskId_86 string `gorm:"column:task_id_86"`
	TaskId_87 string `gorm:"column:task_id_87"`
	TaskId_88 string `gorm:"column:task_id_88"`
	TaskId_89 string `gorm:"column:task_id_89"`
	TaskId_a6 string `gorm:"column:task_id_a6"`
	TaskId_b4 string `gorm:"column:task_id_b4"`
	TaskId_5e string `gorm:"column:task_id_5e"`
}

type INFORMARTIONALLSUN struct {
	//Id       int    `gorm:"column:id"`
	TaskId   string `gorm:"column:task_id"`
	GetTime  string `gorm:"column:get_time"`
	CountSum int    `gorm:"column:count_sum"`
}

type INFORMARTIONALL struct {
	//ID               int    `gorm:"column:id"`
	Url              string `gorm:"column:url"`
	Title            string `gorm:"column:title"`
	Author           string `gorm:"column:author"`
	Source           string `gorm:"column:source"`
	Release_datetime string `gorm:"column:release_datetime"`
	Content          string `gorm:"column:content"`
	Media_type       string `gorm:"column:media_type"`
	Original_title   string `gorm:"column:original_title"`
	Editor           string `gorm:"column:editor"`
	Reporter         string `gorm:"column:reporter"`
	Contents         string `gorm:"column:contents"`
	Reading_times    string `gorm:"column:reading_times"`
	Abstract_data    string `gorm:"column:abstract_data"`
	Media            string `gorm:"column:media"`
	Media_channel    string `gorm:"column:media_channel"`
	Location         string `gorm:"column:location"`
	Location_path    string `gorm:"column:location_path"`
	Collection_tool  string `gorm:"column:collection_tool"`
	User             string `gorm:"column:user"`
	Site_url         string `gorm:"column:site_url"`
	Leaf_id          string `gorm:"column:leaf_id"`
	Task_id          string `gorm:"column:task_id"`
	Task_name        string `gorm:"column:task_name"`
	Get_time         string `gorm:"column:get_time"`
	Keyword          string `gorm:"column:keyword"`
	Pub_time         string `gorm:"column:pub_time"`
	Project_id       string `gorm:"column:project_id"`
	Error_msg        string `gorm:"column:error_msg"`
}

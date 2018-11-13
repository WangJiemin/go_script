package models

type INFORALL struct {
	ID               int    `gorm:"column:id"`
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

type TASKINFOS struct {
	TASKID string `gorm:"column:task_id"`
	NAME   string `gorm:"column:name"`
}

type TASKCOUNT struct {
	COUNT int
}

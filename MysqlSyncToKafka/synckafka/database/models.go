package database

/*
import (
	"github.com/WangJiemin/gocomm/json"
)
*/

type InForMationTwo struct {
	Id          int    `gorm:"column:id"`
	Url         string `gorm:"column:url"`
	Content     string `gorm:"column:content"`
	DataSource  string `gorm:"column:data_source"`
	CrawlSource string `gorm:"column:crawl_source"`
	Keywork     string `gorm:"column:keywork"`
	Pubdatetime string `gorm:"column:pubdatetime"`
	Gettime     string `gorm:"column:gettime"`
	ContentType string `gorm:"column:content_type"`
	AnalysisFlg string `gorm:"column:analysis_flg"`
	SyncStatus  string `gorm:"column:sync_status"`
	BackupFlg   string `gorm:"column:backup_flg"`
}

/*
type INFORMATRIONALL struct {
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

type INFOallForJson struct {
	ID               int    `json:"id"`
	Url              string `json:"url"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	Source           string `json:"source"`
	Release_datetime string `json:"release_datetime"`
	Content          string `json:"content"`
	Media_type       string `json:"media_type"`
	Original_title   string `json:"original_title"`
	Editor           string `json:"editor"`
	Reporter         string `json:"reporter"`
	Contents         string `json:"contents"`
	Reading_times    string `json:"reading_times"`
	Abstract_data    string `json:"abstract_data"`
	Media            string `json:"media"`
	Media_channel    string `json:"media_channel"`
	Location         string `json:"location"`
	Location_path    string `json:"location_path"`
	Collection_tool  string `json:"collection_tool"`
	User             string `json:"user"`
	Site_url         string `json:"site_url"`
	Leaf_id          string `json:"leaf_id"`
	Task_id          string `json:"task_id"`
	Task_name        string `json:"task_name"`
	Get_time         string `json:"get_time"`
	Keyword          string `json:"keyword"`
	Pub_time         string `json:"pub_time"`
	Project_id       string `json:"project_id"`
	Error_msg        string `json:"error_msg"`
}

func FromInfoallToJson(info *INFORMATRIONALL) *INFOallForJson {
	infoallForJson := infoallForJson{
		ID:               info.ID,
		Url:              info.Url,
		Title:            info.Title,
		Author:           info.Author,
		Source:           info.Source,
		Release_datetime: info.Release_datetime,
		Content:          info.Content,
		Media_type:       info.Media_type,
		Original_title:   info.Original_title,
		Editor:           info.Editor,
		Reporter:         info.Reporter,
		Contents:         info.Contents,
		Reading_times:    info.Reading_times,
		Abstract_data:    info.Abstract_data,
		Media:            info.Media,
		Media_channel:    info.Media_channel,
		Location:         info.Location,
		Location_path:    info.Location_path,
		Collection_tool:  info.Collection_tool,
		User:             info.User,
		Site_url:         info.Site_url,
		Leaf_id:          info.Leaf_id,
		Task_id:          info.Task_id,
		Task_name:        info.Task_name,
		Get_time:         info.Get_time,
		Keyword:          info.Keyword,
		Pub_time:         info.Pub_time,
		Project_id:       info.Project_id,
		Error_msg:        info.Error_msg,
	}
}
*/

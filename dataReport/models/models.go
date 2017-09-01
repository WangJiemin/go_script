package models

const (
	TimeFormart = "2006-01-02 15:04:05"
)

type INFOTEMP struct {
	TASKID          string `gorm:"type:varchar(128);"`
	COUNT           int
	COLLECTED_COUNT int
	INSERT_COUNT    int
	PAGE_COUNT      int
	//COUNT_NULL      int
}

type INFO struct {
	//ID               int    `gorm:"cloumn:id"`
	//URL              string `gorm:"cloumn:url"`
	//TITLE            string `gorm:"cloumn:title"`
	//AUTHOR           string `gorm:"cloumn:author"`
	//SOURCE           string `gorm:"cloumn:source"`
	//RELEASE_DATETIME string `gorm:"cloumn:release_datetime"`
	//CONTENT          string `gorm:"cloumn:content"`
	//MEDIA_TYPE       string `gorm:"cloumn:media_type"`
	//ORIGINAL_TITLE   string `gorm:"cloumn:original_title"`
	//EDITOR           string `gorm:"cloumn:editor"`
	//REPORTER         string `gorm:"cloumn:reporter"`
	//CONTENTS         string `gorm:"cloumn:contents"`
	//READING_TIMES    string `gorm:"cloumn:reading_times"`
	//ABSTRACT_DATA    string `gorm:"cloumn:abstract_data"`
	//MEDIA            string `gorm:"cloumn:media"`
	//MEDIA_CHANNEL    string `gorm:"cloumn:media_channel"`
	//LOCATION         string `gorm:"cloumn:location"`
	//LOCATION_PATH    string `gorm:"cloumn:location_path"`
	//COLLECTION_TOOL  string `gorm:"cloumn:collection_tool"`
	//USER             string `gorm:"cloumn:user"`
	//SITE_URL         string `gorm:"cloumn:site_url"`
	//LEAF_ID          string `gorm:"cloumn:leaf_id"`
	TASK_ID string `gorm:"cloumn:task_id"`
	//TASK_NAME        string `gorm:"cloumn:task_name"`
	//GET_TIME         string `gorm:"cloumn:get_time"`
	//KEYWORD          string `gorm:"cloumn:keyword"`
	//PUB_TIME         string `gorm:"cloumn:pub_time"`
	//PROJECT_ID       string `gorm:"cloumn:project_id"`
	//ERROR_MSG        string `gorm:"cloumn:error_msg"`
	COUNT int
}

type TASKINFOS struct {
	//ID                int
	TASK_ID string
	//CONFIG_INFO       string
	CREATE_TIME string
	//IS_INCREMENTAL    int
	//IS_TIMER_START    int
	NAME string
	//NEED_SCREENSHOT   int
	//PROJECT_ID        string
	//PROJECT_NAME      string
	//OWNER_ID          int
	OWNER_NAME string
	//START_TIME        time.Time
	//STATUS            int
	//STOP_TIME         time.Time
	//TYPE_REPEAT       int
	//TYPE_TASK         int
	//UPDATE_TIME       time.Time
	//VERSION           int
	//IS_FAVORITE       int
	//STEP_PROGRESS     int
	//IS_START          int
	//HAS_STOP_TIME     int
	//IS_KEEP_LAST_DATA int
	//RUN_PERIODICALLY  int
	//INTERVAL          int
	//INTERVAL_TYPE     int
	//RULE_INFO         string
	//REMARKS           string
	//CHANNEL           string
	//NEED_SREENSHOT    int
	ENTRY_LINK      string
	PAGE_COUNT      string
	COLLECTED_COUNT string
	INSERT_COUNT    string
	COUNT           string
}

type TASKLOGS struct {
	//ID              int
	//OWNER_ID        int
	TASK_ID string
	//TASK_NAME       string
	//PROJECT_ID      string
	COLLECTED_COUNT string
	INSERT_COUNT    string
	PAGE_COUNT      string
	//SERVICE_NAME    string
	//START_TIME      string
	//FINISH_TIME     string
	//UPDATE_TIME     string
}

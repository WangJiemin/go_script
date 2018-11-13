package models

import "time"

func DBAction(db string) {
	creatrtable := "CREATE TABLE myinfo " +
		"( id int(11) NOT NULL AUTO_INCREMENT COMMENT 'id', " +
		"name varchar(10) NOT NULL DEFAULT 'baabb' COMMENT 'agedddd', " +
		"age2 int(11) NOT NULL COMMENT 'agddd', " +
		"age int(11) NOT NULL COMMENT 'age', " +
		"PRIMARY KEY (id), " +
		"KEY idx_name (id,name))" +
		"ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='t'"

	altertable := "alter table myinfo rename to myinfo1," +
		"add column age3 int not null comment 'age'," +
		"add index idx_age2(age2)," +
		"drop column age"

	inserttable := "INSERT INTO myinfo1" +
		"(name,age2,age) " +
		"SELECT DataType,idint,idint FROM ccb_news"

	updatetable := "UPDATE myinfo1 SET age2= 1000,age = 4567 WHERE id > 0 AND id < 10"

	deletetable := "DELETE FROM myinfo1 WHERE id > 101 AND id < 200"

	DBhandler(db, creatrtable)
	time.Sleep(60 * time.Second)

	DBhandler(db, altertable)
	time.Sleep(60 * time.Second)

	DBhandler(db, inserttable)
	time.Sleep(60 * time.Second)

	DBhandler(db, updatetable)
	time.Sleep(60 * time.Second)

	DBhandler(db, deletetable)
	time.Sleep(60 * time.Second)
}

package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jonsen/gotld"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type industry struct {
	ID        int
	Name      sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

type customer struct {
	ID                int
	Name              sql.NullString
	ArtificialPerson  sql.NullString
	Contact           sql.NullString
	Province          sql.NullString
	City              sql.NullString
	District          sql.NullString
	Address           sql.NullString
	Email             sql.NullString
	URL               sql.NullString
	Domain            sql.NullString
	BusinessScope     sql.NullString
	Type              sql.NullString
	RegisteredCapital sql.NullString
	FoundTime         time.Time
	SocialCrediCode   sql.NullString
	IndustryID        sql.NullInt64
	InsuredNumber     sql.NullInt64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func main() {
	var MysqlHost string = "rdsfjnifbfjnifbo.mysql.rds.aliyuncs.com"
	var MysqlDbname string = "bigbusiness"
	var MysqLUser string = "bigbusiness"
	var MysqlPasswd string = "LiuRui123$%^"
	linkStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", MysqLUser, MysqlPasswd, MysqlHost, MysqlDbname)
	fmt.Println(linkStr)
	db, err := gorm.Open("mysql", linkStr)
	db.SingularTable(true)
	if err != nil {
		fmt.Errorf("创建数据库连接失败:%v", err)
	}
	defer db.Close()
	//准备读取文件
	str, _ := os.Getwd()
	fileName := str + "\\tianjinbigbusiness.csv"
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	/*industryMap := make(map[string]industry)
	for {
		row, err := r.Read()
		fmt.Println(row)
		if err != nil && err != io.EOF {
			//log.Fatalf("can not read, err is %+v", err)
			continue
		}
		if err == io.EOF {
			break
		}
		if row[0] == "企业名称" {
			continue
		}
		ind := row[15]
		industryMap[ind] = industry{
			Name: sql.NullString{String: ind, Valid: true},
		}
	}*/
	/*
		fmt.Println(industryMap)
		for indname := range industryMap {
			indstruct := industryMap[indname]
			var inds industry
			if err := db.Where("name = ?", indname).First(&inds).Error; err != nil {
				if inds == (industry{}) {
					db.Create(&indstruct)
				}
			}
		}
	*/
	customers := make([]customer, 0)
	canDo := true
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			//log.Fatalf("can not read, err is %+v", err)
			continue
		}
		if err == io.EOF {
			break
		}
		if row[0] == "企业名称" {
			continue
		}
		//if row[0] == "超莲知识产权服务有限公司" {
		//	canDo = true
		//	continue
		//}
		if !canDo {
			continue
		}
		var inds industry
		indname := row[15]
		if err := db.Where("name = ?", indname).First(&inds).Error; err != nil {

		}
		indid := int64(inds.ID)
		domains := subDmmain(row[8])
		domain := ""
		if len(domains) >= 1 {
			domain = domains[0]
		}
		InsuredNumber, err := strconv.ParseInt(row[20], 10, 64)
		nowUnix, _ := dateStrToUnix(row[13], "2006/1/02")
		cus := customer{
			Name:              sql.NullString{String: row[0], Valid: true},
			ArtificialPerson:  sql.NullString{String: row[1], Valid: true},
			Contact:           sql.NullString{String: row[2], Valid: true},
			Province:          sql.NullString{String: row[3], Valid: true},
			City:              sql.NullString{String: row[4], Valid: true},
			District:          sql.NullString{String: row[5], Valid: true},
			Address:           sql.NullString{String: row[6], Valid: true},
			Email:             sql.NullString{String: row[7], Valid: true},
			URL:               sql.NullString{String: row[8], Valid: true},
			Domain:            sql.NullString{String: domain, Valid: true},
			BusinessScope:     sql.NullString{String: row[9], Valid: true},
			Type:              sql.NullString{String: row[10], Valid: true},
			RegisteredCapital: sql.NullString{String: row[11], Valid: true},
			SocialCrediCode:   sql.NullString{String: row[14], Valid: true},
			IndustryID:        sql.NullInt64{Int64: indid, Valid: true},
			InsuredNumber:     sql.NullInt64{Int64: InsuredNumber, Valid: true},
			FoundTime:         time.Unix(nowUnix, 0),
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}
		db.Create(&cus)
		//customers = append(customers, cus)
		if len(customers) >= 50 {
			//fmt.Println(customers)
			//db.Create(&customers)
			customers = customers[:0]
		}
	}
	//db.Create(&customers)
}

// 截取域名
func subDmmain(url string) []string {
	if url == "" {
		return nil
	}
	pattern := "[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?"
	r, _ := regexp.Compile(pattern)
	urls := r.FindAllString(url, -1)
	domains := make([]string, 0)
	tempMap := map[string]byte{} // 存放不重复主键
	for _, value := range urls {
		domain := getUrlTldDomain(value)
		if domain == "" {
			continue
		}
		l := len(tempMap)
		tempMap[domain] = 0    //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			domains = append(domains, domain)
		}
	}
	return domains
}

// 截取url域名
func getUrlTldDomain(urls string) string {
	_, domain, err := gotld.GetTld(urls)
	if nil != err {
		fmt.Println(err)
		return ""
	}
	return domain
}

//时间转时间戳
func dateStrToUnix(timeStr, layout string) (int64, error) {
	local, err := time.LoadLocation("Asia/Shanghai") //设置时区
	if err != nil {
		return 0, err
	}
	tt, err := time.ParseInLocation(layout, timeStr, local)
	if err != nil {
		return 0, err
	}
	timeUnix := tt.Unix()
	return timeUnix, nil
}

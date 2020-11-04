package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)
var DB *gorm.DB
type DEPARTMENT struct {
	DEPT_ID   int    `gorm:"column:DEPT_ID;"`
	DEPT_NAME string `gorm:"column:DEPT_NAME;"`
	HEAD_ID   int    `gorm:"column:HEAD_ID;"`
	DEPT_HEAD string `gorm:"column:DEPT_HEAD;"`
} //Department

func main() {
	r := gin.Default()
	r.POST("/insert", insert)

	server, port, user, password, database := connectdb()
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)
	DB, err := gorm.Open("mssql", connectionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer DB.Close()


}

func connectdb() (string, int, string, string, string) {
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fetal error config file: %s", err))
	}

	server1 := viper.GetString("mssql.server")
	port1 := viper.GetInt("mssql.port")
	user1 := viper.GetString("mssql.user")
	password1 := viper.GetString("mssql.password")
	database1 := viper.GetString("mssql.database")

	return server1, port1, user1, password1, database1
} //connectdatabase


func insert(c *gin.Context) {
	var deptInfo []DEPARTMENT
	c.BindJSON(&deptInfo)
	fmt.Printf("deptInfo = %+v", deptInfo)
	err := DB.Select("DEPT_ID,DEPT_NAME,HEAD_ID,DEPT_HEAD").Table("DEPARTMENT").Create(&deptInfo).Error
	if err != nil {
		fmt.Print(err)
		panic("failed to insert database")
	}
	c.JSON(200, gin.H{
		"Status": deptInfo,
	})
} //databa


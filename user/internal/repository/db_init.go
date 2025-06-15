package repository

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=", charset + "&parseTime=true&loc=Local&timeout=10s"}, "")
	err := BuildDatabase(dsn)
	if err != nil {
		panic(err)
	}
}

func BuildDatabase(dsn string) error {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,  // 禁用用datetime的精度，mysql5.6之前数据不支持
		DontSupportRenameIndex:    true,  // 重命名索引方式采用删除后创建 mysql5.7 之前不支持
		DontSupportRenameColumn:   true,  // 用 change 重命名，mysql8 之前数据库不支持重命列
		SkipInitializeWithVersion: false, // 根据版本不能自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: ormLogger,
	})

	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(20) // 设置连接池运行空闲连接数量
	sqlDB.SetMaxOpenConns(20)    // 最大打开数量
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
	return err
}

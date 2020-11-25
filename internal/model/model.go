package model

import (
	"fmt"
	otgorm "github.com/eddycjy/opentracing-gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gotour/blog-service/global"
	"gotour/blog-service/pkg/setting"
	"time"
)

// 公共属性
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `grom:"<-:create" json:"created_by"`  // allow read and create
	ModifiedBy string `grom:"<-:update" json:"modified_by"` // allow read and update
	CreatedOn  uint32 `gorm:"<-:create,autoCreateTime" json:"created_on"`
	ModifiedOn uint32 `gorm:"<-:update,autoUpdateTIme" json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `gorm:"default=0" json:"is_del"`
}

// NewDBEngine
func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.Port,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	db, err := gorm.Open(databaseSetting.DBType, s)

	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)

	// 替换ORM框架默认的回调函数
	db.Callback().Create().Replace("gorm:update_time_stamp",
		updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp",
		updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete",
		deleteCallback)

	// 添加回调， 注册链路跟踪
	otgorm.AddGormCallbacks(db)

	return db, nil
}

func setTimeField(scope *gorm.Scope, timeField string, nowTime int64) (err error) {
	if timeField, ok := scope.FieldByName(timeField); ok {
		if timeField.IsBlank {
			err = timeField.Set(nowTime)
		}
	}
	return
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		_ = setTimeField(scope, "CreateOn", nowTime)
		_ = setTimeField(scope, "ModifyOn", nowTime)
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		_ = setTimeField(scope, "ModifyOn", nowTime)
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnFiled, hasDeletedOnField := scope.FieldByName("DeleteOn")
		isDelField, hasIsDelField := scope.FieldByName("is_del")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v, %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnFiled.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

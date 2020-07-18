package service

import (
	"errors"
	"gf-server/app/api/request"
	"gf-server/app/model/dictionaries"

	"github.com/gogf/gf/frame/g"
)

// CreateDictionary create a Dictionary
// CreateDictionary 创建Dictionary
func CreateDictionary(dict request.CreateDictionary) (err error) {
	if !dictionaries.RecordNotFound(g.Map{"type": dict.Type}) {
		return errors.New("存在相同的type，不允许创建")
	}
	if _, err = dictionaries.Insert(&dict); err != nil {
		return errors.New("创建Dictionary失败")
	}
	return
}

// DeleteDictionary delete a Dictionary
// DeleteDictionary 删除Dictionary
func DeleteDictionary(dict request.DeleteDictionary) (err error) {
	if _, err = dictionaries.Delete(g.Map{"id": dict.ID}); err != nil {
		return errors.New("删除Dictionary失败")
	}
	// TODO 删除DictionaryDetails
	return
}

// UpdateDictionary update Dictionary
// UpdateDictionary 更新 Dictionary
func UpdateDictionary(dict request.UpdateDictionary) (err error) {
	return
}

// GetDictionary get the info of a Dictionary
// GetDictionary 获取Dictionary的信息
func GetDictionary(dict request.GetDictionary) (dictReturn *dictionaries.Entity, err error) {
	return dictionaries.FindOne("id = ? OR type = ?", dict.ID, dict.Type)
}

func GetDictionaryInfoList(info request.DictionarySearch) {

}

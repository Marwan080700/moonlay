package repository

import (
	"moonlay/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindTodos() ([]models.DataList, error)
	// Todo (Datalist)
	GetDataList(ID int)(models.DataList,error)
	CreateDataList(todo models.DataList)(models.DataList, error)
	UpdateDataList(dataList models.DataList) (models.DataList, error)
	DeleteDataListByID(id int) error

	// Todo (SubDataList)
	GetAllSubDataListByDataListID(dataListID int) ([]models.DataSublist, error)
	CreateSubDataList(subTodo models.DataSublist)(models.DataSublist, error)
	GetSubDataListByID(id int) (models.DataSublist, error)
	UpdateSubDataList(SubDataList models.DataSublist) (models.DataSublist, error)
	DeleteAllSubDataListsByDataListID(dataListID int) error
	DeleteSubDataListByID(id int) error
}

type repository struct {
	db *gorm.DB
}

func RepositoryTodo(db *gorm.DB) *repository {
	return &repository{db}
}

// Find all To-Do-List

func (r *repository) FindTodos() ([]models.DataList, error){
	var todos []models.DataList
	err := r.db.Preload("Sublist").Find(&todos).Error

	return todos, err
}

// To-Do-List (datalist) ----------------------------------------------------------------------------------------

// GET To-Do-List show (datalist)
func(r *repository) GetDataList(ID int)(models.DataList,error){
	var dataList models.DataList
	err := r.db.Preload("Sublist").First(&dataList, ID).Error

	return dataList, err
}

// Create To-Do-List (datalist)
func (r *repository) CreateDataList(todo models.DataList)(models.DataList, error){
	err := r.db.Create(&todo).Error

	return todo, err
}

// update To-Do-List (datalist)
func (r *repository) UpdateDataList(dataList models.DataList) (models.DataList, error) {
    err := r.db.Preload("list").Save(&dataList).Error
    return dataList, err
}

// delete To-Do-List (datalist)
func (r *repository) DeleteDataListByID(id int) error {
    return r.db.Where("id = ?", id).Delete(&models.DataList{}).Error
}


// To-Do-List show (subdatalist) ------------------------------------------------------------------------------

func (r *repository) GetAllSubDataListByDataListID(dataListID int) ([]models.DataSublist, error) {
    var subDataList []models.DataSublist
    err := r.db.Find(&subDataList, "data_list_id = ?", dataListID).Error 

    return subDataList, err
}

// get To-Do-list (subdatalist)
func (r *repository) GetSubDataListByID(id int) (models.DataSublist, error) {
	var subDataList models.DataSublist
	err := r.db.First(&subDataList, id).Error

	return subDataList, err
}

// Create To-Do-List (subdatalist)
func (r *repository) CreateSubDataList(subTodo models.DataSublist)(models.DataSublist, error){
	err := r.db.Create(&subTodo).Error

	return subTodo, err
}

// update To-Do-List (subdatalist)
func (r *repository) UpdateSubDataList(SubDataList models.DataSublist) (models.DataSublist, error) {
    err := r.db.Save(&SubDataList).Error
    return SubDataList, err
}

//delete all subdatalist
func (r *repository) DeleteAllSubDataListsByDataListID(dataListID int) error {
    return r.db.Where("data_list_id = ?", dataListID).Delete(&models.DataSublist{}).Error
}

//detlete single subdatalist
func (r *repository) DeleteSubDataListByID(id int) error {
    return r.db.Where("id = ?", id).Delete(&models.DataSublist{}).Error
}
package user

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	GetAll(filters filter) ([]User, error)
	GetById(id string) (*User, error)
	Delete(id string) (bool, error)
	Update(id string, firstName *string, lastName *string, email *string, telefono *string) error
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (repo *repo) Create(user *User) error {
	user.ID = uuid.New().String()
	result := repo.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	repo.log.Println("Se ha creado un nuevo usuario con el ID: ", user.ID)
	return nil
}

func (repo *repo) GetAll(filsters filter) ([]User, error) {
	var users []User
	db := repo.db.Model(&users)
	db = ApplyFilters(db, filsters)
	result := db.Order("created_at desc").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (repo *repo) GetById(id string) (*User, error) {
	user := User{ID: id}
	result := repo.db.First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *repo) Delete(id string) (bool, error) {
	user := User{ID: id}
	result := repo.db.Delete(user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (repo *repo) Update(id string, firstName *string, lastName *string, email *string, telefono *string) error {
	values := make(map[string]interface{})
	if firstName != nil {
		values["first_name"] = *firstName
	}

	if lastName != nil {
		values["last_name"] = *lastName
	}

	if email != nil {
		values["email"] = *email
	}

	if telefono != nil {
		values["telefono"] = *telefono
	}

	if err := repo.db.Model(&User{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func ApplyFilters(bd *gorm.DB, filters filter) *gorm.DB {
	if filters.First_name != "" {
		filters.First_name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.First_name))
		bd = bd.Where("lower(first_name) like ?", filters.First_name)
	}

	if filters.Last_name != "" {
		filters.Last_name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Last_name))
		bd = bd.Where("lower(last_name) like ?", filters.Last_name)
	}

	return bd
}

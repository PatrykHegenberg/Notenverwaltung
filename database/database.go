package database

import (
	"log"
	"time"

	"github.com/PatrykHegenberg/Notenverwaltung/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AutoMigrate() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(
		&model.Class{},
		&model.ExamType{},
		&model.Exam{},
		&model.ExamScoreStudent{},
		&model.Grade{},
		&model.Role{},
		&model.School{},
		&model.Score{},
		&model.User{},
		&model.Student{},
		&model.Subject{},
		&model.SubjectExam{},
		&model.SubjectTeacherClass{},
		&model.Teacher{},
	)
	var adminRole model.Role
	var roles []*model.Role
	if err := db.Where("name = ?", "Admin").First(&adminRole).Error; err != nil {
		roles = []*model.Role{
			{
				Name: "Admin",
			},
			{
				Name: "Lehrer",
			},
			{
				Name: "Schueler",
			},
		}
		db.Create(&roles)
		log.Println("Admin Rolle wurde angelegt.")
	} else {
		log.Println("Admin Rolle bereits vorhande.")
	}
	var testSchool model.School
	if err := db.Where("name = ?", "TestSchule").First(&testSchool).Error; err != nil {
		testSchool = model.School{
			Name: "TestSchule",
		}
		db.Create(&testSchool)
		log.Println("Schule 'TestSchule' wurde erstellt.")
	} else {
		log.Println("Schule 'TestSchule' bereits vorhanden.")
	}
	var adminUser model.User
	if err := db.Where("username = ?", "test_admin").First(&adminUser).Error; err != nil {
		// Fehler bedeutet, dass der Benutzer nicht gefunden wurde, also erstellen wir ihn
		adminUser = model.User{
			Email:     "test_admin@example.com",
			Username:  "test_admin",
			Password:  "Password",
			CreatedAt: time.Now().String(),
			RoleID:    1,
			SchoolID:  1,
		}
		db.Create(&adminUser)
		log.Println("Benutzer 'test_admin' wurde erstellt.")
	} else {
		log.Println("Benutzer 'test_admin' ist bereits vorhanden.")
	}
	var examTypes []*model.ExamType
	var examType *model.ExamType
	if err := db.Where("name = ?", "schriftlich").First(&examType).Error; err != nil {
		examTypes = []*model.ExamType{
			{
				Name: "schriftlich",
			},
			{
				Name: "mündlich",
			},
			{
				Name: "sonstige",
			},
		}
		db.Create(examTypes)
		log.Println("ExamTypes wurden erstellt.")
	} else {
		log.Println("ExamTypes existieren bereits")
	}
	var grades = []*model.Grade{
		{
			Name: 1,
		},
		{
			Name: 2,
		},
		{
			Name: 3,
		},
		{
			Name: 4,
		},
		{
			Name: 5,
		},
		{
			Name: 6,
		},
	}
	db.Create(grades)
}

// Funktion zum Abrufen der Datenbankinstanz
func GetDBInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Established database connection")
	return db
}

func GetRoleIDByName(roleName string) (uint, error) {
	var role model.Role
	if err := GetDBInstance().Where("name = ?", roleName).First(&role).Error; err != nil {
		return 0, err
	}
	return role.ID, nil
}

func GetSchoolIDByName(schoolName string) (uint, error) {
	var school model.School
	if err := GetDBInstance().Where("name = ?", schoolName).First(&school).Error; err != nil {
		return 0, err
	}
	return school.ID, nil
}

func GetUserByName(username string) (*model.User, error) {
	var user model.User
	if err := GetDBInstance().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetRoleById(id uint) (*model.Role, error) {
	var role model.Role
	if err := GetDBInstance().Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func GetSchoolById(id uint) (*model.School, error) {
	var role model.School
	if err := GetDBInstance().Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
func GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := GetDBInstance().Preload("Role").Preload("School").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

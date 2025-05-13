package models

type UserRole string

const (
	Student          UserRole = "STUDENT"
	TeachersFavorite UserRole = "TEACHERSFAVORITE"
	Teacher          UserRole = "TEACHER"
)

type User struct {
	ID        uint     `gorm:primaryKey`
	Email     string   `gorm:uniqueIndex;not null`
	Password  string   `gorm:not null`
	FirstName string   `gorm:not null`
	LastName  string   `gorm:not null`
	Role      UserRole `gorm:type:text;not null`
}

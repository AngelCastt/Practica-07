package controler

import (
	"fmt"

	"github.com/AngelCastt/Practica-07/database"
	"gorm.io/gorm"
)

// UserControler interface define los métodos disponibles.
type UserControler interface {
	CreateUser(user database.Usuarios) error
	GetUser() ([]database.Usuarios, error)
	DeleteUser(userID uint) error
	UpdateUser(user database.Usuarios) error
}

// contiene la conexión de base de datos
type userControler struct {
	db *gorm.DB
}

// Consultar
func (i *userControler) GetUser() ([]database.Usuarios, error) {
	var users []database.Usuarios
	// Find obtiene los usuarios
	result := i.db.Find(&users)

	if result.Error != nil {
		fmt.Println("Error al obtener usuarios:", result.Error)
		return nil, result.Error
	}
	fmt.Println(users)
	return users, nil
}

// Crear
func (i *userControler) CreateUser(user database.Usuarios) error {
	result := i.db.Create(&user)
	if result.Error != nil {
		fmt.Println("Error al crear el usuario:", result.Error)
		return result.Error
	}
	return nil
}

// Eliminar
func (i *userControler) DeleteUser(ID uint) error {
	result := i.db.Delete(&database.Usuarios{}, ID)

	if result.Error != nil {
		fmt.Println("Error al eliminar el usuario:", result.Error)
		return result.Error
	}

	fmt.Println("Usuario eliminado correctamente")
	return nil
}

// Actualizar
func (i *userControler) UpdateUser(user database.Usuarios) error {
	// Buscamos el usuario por su ID y actualizamos sus datos.
	result := i.db.Model(&database.Usuarios{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		fmt.Println("Error al actualizar el usuario:", result.Error)
		return result.Error
	}

	fmt.Println("Usuario actualizado correctamente")
	return nil
}

func NewUserControler(db *gorm.DB) UserControler {
	return &userControler{db}
}

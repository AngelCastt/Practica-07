package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AngelCastt/Practica-07/controler"
	"github.com/AngelCastt/Practica-07/database"
	"github.com/gin-gonic/gin"
)

type Usuarios struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	db, err := database.NewDataBaseDriver()
	if err != nil {
		fmt.Println("Error al conectar a la base de datos")
	}
	db.AutoMigrate(&database.Usuarios{}) //crear tablas
	userControler := controler.NewUserControler(db)
	userControler.CreateUser(database.Usuarios{
		Name:  "Veronica Merari",
		Email: "merari@gmail.com",
	})
	userControler.CreateUser(database.Usuarios{
		Name: "Ivi Dominguez", Email: "ivi@gmail.com",
	})
	userControler.CreateUser(database.Usuarios{
		Name: "Angel Castrejon", Email: "angel@gmail.com",
	})
	if err != nil {
		fmt.Println("Error al crear usuario")
	}
	//consultar
	usuarios, err := userControler.GetUser()
	if err != nil {
		fmt.Println("Error al obtener usuarios: ", err)
	} else {
		fmt.Println("Lista de usuarios:", usuarios)
	}

	//Actualizar
	if len(usuarios) > 0 {
		usuarioActualizado := usuarios[2]
		usuarioActualizado.Name = "Uriel Castrejon"
		usuarioActualizado.Email = "uriel@gmail.com"
		err = userControler.UpdateUser(usuarioActualizado)
		//evaluamos errores
		if err != nil {
			fmt.Println("Error al actualizar el usuario: ", err)
		} else {
			fmt.Println("Usuario actualizado correctamente")
		}
	}

	//Borrar
	if len(usuarios) > 0 {
		err = userControler.DeleteUser(usuarios[2].ID)
		if err != nil {
			fmt.Println("Error al eliminar el usuario: ", err)
		} else {
			fmt.Println("Usuario eliminado correctamente")
		}
	}

	return

	router := gin.Default()
	users := []Usuarios{}
	indexUser := 1
	fmt.Println("Running app")
	// Tomar archivos de la carpeta template
	router.LoadHTMLGlob("templates/*")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title":       "Main website",
			"total_users": len(users),
			"users":       []Usuarios{},
		})
	})
	// API URLs
	//Obtener usuarios
	router.GET("/api/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	//Creación de usuarios
	router.POST("/api/users", func(c *gin.Context) {
		var user Usuarios
		if c.BindJSON(&user) == nil {
			user.Id = indexUser
			users = append(users, user)
			indexUser++
			c.JSON(200, user)
		} else {
			c.JSON(400, gin.H{
				"error": "Invalid payload",
			})
		}
	})
	//Eliminación de usuarios
	router.DELETE("/api/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idParsed, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid id",
			})
			return
		}
		fmt.Println("Id a borrar: ", id)
		for i, user := range users {
			if user.Id == idParsed {
				users = append(users[:i], users[i+1:]...)
				c.JSON(200, gin.H{
					"message": "User Deleted",
				})
				return
			}
		}
		c.JSON(201, gin.H{})
	})
	//ACtualizar usuarios
	router.PUT("/api/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idParsed, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid id",
			})
			return
		}
		var user Usuarios
		err = c.BindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid payload",
			})
			return
		}
		fmt.Println("Id a actualizar: ", id)
		for i, u := range users {
			if u.Id == idParsed {
				users[i] = user
				users[i].Id = idParsed
				c.JSON(200, users[i])
				return
			}
		}
		c.JSON(201, gin.H{})
	})
	router.Run(":8001")
}

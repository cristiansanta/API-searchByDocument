package main

import (
    "database/sql"
    "log"
    "SEARCHBYDOCUMENT/cmd/api/routes"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

var db *sql.DB

func main() {
    // Inicializar la conexión a la base de datos
    var err error
    db, err = sql.Open("postgres", "host=localhost port=5432 user=admin password=gestion72 dbname=db_gestion_usuarios_victimas sslmode=disable")
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }
    defer db.Close()

    // Verificar la conexión
    if err = db.Ping(); err != nil {
        log.Fatal("Error al hacer ping a la base de datos:", err)
    }

    log.Println("Conexión exitosa a la base de datos PostgreSQL")

    // Instancia de Gin
    g := gin.Default()

    // Configuración CORS simplificada
    g.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * 3600, // Cache preflight requests for 12 hours
    }))

    // Registrar rutas
    routes.RegisterRoutes(g, db)

    // Ejecutar el servidor
    if err := g.Run(":8082"); err != nil {
        log.Fatal("Error al iniciar el servidor:", err)
    }
}

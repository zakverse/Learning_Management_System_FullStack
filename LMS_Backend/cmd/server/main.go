package main

import (
	"log"

	"LMS_Backend/internal/delivery/http"
	"LMS_Backend/internal/infrastructure"	
	"LMS_Backend/internal/repository"
	"LMS_Backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Gagal load .env")
	}

	// init DB
	db, err := infrastructure.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// dependency injection
	userRepo := repository.NewUserRepository(db)
	authUC := usecase.NewAuthUsecase(userRepo)
	http.NewAuthHandler(r, authUC)

	courseRepo := repository.NewCourseRepository(db)
	courseUC := usecase.NewCourseUsecase(courseRepo)
	http.NewCourseHandler(r, courseUC)

	chapterRepo := repository.NewChapterRepository(db)
	chapterUC := usecase.NewChapterUsecase(chapterRepo)
	http.NewChapterHandler(r, chapterUC)

	r.Run(":8080")
}

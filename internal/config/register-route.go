package config

import (
	hdl "github.com/B1gdawg0/real_course_review_backend/internal/handler"
	"github.com/B1gdawg0/real_course_review_backend/internal/middleware"
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	repo "github.com/B1gdawg0/real_course_review_backend/internal/repository"
	uc "github.com/B1gdawg0/real_course_review_backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutesV1(app *fiber.App, db *gorm.DB, cfg *m.Config) {
	api := app.Group("/api/v1")

	authRepo := repo.NewAuthRepository(db)
	authUC := uc.NewAuthUsecase(authRepo, cfg.JWT_SECRET)
	authHDL := hdl.NewAuthHandler(authUC)

	auth := api.Group("/auth")
	auth.Post("", authHDL.HandleAuth)

	api.Use(middleware.JWTMiddleware(cfg.JWT_SECRET))
	// -- after this line, jwt token required --

	courseRepo := repo.NewClassRepository(db)
	courseUC := uc.NewCourseUseCase(courseRepo)
	courseHDL := hdl.NewClassHandler(courseUC)

	profRepo := repo.NewProfessorRepository(db)
	profUC := uc.NewProfessorUseCase(profRepo)
	profHDL := hdl.NewProfessorHandler(profUC)

	reviewRepo := repo.NewReviewRepository(db)
	reviewUC := uc.NewReviewUseCase(reviewRepo, courseRepo)
	reviewHDL := hdl.NewReviewHandler(reviewUC)

	tagRepo := repo.NewTagRepository(db)
	tagUC := uc.NewTagUseCase(tagRepo)
	tagHDL := hdl.NewTagHandler(tagUC)

	course := api.Group("/course")
	course.Get("", courseHDL.GetAllOrOne)

	prof := api.Group("/professor")
	prof.Get("",profHDL.GetAllOrOne)

	review := api.Group("/review")
	review.Get("/:filterType/:id", reviewHDL.GetReviews)
	review.Post("", reviewHDL.CreateReview)

	tag := api.Group("/tag")
	tag.Get("", tagHDL.GetAllOrOne)
}
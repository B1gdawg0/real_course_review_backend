package config

import (
	hdl "github.com/B1gdawg0/real_course_review_backend/internal/handler"
	repo "github.com/B1gdawg0/real_course_review_backend/internal/repository"
	uc "github.com/B1gdawg0/real_course_review_backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutesV1(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api/v1")

	courseRepo := repo.NewClassRepository(db)
	courseUC := uc.NewCourseUseCase(courseRepo)
	courseHDL := hdl.NewClassHandler(courseUC)

	reviewRepo := repo.NewReviewRepository(db)
	reviewUC := uc.NewReviewUseCase(reviewRepo)
	reviewHDL := hdl.NewReviewHandler(reviewUC)

	tagRepo := repo.NewTagRepository(db)
	tagUC := uc.NewTagUseCase(tagRepo)
	tagHDL := hdl.NewTagHandler(tagUC)

	course := api.Group("/course")
	course.Get("", courseHDL.GetAllOrOne)

	review := api.Group("/review")
	review.Get("/:filterType/:id", reviewHDL.GetReviews)
	review.Post("", reviewHDL.CreateReview)

	tag := api.Group("/tag")
	tag.Get("", tagHDL.GetAllOrOne)
}
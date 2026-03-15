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

	profRepo := repo.NewProfessorRepository(db)
	profUC := uc.NewProfessorUseCase(profRepo)
	profHDL := hdl.NewProfessorHandler(profUC)

	courseRepo := repo.NewClassRepository(db)
	reviewRepo := repo.NewReviewRepository(db)

	courseUC := uc.NewCourseUseCase(courseRepo, reviewRepo, cfg.N8N_BASE_URL)
	courseHDL := hdl.NewClassHandler(courseUC)
	
	reviewUC := uc.NewReviewUseCase(reviewRepo, courseRepo)
	reviewHDL := hdl.NewReviewHandler(reviewUC)

	tagRepo := repo.NewTagRepository(db)
	tagUC := uc.NewTagUseCase(tagRepo)
	tagHDL := hdl.NewTagHandler(tagUC)

	voteRepo := repo.NewVoteRepository(db)
	voteUC := uc.NewVoteUseCase(voteRepo, authRepo, reviewRepo)
	voteHDL := hdl.NewVoteHandler(voteUC)

	reportRepo := repo.NewReportRepository(db)
	reportUC := uc.NewReportUseCase(reportRepo)
	reportHDL := hdl.NewReportHandler(reportUC)

	vote := api.Group("/vote")
	course := api.Group("/course")
	prof := api.Group("/professor")
	review := api.Group("/review")
	tag := api.Group("/tag")
	report := api.Group("/report")

	course.Get("", courseHDL.GetCourses)
	course.Get("/compare/:first/:second", courseHDL.CompareCourseByIds)
	course.Get("/summary/:first/:second", courseHDL.GetAISummary)

	prof.Get("",profHDL.GetAllOrOne)

	review.Get("/:filterType/:id", reviewHDL.GetReviews)

	tag.Get("", tagHDL.GetAllOrOne)

	report.Post("", reportHDL.CreateReport)

	// -- after this line, jwt token required --
	api.Use(middleware.JWTMiddleware(cfg.JWT_SECRET))

	course.Put("", courseHDL.UpdateCourseRecStatus)

	vote.Post("", voteHDL.Vote)

	review.Post("", reviewHDL.CreateReview)
	
	report.Get("", reportHDL.GetAll)
}
package config

import (
	"fmt"
	"math/rand"
	"time"

	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// --- SystemConfig ---
	var configCount int64
	db.Model(&m.SystemConfig{}).Count(&configCount)
	if configCount == 0 {
		config := []m.SystemConfig{
			{Key: "m", Value: "0"},
			{Key: "BAYMN", Value: "5"},
			{Key: "max_reviews_per_user", Value: "10"},
			{Key: "review_cooldown_hours", Value: "24"},
		}
		db.Create(&config)
	}

	// --- Users ---
	var userCount int64
	db.Model(&m.User{}).Count(&userCount)
	if userCount == 0 {
		users := []m.User{
			{
				Name:    "Alice Johnson",
				Email:   "alice.johnson@student.edu",
				Role:    "USER",
				UniYear: "3rd Year",
			},
			{
				Name:    "Bob Chen",
				Email:   "bob.chen@student.edu",
				Role:    "USER",
				UniYear: "2nd Year",
			},
		}
		db.Create(&users)
	}

	// --- Professors ---
	var profCount int64
	db.Model(&m.Professor{}).Count(&profCount)
	if profCount == 0 {
		professors := []m.Professor{
			{
				Name:           "Dr. John Smith",
				Email:          "john.smith@university.edu",
				Title:          "Associate Professor",
				ImageURL:       "https://example.com/john.jpg",
				Phone:          "123-456-7890",
				Description:    "Specialist in Database Systems.",
				UniRoomAddress: "Room 301, CS Building",
			},
			{
				Name:           "Dr. Emily Johnson",
				Email:          "emily.johnson@university.edu",
				Title:          "Assistant Professor",
				ImageURL:       "https://example.com/emily.jpg",
				Phone:          "987-654-3210",
				Description:    "Expert in Machine Learning and AI.",
				UniRoomAddress: "Room 405, AI Building",
			},
			{
				Name:           "Dr. Michael Brown",
				Email:          "michael.brown@university.edu",
				Title:          "Professor",
				ImageURL:       "https://example.com/michael.jpg",
				Phone:          "555-123-4567",
				Description:    "Software Engineering and System Design expert.",
				UniRoomAddress: "Room 203, Engineering Building",
			},
		}
		db.Create(&professors)
	}

	// --- Courses ---
	var courseCount int64
	db.Model(&m.Course{}).Count(&courseCount)
	if courseCount == 0 {
		// fetch professors
		var john, emily, michael m.Professor
		db.Where("email = ?", "john.smith@university.edu").First(&john)
		db.Where("email = ?", "emily.johnson@university.edu").First(&emily)
		db.Where("email = ?", "michael.brown@university.edu").First(&michael)

		courses := []m.Course{
			{
				Name:        "Database Systems",
				Description: "Advanced concepts in relational and NoSQL databases.",
				Semester:    "Fall 2025",
				Code:        "CS303",
				Credit:      3,
				ReviewCount: 0,
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Introduction to Machine Learning",
				Description: "Supervised and unsupervised learning algorithms.",
				Semester:    "Spring 2026",
				Code:        "CS410",
				Credit:      4,
				ReviewCount: 0,
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Software Engineering",
				Description: "Principles and practices of software development.",
				Semester:    "Fall 2025",
				Code:        "CS350",
				Credit:      3,
				ReviewCount: 0,
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Data Structures and Algorithms",
				Description: "Fundamental data structures and algorithm analysis.",
				Semester:    "Spring 2026",
				Code:        "CS201",
				Credit:      4,
				ReviewCount: 0,
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john, michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Computer Networks",
				Description: "Network protocols, architecture, and security.",
				Semester:    "Fall 2025",
				Code:        "CS420",
				Credit:      3,
				ReviewCount: 0,
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily, michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}
		db.Create(&courses)
	}

	// --- Reviews ---
	var reviewCount int64
	db.Model(&m.Review{}).Count(&reviewCount)
	if reviewCount == 0 {
		// fetch users and courses
		var users []m.User
		var courses []m.Course
		db.Find(&users)
		db.Find(&courses)

		// Create 10 reviews
		reviews := []m.Review{
			{
				UserID:      users[0].ID, // Alice
				CourseID:    courses[0].ID, // Database Systems
				Status:      "publish",
				Description: "Great course! The professor explains complex database concepts very clearly. The assignments are challenging but fair.",
				AvgRate:     4.5,
				Rate:        "4.0,5.0,4.0,5.0", // difficulty, workload, teaching, overall
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[1].ID, // Bob
				CourseID:    courses[0].ID, // Database Systems
				Status:      "publish",
				Description: "The material is interesting but the pace is quite fast. Make sure to keep up with readings.",
				AvgRate:     3.5,
				Rate:        "4.0,4.0,3.0,3.0",
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[0].ID, // Alice
				CourseID:    courses[1].ID, // Machine Learning
				Status:      "publish",
				Description: "Excellent introduction to ML concepts. Dr. Johnson is very knowledgeable and helpful during office hours.",
				AvgRate:     4.8,
				Rate:        "3.0,4.0,5.0,5.0",
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[1].ID, // Bob
				CourseID:    courses[1].ID, // Machine Learning
				Status:      "publish",
				Description: "Very math-heavy course. Be prepared to spend a lot of time on assignments. Worth it though!",
				AvgRate:     4.0,
				Rate:        "5.0,5.0,3.0,4.0",
				Score:       0.0,
				IsAnonymous: true,
			},
			{
				UserID:      users[0].ID, // Alice
				CourseID:    courses[2].ID, // Software Engineering
				Status:      "publish",
				Description: "Practical course with real-world applications. The group project was particularly valuable.",
				AvgRate:     4.3,
				Rate:        "3.0,3.0,5.0,5.0",
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[1].ID, // Bob
				CourseID:    courses[2].ID, // Software Engineering
				Status:      "publish",
				Description: "Good course overall but the workload can be overwhelming during project weeks.",
				AvgRate:     3.8,
				Rate:        "4.0,5.0,3.0,4.0",
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[0].ID, // Alice
				CourseID:    courses[3].ID, // Data Structures
				Status:      "publish",
				Description: "Fundamental course for CS students. The concepts are essential and well-taught.",
				AvgRate:     4.5,
				Rate:        "4.0,4.0,5.0,4.0",
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[1].ID, // Bob
				CourseID:    courses[3].ID, // Data Structures
				Status:      "publish",
				Description: "Challenging but rewarding. Make sure to practice coding problems regularly.",
				AvgRate:     4.0,
				Rate:        "5.0,4.0,4.0,3.0",
				Score:       0.0,
				IsAnonymous: true,
			},
			{
				UserID:      users[0].ID, // Alice
				CourseID:    courses[4].ID, // Computer Networks
				Status:      "publish",
				Description: "Interesting topic but could use more hands-on labs. Theory is solid though.",
				AvgRate:     3.5,
				Rate:        "3.0,3.0,4.0,3.0",
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[1].ID, // Bob
				CourseID:    courses[4].ID, // Computer Networks
				Status:      "publish",
				Description: "Great coverage of network protocols. The final project was very educational.",
				AvgRate:     4.2,
				Rate:        "4.0,4.0,4.0,5.0",
				Score:       0.0,
				IsAnonymous: false,
			},
		}
		
		result := db.Create(&reviews)
		if result.Error != nil {
			fmt.Printf("Error creating reviews: %v\n", result.Error)
		} else {
			fmt.Printf("Created %d reviews\n", len(reviews))
		}
	}

	// --- Votes ---
	var voteCount int64
	db.Model(&m.Vote{}).Count(&voteCount)
	if voteCount == 0 {
		// fetch users and reviews for votes
		var users []m.User
		var reviews []m.Review
		db.Find(&users)
		db.Find(&reviews)

		rand.Seed(time.Now().UnixNano())
		
		var votes []m.Vote
		
		// Create 2-3 votes for each review
		for _, review := range reviews {
			numVotes := 2 + rand.Intn(2) // 2 or 3 votes
			
			for i := 0; i < numVotes; i++ {
				userIndex := rand.Intn(len(users))
				voteValue := 1
				if rand.Float32() < 0.3 { // 30% chance of downvote
					voteValue = -1
				}
				
				vote := m.Vote{
					UserID:   users[userIndex].ID,
					ReviewID: review.ID,
					Vote:     voteValue,
				}
				votes = append(votes, vote)
			}
		}
		
		result := db.Create(&votes)
		if result.Error != nil {
			fmt.Printf("Error creating votes: %v\n", result.Error)
		} else {
			fmt.Printf("Created %d votes\n", len(votes))
		}
		
		// Update vote counts in reviews
		for _, review := range reviews {
			var upCount, downCount int64
			db.Model(&m.Vote{}).Where("review_id = ? AND vote = ?", review.ID, 1).Count(&upCount)
			db.Model(&m.Vote{}).Where("review_id = ? AND vote = ?", review.ID, -1).Count(&downCount)
			
			db.Model(&review).Updates(map[string]interface{}{
				"up_count":   upCount,
				"down_count": downCount,
			})
		}
		
		// Update course review counts and rates
		var allCourses []m.Course
		db.Find(&allCourses)
		for _, course := range allCourses {
			var courseReviews []m.Review
			db.Where("course_id = ?", course.ID).Find(&courseReviews)
			
			if len(courseReviews) > 0 {
				var rates []string
				var totalScore float64
				
				for _, review := range courseReviews {
					rates = append(rates, review.Rate)
					totalScore += review.AvgRate
				}
				
				avgScore := totalScore / float64(len(courseReviews))
				rateString := ""
				for i, rate := range rates {
					if i > 0 {
						rateString += ","
					}
					rateString += rate
				}
				
				db.Model(&course).Updates(map[string]interface{}{
					"review_count": len(courseReviews),
					"rate":         rateString,
					"score":        avgScore,
				})
			}
		}
	}
}
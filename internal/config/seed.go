package config

import (
	"fmt"
	"log"
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
				Name: "Database Systems",
				Description: "Advanced concepts in relational and NoSQL databases.",
				Semester: "Fall 2025",
				Code: "CS303",
				Credit: 3,
				ReviewCount: 0,
				Rate: "0",
				Score: 0.0,
				Professors: []m.Professor{john},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Name: "Introduction to Machine Learning",
				Description: "Supervised and unsupervised learning algorithms.",
				Semester: "Spring 2026",
				Code: "CS410",
				Credit: 4,
				ReviewCount: 0,
				Rate: "0",
				Score: 0.0,
				Professors: []m.Professor{emily},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Name: "Software Engineering",
				Description: "Principles and practices of software development.",
				Semester: "Fall 2025",
				Code: "CS350",
				Credit: 3,
				ReviewCount: 0,
				Rate: "0",
				Score: 0.0,
				Professors: []m.Professor{michael},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Name: "Data Structures and Algorithms",
				Description: "Fundamental data structures and algorithm analysis.",
				Semester: "Spring 2026",
				Code: "CS201",
				Credit: 4,
				ReviewCount: 0,
				Rate: "0",
				Score: 0.0,
				Professors: []m.Professor{john, michael},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Name: "Computer Networks",
				Description: "Network protocols, architecture, and security.",
				Semester: "Fall 2025",
				Code: "CS420",
				Credit: 3,
				ReviewCount: 0,
				Rate: "0",
				Score: 0.0,
				Professors: []m.Professor{emily, michael},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		db.Create(&courses)
	}

	// --- Centralized Tags ---
	var tagCount int64
	db.Model(&m.Tag{}).Count(&tagCount)
	if tagCount == 0 {
		// Create unique tags for both courses and reviews
		tagNames := []string{
			// Course/Professor related tags
			"Fun", "Hard", "Helpful", "Strict", "Knowledgeable", "Friendly", 
			"Clear", "Inspiring", "Fair", "Engaging",
			// Review specific tags
			"Easy", "Difficult", "Boring", "Interesting", "Time-consuming", 
			"Rewarding", "Practical", "Theoretical", "Group Work", "Individual",
			"Good Textbook", "Poor Textbook", "Online Resources", "Attendance Required",
			"Pop Quizzes", "Final Project", "Midterm Heavy", "Participation Matters",
			"Math Heavy", "Coding Intensive", "Writing Intensive", "Lab Work",
			"Fast Paced", "Well Organized", "Confusing", "Outdated Material",
			"Fundamental", "Challenging", // Added missing tags
		}
		
		var tags []m.Tag
		for _, tagName := range tagNames {
			tags = append(tags, m.Tag{
				Name: tagName,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})
		}
		
		// Create all unique tags
		db.Create(&tags)
	}

	// --- Create Course-Tag Junction Records ---
	var courseTagCount int64
	db.Model(&m.CourseTag{}).Count(&courseTagCount)
	if courseTagCount == 0 {
		// Get all courses and all tags
		var courses []m.Course
		var tags []m.Tag
		db.Find(&courses)
		db.Find(&tags)
		
		if len(courses) > 0 && len(tags) > 0 {
			// Create course-tag relationships using junction table
			for _, course := range courses {
				// Filter tags appropriate for courses (professor/course characteristics)
				courseAppropriateTagNames := []string{
					"Fun", "Hard", "Helpful", "Strict", "Knowledgeable", 
					"Friendly", "Clear", "Inspiring", "Fair", "Engaging",
				}
				
				var courseAppropriateTags []m.Tag
				for _, tag := range tags {
					for _, appropriateName := range courseAppropriateTagNames {
						if tag.Name == appropriateName {
							courseAppropriateTags = append(courseAppropriateTags, tag)
							break
						}
					}
				}
				
				// Shuffle and assign 1-3 random appropriate tags
				shuffledTags := make([]m.Tag, len(courseAppropriateTags))
				copy(shuffledTags, courseAppropriateTags)
				rand.Shuffle(len(shuffledTags), func(i, j int) { 
					shuffledTags[i], shuffledTags[j] = shuffledTags[j], shuffledTags[i] 
				})
				
				numTags := 1 + rand.Intn(3)
				if numTags > len(shuffledTags) {
					numTags = len(shuffledTags)
				}
				selectedTags := shuffledTags[:numTags]
				
				// Create CourseTag junction records directly
				for _, tag := range selectedTags {
					courseTag := m.CourseTag{
						CourseID: course.ID,
						TagID:    tag.ID,
					}
					result := db.Create(&courseTag)
					if result.Error != nil {
						log.Printf("Error creating course-tag relationship for course %s and tag %s: %v", 
							course.Name, tag.Name, result.Error)
					}
				}
			}
		}
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

	// --- Create Review-Tag Junction Records ---
	var reviewTagCount int64
	db.Model(&m.ReviewTag{}).Count(&reviewTagCount)
	if reviewTagCount == 0 {
		// Get all reviews and all tags
		var reviews []m.Review
		var tags []m.Tag
		db.Find(&reviews)
		db.Find(&tags)
		
		if len(reviews) > 0 && len(tags) > 0 {
			// Define review-appropriate tag mappings based on review content
			reviewTagMappings := map[int][]string{
				0: {"Clear", "Fair", "Practical"},                    // Database Systems - Alice
				1: {"Fast Paced", "Time-consuming", "Interesting"},   // Database Systems - Bob
				2: {"Helpful", "Knowledgeable", "Rewarding"},         // Machine Learning - Alice  
				3: {"Math Heavy", "Time-consuming", "Difficult"},     // Machine Learning - Bob
				4: {"Practical", "Group Work", "Rewarding"},          // Software Engineering - Alice
				5: {"Time-consuming", "Group Work", "Difficult"},     // Software Engineering - Bob
				6: {"Fundamental", "Well Organized", "Clear"},        // Data Structures - Alice
				7: {"Challenging", "Coding Intensive", "Rewarding"},  // Data Structures - Bob
				8: {"Theoretical", "Boring", "Outdated Material"},    // Computer Networks - Alice
				9: {"Final Project", "Rewarding", "Well Organized"},  // Computer Networks - Bob
			}
			
			// Create review-tag relationships using junction table
			for i, review := range reviews {
				var selectedTagNames []string
				
				if tagNames, exists := reviewTagMappings[i]; exists {
					selectedTagNames = tagNames
				} else {
					// If specific tags not found, assign random appropriate ones
					reviewAppropriateTagNames := []string{
						"Easy", "Difficult", "Interesting", "Time-consuming", "Rewarding",
						"Practical", "Theoretical", "Group Work", "Individual", "Final Project",
						"Math Heavy", "Coding Intensive", "Fast Paced", "Well Organized",
					}
					
					// Shuffle and select 2-3 random tags
					shuffled := make([]string, len(reviewAppropriateTagNames))
					copy(shuffled, reviewAppropriateTagNames)
					rand.Shuffle(len(shuffled), func(i, j int) {
						shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
					})
					
					numTags := 2 + rand.Intn(2) // 2-3 tags
					if numTags > len(shuffled) {
						numTags = len(shuffled)
					}
					selectedTagNames = shuffled[:numTags]
				}
				
				// Find matching tags and create ReviewTag junction records
				for _, tagName := range selectedTagNames {
					for _, tag := range tags {
						if tag.Name == tagName {
							reviewTag := m.ReviewTag{
								ReviewID: review.ID,
								TagID:    tag.ID,
							}
							result := db.Create(&reviewTag)
							if result.Error != nil {
								log.Printf("Error creating review-tag relationship for review %d and tag %s: %v", 
									i, tag.Name, result.Error)
							}
							break
						}
					}
				}
			}
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
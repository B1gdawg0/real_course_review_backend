package config

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	"github.com/B1gdawg0/real_course_review_backend/internal/utils"
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
			// General Education - Citizenship
			{
				Name:        "Active Citizenship",
				Description: "Thai and Global Citizenship development course.",
				Semester:    "1/2025",
				Code:        "01999111",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Core Courses
			{
				Name:        "Calculus I",
				Description: "Introduction to differential and integral calculus.",
				Semester:    "1/2025",
				Code:        "01417111",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Probability and Statistics for Computer Science",
				Description: "Probability theory and statistical methods for CS.",
				Semester:    "2/2025",
				Code:        "01417322",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Calculus II",
				Description: "Continuation of Calculus I with advanced topics.",
				Semester:    "2/2025",
				Code:        "01418131",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Fundamentals of Computing",
				Description: "Introduction to computer systems and programming.",
				Semester:    "1/2025",
				Code:        "01418132",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Required - Software Technology
			{
				Name:        "Computer Programming I",
				Description: "Introduction to programming using Python.",
				Semester:    "1/2025",
				Code:        "01418112",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Computer Programming II",
				Description: "Object-oriented programming and data structures.",
				Semester:    "2/2025",
				Code:        "01418113",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Software Construction",
				Description: "Software design patterns and construction techniques.",
				Semester:    "1/2025",
				Code:        "01418211",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Data Structures",
				Description: "Fundamental data structures and algorithms.",
				Semester:    "1/2025",
				Code:        "01418231",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Algorithm Design and Analysis",
				Description: "Design and analysis of efficient algorithms.",
				Semester:    "2/2025",
				Code:        "01418232",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Required - System Infrastructure
			{
				Name:        "Discrete Mathematics",
				Description: "Mathematical foundations for computer science.",
				Semester:    "1/2025",
				Code:        "01418111",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Operating Systems",
				Description: "Principles and design of operating systems.",
				Semester:    "1/2025",
				Code:        "01418331",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "System Programming",
				Description: "Low-level programming and system calls.",
				Semester:    "2/2025",
				Code:        "01418332",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Computer Networks I",
				Description: "Network protocols and architecture.",
				Semester:    "1/2025",
				Code:        "01418351",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Required - Hardware and Architecture
			{
				Name:        "Computer Architecture",
				Description: "Computer organization and architecture principles.",
				Semester:    "2/2025",
				Code:        "01418233",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Required - Application Technology
			{
				Name:        "Database Systems",
				Description: "Database design and management systems.",
				Semester:    "1/2025",
				Code:        "01418221",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Intelligent Systems",
				Description: "Artificial intelligence and intelligent systems.",
				Semester:    "2/2025",
				Code:        "01418321",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Cooperative Education Preparation",
				Description: "Preparation for cooperative education.",
				Semester:    "1/2025",
				Code:        "01418390",
				Credit:      1,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Cooperative Education",
				Description: "Work experience in industry.",
				Semester:    "Summer/2025",
				Code:        "01418490",
				Credit:      6,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john, emily, michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Special Topics in Computer Science",
				Description: "Selected advanced topics in computer science.",
				Semester:    "2/2025",
				Code:        "01418497",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Senior Project",
				Description: "Capstone project for senior students.",
				Semester:    "2/2025",
				Code:        "01418499",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john, emily, michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Elective - Software Development
			{
				Name:        "Software Engineering Principles",
				Description: "Principles and practices of software engineering.",
				Semester:    "1/2025",
				Code:        "01418311",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Compiler Design",
				Description: "Theory and implementation of compilers.",
				Semester:    "2/2025",
				Code:        "01418471",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Elective - Database and Information Systems
			{
				Name:        "Database Programming",
				Description: "Advanced database programming techniques.",
				Semester:    "2/2025",
				Code:        "01418222",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Database System Implementation",
				Description: "Internal implementation of database systems.",
				Semester:    "1/2025",
				Code:        "01418322",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Information Retrieval",
				Description: "Techniques for information retrieval and search.",
				Semester:    "2/2025",
				Code:        "01418341",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Elective - Network and Systems
			{
				Name:        "Computer Networks II",
				Description: "Advanced networking concepts and protocols.",
				Semester:    "2/2025",
				Code:        "01418352",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Network Programming",
				Description: "Programming for networked applications.",
				Semester:    "1/2025",
				Code:        "01418353",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Computer Security",
				Description: "Security principles and cryptography.",
				Semester:    "2/2025",
				Code:        "01418451",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Elective - AI and Machine Learning
			{
				Name:        "Introduction to Data Science",
				Description: "Data analysis and visualization techniques.",
				Semester:    "1/2025",
				Code:        "01418282",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Artificial Intelligence",
				Description: "AI techniques and applications.",
				Semester:    "1/2025",
				Code:        "01418361",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Machine Learning",
				Description: "Machine learning algorithms and applications.",
				Semester:    "2/2025",
				Code:        "01418362",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Deep Learning",
				Description: "Neural networks and deep learning techniques.",
				Semester:    "1/2025",
				Code:        "01418363",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{emily},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Elective - Computer Graphics and Multimedia
			{
				Name:        "Multimedia Systems",
				Description: "Multimedia data processing and systems.",
				Semester:    "1/2025",
				Code:        "01418281",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Computer Graphics I",
				Description: "Fundamentals of computer graphics.",
				Semester:    "1/2025",
				Code:        "01418381",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Elective - Web and Applications
			{
				Name:        "Web Application Development",
				Description: "Modern web application development.",
				Semester:    "1/2025",
				Code:        "01418342",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Mobile Application Development",
				Description: "iOS and Android app development.",
				Semester:    "2/2025",
				Code:        "01418421",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},

			// Elective for Non-CS Students
			{
				Name:        "Introduction to Computers and Programming",
				Description: "Basic computer literacy and programming.",
				Semester:    "1/2025",
				Code:        "01418101",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{john},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Name:        "Computer and Information Literacy",
				Description: "Information technology and digital literacy.",
				Semester:    "1/2025",
				Code:        "01418102",
				Credit:      3,
				ReviewCount: "0,0,0",
				Rate:        "0",
				Score:       0.0,
				Professors:  []m.Professor{michael},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
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
				Name:      tagName,
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
			// --- 01418112: Computer Programming I (courses[0]) ---
			{
				UserID:      users[0].ID,
				CourseID:    courses[0].ID,
				Status:      "publish",
				Description: "วิชานี้เป็นพื้นฐานที่สำคัญมาก อาจารย์สอนเข้าใจง่าย แต่ต้องฝึกเขียนโค้ดเองเยอะๆ ไม่งั้นทำข้อสอบไม่ได้",
				AvgRate:     5.0,
				Rate:        "4.0,4.0,5.0,5.0",
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[1].ID,
				CourseID:    courses[0].ID,
				Status:      "publish",
				Description: "Hard for beginners. The logic is confusing at first. You need to practice Python every day.",
				AvgRate:     3.0,
				Rate:        "5.0,5.0,3.0,3.0", // Hard difficulty
				Score:       0.0,
				IsAnonymous: true,
			},
			{
				UserID:      users[0].ID,
				CourseID:    courses[0].ID,
				Status:      "publish",
				Description: "แลปโหดมาก ตัดเกรดอิงกลุ่ม คนเก่งเยอะ ต้องขยันสุดๆ",
				AvgRate:     4.0,
				Rate:        "5.0,5.0,4.0,4.0", // High workload
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[1].ID,
				CourseID:    courses[0].ID,
				Status:      "publish",
				Description: "Good intro to CS. If you pass this with an A, you are ready for the rest.",
				AvgRate:     5.0,
				Rate:        "3.0,3.0,5.0,5.0",
				Score:       0.0,
				IsAnonymous: false,
			},

			// --- 01418113: Computer Programming II (courses[1]) ---
			{
				UserID:      users[0].ID,
				CourseID:    courses[1].ID,
				Status:      "publish",
				Description: "เรียน Java/OOP สนุกดี แต่เริ่มยากตรงเรื่อง Class/Object ใครพื้นฐาน 112 ไม่แน่น เหนื่อยแน่นอน",
				AvgRate:     4.0,
				Rate:        "4.0,4.0,4.0,4.0",
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[1].ID,
				CourseID:    courses[1].ID,
				Status:      "publish",
				Description: "More complex than Prog I. Pointers and memory management in C part are tricky.",
				AvgRate:     3.0,
				Rate:        "5.0,4.0,3.0,3.0",
				Score:       0.0,
				IsAnonymous: true,
			},

			// --- 01418232: Algorithm Design (courses[2]) ---
			{
				UserID:      users[0].ID,
				CourseID:    courses[2].ID,
				Status:      "publish",
				Description: "The most difficult math-heavy class. Dynamic Programming makes me cry.",
				AvgRate:     2.0,
				Rate:        "5.0,5.0,2.0,2.0", // Very Hard
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[1].ID,
				CourseID:    courses[2].ID,
				Status:      "publish",
				Description: "ยากมากกกก เนื้อหาเยอะสุดๆ ต้องเข้าใจ Big O ให้แม่นๆ ข้อสอบเขียนมือจนเมื่อย",
				AvgRate:     3.0,
				Rate:        "5.0,5.0,4.0,3.0",
				Score:       0.0,
				IsAnonymous: true,
			},
			{
				UserID:      users[0].ID,
				CourseID:    courses[2].ID,
				Status:      "publish",
				Description: "Essential for job interviews. Pay attention to graph algorithms.",
				AvgRate:     5.0,
				Rate:        "4.0,4.0,5.0,5.0",
				Score:       0.0,
				IsAnonymous: false,
			},

			// --- 01418331: Operating Systems (courses[3]) ---
			{
				UserID:      users[1].ID,
				CourseID:    courses[3].ID,
				Status:      "publish",
				Description: "โปรเจคเขียน OS จำลองคือตำนาน อดนอน 3 คืนติดเพื่อแก้ Bug ตัวเดียว",
				AvgRate:     4.0,
				Rate:        "5.0,5.0,4.0,4.0", // Extreme Workload
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[0].ID,
				CourseID:    courses[3].ID,
				Status:      "publish",
				Description: "Concurrency, Semaphores, Deadlocks. Very abstract concepts but cool when you understand them.",
				AvgRate:     4.0,
				Rate:        "5.0,3.0,4.0,4.0",
				Score:       0.0,
				IsAnonymous: true,
			},
			{
				UserID:      users[1].ID,
				CourseID:    courses[3].ID,
				Status:      "publish",
				Description: "วิชาปราบเซียน ใครผ่านวิชานี้ไปได้คือจบปี 3 อย่างภาคภูมิใจ",
				AvgRate:     5.0,
				Rate:        "5.0,5.0,5.0,5.0",
				Score:       0.0,
				IsAnonymous: false,
			},

			// --- 01418342: Web Application Development (courses[4]) ---
			{
				UserID:      users[0].ID,
				CourseID:    courses[4].ID,
				Status:      "publish",
				Description: "Practical and fun. We built a full-stack MERN app. Portfolio ready!",
				AvgRate:     5.0,
				Rate:        "3.0,4.0,5.0,5.0",
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[1].ID,
				CourseID:    courses[4].ID,
				Status:      "publish",
				Description: "งานเยอะมากกก งานกลุ่มต้องเลือกเพื่อนดีๆ ไม่งั้นแบกหลังหัก",
				AvgRate:     3.0,
				Rate:        "3.0,5.0,3.0,3.0", // High Workload
				Score:       0.0,
				IsAnonymous: true,
			},
			{
				UserID:      users[0].ID,
				CourseID:    courses[4].ID,
				Status:      "publish",
				Description: "Best elective if you want to be a web dev. Teaches React, Node, and deployment.",
				AvgRate:     5.0,
				Rate:        "3.0,3.0,5.0,5.0",
				Score:       0.0,
				IsAnonymous: false,
			},

			// --- 01999111: Thai and Global Citizenship (GenEd) (courses[5]) ---
			{
				UserID:      users[1].ID,
				CourseID:    courses[5].ID,
				Status:      "publish",
				Description: "Easy A. Just attend the class and submit the group video project.",
				AvgRate:     5.0,
				Rate:        "1.0,2.0,5.0,5.0", // Very Easy
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[0].ID,
				CourseID:    courses[5].ID,
				Status:      "publish",
				Description: "น่าเบื่อหน่อยๆ นั่งฟังบรรยายยาวๆ แต่เกรดสวย แนะนำให้ลงเก็บเกรด",
				AvgRate:     4.0,
				Rate:        "1.0,1.0,3.0,4.0",
				Score:       0.0,
				IsAnonymous: true,
			},

			// --- 01418499: Senior Project (courses[6]) ---
			{
				UserID:      users[1].ID,
				CourseID:    courses[6].ID,
				Status:      "publish",
				Description: "Depends heavily on your advisor. Choose wisely or you will suffer.",
				AvgRate:     3.0,
				Rate:        "5.0,5.0,3.0,3.0",
				Score:       0.0,
				IsAnonymous: false,
			},
			{
				UserID:      users[0].ID,
				CourseID:    courses[6].ID,
				Status:      "publish",
				Description: "เครียดมาก ต้องแบ่งเวลาดีๆ อย่าดองงาน ทำเล่มวิจัยเหนื่อยกว่าเขียนโค้ดอีก",
				AvgRate:     3.0,
				Rate:        "5.0,5.0,3.0,3.0",
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
				0: {"Clear", "Fair", "Practical"},                   // Database Systems - Alice
				1: {"Fast Paced", "Time-consuming", "Interesting"},  // Database Systems - Bob
				2: {"Helpful", "Knowledgeable", "Rewarding"},        // Machine Learning - Alice
				3: {"Math Heavy", "Time-consuming", "Difficult"},    // Machine Learning - Bob
				4: {"Practical", "Group Work", "Rewarding"},         // Software Engineering - Alice
				5: {"Time-consuming", "Group Work", "Difficult"},    // Software Engineering - Bob
				6: {"Fundamental", "Well Organized", "Clear"},       // Data Structures - Alice
				7: {"Challenging", "Coding Intensive", "Rewarding"}, // Data Structures - Bob
				8: {"Theoretical", "Boring", "Outdated Material"},   // Computer Networks - Alice
				9: {"Final Project", "Rewarding", "Well Organized"}, // Computer Networks - Bob
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

			if len(courseReviews) == 0 {
				continue
			}

			reviewCount := dtos.ReviewCountResponse{Happiness: 0, Easiness: 0, Quality: 0}
			totalRate := dtos.RatingResponse{Happiness: 0, Easiness: 0, Quality: 0}
			totalScore := 0.0

			for _, review := range courseReviews {
				reviewCount.Happiness++
				reviewCount.Easiness++
				reviewCount.Quality++

				rate, _ := utils.ParseRate(review.Rate)
				totalRate.Happiness += rate.Happiness
				totalRate.Easiness += rate.Easiness
				totalRate.Quality += rate.Quality

				totalScore += review.AvgRate
			}

			numReviews := float64(len(courseReviews))
			avgRate := dtos.RatingResponse{
				Happiness: totalRate.Happiness / numReviews,
				Easiness:  totalRate.Easiness / numReviews,
				Quality:   totalRate.Quality / numReviews,
			}

			// Update course
			db.Model(&course).Updates(map[string]interface{}{
				"review_count": utils.ReviewCountToString(reviewCount),
				"rate":         utils.RateToString(avgRate),
				"score":        totalScore / numReviews,
			})
		}

	}
}

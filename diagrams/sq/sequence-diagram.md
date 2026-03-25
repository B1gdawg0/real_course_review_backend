# Sequence Diagrams - Real Course Review Backend

## 1. Auth Flow - Login/Register
**หน้าบ้าน:** Login Page

```mermaid
sequenceDiagram
autonumber
actor User as ผู้ใช้งาน (User)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as AuthHandler
participant UseCase as AuthUseCase
participant Repo as AuthRepository
participant DB as ฐานข้อมูล (Database)

    User->>UI: เข้าสู่หน้า Login & กรอก Email (@ku.th)
    UI->>Handler: POST /api/v1/auth
    Handler->>UseCase: authUsecase.LoginOrRegister()
    UseCase->>Repo: authRepo.GetUserByEmail()
    Repo->>DB: SELECT * FROM users WHERE email = ?
    DB-->>Repo: Return user หรือ Not Found
    alt User ไม่มีในระบบ
        Repo-->>UseCase: User Not Found
        UseCase->>Repo: authRepo.CreateUser()
        Repo->>DB: INSERT INTO users (name, email, role)
        DB-->>Repo: User Created
        Repo-->>UseCase: Return New User
    else User มีในระบบ
        Repo-->>UseCase: Return Existing User
    end
    UseCase->>UseCase: Generate JWT Token
    UseCase-->>Handler: Return AuthResponse (Token + User)
    Handler-->>UI: Return JSON (200 OK)
    UI-->>User: Redirect ไปหน้า Home พร้อม Token
```

---

## 2. Course Flow - Get All Courses
**หน้าบ้าน:** Home Page

```mermaid
sequenceDiagram
autonumber
actor User as ผู้ใช้งาน (User)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as CourseHandler
participant UseCase as CourseUseCase
participant Repo as CourseRepository
participant DB as ฐานข้อมูล (Database)

    User->>UI: เข้าสู่หน้า Home
    UI->>Handler: GET /api/v1/course?page=1&size=10<br/>(with JWT Token)
    Handler->>UseCase: courseUseCase.GetAll(page, size)
    UseCase->>Repo: courseRepo.GetAll()
    Repo->>DB: SELECT * FROM courses<br/>LEFT JOIN course_tags<br/>LEFT JOIN tags
    DB-->>Repo: Return Course List []
    Repo-->>UseCase: Return Course Entities
    UseCase->>UseCase: Calculate Rate & Score<br/>Sort by Score
    UseCase-->>Handler: Return CourseShortResponse[]<br/>+ Pagination Meta
    Handler-->>UI: Return JSON (200 OK)
    UI-->>User: แสดงรายวิชาทั้งหมด (พร้อม Tags, Rate)
```

---

## 3. Course Flow - Search Courses
**หน้าบ้าน:** Home Page

```mermaid
sequenceDiagram
autonumber
actor User as ผู้ใช้งาน (User)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as CourseHandler
participant UseCase as CourseUseCase
participant Repo as CourseRepository
participant DB as ฐานข้อมูล (Database)

    User->>UI: พิมพ์คำค้นหาในช่อง Search
    UI->>Handler: GET /api/v1/course?q=keyword&page=1&size=10<br/>(with JWT Token)
    Handler->>UseCase: courseUseCase.Search(keyword, page, size)
    UseCase->>Repo: courseRepo.Search(keyword, page, size)
    Repo->>DB: Full Text Search Query<br/>WHERE fts @@ plainto_tsquery(?)<br/>OR code ILIKE ?<br/>ORDER BY ts_rank DESC
    DB-->>Repo: Return Filtered Courses []
    Repo-->>UseCase: Return Search Results
    UseCase->>UseCase: Calculate Rate & Score
    UseCase-->>Handler: Return CourseShortResponse[]<br/>+ Pagination Meta
    Handler-->>UI: Return JSON (200 OK)
    UI-->>User: แสดงผลลัพธ์การค้นหา
```

---

## 4. Course Flow - Get Course By ID
**หน้าบ้าน:** Course Detail Page

```mermaid
sequenceDiagram
autonumber
actor User as ผู้ใช้งาน (User)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as CourseHandler
participant UseCase as CourseUseCase
participant Repo as CourseRepository
participant DB as ฐานข้อมูล (Database)

    User->>UI: คลิกที่วิชาเพื่อดูรายละเอียด
    UI->>Handler: GET /api/v1/course?id={courseId}<br/>(with JWT Token)
    Handler->>UseCase: courseUseCase.GetCourseById(id)
    UseCase->>Repo: courseRepo.GetCourseById(id)
    Repo->>DB: SELECT * FROM courses<br/>WHERE id = ?<br/>PRELOAD professors, tags
    DB-->>Repo: Return Course Entity
    Repo-->>UseCase: Return Course
    UseCase->>UseCase: Calculate Rate & Review Count
    UseCase-->>Handler: Return CourseFullResponse
    Handler-->>UI: Return JSON (200 OK)
    UI-->>User: แสดงรายละเอียดวิชา (ชื่อ, รายละเอียด, อาจารย์, Rate, Tags)
```

---

## 5. Course Flow - Compare Courses
**หน้าบ้าน:** Compare Page

```mermaid
sequenceDiagram
autonumber
actor User as ผู้ใช้งาน (User)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as CourseHandler
participant UseCase as CourseUseCase
participant CourseRepo as CourseRepository
participant ReviewRepo as ReviewRepository
participant DB as ฐานข้อมูล (Database)

    User->>UI: เปรียบเทียบ 2 วิชา
    UI->>Handler: GET /api/v1/course/compare/{courseId1}/{courseId2}<br/>(with JWT Token)
    Handler->>UseCase: courseUseCase.CompareCoursesById(userId, first, second)
    UseCase->>CourseRepo: courseRepo.GetCourseById(first)
    CourseRepo->>DB: SELECT * FROM courses WHERE id = ?
    DB-->>CourseRepo: Return Course 1
    CourseRepo-->>UseCase: Return Course 1
    UseCase->>CourseRepo: courseRepo.GetCourseById(second)
    CourseRepo->>DB: SELECT * FROM courses WHERE id = ?
    DB-->>CourseRepo: Return Course 2
    CourseRepo-->>UseCase: Return Course 2
    UseCase->>ReviewRepo: reviewRepo.GetHotReviewsByCourseID(userId, first, 1, 0)
    ReviewRepo->>DB: SELECT * FROM reviews<br/>WHERE course_id = ?<br/>ORDER BY score DESC LIMIT 1
    DB-->>ReviewRepo: Return Hot Review 1
    ReviewRepo-->>UseCase: Return Hot Review 1
    UseCase->>ReviewRepo: reviewRepo.GetHotReviewsByCourseID(userId, second, 1, 0)
    ReviewRepo->>DB: SELECT * FROM reviews<br/>WHERE course_id = ?<br/>ORDER BY score DESC LIMIT 1
    DB-->>ReviewRepo: Return Hot Review 2
    ReviewRepo-->>UseCase: Return Hot Review 2
    UseCase->>UseCase: Build Comparison Data
    UseCase-->>Handler: Return CourseCompareResponse[]
    Handler-->>UI: Return JSON (200 OK)
    UI-->>User: แสดงการเปรียบเทียบ 2 วิชา
```

---

## 6. Review Flow - Get Reviews by Course ID
**หน้าบ้าน:** Course Detail Page

```mermaid
sequenceDiagram
autonumber
actor User as ผู้ใช้งาน (User)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as ReviewHandler
participant UseCase as ReviewUseCase
participant Repo as ReviewRepository
participant DB as ฐานข้อมูล (Database)

    User->>UI: ดูรีวิวของวิชา
    UI->>Handler: GET /api/v1/review/c/{courseId}?page=1&size=10<br/>(with JWT Token)
    Handler->>UseCase: reviewUseCase.GetReviewsByCourseID(userId, courseId, page, size)
    UseCase->>Repo: reviewRepo.GetReviewsByCourseId(userId, courseId, page, size)
    Repo->>DB: SELECT * FROM reviews<br/>WHERE course_id = ?<br/>LEFT JOIN users, tags, votes<br/>ORDER BY score DESC
    DB-->>Repo: Return Review List []
    Repo-->>UseCase: Return Reviews
    UseCase->>UseCase: Map to DTO (calculate AvgRate, user vote status)
    UseCase-->>Handler: Return ReviewFullResponse[]<br/>+ Pagination Meta
    Handler-->>UI: Return JSON (200 OK)
    UI-->>User: แสดงรีวิวทั้งหมดของวิชา
```

---

## 7. Review Flow - Create Review
**หน้าบ้าน:** Course Detail Page

```mermaid
sequenceDiagram
autonumber
actor User as ผู้ใช้งาน (User)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as ReviewHandler
participant UseCase as ReviewUseCase
participant CourseRepo as CourseRepository
participant ReviewRepo as ReviewRepository
participant DB as ฐานข้อมูล (Database)

    User->>UI: เขียนรีวิว & กดส่ง
    UI->>Handler: POST /api/v1/review<br/>(with JWT Token, review data)
    Handler->>UseCase: reviewUseCase.CreateReview(req, userId)
    UseCase->>CourseRepo: courseRepo.GetCourseById(courseId)
    CourseRepo->>DB: SELECT * FROM courses WHERE id = ?
    DB-->>CourseRepo: Return Course
    CourseRepo-->>UseCase: Return Course
    UseCase->>UseCase: Recalculate Course Rate & Review Count
    UseCase->>ReviewRepo: reviewRepo.CreateReview(review, updatedCourse)
    ReviewRepo->>DB: INSERT INTO reviews (user_id, course_id, description, rate...)<br/>UPDATE courses SET rate = ?, review_count = ?
    DB-->>ReviewRepo: Review Created
    ReviewRepo-->>UseCase: Return Review
    UseCase-->>Handler: Return ReviewFullResponse
    Handler-->>UI: Return JSON (201 Created)
    UI-->>User: แสดงรีวิวที่สร้างสำเร็จ
```

---

## 8. Vote Flow - Vote Review
**หน้าบ้าน:** Course Detail Page

```mermaid
sequenceDiagram
autonumber
actor User as ผู้ใช้งาน (User)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as VoteHandler
participant UseCase as VoteUseCase
participant AuthRepo as AuthRepository
participant ReviewRepo as ReviewRepository
participant VoteRepo as VoteRepository
participant DB as ฐานข้อมูล (Database)

    User->>UI: กด Upvote/Downvote ในรีวิว
    UI->>Handler: POST /api/v1/vote<br/>(with JWT Token, reviewId, vote)
    Handler->>UseCase: voteUseCase.AddVoteToReview(userId, reviewId, vote)
    UseCase->>AuthRepo: authRepo.VerifyUserById(userId)
    AuthRepo->>DB: SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)
    DB-->>AuthRepo: User Exists = true/false
    AuthRepo-->>UseCase: Return boolean
    UseCase->>ReviewRepo: reviewRepo.VerifyReviewById(reviewId)
    ReviewRepo->>DB: SELECT EXISTS(SELECT 1 FROM reviews WHERE id = ?)
    DB-->>ReviewRepo: Review Exists = true/false
    ReviewRepo-->>UseCase: Return boolean
    UseCase->>VoteRepo: voteRepo.GetVoteByUserAndReviewId(userId, reviewId)
    VoteRepo->>DB: SELECT * FROM votes<br/>WHERE user_id = ? AND review_id = ?
    DB-->>VoteRepo: Return Vote หรือ nil
    alt มี Vote อยู่แล้วและค่าโหวตเท่าเดิม
        VoteRepo-->>UseCase: Return Existing Vote
        UseCase->>VoteRepo: voteRepo.DeleteVote(userId, reviewId)
        VoteRepo->>DB: DELETE FROM votes<br/>WHERE user_id = ? AND review_id = ?
        DB-->>VoteRepo: Vote Deleted
    else มี Vote อยู่แล้วแต่คนละค่า
        VoteRepo-->>UseCase: Return Existing Vote
        UseCase->>VoteRepo: voteRepo.UpdateVote(userId, reviewId, vote)
        VoteRepo->>DB: UPDATE votes SET vote = ?<br/>WHERE user_id = ? AND review_id = ?
        DB-->>VoteRepo: Vote Updated
    else ยังไม่มี Vote
        VoteRepo-->>UseCase: Return nil
        UseCase->>VoteRepo: voteRepo.AddVoteToReview(userId, reviewId, vote)
        VoteRepo->>DB: INSERT INTO votes (user_id, review_id, vote)
        DB-->>VoteRepo: Vote Created
    end
    UseCase->>VoteRepo: voteRepo.GetVotesByReviewId(reviewId)
    VoteRepo->>DB: SELECT count(*) FILTER (vote = 1/-1)
    DB-->>VoteRepo: upCount, downCount
    UseCase->>UseCase: utils.WilsonScoreFromVotes(up, down)
    UseCase->>ReviewRepo: reviewRepo.UpdateScoreForReview(reviewId, score)
    ReviewRepo->>DB: UPDATE reviews SET score = ? WHERE id = ?
    DB-->>ReviewRepo: Review Score Updated
    UseCase-->>Handler: Return Success
    Handler-->>UI: Return JSON (200 OK)
    UI-->>User: อัพเดท UI แสดงสถานะ Vote
```

---

## 9. Professor Flow - Get All/One Professor

```mermaid
sequenceDiagram
autonumber
actor User as ผู้ใช้งาน (User)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as ProfessorHandler
participant UseCase as ProfessorUseCase
participant Repo as ProfessorRepository
participant DB as ฐานข้อมูล (Database)

    User->>UI: ดูรายชื่ออาจารย์หรือรายละเอียดอาจารย์
    UI->>Handler: GET /api/v1/professor?id={professorId}<br/>หรือ GET /api/v1/professor (all)<br/>(with JWT Token)
    Handler->>UseCase: professorUseCase.GetProfessorById(id)<br/>หรือ professorUseCase.GetAll()
    UseCase->>Repo: professorRepo.GetProfessorById(id)<br/>หรือ professorRepo.GetAll()
    Repo->>DB: SELECT * FROM professors<br/>WHERE id = ? หรือ SELECT * FROM professors
    DB-->>Repo: Return Professor(s)
    Repo-->>UseCase: Return Professor(s)
    UseCase-->>Handler: Return ProfessorResponse
    Handler-->>UI: Return JSON (200 OK)
    UI-->>User: แสดงข้อมูลอาจารย์
```

---

## 10. Tag Flow - Get All/One Tag

```mermaid
sequenceDiagram
autonumber
actor User as ผู้ใช้งาน (User)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as TagHandler
participant UseCase as TagUseCase
participant Repo as TagRepository
participant DB as ฐานข้อมูล (Database)

    User->>UI: ดู Tags (สำหรับ filter หรือเลือกเวลาเขียนรีวิว)
    UI->>Handler: GET /api/v1/tag?id={tagId}<br/>หรือ GET /api/v1/tag (all)<br/>(with JWT Token)
    Handler->>UseCase: tagUseCase.GetTagById(id)<br/>หรือ tagUseCase.GetAll()
    UseCase->>Repo: tagRepo.GetTagById(id)<br/>หรือ tagRepo.GetAll()
    Repo->>DB: SELECT * FROM tags<br/>WHERE id = ? หรือ SELECT * FROM tags
    Repo-->>UseCase: Return Tag(s)
    UseCase-->>Handler: Return TagResponse
    Handler-->>UI: Return JSON (200 OK)
    UI-->>User: แสดง Tags ทั้งหมด
```

---

## 11. Report Flow - Create Report
**หน้าบ้าน:** Course Detail Page

```mermaid
sequenceDiagram
autonumber
actor User as ผู้ใช้งาน (User)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as ReportHandler
participant UseCase as ReportUseCase
participant Repo as ReportRepository
participant DB as ฐานข้อมูล (Database)

    User->>UI: รายงานรีวิวที่ไม่เหมาะสม
    UI->>Handler: POST /api/v1/report<br/>(with JWT Token, reviewId, reason)
    Handler->>UseCase: reportUseCase.AddReportToReview(req)
    UseCase->>UseCase: Validate report type (1-4)
    alt ReportType ไม่ถูกต้อง
        UseCase-->>Handler: Error "invalid report type"
        Handler-->>UI: Return JSON (400)
    else ReportType ถูกต้อง
        UseCase->>Repo: reportRepo.AddReportToReview(req)
        Repo->>DB: INSERT reports + UPDATE reviews.report_count
        DB-->>Repo: Report Created
        Repo-->>UseCase: Success
        UseCase-->>Handler: Success
        Handler-->>UI: Return JSON (200 OK)
        UI-->>User: แสดงข้อความ "รายงานสำเร็จ"
    end
```

---

## 12. Admin Flow - Get All Reports
**หน้าบ้าน:** Admin Report Management Page

```mermaid
sequenceDiagram
autonumber
actor Admin as ผู้ดูแลระบบ (Admin)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as ReportHandler
participant UseCase as ReportUseCase
participant Repo as ReportRepository
participant DB as ฐานข้อมูล (Database)

    Admin->>UI: เข้าสู่หน้าจัดการรายงาน
    UI->>Handler: GET /api/v1/report<br/>(with JWT role=ADMIN)
    Handler->>Handler: Validate role == ADMIN
    alt ไม่ใช่ ADMIN
        Handler-->>UI: Return JSON (403 Forbidden)
        UI-->>Admin: แจ้งเตือนไม่มีสิทธิ์เข้าถึง
    else เป็น ADMIN
        Handler->>UseCase: reportUseCase.GetAllReports()
        UseCase->>Repo: reportRepo.GetAllReports()
        Repo->>DB: SELECT * FROM reports<br/>LEFT JOIN users, reviews
        DB-->>Repo: Report Entities
        Repo-->>UseCase: Return Report List
        UseCase-->>Handler: ReportShortResponse[]
        Handler-->>UI: Return JSON (200 OK)
        UI-->>Admin: แสดงรายการรายงานทั้งหมด
    end
```

---

## 13. Admin Flow - Manage Report (Solve/Reject)
**หน้าบ้าน:** Admin Report Management Page

```mermaid
sequenceDiagram
autonumber
actor Admin as ผู้ดูแลระบบ (Admin)
participant UI as หน้าเว็บไซต์ (Frontend)
participant Handler as ReportHandler
participant UseCase as ReportUseCase
participant ReportRepo as ReportRepository
participant ReviewRepo as ReviewRepository
participant DB as ฐานข้อมูล (Database)

    Admin->>UI: กดปุ่ม "Solve" หรือ "Reject" ที่รายงาน
    UI->>Handler: PATCH /api/v1/report/{reportId}/{action}<br/>(action: "solve" or "reject")<br/>(with JWT role=ADMIN)
    Handler->>Handler: Validate role == ADMIN
    Handler->>UseCase: reportUseCase.ManageReport(reportId, action)
    UseCase->>ReportRepo: reportRepo.GetReportById(reportId)
    ReportRepo->>DB: SELECT * FROM reports WHERE id = ?
    DB-->>ReportRepo: Return Report
    ReportRepo-->>UseCase: Return Report (with reviewId)

    alt action is "solve"
        UseCase->>ReportRepo: reportRepo.UpdateReportStatus(reportId, "SOLVED")
        ReportRepo->>DB: UPDATE reports SET status = 'SOLVED' WHERE id = ?
        DB-->>ReportRepo: Status Updated
        UseCase->>ReviewRepo: reviewRepo.DeleteReview(reviewId)
        ReviewRepo->>DB: DELETE FROM reviews WHERE id = ?
        DB-->>ReviewRepo: Review Deleted
    else action is "reject"
        UseCase->>ReportRepo: reportRepo.UpdateReportStatus(reportId, "REJECTED")
        ReportRepo->>DB: UPDATE reports SET status = 'REJECTED' WHERE id = ?
        DB-->>ReportRepo: Status Updated
    end

    UseCase-->>Handler: Return Success
    Handler-->>UI: Return JSON (200 OK)
    UI-->>Admin: อัพเดทสถานะรายงานใน UI
```

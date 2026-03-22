# Class Diagram - Real Course Review Backend

```mermaid
classDiagram
    class User {
        +UUID ID
        +String Name
        +String Email
        +String Role
        +String UniYear
        +DateTime CreatedAt
        +DateTime UpdatedAt
    }

    class Course {
        +UUID ID
        +String Name
        +String Description
        +String Semester
        +String Code
        +Integer Credit
        +String CourseType
        +String ReviewCount
        +String Rate
        +Float Score
        +Boolean RecStatus
        +String FTS
        +DateTime CreatedAt
        +DateTime UpdatedAt
    }

    class Professor {
        +UUID ID
        +String Name
        +String Email
        +String Title
        +String ImageURL
        +String Phone
        +String Description
        +String UniRoomAddress
        +DateTime CreatedAt
        +DateTime UpdatedAt
    }

    class Review {
        +UUID ID
        +UUID UserID
        +UUID CourseID
        +String Status
        +String Description
        +Float AvgRate
        +Float Score
        +Float DecayedScore
        +Integer UpCount
        +Integer DownCount
        +Integer? UserVote
        +Integer ReportCount
        +Boolean IsAnonymous
        +String Grade
        +String Year
        +String Sec
        +Boolean RecStatus
        +String Rate
        +DateTime CreatedAt
        +DateTime UpdatedAt
    }

    class Tag {
        +UUID ID
        +String Name
        +DateTime CreatedAt
        +DateTime UpdatedAt
    }

    class Vote {
        +UUID ID
        +UUID UserID
        +UUID ReviewID
        +Integer Vote
        +DateTime CreatedAt
        +DateTime UpdatedAt
    }

    class Report {
        +UUID ID
        +UUID UserID
        +UUID ReviewID
        +ReportType ReportType
        +Boolean RecStatus
        +String SolveReason
        +DateTime CreatedAt
        +DateTime UpdatedAt
    }

    class CourseTag {
        +UUID CourseID
        +UUID TagID
    }

    class ReviewTag {
        +UUID ReviewID
        +UUID TagID
    }

    %% User Relationships
    User "1" --> "0..*" Review : writes
    User "1" --> "0..*" Vote : casts
    User "1" --> "0..*" Report : creates

    %% Course Relationships
    Course "1" --> "0..*" Review : has
    Course "0..*" --> "0..*" Professor : taught by
    Course "0..*" --> "0..*" Tag : categorized by

    %% Review Relationships
    Review "1" --> "0..*" Vote : receives
    Review "1" --> "0..*" Report : reported by
    Review "0..*" --> "0..*" Tag : tagged with

    %% Junction Tables
    Course "1" --> "0..*" CourseTag : ""
    Tag "1" --> "0..*" CourseTag : ""
    
    Review "1" --> "0..*" ReviewTag : ""
    Tag "1" --> "0..*" ReviewTag : ""
    
    %% Vote and Report References
    Vote --> User : voted by
    Vote --> Review : on
    
    Report --> User : reported by
    Report --> Review : on
```

---

## Entity Descriptions

### User
- ผู้ใช้งานในระบบ (นักศึกษา/ผู้ดูแลระบบ)
- สามารถเขียนรีวิว, โหวต, และรายงานรีวิวได้
- มี Role สำหรับแยกสิทธิ์ (USER, ADMIN)

### Course
- รายวิชาในระบบ
- มีข้อมูล Code, Credit, Semester, CourseType
- คำนวณคะแนนเฉลี่ย (Rate) และจำนวนรีวิว (ReviewCount)
- มีสถานะการแนะนำวิชา (RecStatus)
- มี Full Text Search (FTS) สำหรับการค้นหา

### Professor
- อาจารย์ผู้สอน
- มีข้อมูลติดต่อและที่ตั้งห้องทำงาน
- สอนได้หลายวิชา (Many-to-Many กับ Course)

### Review
- รีวิวของวิชาที่เขียนโดยผู้ใช้
- มีคะแนนเฉลี่ย (AvgRate) จาก Rate (Happiness, Easiness, Quality)
- มีระบบ Vote (UpCount, DownCount, UserVote)
- รองรับคะแนนที่ใช้จัดอันดับตามเวลา (DecayedScore)
- สามารถตั้งค่าเป็นไม่เปิดเผยตัวตน (IsAnonymous)
- มีสถานะเนื้อหา (publish, ban, delete, hidden) และสถานะการแนะนำ (RecStatus)

### Tag
- แท็กสำหรับจัดหมวดหมู่ Course และ Review
- ใช้ Many-to-Many relationship

### Vote
- การโหวตรีวิว (Upvote = 1, Downvote = -1)
- ผู้ใช้หนึ่งคนสามารถโหวตรีวิวหนึ่งรีวิวได้เพียงครั้งเดียว

### Report
- การรายงานรีวิวที่ไม่เหมาะสม
- มี ReportType เพื่อระบุประเภทการรายงาน (1-4)
- มีสถานะการจัดการรายงาน (RecStatus) และเหตุผลการปิดรายงาน (SolveReason)

### Junction Tables
- **CourseTag**: เชื่อม Course กับ Tag (Many-to-Many)
- **ReviewTag**: เชื่อม Review กับ Tag (Many-to-Many)

---

## Relationships Summary

| From | Relationship | To | Description |
|------|-------------|-----|-------------|
| User | 1 : N | Review | ผู้ใช้เขียนได้หลายรีวิว |
| User | 1 : N | Vote | ผู้ใช้โหวตได้หลายรีวิว |
| User | 1 : N | Report | ผู้ใช้รายงานได้หลายรีวิว |
| Course | 1 : N | Review | วิชามีได้หลายรีวิว |
| Course | N : M | Professor | วิชาสอนโดยหลายอาจารย์ |
| Course | N : M | Tag | วิชามีได้หลายแท็ก |
| Review | 1 : N | Vote | รีวิวได้รับหลายโหวต |
| Review | 1 : N | Report | รีวิวถูกรายงานได้หลายครั้ง |
| Review | N : M | Tag | รีวิวมีได้หลายแท็ก |

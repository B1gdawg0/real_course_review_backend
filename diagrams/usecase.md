# Use Case Descriptions - Real Course Review Backend

## 1. Login/Register

| **Usecase ID** | UC-001 |
|----------------|---------|
| **Usecase Name** | เข้าสู่ระบบ/สมัครสมาชิก |
| **Actor** | ผู้ใช้งาน (นักศึกษา) |
| **Pre-Condition** | - ผู้ใช้ต้องมี Email ที่ลงท้ายด้วย @ku.th |
| **Post-Condition** | - ผู้ใช้ได้รับ JWT Token<br>- ระบบบันทึกข้อมูลผู้ใช้ (กรณีสมัครใหม่)<br>- ผู้ใช้สามารถเข้าถึงฟังก์ชันต่างๆ ในระบบได้ |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. ผู้ใช้เข้าสู่หน้า Login | |
| 2. ผู้ใช้กรอก Name และ Email (@ku.th) | |
| 3. ผู้ใช้กดปุ่ม Login/Register | |
| | 4. ระบบตรวจสอบว่า Email ลงท้ายด้วย @ku.th<br>`if !strings.HasSuffix(req.Email, "@ku.th")` |
| | 5. ระบบค้นหาผู้ใช้ในฐานข้อมูล<br>`authRepo.GetUserByEmail()`<br>`SELECT * FROM users WHERE email = ?` |
| | 6. ถ้าไม่พบผู้ใช้ในระบบ:<br>`authRepo.CreateUser()`<br>`INSERT INTO users (name, email, role)` |
| | 7. ระบบสร้าง JWT Token<br>`jwt.NewWithClaims()` พร้อม claims: sub, email, role, exp |
| | 8. ระบบส่ง Token และข้อมูลผู้ใช้กลับไป<br>`Return { token, user }` |
| 9. ผู้ใช้ถูก redirect ไปหน้า Home พร้อม Token | |

| **Alternative Flow** | |
|---------------------|---|
| **4a. Email ไม่ใช่ @ku.th** | ระบบแสดงข้อความ "unauthorized domain" |

---

## 2. ดูรายวิชาทั้งหมด

| **Usecase ID** | UC-002 |
|----------------|---------|
| **Usecase Name** | ดูรายวิชาทั้งหมด |
| **Actor** | ผู้ใช้งาน |
| **Pre-Condition** | - ผู้ใช้ Login เข้าสู่ระบบแล้ว (มี JWT Token) |
| **Post-Condition** | - แสดงรายวิชาทั้งหมดพร้อม Tags, Rate, Review Count<br>- รายวิชาถูกเรียงลำดับตาม Score |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. ผู้ใช้เข้าสู่หน้า Home | |
| | 2. ระบบดึงข้อมูลรายวิชาทั้งหมด<br>`courseRepo.GetAll()`<br>`SELECT * FROM courses`<br>`LEFT JOIN course_tags`<br>`LEFT JOIN tags` |
| | 3. ระบบคำนวณ Rate และ Score<br>`utils.ParseRate(entity.Rate)`<br>`utils.ParseReviewCount(entity.ReviewCount)` |
| | 4. ระบบเรียงลำดับตาม Score<br>`sort.Slice(res, func(i, j int) bool { return res[i].Score > res[j].Score })` |
| | 5. ระบบส่งข้อมูลพร้อม Pagination<br>`Return CourseShortResponse[], page, size, total, totalPages` |
| 6. ผู้ใช้เห็นรายวิชาทั้งหมดพร้อม Tags และ Rate | |

---

## 3. ค้นหารายวิชา

| **Usecase ID** | UC-003 |
|----------------|---------|
| **Usecase Name** | ค้นหารายวิชา |
| **Actor** | ผู้ใช้งาน |
| **Pre-Condition** | - ผู้ใช้ Login เข้าสู่ระบบแล้ว<br>- ผู้ใช้อยู่ในหน้า Home |
| **Post-Condition** | - แสดงรายวิชาที่ตรงกับคำค้นหา<br>- ผลลัพธ์ถูกเรียงตาม relevance (ts_rank) |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. ผู้ใช้พิมพ์คำค้นหาในช่อง Search | |
| 2. ผู้ใช้กด Enter หรือปุ่มค้นหา | |
| | 3. ระบบทำ Full Text Search<br>`courseRepo.Search(keyword, page, size)`<br>`WHERE fts @@ plainto_tsquery('english', ?)`<br>`OR name ILIKE ?`<br>`OR code ILIKE ?`<br>`ORDER BY ts_rank(fts, plainto_tsquery('english', ?)) DESC` |
| | 4. ระบบคำนวณ Rate และ Score |
| | 5. ระบบส่งผลลัพธ์การค้นหาพร้อม Pagination<br>`Return CourseShortResponse[], page, size, total, totalPages` |
| 6. ผู้ใช้เห็นผลลัพธ์การค้นหา | |

| **Alternative Flow** | |
|---------------------|---|
| **3a. ไม่พบรายวิชาที่ตรงกับคำค้นหา** | ระบบแสดงข้อความ "ไม่พบรายวิชา" และแสดง array ว่าง |

---

## 4. ดูรายละเอียดรายวิชาและรีวิว

| **Usecase ID** | UC-004 |
|----------------|---------|
| **Usecase Name** | ดูรายละเอียดรายวิชาและรีวิว |
| **Actor** | ผู้ใช้งาน |
| **Pre-Condition** | - ผู้ใช้ Login เข้าสู่ระบบแล้ว<br>- มี Course ID ที่ต้องการดู |
| **Post-Condition** | - แสดงรายละเอียดวิชาแบบเต็ม (ชื่อ, รายละเอียด, อาจารย์, Tags, Rate)<br>- แสดงรีวิวทั้งหมดของวิชา พร้อม Vote count และ Tags<br>- รีวิวถูกเรียงตาม Score |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. ผู้ใช้คลิกที่รายวิชาที่สนใจ | |
| | 2. ระบบดึงข้อมูลรายวิชาตาม ID<br>`courseRepo.GetCourseById(id)`<br>`SELECT * FROM courses WHERE id = ?`<br>`PRELOAD professors, tags` |
| | 3. ระบบคำนวณ Rate และ Review Count<br>`utils.ParseRate(entity.Rate)`<br>`utils.ParseReviewCount(entity.ReviewCount)` |
| | 4. ระบบดึงรีวิวตาม Course ID<br>`reviewRepo.GetReviewsByCourseId(userId, courseId, page, size)`<br>`SELECT * FROM reviews WHERE course_id = ?`<br>`LEFT JOIN users, tags, votes`<br>`ORDER BY score DESC` |
| | 5. ระบบคำนวณ AvgRate และสถานะการ Vote ของผู้ใช้<br>`Map to DTO (calculate AvgRate, user vote status)` |
| | 6. ระบบส่งข้อมูลรายวิชาพร้อมรีวิว<br>`Return { course: CourseFullResponse, reviews: ReviewFullResponse[], page, size, total, totalPages }` |
| 7. ผู้ใช้เห็นรายละเอียดวิชาพร้อมอาจารย์ผู้สอน, Tags และรีวิวทั้งหมด | |

| **Alternative Flow** | |
|---------------------|---|
| **2a. ไม่พบรายวิชา** | ระบบแสดง Error "Course not found" (404) |

---

## 5. เปรียบเทียบรายวิชา

| **Usecase ID** | UC-005 |
|----------------|---------|
| **Usecase Name** | เปรียบเทียบรายวิชา 2 วิชา |
| **Actor** | ผู้ใช้งาน |
| **Pre-Condition** | - ผู้ใช้ Login เข้าสู่ระบบแล้ว<br>- มี Course ID 2 วิชาที่ต้องการเปรียบเทียบ |
| **Post-Condition** | - แสดงข้อมูลเปรียบเทียบ 2 วิชา (Rate, Tags, Hot Reviews, AI Summary) |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. ผู้ใช้เลือกรายวิชา 2 วิชาเพื่อเปรียบเทียบ | |
| 2. ผู้ใช้กดปุ่มเปรียบเทียบ | |
| | 3. ระบบดึงข้อมูลวิชาที่ 1<br>`courseRepo.GetCourseById(first)` |
| | 4. ระบบดึงข้อมูลวิชาที่ 2<br>`courseRepo.GetCourseById(second)` |
| | 5. ระบบดึงรีวิวยอดนิยมของวิชาที่ 1<br>`reviewRepo.GetHotReviewsByCourseID(userId, first, 1, 0)` |
| | 6. ระบบดึงรีวิวยอดนิยมของวิชาที่ 2<br>`reviewRepo.GetHotReviewsByCourseID(userId, second, 1, 0)` |
| | 7. ระบบสร้างข้อมูลเปรียบเทียบพร้อม AI Summary<br>`Return CourseCompareResponse[]` |
| 8. ผู้ใช้เห็นการเปรียบเทียบ 2 วิชาแบบ side-by-side | |

| **Alternative Flow** | |
|---------------------|---|
| **3a/4a. ไม่พบรายวิชาใดวิชาหนึ่ง** | ระบบแสดง Error "Course not found" |

---

## 6. เขียนรีวิว

| **Usecase ID** | UC-006 |
|----------------|---------|
| **Usecase Name** | เขียนรีวิววิชา |
| **Actor** | ผู้ใช้งาน |
| **Pre-Condition** | - ผู้ใช้ Login เข้าสู่ระบบแล้ว<br>- อยู่ในหน้ารายละเอียดวิชา |
| **Post-Condition** | - รีวิวถูกบันทึกในระบบ<br>- คะแนนเฉลี่ยของวิชาถูกอัพเดท<br>- จำนวนรีวิวของวิชาเพิ่มขึ้น |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. ผู้ใช้กดปุ่ม "เขียนรีวิว" | |
| 2. ผู้ใช้กรอกข้อมูล:<br>- เลือกอาจารย์<br>- ให้คะแนน (Happiness, Easiness, Quality)<br>- เขียนความคิดเห็น<br>- เลือก Tags<br>- กรอก Grade, Year, Section<br>- เลือกเปิดเผยตัวตนหรือไม่ | |
| 3. ผู้ใช้กดปุ่มส่งรีวิว | |
| | 4. ระบบดึงข้อมูลวิชา<br>`courseRepo.GetCourseById(courseId)` |
| | 5. ระบบคำนวณคะแนนเฉลี่ยใหม่ของวิชา<br>`utils.RecalculateCourseReview(course, &req)` |
| | 6. ระบบบันทึกรีวิวและอัพเดทวิชา<br>`reviewRepo.CreateReview(review, updatedCourse)`<br>`INSERT INTO reviews (...)`<br>`UPDATE courses SET rate = ?, review_count = ?` |
| | 7. ระบบส่งข้อมูลรีวิวที่สร้างกลับไป<br>`Return ReviewFullResponse` |
| 8. ผู้ใช้เห็นรีวิวที่เขียนปรากฏในรายการรีวิว | |

| **Alternative Flow** | |
|---------------------|---|
| **3a. ข้อมูลไม่ครบถ้วน** | ระบบแสดง Error "missing required fields" |
| **4a. ไม่พบรายวิชา** | ระบบแสดง Error "Course not found" |

---

## 7. โหวตรีวิว

| **Usecase ID** | UC-007 |
|----------------|---------|
| **Usecase Name** | โหวตรีวิว (Upvote/Downvote) |
| **Actor** | ผู้ใช้งาน |
| **Pre-Condition** | - ผู้ใช้ Login เข้าสู่ระบบแล้ว<br>- มีรีวิวที่ต้องการโหวต |
| **Post-Condition** | - Vote ถูกบันทึกหรืออัพเดทในระบบ<br>- จำนวน Upvote/Downvote ของรีวิวเปลี่ยนแปลง |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. ผู้ใช้กดปุ่ม Upvote (👍) หรือ Downvote (👎) | |
| | 2. ระบบตรวจสอบว่าผู้ใช้มีอยู่ในระบบ<br>`authRepo.VerifyUserById(userId)`<br>`SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)` |
| | 3. ระบบตรวจสอบว่ารีวิวมีอยู่ในระบบ<br>`reviewRepo.VerifyReviewById(reviewId)`<br>`SELECT EXISTS(SELECT 1 FROM reviews WHERE id = ?)` |
| | 4. ระบบตรวจสอบว่าผู้ใช้เคย Vote รีวิวนี้หรือไม่<br>`voteRepo.GetVoteByUserAndReviewId(userId, reviewId)`<br>`SELECT * FROM votes WHERE user_id = ? AND review_id = ?` |
| | 5a. ถ้าเคย Vote แล้ว และเป็น Vote เดิม:<br>แสดง Error "can't vote the same vote" |
| | 5b. ถ้าเคย Vote แล้ว แต่เป็น Vote ตรงข้าม:<br>`voteRepo.UpdateVote(userId, reviewId, vote)`<br>`UPDATE votes SET vote = ? WHERE user_id = ? AND review_id = ?` |
| | 5c. ถ้ายังไม่เคย Vote:<br>`voteRepo.AddVoteToReview(userId, reviewId, vote)`<br>`INSERT INTO votes (user_id, review_id, vote)` |
| | 6. ระบบส่งสถานะสำเร็จกลับไป |
| 7. ผู้ใช้เห็น UI แสดงสถานะ Vote ที่อัพเดทแล้ว | |

| **Alternative Flow** | |
|---------------------|---|
| **2a. ผู้ใช้ไม่พบในระบบ** | ระบบแสดง Error "user not found" |
| **3a. รีวิวไม่พบในระบบ** | ระบบแสดง Error "review not found" |
| **5a. Vote ซ้ำ** | ระบบแสดง Error "can't vote the same vote" |

---

## 8. ดูข้อมูลอาจารย์

| **Usecase ID** | UC-008 |
|----------------|---------|
| **Usecase Name** | ดูข้อมูลอาจารย์ |
| **Actor** | ผู้ใช้งาน |
| **Pre-Condition** | - ผู้ใช้ Login เข้าสู่ระบบแล้ว |
| **Post-Condition** | - แสดงข้อมูลอาจารย์ทั้งหมด หรือข้อมูลอาจารย์เฉพาะคน |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. ผู้ใช้เข้าสู่หน้าดูข้อมูลอาจารย์<br>หรือค้นหาอาจารย์ | |
| | 2. ระบบดึงข้อมูลอาจารย์<br>`professorRepo.GetProfessorById(id)` (ถ้ามี id)<br>หรือ `professorRepo.GetAll()` (ถ้าไม่มี id)<br>`SELECT * FROM professors WHERE id = ?`<br>หรือ `SELECT * FROM professors` |
| | 3. ระบบส่งข้อมูลอาจารย์กลับไป<br>`Return ProfessorResponse` |
| 4. ผู้ใช้เห็นข้อมูลอาจารย์ (ชื่อ, Email, Title, เบอร์โทร, ห้องทำงาน) | |

---

## 9. รายงานรีวิว

| **Usecase ID** | UC-009 |
|----------------|---------|
| **Usecase Name** | รายงานรีวิวที่ไม่เหมาะสม |
| **Actor** | ผู้ใช้งาน |
| **Pre-Condition** | - ผู้ใช้ Login เข้าสู่ระบบแล้ว<br>- มีรีวิวที่ต้องการรายงาน |
| **Post-Condition** | - Report ถูกบันทึกในระบบ<br>- จำนวน Report ของรีวิวเพิ่มขึ้น |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. ผู้ใช้กดปุ่ม "รายงาน" ที่รีวิว | |
| 2. ผู้ใช้เลือกประเภทการรายงาน (ReportType 1-4) | |
| 3. ผู้ใช้กดยืนยันการรายงาน | |
| | 4. ระบบบันทึกการรายงานและอัพเดทจำนวน Report<br>`reportRepo.AddReportToReview(req)`<br>`INSERT INTO reports (user_id, review_id, reason)`<br>`UPDATE reviews SET report_count = report_count + 1` |
| | 5. ระบบส่งสถานะสำเร็จกลับไป |
| 6. ผู้ใช้เห็นข้อความ "รายงานสำเร็จ" | |

| **Alternative Flow** | |
|---------------------|---|
| **3a. ไม่ระบุ Review ID** | ระบบแสดง Error "missing required fields" |

---

## 10. ดูรายงานทั้งหมด (Admin)

| **Usecase ID** | UC-010 |
|----------------|---------|
| **Usecase Name** | ดูรายงานทั้งหมด |
| **Actor** | ผู้ดูแลระบบ (Admin) |
| **Pre-Condition** | - ผู้ใช้ Login ด้วย Role = ADMIN |
| **Post-Condition** | - แสดงรายงานทั้งหมดที่ผู้ใช้รายงานมา |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. Admin เข้าสู่หน้าจัดการรายงาน | |
| | 2. ระบบดึงรายงานทั้งหมด<br>`reportRepo.GetAllReports()`<br>`SELECT * FROM reports`<br>`LEFT JOIN users, reviews` |
| | 3. ระบบส่งข้อมูลรายงานกลับไป<br>`Return ReportResponse[]` |
| 4. Admin เห็นรายงานทั้งหมดพร้อมข้อมูลผู้รายงานและรีวิวที่ถูกรายงาน | |

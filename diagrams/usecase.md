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
| **Post-Condition** | - แสดงรายวิชาทั้งหมดพร้อม Tags, Rate, Review Count<br>- รองรับ Pagination และ Filter (semester, course_type, tag_ids) |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. ผู้ใช้เข้าสู่หน้า Home | |
| | 2. ระบบดึงข้อมูลรายวิชาทั้งหมด<br>`courseRepo.GetAll()`<br>`SELECT * FROM courses`<br>`LEFT JOIN course_tags`<br>`LEFT JOIN tags` |
| | 3. ระบบคำนวณ Rate และ Score<br>`utils.ParseRate(entity.Rate)`<br>`utils.ParseReviewCount(entity.ReviewCount)` |
| | 4. ระบบส่งข้อมูลพร้อม Pagination<br>`Return CourseShortResponse[], page, size, total, totalPages` |
| 5. ผู้ใช้เห็นรายวิชาทั้งหมดพร้อม Tags และ Rate | |

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
| | 4. ระบบดึงรีวิวตาม Course ID<br>`reviewRepo.GetReviewsByCourseId(userId, courseId, page, size)`<br>`SELECT * FROM reviews WHERE course_id = ?`<br>`LEFT JOIN users, tags, votes`<br>`ORDER BY decayed_score DESC` |
| | 5. ระบบคำนวณ AvgRate และสถานะการ Vote ของผู้ใช้<br>`Map to DTO (calculate AvgRate, user vote status)` |
| | 6. ระบบส่งข้อมูลรายวิชาและรีวิวแยกตาม endpoint<br>`Return CourseFullResponse` และ `ReviewFullResponse[], page, size, total, totalPages` |
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
| **Post-Condition** | - รีวิวถูกบันทึกในระบบ<br>- คะแนนเฉลี่ยของวิชาถูกอัพเดท<br>- จำนวนรีวิวของวิชาเพิ่มขึ้น<br>- ระบบคำนวณคะแนนรีวิว (score/decayed score) สำหรับการจัดอันดับ |

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
| | 5a. ถ้าเคย Vote แล้ว และเป็น Vote เดิม:<br>`voteRepo.DeleteVote(userId, reviewId)`<br>`DELETE FROM votes WHERE user_id = ? AND review_id = ?` (toggle off) |
| | 5b. ถ้าเคย Vote แล้ว แต่เป็น Vote ตรงข้าม:<br>`voteRepo.UpdateVote(userId, reviewId, vote)`<br>`UPDATE votes SET vote = ? WHERE user_id = ? AND review_id = ?` |
| | 5c. ถ้ายังไม่เคย Vote:<br>`voteRepo.AddVoteToReview(userId, reviewId, vote)`<br>`INSERT INTO votes (user_id, review_id, vote)` |
| | 6. ระบบคำนวณคะแนนรีวิวใหม่จากจำนวนโหวต<br>`voteRepo.GetVotesByReviewId(reviewId)` + `utils.WilsonScoreFromVotes(up, down)` |
| | 7. ระบบอัพเดท score ของรีวิว<br>`reviewRepo.UpdateScoreForReview(reviewId, score)` |
| | 8. ระบบส่งสถานะสำเร็จกลับไป |
| 9. ผู้ใช้เห็น UI แสดงสถานะ Vote ที่อัพเดทแล้ว | |

| **Alternative Flow** | |
|---------------------|---|
| **2a. ผู้ใช้ไม่พบในระบบ** | ระบบแสดง Error "user not found" |
| **3a. รีวิวไม่พบในระบบ** | ระบบแสดง Error "review not found" |
| **1a. vote ไม่ใช่ 1 หรือ -1** | ระบบแสดง Error "invalid vote value" |

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
| | 4. ระบบตรวจสอบประเภทการรายงานว่าอยู่ในช่วง 1-4<br>`if rq.ReportType < 1 || rq.ReportType > 4` |
| | 5. ระบบบันทึกการรายงานและอัพเดทจำนวน Report<br>`reportRepo.AddReportToReview(req)`<br>`INSERT INTO reports (user_id, review_id, reason)`<br>`UPDATE reviews SET report_count = report_count + 1` |
| | 6. ระบบส่งสถานะสำเร็จกลับไป |
| 7. ผู้ใช้เห็นข้อความ "รายงานสำเร็จ" | |

| **Alternative Flow** | |
|---------------------|---|
| **3a. ไม่ระบุ Review ID** | ระบบแสดง Error "missing required fields" |
| **4a. ReportType ไม่ถูกต้อง** | ระบบแสดง Error "invalid report type" |

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

---

## 11. จัดการรายงาน (Admin)

| **Usecase ID** | UC-011 |
|----------------|---------|
| **Usecase Name** | จัดการรายงาน (อนุมัติ/ไม่อนุมัติ) |
| **Actor** | ผู้ดูแลระบบ (Admin) |
| **Pre-Condition** | - ผู้ใช้ Login ด้วย Role = ADMIN<br>- มี Report ID ที่ต้องการจัดการ |
| **Post-Condition** | - สถานะของ Report ถูกเปลี่ยนเป็น `SOLVED` หรือ `REJECTED`<br>- หาก `SOLVED` รีวิวที่ถูกรายงานจะถูกลบ |

| **Normal Flow (Solve)** | |
|-----------------|---|
| **Actor** | **System** |
| 1. Admin กดปุ่ม "Solve" ที่รายงาน | |
| | 2. ระบบอัพเดทสถานะ Report เป็น `SOLVED`<br>`reportRepo.UpdateReportStatus(reportId, "SOLVED")`<br>`UPDATE reports SET status = 'SOLVED' WHERE id = ?` |
| | 3. ระบบทำการลบรีวิวที่เกี่ยวข้อง<br>`reviewRepo.DeleteReview(reviewId)`<br>`DELETE FROM reviews WHERE id = ?` |
| | 4. ระบบส่งสถานะสำเร็จกลับไป |
| 5. UI อัพเดทสถานะของรายงาน และรีวิวหายไป | |

| **Normal Flow (Reject)** | |
|-----------------|---|
| **Actor** | **System** |
| 1. Admin กดปุ่ม "Reject" ที่รายงาน | |
| | 2. ระบบอัพเดทสถานะ Report เป็น `REJECTED`<br>`reportRepo.UpdateReportStatus(reportId, "REJECTED")`<br>`UPDATE reports SET status = 'REJECTED' WHERE id = ?` |
| | 3. ระบบส่งสถานะสำเร็จกลับไป |
| 4. UI อัพเดทสถานะของรายงาน | |

| **Alternative Flow** | |
|---------------------|---|
| **1a. ไม่ใช่ Admin** | ระบบแสดง Error "Forbidden" (403) |
| **1b. ไม่พบ Report** | ระบบแสดง Error "Report not found" (404) |
| **3a. ไม่พบ Review ที่จะลบ** | ระบบแสดง Error "Review not found" (404) |

---

## 12. ปรับสถานะแนะนำรายวิชา (Admin)

| **Usecase ID** | UC-012 |
|----------------|---------|
| **Usecase Name** | ปรับสถานะแนะนำรายวิชา (เปิด/ปิดการแสดงผล) |
| **Actor** | ผู้ดูแลระบบ (Admin) |
| **Pre-Condition** | - ผู้ใช้ Login ด้วย Role = ADMIN<br>- มี JWT Token ที่ถูกต้อง |
| **Post-Condition** | - ค่า `rec_status` ของรายวิชาถูกอัพเดทตามที่กำหนด |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. Admin เลือกรายวิชาที่ต้องการปรับสถานะ | |
| 2. Admin กดเปิด/ปิดสถานะแนะนำ | |
| | 3. Frontend เรียก `PUT /api/v1/course` พร้อม body `{ id, rec_status }` |
| | 4. ระบบตรวจสอบสิทธิ์ว่าเป็น ADMIN (จาก JWT) |
| | 5. ระบบ parse และ validate request body (ต้องมี `id`) |
| | 6. ระบบอัพเดทสถานะรายวิชา<br>`courseUseCase.UpdateCourseRecStatus(id, recStatus)`<br>`courseRepo.UpdateCourseRecStatus(id, recStatus)`<br>`UPDATE courses SET rec_status = ? WHERE id = ?` |
| | 7. ระบบตอบกลับสำเร็จ (No Content) |
| 8. UI อัพเดทสถานะรายวิชา | |

| **Alternative Flow** | |
|---------------------|---|
| **4a. ไม่ใช่ Admin** | ระบบตอบกลับ Forbidden (403) |
| **5a. body ไม่ถูกต้อง / parse ไม่ได้** | ระบบตอบกลับ Bad Request (400) |
| **5b. ไม่ส่ง `id`** | ระบบตอบกลับ Bad Request (400) |

---

## 13. นำเข้ารายวิชา (CSV) แบบกลุ่ม (Admin)

| **Usecase ID** | UC-013 |
|----------------|---------|
| **Usecase Name** | นำเข้ารายวิชาแบบกลุ่มผ่านไฟล์ CSV |
| **Actor** | ผู้ดูแลระบบ (Admin) |
| **Pre-Condition** | - ผู้ใช้ Login ด้วย Role = ADMIN<br>- มีไฟล์ CSV ที่ format ถูกต้อง (มี header และอย่างน้อย 8 คอลัมน์) |
| **Post-Condition** | - รายวิชาใหม่ถูกสร้างในฐานข้อมูล (เฉพาะแถวที่ข้อมูลถูกต้อง)<br>- ความสัมพันธ์กับอาจารย์ถูกผูกให้กับรายวิชา (ตาม professor IDs ในไฟล์)<br>- ระบบส่งผลลัพธ์เป็นจำนวนสำเร็จ/ล้มเหลวและเลขแถวที่ล้มเหลว |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. Admin ไปที่หน้าจอนำเข้ารายวิชา | |
| 2. Admin อัปโหลดไฟล์ CSV | |
| | 3. Frontend เรียก `POST /api/v1/course/import` แบบ `multipart/form-data` (field `file`) |
| | 4. ระบบตรวจสอบสิทธิ์ว่าเป็น ADMIN (จาก JWT) |
| | 5. ระบบอ่านไฟล์ CSV และแปลงเป็น rows[][]string |
| | 6. ระบบรวบรวม professor IDs ทั้งหมดจากคอลัมน์อาจารย์ และตรวจสอบว่ามีอยู่จริง<br>`profRepo.GetExistingProfessorIDs(allProfIDs)` |
| | 7. ระบบไล่ทีละแถวเพื่อสร้าง course model และ validate ข้อมูล (เช่น credit ต้องเป็นตัวเลข) |
| | 8. ถ้าแถวใดมี professor ID ที่ไม่มีอยู่จริง ระบบจะข้ามแถวนั้นและบันทึกเลขแถวเป็น failed |
| | 9. ระบบ bulk create รายวิชาที่ผ่านการตรวจสอบ พร้อมผูกความสัมพันธ์อาจารย์<br>`courseRepo.BulkCreateCourses(courses)` |
| | 10. ระบบตอบกลับผลลัพธ์ `{ success_count, failed_count, failed_rows }` |
| 11. UI แสดงสรุปผลการนำเข้า | |

| **Alternative Flow** | |
|---------------------|---|
| **4a. ไม่ใช่ Admin** | ระบบตอบกลับ Forbidden (403) |
| **3a. ไม่ส่งไฟล์ `file`** | ระบบตอบกลับ Bad Request (400) และข้อความ "file is required" |
| **5a. ไฟล์อ่านไม่ได้/CSV parse ไม่ได้** | ระบบตอบกลับ Error |
| **7a. แถวข้อมูลคอลัมน์ไม่พอ หรือ credit ไม่ใช่ตัวเลข** | ระบบตอบกลับ Error |

---

## 14. แก้ไขข้อมูลรายวิชาตาม ID (Admin)

| **Usecase ID** | UC-014 |
|----------------|---------|
| **Usecase Name** | แก้ไขข้อมูลรายวิชา (แก้เฉพาะ field ที่ส่งมา) |
| **Actor** | ผู้ดูแลระบบ (Admin) |
| **Pre-Condition** | - ผู้ใช้ Login ด้วย Role = ADMIN<br>- ระบุ `courseId` ใน path parameter ได้ถูกต้อง |
| **Post-Condition** | - ข้อมูลรายวิชาถูกอัพเดทตาม field ที่ส่งมา (name/description/semester/code/credit/course_type)<br>- หากส่ง `professor_ids` ระบบจะ replace ความสัมพันธ์อาจารย์ของรายวิชานั้น |

| **Normal Flow** | |
|-----------------|---|
| **Actor** | **System** |
| 1. Admin เปิดหน้าจอแก้ไขรายวิชา | |
| 2. Admin แก้ไขข้อมูลและกดบันทึก | |
| | 3. Frontend เรียก `PUT /api/v1/course/update/{courseId}` พร้อม JSON body (field เป็น optional) |
| | 4. ระบบตรวจสอบสิทธิ์ว่าเป็น ADMIN (จาก JWT) |
| | 5. ระบบ parse request body เป็น `UpdateCourseRequest` |
| | 6. ระบบอ่านข้อมูลรายวิชาเดิมจากฐานข้อมูล<br>`courseRepo.GetCourseById(courseId)` |
| | 7. ระบบ apply เฉพาะ field ที่ถูกส่งมา และเตรียมข้อมูล professors ตาม `professor_ids` (ถ้ามี) |
| | 8. ระบบอัพเดทข้อมูลใน transaction และ replace association ของ professors (ถ้าส่งมา)<br>`courseRepo.UpdateCourseById(courseId, updatedCourse)` |
| | 9. ระบบตอบกลับสำเร็จ (No Content) |
| 10. UI แสดงผลการบันทึกสำเร็จ | |

| **Alternative Flow** | |
|---------------------|---|
| **4a. ไม่ใช่ Admin** | ระบบตอบกลับ Forbidden (403) |
| **5a. body ไม่ถูกต้อง / parse ไม่ได้** | ระบบตอบกลับ Bad Request (400) |
| **3a. ไม่ระบุ courseId ใน path** | ระบบตอบกลับ Bad Request (400) |
| **6a. ไม่พบรายวิชา/อ่านข้อมูลไม่ได้** | ระบบตอบกลับ Error |

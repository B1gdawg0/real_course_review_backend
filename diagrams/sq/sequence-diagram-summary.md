# Sequence Diagram Summary (Brief)

เอกสารนี้สรุปความหมายของแต่ละ Sequence Diagram ใน [diagrams/sq/sequence-diagram.md](diagrams/sq/sequence-diagram.md) แบบคร่าวๆ รูปละ 1 ย่อหน้า โดยโฟกัสที่ actor หลัก, endpoint, และลำดับการทำงานระหว่าง Handler → UseCase → Repository → Database (รวมถึงบริการภายนอกถ้ามี)

## 1. Auth Flow - Login/Register
ผู้ใช้งานเข้าหน้า Login และส่งคำขอ `POST /api/v1/auth` จาก Frontend ไปที่ `AuthHandler` จากนั้น `AuthUseCase` จะตรวจสอบผู้ใช้ผ่าน `AuthRepository` ด้วยอีเมล หากไม่พบจะสร้างผู้ใช้ใหม่ในฐานข้อมูล แต่ถ้ามีอยู่แล้วจะดึงข้อมูลเดิมกลับมา แล้ว UseCase จะสร้าง JWT และส่งผลลัพธ์ (token + user) กลับไปให้ Frontend เพื่อใช้งานต่อในระบบ.

## 2. Course Flow - Get All Courses
เมื่อผู้ใช้งานเข้าหน้า Home ระบบเรียก `GET /api/v1/course?page=&size=` พร้อม JWT ไปที่ `CourseHandler` แล้วส่งต่อไป `CourseUseCase` เพื่อดึงรายวิชาทั้งหมดผ่าน `CourseRepository` ซึ่งอ่านข้อมูลจากตาราง courses และความสัมพันธ์ (tags) จากนั้น UseCase จะคำนวณ rate/score และจัดเรียงก่อนตอบกลับเป็นรายการวิชาแบบย่อพร้อมข้อมูล pagination ให้ Frontend แสดงผล.

## 3. Course Flow - Search Courses
ผู้ใช้งานพิมพ์คำค้นหาแล้ว Frontend เรียก `GET /api/v1/course?q=keyword&page=&size=` ไปที่ `CourseHandler` ซึ่งให้ `CourseUseCase` ทำการค้นหาผ่าน `CourseRepository` โดยใช้เงื่อนไขค้นหา (เช่น full-text search/ILIKE และจัดอันดับผลลัพธ์) เมื่อได้รายวิชาที่ตรงเงื่อนไขแล้ว UseCase จะคำนวณ rate/score และส่งผลลัพธ์พร้อม pagination กลับให้ Frontend.

## 4. Course Flow - Get Course By ID
เมื่อผู้ใช้งานกดดูรายละเอียดวิชา Frontend เรียก `GET /api/v1/course?id={courseId}` ไปที่ `CourseHandler` แล้วให้ `CourseUseCase` ดึงข้อมูลรายวิชาแบบเต็มจาก `CourseRepository` (รวม professors, tags, reviews) จากนั้น UseCase คำนวณค่า rate และจำนวนรีวิว ก่อนตอบกลับให้ Frontend; หลังจากได้ข้อมูลแล้ว Frontend จะส่งรีวิวไปที่ n8n webhook เพื่อสรุปรีวิว และนำข้อความสรุปมาแสดงร่วมกับรายละเอียดวิชา.

## 5. Course Flow - Compare Courses
ในหน้า Compare ผู้ใช้งานเลือก 2 วิชาแล้ว Frontend เรียก `GET /api/v1/course/compare/{courseId1}/{courseId2}` ไปที่ `CourseHandler` ซึ่ง `CourseUseCase` จะดึงข้อมูลรายวิชา 2 ตัวผ่าน `CourseRepository` และดึง “hot review” ของแต่ละวิชาผ่าน `ReviewRepository` (เช่นเรียงตาม score) จากนั้นประกอบข้อมูลเปรียบเทียบและส่งกลับเป็น `CourseCompareResponse[]` ให้ Frontend แสดงผล.

## 6. Review Flow - Get Reviews by Course ID
เมื่อผู้ใช้งานต้องการดูรีวิวของวิชา Frontend เรียก `GET /api/v1/review/c/{courseId}?page=&size=` ไปที่ `ReviewHandler` แล้วส่งต่อไป `ReviewUseCase` เพื่อดึงรีวิวผ่าน `ReviewRepository` (รวมข้อมูลผู้ใช้, tags, votes และเรียงตาม score) จากนั้น UseCase จะแปลงเป็น DTO พร้อมคำนวณค่าเพิ่มเติมเช่น avg rate และสถานะการโหวตของผู้ใช้ ก่อนส่งกลับให้ Frontend พร้อม pagination.

## 7. Review Flow - Create Review
ผู้ใช้งานเขียนรีวิวและกดส่ง Frontend เรียก `POST /api/v1/review` ไปที่ `ReviewHandler` และ `ReviewUseCase` จะตรวจสอบว่ารายวิชามีอยู่จริงผ่าน `CourseRepository` จากนั้นคำนวณ/อัปเดตค่า rate และจำนวนรีวิวของรายวิชา แล้วบันทึกรีวิวใหม่ผ่าน `ReviewRepository` (พร้อมอัปเดตตาราง courses) ก่อนตอบกลับเป็นข้อมูลรีวิวที่สร้างสำเร็จให้ Frontend.

## 8. Vote Flow - Vote Review
เมื่อผู้ใช้งานกด upvote/downvote ที่รีวิว Frontend เรียก `POST /api/v1/vote` ไปที่ `VoteHandler` แล้ว `VoteUseCase` จะตรวจสอบความถูกต้องของผู้ใช้และรีวิวผ่าน `AuthRepository` และ `ReviewRepository` จากนั้นอ่านสถานะโหวตเดิมผ่าน `VoteRepository` เพื่อเลือกทำงานแบบลบโหวต (ถ้าโหวตซ้ำค่าเดิม), อัปเดตโหวต (ถ้าคนละค่า), หรือเพิ่มโหวตใหม่ (ถ้ายังไม่เคยโหวต) เสร็จแล้วคำนวณคะแนนรีวิวใหม่จากจำนวนโหวต (เช่น Wilson score) และอัปเดต score ของรีวิวในฐานข้อมูล ก่อนส่งผลลัพธ์ให้ Frontend อัปเดต UI.

## 9. Professor Flow - Get All/One Professor
ผู้ใช้งานขอดูรายชื่ออาจารย์ทั้งหมดหรือรายละเอียดอาจารย์รายคน Frontend เรียก `GET /api/v1/professor` หรือ `GET /api/v1/professor?id={professorId}` ไปที่ `ProfessorHandler` แล้ว `ProfessorUseCase` จะเรียก `ProfessorRepository` เพื่ออ่านข้อมูลจากตาราง professors จากนั้นส่ง response กลับให้ Frontend เพื่อแสดงข้อมูล.

## 10. Tag Flow - Get All/One Tag
ผู้ใช้งานต้องการดูรายการ tag ทั้งหมดหรือ tag รายตัว Frontend เรียก `GET /api/v1/tag` หรือ `GET /api/v1/tag?id={tagId}` ไปที่ `TagHandler` แล้ว `TagUseCase` จะดึงข้อมูลผ่าน `TagRepository` จากตาราง tags และส่งผลลัพธ์กลับให้ Frontend เพื่อใช้แสดงหรือใช้ประกอบการเลือกตอนเขียนรีวิว.

## 11. Report Flow - Create Report
เมื่อผู้ใช้งานรายงานรีวิวที่ไม่เหมาะสม Frontend เรียก `POST /api/v1/report` ไปที่ `ReportHandler` แล้ว `ReportUseCase` จะตรวจสอบความถูกต้องของประเภทการรายงาน (เช่นค่า 1–4) หากไม่ถูกต้องจะตอบ error 400 แต่ถ้าถูกต้องจะเรียก `ReportRepository` เพื่อบันทึก report และอัปเดตจำนวนรายงานของรีวิว ก่อนตอบกลับว่าแจ้งรายงานสำเร็จให้ Frontend.

## 12. Admin Flow - Get All Reports
ผู้ดูแลระบบเข้าหน้าจัดการรายงาน Frontend เรียก `GET /api/v1/report` พร้อม JWT role=ADMIN ไปที่ `ReportHandler` ซึ่งจะตรวจสอบสิทธิ์ก่อน ถ้าไม่ใช่ admin จะตอบ 403 แต่ถ้าถูกต้องจะให้ `ReportUseCase` ดึงรายการรายงานทั้งหมดผ่าน `ReportRepository` (รวม join ข้อมูล user/review ตามที่ต้องใช้แสดงผล) แล้วส่งกลับเป็นรายการรายงานให้ Frontend.

## 13. Admin Flow - Manage Report (Solve/Reject)
เมื่อแอดมินกด “Solve” หรือ “Reject” สำหรับรายงานหนึ่งรายการ Frontend เรียก `PATCH /api/v1/report/{reportId}/{action}` ไปที่ `ReportHandler` และตรวจสิทธิ์ admin จากนั้น `ReportUseCase` จะดึงข้อมูล report ผ่าน `ReportRepository` แล้วแตกแขนงตาม action: ถ้า solve จะอัปเดตสถานะรายงานเป็น SOLVED และสั่ง `ReviewRepository` ลบรีวิวที่ถูกรายงาน แต่ถ้า reject จะอัปเดตสถานะเป็น REJECTED อย่างเดียว สุดท้ายตอบกลับ success เพื่อให้ Frontend อัปเดตสถานะบนหน้าจอ.

## 14. Admin Course Flow - Update Course (Status + Detail)
แอดมินแก้ไขข้อมูลรายวิชาได้ 2 แบบภายใต้สิทธิ์ ADMIN: (1) ปรับสถานะแนะนำโดยเรียก `PUT /api/v1/course` พร้อม body `{ id, rec_status }` เพื่ออัปเดตค่า `rec_status` ในตาราง courses หรือ (2) แก้ไขรายละเอียดรายวิชาตาม ID โดยเรียก `PUT /api/v1/course/update/{courseId}` ซึ่ง `CourseUseCase` จะโหลดข้อมูลเดิม, apply เฉพาะ field ที่ส่งมา และถ้ามี `professor_ids` จะ replace ความสัมพันธ์อาจารย์ ก่อนอัปเดตใน transaction แล้วตอบกลับแบบ No Content ให้ UI อัปเดตผลลัพธ์.

## 15. Admin Course Flow - Import Courses (CSV)
แอดมินอัปโหลดไฟล์ CSV ผ่าน `POST /api/v1/course/import` (multipart/form-data field `file`) ไปที่ `CourseHandler` หลังตรวจสิทธิ์ admin แล้ว `CourseUseCase` จะอ่านและ parse CSV, รวบรวม professor IDs เพื่อเช็คว่ามีอยู่จริงผ่าน `ProfessorRepository` จากนั้นสร้างรายการ courses ที่ valid และเรียก `CourseRepository` ทำ bulk create พร้อมผูกความสัมพันธ์อาจารย์ในฐานข้อมูล สุดท้ายตอบกลับเป็นผลสรุป `{ success_count, failed_count, failed_rows }` ให้ UI แสดงผลการนำเข้า.

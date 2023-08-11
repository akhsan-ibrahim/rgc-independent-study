SELECT reports.id,fullname,class,status,study,score
FROM students
INNER JOIN reports
ON students.id = reports.student_id
WHERE status = 'active' AND score < 70
ORDER BY score;
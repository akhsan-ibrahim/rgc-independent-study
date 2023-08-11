SELECT
id,
CONCAT(first_name,' ',last_name) AS fullname,
SPLIT_PART(exam_id,'-',1) AS class,
(bahasa_indonesia + bahasa_inggris + matematika + ipa)/4 AS average_score
FROM final_scores
WHERE exam_status = 'pass' AND (fee_status = 'installment' OR fee_status = 'full')
ORDER BY average_score DESC
LIMIT 5;
-- Check if all students in a group have completed their payments
SELECT g.id AS group_id, g.class_code, 
    CASE 
        WHEN COUNT(p.id) = COUNT(CASE WHEN p.payment_status = 'Accepted' THEN 1 END) 
        THEN TRUE 
        ELSE FALSE 
    END AS full_payment_completed
FROM groups g
JOIN users u ON u.group_id = g.id
LEFT JOIN payments p ON p.student_id = u.id
GROUP BY g.id, g.class_code;

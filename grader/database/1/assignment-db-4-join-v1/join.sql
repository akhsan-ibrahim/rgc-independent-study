SELECT orders.id,fullname,email,product_name,unit_price,quantity,order_date
FROM users
INNER JOIN orders
ON users.id = orders.user_id
WHERE status = 'active' AND (unit_price*quantity > 500000 OR quantity > 20);
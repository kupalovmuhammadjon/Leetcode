SELECT s.name
FROM Orders o
RIGHT JOIN Company c ON o.com_id = c.com_id AND c.name = 'RED'
RIGHT JOIN SalesPerson s ON o.sales_id = s.sales_id
WHERE o.order_id IS NULL;

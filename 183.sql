# Write your MySQL query statement below
SELECT t1.name AS Customers
FROM Customers AS t1
LEFT JOIN Orders AS t2 ON t1.id = t2.customerId
WHERE t2.customerId IS NULL;
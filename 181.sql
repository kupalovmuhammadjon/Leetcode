# Write your MySQL query statement below
SELECT t1.name as Employee
FROM Employee as t1
INNER JOIN Employee as t2 ON t2.id = t1.managerId and t1.salary > t2.salary;

-- Write your PostgreSQL query statement below
select unique_id, name from EmployeeUNI as eu right join Employees as e on eu.id = e.id
select 
    employee_id
from 
    Employees
where
    manager_id is not null and
    salary < 30000 and 
    manager_id not in(
        select employee_id
        from Employees
    )
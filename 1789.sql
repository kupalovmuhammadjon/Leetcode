
select 
    e.employee_id,
    e.department_id
from 
    Employee as e
where
    primary_flag = 'Y'
union
select 
    e.employee_id,
    e.department_id
from 
    Employee as e
where e.employee_id in(
    select 
        employee_id
    from 
        Employee
    group by
        employee_id
    having 
        count(employee_id) = 1
)



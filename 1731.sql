-- Write your PostgreSQL query statement below
with managers as (
    select 
        reports_to as manager_id,
        count(reports_to) as reports_count,
        round(avg(age)::numeric) as average_age
    from
        Employees
    group by
        reports_to 
    having
        reports_to is not null
        
)

select 
    manager_id as employee_id,
    e.name,
    reports_count,
    average_age
from 
    managers
left join
    Employees as e
on 
    manager_id = e.employee_id
order by 
    manager_id


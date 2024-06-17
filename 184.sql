select 
    d.name as Department,
    e.name as Employee,
    e.salary as Salary
from
    employee as e
join
    department as d 
on 
    d.id = e.departmentId
where
    e.salary in (
        select distinct
            salary
        from
            employee as e2
        where
            e2.departmentId = e.departmentId
        group by
            salary
        order by
            salary desc
        limit 1
    );

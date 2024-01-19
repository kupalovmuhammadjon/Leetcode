select project_id, round(avg(experience_years), 2) as average_years
from Employee as e 
join Project as p on p.employee_id = e.employee_id 
group by project_id
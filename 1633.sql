-- Write your PostgreSQL query statement below
select 
    r.contest_id,
    round((count(r.user_id)::numeric / (select count(*) from Users)) * 100, 2) 
    as percentage
from
    Register as r
left join 
    users as u 
on
    r.user_id = u.user_id
group by 
    r.contest_id
order by
    percentage desc, r.contest_id asc
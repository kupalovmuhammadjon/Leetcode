-- Write your PostgreSQL query statement below
select 
    *
from
    Users
where
    mail like '%@leetcode.com' and
    left(mail, length(mail)-13) ~ '^[A-Za-z]+[A-Za-z0-9_.-]*$';

    
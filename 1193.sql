# Write your MySQL query statement below
--  month    | country | trans_count | approved_count | trans_total_amount | approved_total_amount
--  make_date(year(trans_date), month(trans_date)) as month
select concat(year(trans_date), case when month(trans_date) < 10 then '-0' else '-' end, month(trans_date)) as month, 
    country, count(month(trans_date)) as trans_count, 
    count(case when state = 'approved' then 1 end) as approved_count,
    sum(amount) as trans_total_amount, 
    coalesce(sum(case when state = 'approved' then amount end), 0) as approved_total_amount
from Transactions
group by country, month


-- Write your PostgreSQL query statement below
--  month    | country | trans_count | approved_count | trans_total_amount | approved_total_amount
--  make_date(year(trans_date), month(trans_date)) as month
select to_char(trans_date, 'YYYY-MM') as month, 
    country, count(*) as trans_count, 
    count(case when state = 'approved' then 1 end) as approved_count,
    sum(amount) as trans_total_amount, 
    coalesce(sum(case when state = 'approved' then amount end), 0) as approved_total_amount
from Transactions
group by country, month
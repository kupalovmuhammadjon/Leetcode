with categories as(
    select 'Low Salary' as category
    union all
    select 'Average Salary'
    union all 
    select 'High Salary'
)

select 
    ct.category,
    coalesce(count(income), 0) as accounts_count
from 
    categories as ct
left join
    Accounts as a
on 
    (ct.category = 'Low Salary' AND a.income < 20000) OR
    (ct.category = 'Average Salary' AND a.income BETWEEN 20000 AND 50000) OR
    (ct.category = 'High Salary' AND a.income > 50000)
group by
    ct.category

select 
    unnest(array['Low Salary', 'Average Salary', 'High Salary']) as category,
    unnest(array[
        sum(case when income < 20000 then 1 else 0 end),
        sum(case when income >= 20000 AND income <= 50000 then 1 else 0 end),
        sum(case when income > 50000 then 1 else 0 end)
    ]) as accounts_count
from 
    Accounts

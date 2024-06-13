# Write your MySQL query statement below
with balance as(
    select 
        account, sum(amount) as balance
    from
        Transactions
    group by
        account
    having
        sum(amount) > 10000
)

select
    u.name, b.balance
from
    balance as b
join 
    Users as u
on
    b.account = u.account
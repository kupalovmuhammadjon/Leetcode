-- Write your PostgreSQL query statement below
with total as(
    select 
        count(*) as numberofProducts
    from 
        Product
),
count_per_customer as(
    select 
        customer_id, count(distinct product_key)
    from 
        Customer
    group by
        customer_id

)

select 
    cp.customer_id
from
    total as t
join 
    count_per_customer as cp
on 
    cp.count = t.numberofProducts



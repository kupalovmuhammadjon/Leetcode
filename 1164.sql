
with recentDates as(
    select 
        product_id,
        new_price,
        change_date
    from
        Products
    where
        change_date <= '2019-08-16'

),
max as (
    select 
        product_id,
        max(change_date) as md
    from
        recentDates 
    group by
        product_id
)

select 
    distinct p.product_id,
    case when m.product_id is not null 
    then 
        p.new_price
    else
        10
    end as price
from 
    Products as p
left join 
    max as m
on
    p.product_id = m.product_id and m.md = p.change_date
where
    m.product_id is not null or
    p.product_id not in ( select product_id from max )




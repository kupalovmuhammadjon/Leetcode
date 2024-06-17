with last_dates as (
    select distinct visited_on
    from 
        Customer
    order by
        visited_on
    offset 6
)

select
    ld.visited_on,
    sum(c.amount) as amount,
    round(sum(c.amount) / count(distinct c.visited_on)::numeric, 2) as average_amount
from
    last_dates as ld
join
    Customer as c
on
    c.visited_on between ld.visited_on - 6 and ld.visited_on 
group by
    ld.visited_on
order by
    ld.visited_on
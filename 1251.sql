
select 
    p.product_id, 
    coalesce(round(sum(p.price*u.units) / sum(u.units::numeric), 2), 0) as average_price
from 
    Prices as p 
left join 
    UnitsSold as u
on
    u.product_id = p.product_id and
    u.purchase_date >= p.start_date and
    u.purchase_date <= p.end_date
group by
    p.product_id

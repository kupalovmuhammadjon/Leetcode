-- Write your PostgreSQL query statement below
with  first_sale as (
    select product_id, min(year) as first_year
    from sales
    group by product_id
)
SELECT s.product_id, f.first_year, s.quantity, s.price
from first_sale as f
left join Sales as s
on s.product_id = f.product_id and s.year = f.first_year




select p.product_id, p.product_name
from Product p inner join Sales s
on p.product_id = s.product_id
group by s.product_id
having min(s.sale_date) >= '2019-01-01' and max(sale_date) <= '2019-03-31'

with first_orders as (
    select customer_id, min(order_date) as order_date
    from Delivery
    group by customer_id
)
-- round(count(case when order_date = customer_pref_delivery_date then 1 end) / count(*)::numeric * 100, 2) as immediate_percentage
select round(count(case when f.order_date = d.customer_pref_delivery_date then 1 end) / count(*)::numeric * 100, 2) as immediate_percentage
from first_orders as f
left join Delivery as d
on f.customer_id = d.customer_id and f.order_date = d.order_date
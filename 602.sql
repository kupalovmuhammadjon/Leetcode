# Write your MySQL query statement below
with num_friends as (
    select
        requester_id as id
    from
        RequestAccepted
    union all
    select
        accepter_id as id
    from
        RequestAccepted
)

select
    nf.id,
    count(nf.id) as num
from
    num_friends as nf
group by
    nf.id
order by 
    num desc
limit 1
    
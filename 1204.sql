with ordered as(
    select q1.turn, sum(q2.weight)
    from 
        Queue as q1
    join
        Queue as q2
    on q1.turn >= q2.turn
    group by q1.turn 
    having sum(q2.weight) <= 1000
    order by turn desc
    limit 1
)
select 
    person_name
from 
    Queue as q
join
    ordered as o
on 
    q.turn = o.turn


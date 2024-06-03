select 
    s1.id,
    case 
        when s2.student is not null then s2.student
        when s3.student is not null then s3.student
    else
        s1.student end as student
from
    Seat as s1
left join
    seat as s2
on
    s1.id+1 = s2.id and s1.id%2=1
left join
    seat as s3
on
    s1.id-1 = s3.id and s1.id%2=0
order by
    s1.id
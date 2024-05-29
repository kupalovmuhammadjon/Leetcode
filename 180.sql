select distinct
    l1.num as ConsecutiveNums
from 
    Logs as l1
join 
    Logs as l2
on 
    l1.id+1 = l2.id and l1.num = l2.num
join 
    Logs as l3
on 
    l2.id+1 = l3.id and l3.num = l2.num

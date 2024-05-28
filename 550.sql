-- Write your PostgreSQL query statement below
with first_login as(
    select  
        player_id, min(event_date) as min_date
    from 
        Activity 
    group by 
        player_id
),
logged_back as(
    select  
        a1.player_id as player_id, a1.event_date as event_date
    from 
        Activity as a1
    join 
        first_login as f
    on 
        a1.player_id = f.player_id 
        and 
        a1.event_date = f.min_date + 1 
)
select 
    round(count(distinct lg.player_id) / 
    coalesce(count(distinct f.player_id)::numeric, 1), 2) as fraction
from 
    first_login as f
left join 
    logged_back as lg
on 
    f.player_id = lg.player_id

-- select * 
-- from logged_back
-- round(count(lg.event_date) / count(case whenlg.event_date)::numeric, 2) as fraction
--  round(count(distinct lg.player_id2) / coalesce(count(distinct a.player_id)::numeric, 1), 2) as fraction
-- count(distinct lg.player_id2), coalesce(count(distinct a.player_id), 1)
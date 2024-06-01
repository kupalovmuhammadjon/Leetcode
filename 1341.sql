-- Write your PostgreSQL query statement below
with rated_users as(
    select
        user_id,
        count(movie_id) as votes
    from
        MovieRating
    group by
        user_id 
    
),
rated_movies as(
    select
        movie_id,
        avg(rating) as avg_rating
    from
        MovieRating
    where
        extract(year from created_at) = '2020' and extract(month from created_at) = '2'
    group by
        movie_id 

)
(select 
    u.name AS results
from
    rated_users as ru
join 
    Users as u
on
    u.user_id = ru.user_id
order by
    ru.votes desc, u.name asc
limit 1)

union all

(select  
    m.title AS results
from
    rated_movies as rm
join 
    Movies as m
on
    m.movie_id = rm.movie_id
order by
    rm.avg_rating desc, m.title asc
limit 1)
            
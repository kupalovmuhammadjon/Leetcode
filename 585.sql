-- Write your PostgreSQL query statement below
with distinct_locations as (
    select
        *
    from
        Insurance
    where
        (lat, lon) not in(
            select
                lat, lon
            from
                Insurance
            group by
                lat, lon
            having
                count(*) > 1
        )
        and
            tiv_2015 in(
            select
                tiv_2015
            from
                Insurance
            group by
                tiv_2015
            having
                count(tiv_2015) > 1
            )

)
select
    round(sum(tiv_2016)::numeric, 2) as tiv_2016
from
    distinct_locations

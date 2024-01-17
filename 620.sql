select * from Cinema
where description != 'boring' and id % 2
order by rating desc
CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) RETURNS TABLE (Salary INT) AS $$
BEGIN
  RETURN QUERY (
    -- Write your PostgreSQL query statement below.
    select distinct
        e.salary 
    from    
        Employee as e
    where
        N-1 >= 0
    order by
        e.salary desc
    offset case 
        when N-1 < 0 then 0 
        else N-1 end
    limit 1
      
  );
END;
$$ LANGUAGE plpgsql;
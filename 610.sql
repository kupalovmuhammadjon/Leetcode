SELECT x, y, z, 
       CASE WHEN x + y > z AND x < y + z AND y < x + z THEN 'Yes' ELSE 'No' END AS triangle
FROM Triangle;

    res := [][]int{}
    for i := 0; i < len(mat); i++{
        row := []int{}
        for j := 0; j < len(mat[0]); j++{
            row = append(row, mat[i][j])
        }
        res = append(res, row)
    }   
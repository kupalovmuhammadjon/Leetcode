class Solution:
    def frequencySort(self, s: str) -> str:
        dc = {}
        st = set()
        for i in s:
            if i in dc:
                dc[i] += 1
            else:
                st.add(i)
                dc[i] = 1
        
        ls = sorted(dc, key=lambda x: dc[x], reverse=True)

        res = ""
        for k in ls:
            res += k * dc[k]
       
        
        return res


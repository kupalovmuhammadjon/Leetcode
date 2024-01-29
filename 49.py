class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        
        res = []
        st = []
        d = {}
        for i in strs:
            st.append("".join(sorted(list(i))))
            d.update({"".join(sorted(list(i))) : []})
        
        for i in range(len(st)):
            d[st[i]].append(strs[i])
        
        return d.values()
            
            


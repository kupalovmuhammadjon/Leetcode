class Solution:
    def isPathCrossing(self, path: str) -> bool:
        
        st = [[0, 0]]
        crd = [0, 0]
        for i in path:
            cord = crd.copy()
            
            if i == "N":
                cord[1] += 1
                if cord in st:
                    return True

                st.append(cord)
                    
            elif i == "S":
                cord[1] -= 1
                if cord in st:
                    return True
                
                st.append(cord)
                    
            elif i == "E":
                cord[0] += 1
                if cord in st:
                    return True

                st.append(cord)
            
            elif i == "W":
                cord[0] -= 1
                if cord in st:
                    return True
                    
                st.append(cord)
                
            crd = cord.copy() 
                    
        return False
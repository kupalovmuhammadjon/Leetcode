class Solution:
    def slowestKey(self, releaseTimes: List[int], keysPressed: str) -> str:
        
        mx = 0
        ls = []
        for i in range(1, len(releaseTimes), 1):
            if mx < releaseTimes[i] - releaseTimes[i-1]:
                mx = releaseTimes[i] - releaseTimes[i-1]
                
        for i in range(1, len(releaseTimes), 1):
            if mx == releaseTimes[i] - releaseTimes[i-1]:
                ls.append(keysPressed[i-1])
         
        ls.sort()
        
        return ls[-1]
                
        
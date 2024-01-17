class Solution:
    def minOperations(self, logs: List[str]) -> int:
        
        count = 0
        
        for log in logs:
            
            if log.count(".") > 0:
                count += 1
                
        return count - 1
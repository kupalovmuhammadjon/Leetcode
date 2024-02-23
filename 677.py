class MapSum:

    def __init__(self):
        self.dc = dict()
        

    def insert(self, key: str, val: int) -> None:
        self.dc[key] = val

    def sum(self, prefix: str) -> int:
        sm = 0
        for i in self.dc.keys():
            if i.startswith(prefix):
                sm += self.dc[i]
        
        return sm



# Your MapSum object will be instantiated and called as such:
# obj = MapSum()
# obj.insert(key,val)
# param_2 = obj.sum(prefix)
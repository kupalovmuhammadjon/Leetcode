class RandomizedCollection:

    def __init__(self):
        self.st = []

    def insert(self, val: int) -> bool:
        if val in self.st:
            self.st.append(val)
            return False
        self.st.append(val)
        return True

    def remove(self, val: int) -> bool:
        if val in self.st:
            self.st.remove(val)
            return True
        return False

    def getRandom(self) -> int:
        return random.choice(self.st)

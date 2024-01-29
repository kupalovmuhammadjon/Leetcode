class Solution:
    def checkbox(self, xs, xf, ys, yf):

        arr = []
        for i in range(xs, xf, 1):
            for j in range(ys, yf, 1):
                if self.b[i][j].isdigit():
                    arr.append(int(self.b[i][j]))

        if len(arr) != len(set(arr)):
            return 0
        
        return 1         

    def isValidSudoku(self, board: List[List[str]]) -> bool:
        self.b = board
        l = len(board)
        for i in range(l):
            temp = []
            for j in range(l):
                if board[i][j].isdigit():
                    temp.append(int(board[i][j]))
            if len(temp) != len(set(temp)):
                return 0

        for i in range(l):
            temp = []
            for j in range(l):
                if board[j][i].isdigit():
                    temp.append(int(board[j][i]))
            if len(temp) != len(set(temp)):
                return 0


        x = self.checkbox(0, 3, 0, 3)
        if not x:
            return x
        x = self.checkbox(0, 3, 3, 6)
        if not x:
            return x
        x = self.checkbox(0, 3, 6, 9)
        if not x:
            return x
        x = self.checkbox(3, 6, 0, 3)
        if not x:
            return x
        x = self.checkbox(3, 6, 3, 6)
        if not x:
            return x
        x = self.checkbox(3, 6, 6, 9)
        if not x:
            return x
        x = self.checkbox(6, 9, 0, 3)
        if not x:
            return x
        x = self.checkbox(6, 9, 3, 6)
        if not x:
            return x
        x = self.checkbox(6, 9, 6, 9)
            
        return x
            
                


        
class Solution:
    def reformatDate(self, date: str) -> str:
        months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
        date = date.split()
        
        res = date[-1] + '-'
        
        match months.index(date[1]):
            case 0:
                res += "01"
            case 1:
                res += "02"
            case 2:
                res += "03"
            case 3:
                res += "04"
            case 4:
                res += "05"
            case 5:
                res += "06"
            case 6:
                res += "07"
            case 7:
                res += "08"
            case 8:
                res += "09"
            case 9:
                res += "10"
            case 10:
                res += "11"
            case 11:
                res += "12"
        
        if "t" not in date[0][:2] and "r" not in date[0][:2] and "s" not in date[0][:2] and "n" not in date[0][:2]:
            res += "-" + date[0][:2]
        else:
            res += "-0" + date[0][0]
            

        return res
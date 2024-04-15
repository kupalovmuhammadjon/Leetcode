func findLatestTime(s string) string {
    result := ""
    if s[0] == '?'{
        if s[1] == '?'|| s[1] == '1' || s[1] == '0'{
           result += "1" 
        }else{
            result += "0"
        }
    }else{
        result += string(s[0])
    }
    if s[1] == '?'{
        if s[0] == '0'{
            result += "9"   
        }else{
            result += "1"  
        }
    }else {
        result += string(s[1])
    }
    result += ":"
    if s[3] == '?'{
        result += "5"
    }else{
        result += string(s[3])
    }
    if s[4] == '?'{
        result += "9"
    }else{
        result += string(s[4])
    }
    return result
}
func isValid(s string) bool {
    if (len(s) < 2) {
        return false
    }
    
    stack := make([]rune, 0)
    
    for _, ch := range s {
        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack, ch)
            continue
        } 
        
        if len(stack) < 1  {
            return false
        } else {
            if ch == ']' && stack[len(stack)-1] != '[' {
                return false
            } else if ch == ')' && stack[len(stack)-1] != '(' {
                return false
            } else if ch == '}' && stack[len(stack)-1] != '{' {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    
    if len(stack) > 0 {
        return false
    }
    
    return true
}
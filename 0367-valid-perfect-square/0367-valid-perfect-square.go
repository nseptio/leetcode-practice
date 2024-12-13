func isPerfectSquare(num int) bool {
    if num == 1 {
        return true
    }
    
    for i:= 0; i < num; i++ {
        square := i*i
        if square > num {
            return false
        } else if square == num {
            return true
        } 
    }
    
    return false
}
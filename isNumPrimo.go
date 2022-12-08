package isNumPrimo

func isNumPrimo(num int) bool{
    var div int
    for div=2; div<num; div++{
        if num%div==0{//numero non primo
            break
        }
    }
    if div<num{//numero non primo
        return false
    }else{//numero primo
        return true
    }
}

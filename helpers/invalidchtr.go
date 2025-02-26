package helpers


func Invalide (str string ) bool{

for _, v := range str {
	if v > 126 || v < 32 {
		return true  
	}
}


	return false 
}
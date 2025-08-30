package main

func main(){
	flowerGame(3,2)
}

func flowerGame(n int, m int) int64 {

	x_limit := n
	y_limit := m
	 
	var oddX int = 0 
	var oddY int = 0
	var evenX int = 0 
	var evenY int = 0

	oddX = (x_limit + 1) / 2
	evenX = x_limit / 2 

	oddY = (y_limit + 1) / 2
	evenY = y_limit / 2 

	pairs := (oddX * evenY) + (evenX * oddY)

	return int64(pairs)
	
}

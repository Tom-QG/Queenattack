package main

import "strconv"
import "fmt"
import "reflect"

func Buildempty() [][]string {
	board := make([][]string, 10)

	board[0] = []string{" ", "a", "b", "c", "d", "e", "f", "g", "h", " "}
	board[8] = []string{strconv.Itoa(8), "_", "_", "_", "_", "_", "_", "_", "_", strconv.Itoa(8)}
	board[7] = []string{strconv.Itoa(7), "_", "_", "_", "_", "_", "_", "_", "_", strconv.Itoa(7)}
	board[6] = []string{strconv.Itoa(6), "_", "_", "_", "_", "_", "_", "_", "_", strconv.Itoa(6)}
	board[5] = []string{strconv.Itoa(5), "_", "_", "_", "_", "_", "_", "_", "_", strconv.Itoa(5)}
	board[4] = []string{strconv.Itoa(4), "_", "_", "_", "_", "_", "_", "_", "_", strconv.Itoa(4)}
	board[3] = []string{strconv.Itoa(3), "_", "_", "_", "_", "_", "_", "_", "_", strconv.Itoa(3)}
	board[2] = []string{strconv.Itoa(2), "_", "_", "_", "_", "_", "_", "_", "_", strconv.Itoa(2)}
	board[1] = []string{strconv.Itoa(1), "_", "_", "_", "_", "_", "_", "_", "_", strconv.Itoa(1)}
	board[9] = []string{" ", "a", "b", "c", "d", "e", "f", "g", "h", " "}

	return board
}

type move struct {
    movelet string 
	movenum int 
}

func Addqueenwhite(board [][]string, w  move) [][]string {
for i:= 1 ; i<9 ; i++ {
if board[0][i] == w.movelet {
	for j:=1 ; j<9 ; j++ {
		if board[j][0] == strconv.Itoa(w.movenum) {
			board[j][i] = "W"
		}
	}
}
}
return board
}

func Addqueenblack(board [][]string, b  move) [][]string {
	for i:= 1 ; i<9 ; i++ {
	if board[0][i] == b.movelet {
		for j:=1 ; j<9 ; j++ {
			if board[j][0] == strconv.Itoa(b.movenum) {
				board[j][i] = "B"
			}
		}
	}

}
return board 
}

func toCharStr(i int) string {
    return string('a' - 1 + i)
}


func toNumber(s string ) int {
	var x int 
	if s == "a" {x= 1}
	if s == "b" {x= 2}
	if s == "c" {x= 3}
	if s == "d" {x= 4}
	if s == "e" {x= 5}
	if s == "f" {x= 6}
	if s == "g" {x= 7}
	if s == "h" {x= 8}
	return x 
}
func returnposition(board [][]string) move {
     m  := move{}  
	 for i:= 1; i < 9 ; i++ {
		for j:= 1 ; j < 9 ; j++ {
            if  board[i][j] != "_" {
	            m.movelet = toCharStr(j) 
	            m.movenum = i 
				break 
            }
		}
		
	}
	return m 
}

func attackHorVet(w [][]string, b [][]string) bool {
	attack := false 
    positionw := returnposition(w)
	positionb := returnposition(b)
	if positionw.movelet == positionb.movelet || positionw.movenum == positionb.movenum {
		attack = true 
	}
	return attack 
}



func attackDiag(w [][]string, b [][]string) bool {
	attack := false 
	positionw := returnposition(w)
	positionb := returnposition(b)
	blet := toNumber(positionb.movelet)
	wlet := toNumber(positionw.movelet)
	easyb := []int{blet, positionb.movenum} 
	 for i:= 1 ; i < 8 ; i ++ {
		 if reflect.DeepEqual(easyb, []int{wlet + i , positionw.movenum + i } ) {
			attack = true 
		 }
	 }
	 for j:= 1 ; j < 8 ; j ++ {
		if  reflect.DeepEqual(easyb, []int{wlet + j , positionw.movenum - j } ){
			attack = true 
		}
	 }
	 for k:= 1 ; k < 8 ; k ++ {
		if  reflect.DeepEqual(easyb, []int{wlet - k , positionw.movenum + k } ){
			attack = true
		}
	 }
	 for p:= 1 ; p < 8 ; p ++ {
		if    reflect.DeepEqual(easyb,  []int{wlet -p , positionw.movenum -p } ){
			attack = true 
		}
	 }

return attack 

}

func main() {
	emptyboardw := Buildempty()
	white := move{movelet: "c", movenum: 4}
  boardw :=  Addqueenwhite(emptyboardw, white)

  emptyboardb := Buildempty()
  black := move{movelet: "e", movenum: 4}
  boardb := Addqueenblack(emptyboardb, black)
	
  for i :=0 ; i<10 ; i++{
		fmt.Println(boardw[i])
	}
	for i :=0 ; i<10 ; i++{
		fmt.Println(boardb[i])
	}


	fmt.Println(attackHorVet(boardw, boardb) || attackDiag(boardw,boardb)) 
}
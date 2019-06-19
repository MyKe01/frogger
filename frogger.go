package main 
import . "g2d"
import "strconv"

var frogfinish = LoadImage("./Images/frogfinish.png")
var background = LoadImage("./Images/frogger_bg.png")
var screen = Size{416,512}
var frog_1 = LoadImage("./Images/frog_1.png")
var character = NewFrog()

var camion1 = NewObstacle(Point{0, 416}, 64, 8)
var camion2 = NewObstacle(Point{224, 416}, 64, 8)
var truck_right = LoadImage("./Images/truck_right.png")

var car1 =  NewObstacle(Point{0, 384}, 32, -16 )
var car = LoadImage("./Images/car1-left.png")

var car2 = NewObstacle(Point{32, 352}, 32, 4)
var car3 = NewObstacle(Point{192, 352}, 32, 4)
var car4 = NewObstacle(Point{352, 352}, 32, 4)
var tractor = LoadImage("./Images/car2-right.png")

var car5 = NewObstacle(Point{0, 320}, 32, -8)
var car6 = NewObstacle(Point{256, 320}, 32, -8)
var sportcar = LoadImage("./Images/car3-left.png")

var car7 = NewObstacle(Point{0, 288}, 32, 16)
var car8 = NewObstacle(Point{256, 288}, 32, 16)
var greycar = LoadImage("./Images/car4-right.png")

var log1 = NewObstacle(Point{0, 224}, 96, 4)
var log2 = NewObstacle(Point{256, 224}, 96, 4)

var log3 = NewObstacle(Point{0, 160}, 96, -8)
var log4 = NewObstacle(Point{160, 160}, 96, -8)
var log5 = NewObstacle(Point{320 , 160}, 96, -8)

var log6 = NewObstacle(Point{224 , 96}, 96, -16)
var log = LoadImage("./Images/log.png")

var turtle_3 = NewObstacle(Point{0, 192}, 96, -4)
var turtle4 = NewObstacle(Point{256, 192}, 32, -4)

var turtle5 = NewObstacle(Point{0, 128}, 32, 8)
var turtle_2= NewObstacle(Point{288, 128}, 64, 8)

var turtle = LoadImage("./Images/turtle.png")
var turtles_2 = LoadImage("./Images/2turtles.png") 
var turtles_3 =LoadImage("./Images/3turtles.png")

var time = LoadImage("./Images/Time.png")
var score = LoadImage("./Images/score.png")
var points int = 0 
var count = 0
var hp int = 5
var life = LoadImage("./Images/hp.png")
var gameover= LoadImage("./Images/gameover1.png")
var win = LoadImage("./Images/win.png")

var clock = 0.0
var len = 320

type frog struct { 										//frog functions
    x, y    int
    w, h 	int
    dx, dy  int
}

func NewFrog() *frog {
    b := &frog{192,448, 32, 32, 32, 32}
    return b
}

func (f *frog) Move(){
	if KeyPressed("ArrowLeft"){
		f.x-=f.dx
    }else if KeyPressed("ArrowRight"){
		f.x+=f.dx
    } else if KeyPressed("ArrowUp"){
		if f.y == 0{
    		f.y -= 0
    	}else{
    		f.y-=f.dy
			points += 100
    	}
    } else if KeyPressed("ArrowDown"){
    	if f.y == 448{
    		f.y += 0
    	}else{
    		f.y+=f.dy
    	}
    }
    f.PacManFrog()
}

func (f *frog) PacManFrog (){
	if f.x > screen.W{
		character.x=192
   		character.y=448
   		hp-=1
   		points -=100
	}else if f.x + f.w <= 0 {
		character.x=192
   		character.y=448
   		hp -=1
   		points -=100
	}
}

type obstacle struct{										//obstacles functions
	x, y	int 
	w, h 	int 
	dx		int
}

func NewObstacle(pos Point, w int, dx int) *obstacle {
    b := &obstacle{ pos.X, pos.Y, w, 32, dx}
    return b
}

func (o *obstacle) Contact() {
	fr:= Rect{character.x, character.y, character.w, character.h}
   	if fr.Y > screen.H/2 && ((fr.X >= o.x && fr.X + fr.W <= o.x+o.w) || (fr.X< o.x && fr.X + fr.W > o.x) || (fr.X >= o.x && fr.X + fr.W >= o.x + o.w && fr.X < o.x + o.w)) && (fr.Y==o.y && fr.Y +fr.H == o.y + o.h){
		character.x=192
   		character.y=448
		points -= 100
		hp -= 1
   	}
   	if o.dx > 0 && fr.Y <=288 {
   		if ((fr.X >= o.x && fr.X + fr.W <= o.x+o.w) || (fr.X< o.x && fr.X + fr.W > o.x)) && (fr.Y==o.y){
   			character.x += o.dx
   		}
   	} else if o.dx < 0 && fr.Y <=288{
   		if ((fr.X >= o.x && fr.X + fr.W <= o.x+o.w) || (fr.X >= o.x && fr.X + fr.W >= o.x + o.w && fr.X < o.x + o.w)) && (fr.Y==o.y){
   			character.x += o.dx  
   		}
   	}
}

func NotContact() {
	fr:= Rect{character.x, character.y, character.w, character.h}
	if fr.Y<= 288{
		if (( fr.X + fr.W < log1.x && fr.X > log2.x + log2.w) || (fr.X + fr.W < log2.x && fr.X > log1.x + log1.w) || (fr.X > log2.x + log2.w && fr.X + fr.W < log1.x) || (fr.X + fr.W < log1.x && fr.X +fr.W < log2.x)) && (fr.Y==log1.y) {
			character.x=192
   			character.y=448
			points -= 100
			hp -= 1
		}// turtle_3turtle_3turtle_3               turtle4              
		if (( fr.X + fr.W < turtle4.x && fr.X > turtle_3.x + turtle_3.w) || (fr.X > turtle4.x + turtle4.w && (fr.X + fr.W < turtle_3.x || fr.X > turtle_3.x + turtle_3.w)) || (fr.X + fr.W < turtle4.x && fr.X + fr.W < turtle_3.x)) && (fr.Y==turtle4.y) {
			character.x=192
   			character.y=448
			points -= 100
			hp -= 1
		}//log3 log4 log5
		if (( fr.X + fr.W < log5.x && fr.X > log4.x + log4.w && fr.X > log3.x + log3.w) || (fr.X + fr.W < log3.x && fr.X + fr.W <  log4.x && fr.X + fr.W < log5.x) || (fr.X > log3.x+log3.w && fr.X + fr.W < log4.x && fr.X + fr.W < log5.x) || (fr.X > log5.x + log5.w && fr.X + fr.W < log3.x )) && (fr.Y==log5.y) {
			character.x=192
   			character.y=448
   			points -= 100
   			hp -= 1
		}//turtle5             turtle_2turtle_2
		if (( fr.X + fr.W < turtle_2.x && fr.X > turtle5.x + turtle5.w) || (fr.X > turtle_2.x + turtle_2.w && (fr.X + fr.W < turtle5.x || fr.X > turtle5.x + turtle5.w)) || (fr.X + fr.W < turtle_2.x && fr.X + fr.W < turtle5.x)) && (fr.Y==turtle5.y) {
			character.x=192
   			character.y=448
   			points -= 100
   			hp -= 1
		}//                log6             
		if (( fr.X + fr.W < log6.x) || (fr.X > log6.x + log6.w)) && (fr.Y==log6.y) {
			character.x=192
   			character.y=448
   			points -= 100
   			hp -= 1
		}
	}
}

func finish(){
	if character.y <= 64{
		character.x=192
   		character.y=448
   		count +=1
	}
	if count ==1 {
		DrawImage(frogfinish, Point{-1, 64})
	}else if count == 2 {
		DrawImage(frogfinish, Point{-1, 64})
		DrawImage(frogfinish, Point{94, 64})
	}else if count == 3{
		DrawImage(frogfinish, Point{-1, 64})
		DrawImage(frogfinish, Point{94, 64})
		DrawImage(frogfinish, Point{190, 64})
	}else if count == 4{
		DrawImage(frogfinish, Point{-1, 64})
		DrawImage(frogfinish, Point{94, 64})
		DrawImage(frogfinish, Point{190, 64})
		DrawImage(frogfinish, Point{285, 64})
	}else if count == 5{
		DrawImage(frogfinish, Point{-1, 64})
		DrawImage(frogfinish, Point{94, 64})
		DrawImage(frogfinish, Point{190, 64})
		DrawImage(frogfinish, Point{285, 64})
		DrawImage(frogfinish, Point{382, 64})	
	}
}

func (o *obstacle) MoveObstacle () {
    o.x += o.dx
    o.PacManObstacles()
    o.Contact()
    NotContact()
}

func (o *obstacle) PacManObstacles() {
	if o.x > screen.W + 32 {
		o.x = 0- 96	
	}else if o.x + 96 <= 0 {
		o.x = screen.W
	}
}

func tick(){
	ClearCanvas()
	DrawImage(background, Point{0,0})
	DrawImage(time, Point{0,480})
	DrawImage(score, Point{0,8})
	SetColor(Color{0,0,0})
	p:= strconv.Itoa(points)
	DrawText(p, Point{114, 10}, 24)
	DrawImage(life, Point{screen.W/2,8})
	SetColor(Color{0,0,0})
	h:=strconv.Itoa(hp)
	DrawText(h, Point{screen.W/2 + 57, 10}, 24)
	camion1.MoveObstacle()
	camion2.MoveObstacle()
	car1.MoveObstacle()
	car2.MoveObstacle()
	car3.MoveObstacle()
	car4.MoveObstacle()
	car5.MoveObstacle()
	car6.MoveObstacle()
	car7.MoveObstacle()
	car8.MoveObstacle()
	log1.MoveObstacle()
	log2.MoveObstacle()
	log3.MoveObstacle()
	log4.MoveObstacle()
	log5.MoveObstacle()
	log6.MoveObstacle()
	turtle_2.MoveObstacle()
	turtle_3.MoveObstacle()
	turtle4.MoveObstacle()
	turtle5.MoveObstacle()
	character.Move()
	DrawImage(turtles_2, Point{turtle_2.x, turtle_2.y})
	DrawImage(turtles_3, Point{turtle_3.x, turtle_3.y})
	DrawImage(turtle, Point{turtle4.x, turtle4.y})
	DrawImage(turtle, Point{turtle5.x, turtle5.y})
	DrawImage(log, Point{log1.x, log1.y})
	DrawImage(log, Point{log2.x, log2.y})
	DrawImage(log, Point{log3.x, log3.y})
	DrawImage(log, Point{log4.x, log4.y})
	DrawImage(log, Point{log5.x, log5.y})
	DrawImage(log, Point{log6.x, log6.y})
	DrawImage(truck_right, Point{camion1.x, camion1.y})
	DrawImage(truck_right, Point{camion2.x, camion2.y})
	DrawImage(car, Point{car1.x, car1.y})
	DrawImage(tractor, Point{car2.x, car2.y})
	DrawImage(tractor, Point{car3.x, car3.y})
	DrawImage(tractor, Point{car4.x, car4.y})
	DrawImage(sportcar, Point{car5.x, car5.y})
	DrawImage(sportcar, Point{car6.x, car6.y})
	DrawImage(greycar, Point{car7.x, car7.y})
	DrawImage(greycar, Point{car8.x, car8.y})
	DrawImage(frog_1, Point{character.x, character.y})
	finish()
	SetColor(Color{255, 255, 255})
	FillRect(Rect{screen.W-320,screen.H-22,len,22})
	clock+=0.1
	if(clock >= 1){
		clock = 0
		len--
		}
	if hp <= 0 || len <= 0{
		DrawImage(gameover, Point{0,0})
		points = 0
		count = 0 
		if KeyPressed("Spacebar"){
			hp = 5
			len = 320
			DrawImage(gameover, Point{416,512})
		}
	}
	if count == 5 && hp > 0{
		DrawImage(win, Point{0,0})
		points = 0
		hp = 5
		if KeyPressed("Spacebar"){
			count = 0 
			len = 320
			DrawImage(win, Point{416,512})
		}
	}
}

func main (){
	InitCanvas(screen)
	MainLoop(tick)
}
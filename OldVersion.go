package main

import (
        "bufio"
        "fmt"
        "log"
        "os"

        "code.google.com/p/draw2d/draw2d"
        "image"
        "image/color"
        "image/png"
        "image/gif"
        "math"
  "strconv"
//  "reflect"
)

const (
//        w, h = 128, 128
//        w, h = 32, 32
        w, h = 16, 16
)

var (
  lastTime int64
  folder   = "./"
  a=0.0
//  gc    *draw2d.ImageGraphicContext
//  i     *image.RGBA
  gc    draw2d.GraphicContext
  i     image.Image
)

func initGc(w, h int) (image.Image, draw2d.GraphicContext) {
        i := image.NewRGBA(image.Rect(0, 0, w, h))
        gc := draw2d.NewGraphicContext(i)

        gc.SetStrokeColor(image.Black)
        gc.SetFillColor(image.White)
        // fill the background
        //gc.Clear()

        return i, gc
}

func saveToPngFile(TestName string, m image.Image) {
  filePath := folder + TestName + ".png"
  f, err := os.Create(filePath)
  if err != nil {
    log.Println(err)
    os.Exit(1)
  }
  defer f.Close()
  b := bufio.NewWriter(f)
  err = png.Encode(b, m)
  if err != nil {
    log.Println(err)
    os.Exit(1)
  }
  err = b.Flush()
  if err != nil {
    log.Println(err)
    os.Exit(1)
  }
  fmt.Printf("Wrote %s OK.\n", filePath)
}

func saveToGifFile(TestName string, img image.Image) {
  fileName := folder + TestName + ".gif"
  out, err := os.Create(fileName)
  if err != nil {
    log.Println(err)
    os.Exit(1)
  }
  defer out.Close()
  var opt gif.Options
  opt.NumColors = 256
  err = gif.Encode(out, img, &opt)
  if err != nil {
    log.Println(err)
    os.Exit(1)
  }
  fmt.Printf("Wrote %s OK.\n", fileName)
}

type Coin struct{
  x,y,z float64
}

type Poin struct{
  x,y float64
}
/*
type Plan struct{
  a,b,c,d uint8
}
*/
type Plan [4]uint8

var cube[8] Coin
//var plan[6] Plan
var plan[6] Plan

func (p Plan) ValueIn4(in Plan) bool {
  var i,j,cpte,valeur uint8
  for i=0;i<4;i++{
    valeur=in[i]
    for j=0;j<4;j++{
      if p[j]==valeur{cpte++}
    }
  }
  if cpte==4{return true}else{return false}
}

func (p Plan) ValueIn2(in Plan) bool {
//  fmt.Printf("Comparatif %v - %v\n",in,p)
  var i,j,cpte,valeur uint8
  cpte=0
  for i=0;i<2;i++{
    valeur=in[i]
    for j=0;j<4;j++{
      if p[j]==valeur{cpte++;}
    }
  }
  if cpte==2{return true}else{return false}
//  return false
}

func (p Plan) ValueIn1(in uint8) bool {
//  fmt.Printf("Comparatif %v - %v\n",in,p)
  var j uint8
  for j=0;j<4;j++{
    if p[j]==in{return true}
  }
  return false
}

// http://connor-johnson.com/2014/06/21/using-matrices-in-golang/

func DrawPlan(gc draw2d.GraphicContext,num uint8){
    gc.SetFillColor(color.NRGBA{0xff,0x00,0x00,0xff})
  var couleur color.NRGBA
  switch(num){
    case 0:couleur=color.NRGBA{0xff,0x00,0x00,0xff}
    case 1:couleur=color.NRGBA{0x00,0xff,0x00,0xff}
    case 2:couleur=color.NRGBA{0x00,0x00,0xff,0xff}
    case 3:couleur=color.NRGBA{0x00,0xff,0xff,0xff}
    case 4:couleur=color.NRGBA{0xff,0x00,0xff,0xff}
    case 5:couleur=color.NRGBA{0xff,0xff,0x00,0xff}
  }
//    gc.SetFillColor(color.NRGBA{0xff,0x00,0x00,0xff})
    gc.SetFillColor(couleur)

//fmt.Println("entree")

    var coinNum uint8
    coinNum=plan[num][0];//fmt.Printf("dp %d - %d %v\n",num,coinNum,cube[coinNum]);
      gc.MoveTo(cube[coinNum].x,cube[coinNum].y)
    coinNum=plan[num][1];//fmt.Printf("dp %d - %d %v\n",num,coinNum,cube[coinNum]);
      gc.LineTo(cube[coinNum].x,cube[coinNum].y)
    coinNum=plan[num][2];//fmt.Printf("dp %d - %d %v\n",num,coinNum,cube[coinNum]);
      gc.LineTo(cube[coinNum].x,cube[coinNum].y)
    coinNum=plan[num][3];//fmt.Printf("dp %d - %d %v\n",num,coinNum,cube[coinNum]);
      gc.LineTo(cube[coinNum].x,cube[coinNum].y)
    gc.Close()
    gc.FillStroke()

//fmt.Println("sortie")
}

func Cube(degX,degY,degZ float64,ref int32){
/*
  var degX,degY,degZ float64
  degX=45.0
  degY=45.0
  degZ=00.0
*/
  angX:=degX * (math.Pi / 180.0)
  angY:=degY * (math.Pi / 180.0)
  angZ:=degZ * (math.Pi / 180.0)

//  a=w/2
  a=5
  cube[0]=Coin{-a,-a, a}
  cube[1]=Coin{ a,-a, a}
  cube[2]=Coin{ a, a, a}
  cube[3]=Coin{-a, a, a}

  cube[4]=Coin{-a,-a,-a}
  cube[5]=Coin{ a,-a,-a}
  cube[6]=Coin{ a, a,-a}
  cube[7]=Coin{-a, a,-a}

  plan[0]=Plan{0,1,2,3}
  plan[1]=Plan{3,2,6,7}
  plan[2]=Plan{7,6,5,4}
  plan[3]=Plan{4,5,1,0}
  plan[4]=Plan{1,5,6,2}
  plan[5]=Plan{4,0,3,7}

  for i:=0;i<8;i++{
    var x,y,z float64

    y=cube[i].y*math.Cos(angX)-cube[i].z*math.Sin(angX)   //rotation axe X
    z=cube[i].y*math.Sin(angX)+cube[i].z*math.Cos(angX)
    cube[i].y=y
    cube[i].z=z

    x=cube[i].x*math.Cos(angY)-cube[i].z*math.Sin(angY)   //rotation axe Y
    z=cube[i].x*math.Sin(angY)+cube[i].z*math.Cos(angY)
    cube[i].x=x
    cube[i].z=z

    x=cube[i].x*math.Cos(angZ)-cube[i].y*math.Sin(angZ)   //rotation axe Z
    y=cube[i].x*math.Sin(angZ)+cube[i].y*math.Cos(angZ)
    cube[i].x=x
    cube[i].y=y
/*
//    fmt.Printf("%d - %v\n",i,cube[i])
    x=cube[i].x*(20.0+z)/25.0
    y=cube[i].y*(20.0+z)/25.0
    cube[i].x=x
    cube[i].y=y
//    fmt.Printf("%d - %v\n",i,cube[i])
*/
  }

  // analyse
  var z float64
  var cpte uint8
  var plat Plan //={0,0,0,0}
  z=0
  cpte=0
  var j uint8
  for j=0;j<8;j++{
//    fmt.Printf("%d - %v\n",j,cube[j])

    if cube[j].z>z {z=cube[j].z;plat[0]=j
    cpte=1;}else if z==cube[j].z{z=cube[j].z;plat[cpte]=j;cpte++;}

//    fmt.Printf("%d - %d  %f\n",j,cpte,z)
  }

  i, gc = initGc(w, h)
//on met du fond blanc
  gc.SetLineWidth(0.0)
  gc.SetFillColor(color.NRGBA{0xff,0xff,0xff,0xff})
//  gc.Rect(0, 0, w-1, w-1)
  gc.MoveTo(0,0)
  gc.LineTo(w,0)
  gc.LineTo(w,w)
  gc.LineTo(0,w)
  gc.Close()
  gc.FillStroke()
  //    fmt.Println("gc type:",reflect.TypeOf(i)) //  gc
  gc.Translate(8,8)
  gc.SetLineWidth(0.0)

  if cpte==4{
    fmt.Printf("Valeur4: %v\n",plat)
    for j=0;j<6;j++{
      if plan[j].ValueIn4(plat)==true{break}
    }
    fmt.Printf("On trouve le plan %d - %v\n",j,plan[j])
    DrawPlan(gc,j)
  }else if cpte==2{
    fmt.Printf("Valeur2 %d - %v\n",j,plat)
    for j=0;j<6;j++{
      if plan[j].ValueIn2(plat)==true{
        fmt.Printf("On trouve le plan %d - %v\n",j,plan[j])
        DrawPlan(gc,j)
      }
    }
  }else if cpte==1{
    fmt.Printf("Valeur1 %d - %v\n",j,plat)
    for j=0;j<6;j++{
      if plan[j].ValueIn1(plat[0])==true{
        fmt.Printf("On trouve le plan %d - %v\n",j,plan[j])
        DrawPlan(gc,j)
      }
    }
//    DrawPlan(gc,4)
//    DrawPlan(gc,1)
//    DrawPlan(gc,0)
  }
//  saveToPngFile("cube"+strconv.Itoa(ref), i)
  saveToPngFile(fmt.Sprintf("cube%06d",ref),i)
}

func TestStar() {
//  var deg float64
  var deg int
  step:=5

  for deg = 0; deg < 60; deg+=step {
  i, gc := initGc(w, h)
//  BlackBg(i)

  gc.SetLineWidth(0.0)
  gc.SetFillColor(color.NRGBA{0xff,0xff,0xff,0xff})
//  gc.Rect(0, 0, w-1, w-1)
  gc.MoveTo(0,0)
  gc.LineTo(w,0)
  gc.LineTo(w,w)
  gc.LineTo(0,w)
  gc.Close()
  gc.FillStroke()

  angle:=30.0 * (math.Pi / 180.0)
  gc.Translate(a,a)
  gc.Rotate(float64(deg) * (math.Pi / 180.0))

  gc.MoveTo(0                 ,-a)    // en haut
  gc.RLineTo(a*math.Cos(angle),a*math.Sin(angle))
  gc.RLineTo(0                ,2*a*math.Sin(angle))
  gc.LineTo( 0                ,a)     // En bas
  gc.RLineTo(-a*math.Cos(angle),-a*math.Sin(angle))
  gc.RLineTo(0                ,-2*a*math.Sin(angle))
  gc.Close()
  gc.SetLineWidth(1.0)
  gc.SetFillColor(color.NRGBA{0xff,0x00,0x00,0xff}) // last value: 0xff no transparency
                                                   // bleu acier: 33, 103, 146
//  gc.SetStrokeColor(image.Black)    // couleur du trait
  gc.FillStroke()

  saveToPngFile("tutu"+strconv.Itoa(deg+100), i)
//  saveToGifFile("tata"+strconv.Itoa(deg+100), i)
//  saveToPngFile("tutu"+strconv.FormatFloat(deg,'g', 1, 64), i)
  }
}
/*
func BlackBg(m image.Image) {
//	size := m.Bounds().Size()
	for x := 0; x < w; x++ {
		for y := 0; y < w; y++ {
		  color := color.RGBA{0x00,0x00,0x00,255}
		  m.Set(x, y, color)
		}
	}
}
*/

func TestFillStroke() {
  i, gc := initGc(w, h)
  gc.MoveTo(128.0, 25.6)
  gc.LineTo(230.4, 230.4)
  gc.RLineTo(-102.4, 0.0)
  gc.CubicCurveTo(51.2, 230.4, 51.2, 128.0, 128.0, 128.0)
  gc.Close()

  gc.MoveTo(64.0, 25.6)
  gc.RLineTo(51.2, 51.2)
  gc.RLineTo(-51.2, 51.2)
  gc.RLineTo(-51.2, -51.2)
  gc.Close()

  gc.SetLineWidth(10.0)
  gc.SetFillColor(color.NRGBA{255, 0x33, 0x33, 0x80})
  gc.SetStrokeColor(image.Black)
  gc.FillStroke()
  saveToPngFile("TestFillStroke", i)
}

/*debug
Valeur1 8 - [2 0 0 0]
On trouve le plan 0 - [0 1 2 3] rouge
On trouve le plan 1 - [3 2 6 7] vert
On trouve le plan 4 - [1 5 6 2] cyan

*/

func main() {
//  a=h/math.Sqrt(2)
  a=w/2
//  TestStar()
  var degX,degY,degZ float64
  var ref int32=1
/*
  degX=05.0
  degY=45.0
  degZ=45.0

  for degX = 0; degX < 360; degX+=5 {
    Cube(degX,degY,degZ,ref)
    ref++
  }
*/
  degX,degY,degZ=0,0,0
  for i:= 0; i < 180; i++ {
    degX+=6
    degY+=4
    degZ+=2
    Cube(degX,degY,degZ,ref)
    ref++
  }
//  TestFillStroke()
}


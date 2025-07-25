```go
####################################### Start.. Basic Go #################################################
package main
import "fmt" //import libary เพื่อใช้วงาน fmt.Println(i)
func main() {
	var i int = 20
	fmt.Println(i)
}
การประกาศตัวแปร
	- var ชื่อตัวแปร ชนิดข้อมูล
	- ยกตัวอย่างเช่น var i int
		- จะมีค่าเริ่มต้นเป็น 0 
		- ถ้าอยากให้มีค่า deafult เป็นตัวที่เรากำหนดให้ใส่ดังนี้ var i int = 20
	- รูปแบบย่อของการประกาศตัวแปร
		- i := 20 //ประกาศตัวแปร i , ชนิดข้อมูลเป็น int และ ค่า deafault เป็น 20
			- ไม่สามารถเปลี่ยนชนิดข้อมูลได้
			- ถ้าต้องการเปลี่ยน type ของข้อมูล ให้สร้างตัวแปรใหม่ขึ้นมาดังนี้
				- j:= string(i)
	- Basic Types
		- bool
		- string
		- int
		- uint
		- byte
		- rune
		- float32 , float64
		- complex64 , complex128
	- const , iota
		- ไม่สามารถเปลี่ยนค่าได้
		- const(
			a = 1
			b
			c
			d
		)
			- output = 1 1 1 1
		- const(
			a = iota
			b
			c
			d
		)
			- output = 0 1 2 3
การใช้งาน if else
	- ไม่ต้องใส่ วงเล็บ
	- ยกตัวอย่างเช่น
	os := runtime.GOOS
	if os == "darwin" {
		fmt.Println("Mac os")
	} else if == "windows" {
		fmt.Println("windows")
	} else {
		fmt.Println("O-o")
	}
การใช้งาน for
	- ไม่ต้องใส่ วงเล็บ
	- ใช้ได้ 3 รูปแบบ
		1. ใช้แบบมีเงื่อนไข เช่น for i := 0; i <= 10; i++
		2. ใช้แบบ while loop เช่น 
			i := 0;
			for i <= 10 {
				fmt.Println(i)
				i++
			}
		3. ใช้แบบไม่มีเงื่อนไข หรือวนลูปไม่รู้จบ เช่น for{fmt.Println("Hello")}
การใช้งาน switch case
	- ยกตัวอย่างเช่น
	switch day{
	case "Sunday","Saturday" :
		fmt.Println("Weekend")
	default:
		fmt.Println("workday")
	}
การใช้งาน functions
	- 2 รูปแบบ
		1. แบบมีการ return กลับ func ชื่อฟังก์ชั้น(ตัวแปร ชนิดตัวแปร) (ชนิดข้อมูลที่จะReturnกลับไป) {}
		2. แบบไม่มีการ return กลับ func ชื่อฟังก์ชั้น(ตัวแปร ชนิดตัวแปร) {}
	- ยกตัวอย่างเช่น
	func add(a int,b int) (int,int){
		return a + 1 , b + 2
	}

	func main() {
		x,y := add(1,2)
		smt.Println(x,y)
	}
การใช้งาน pointer
	- ยกตัวอย่างเช่น
	i := 20
	p := &i // var p *int
	// fmt.Println(p) //output memorie address 0xc00008c0a8
	fmt.Println(*p) //output 20
####################################### End.. Basic Go #################################################

####################################### Start.. Array,Slice และ Map #################################################
การใช้งาน Array
	- [จำนวน arr]ชนิดข้อมูล{ข้อมูลที่ต้องการเก็บ}
	a := [5]string{"A","B","C","D","E"}
	a := [...]string{"A","B","C","D","E"}
	- [จำนวน arr]ชนิดข้อมูล{index: "ข้อมูลที่ต้องการเก็บ"}
	a := [5]string{4:"E"} //กำหนดให้ arr ตัวที่ 4 เก็บ string เป็น E
	a := [...]string{4:"E"}
	- ตัวอย่างการใช้งาน
	const (
		TH int = iota
		EN
		JP
	)
	capitals := [...]string{
		TH: "Bangkok",
		EN: "London",
		JP: "Tokyo",
	}
	fmt.Println
การใช้งาน slice
	- ไม่มีการเก็บจำนวนช่องเหมือน arr หรือมี capacity ได้ไม่จำกัด
	- ถ้าค่าใน arr เปลี่ยน slice ก็จะเปลี่ยนเช่นกัน เนื่องจาก การกำหนดถูกกำหนดแบบ pointer หรือเข้าถึงแบบ memories address
	- Slice := a[0:3] หรือ address poiter,len(s)(ความยาวของ slice),cap(s)(ความจุของ slice)
	- ถ้ามีข้อมูล a := [5]string{"A","B","C","D","E"} แล้วต้องการดึงข้อมูล A ถึง C จะเป็นดังนี้
		ตัวแปร arr ที่ต้องการดึง[index เริ่มต้น : index ที่ต้องการ + 1]
		Slice := a[0:3]
	- ถ้าต้องการดึงข้อมูลทั้งหมดของ array ใช้ดังนี้
		- arr[:]
	- ถ้าต้องการดึงตังแต่เริ่มจนถึง index ตัวที่ 3 ใช้ดังนี้
		- arr[0:3]
		- arr[:3]
	- ถ้าต้องการดึงตังแต่ index ตัวที่ 2 จนถึง index ตัวสุดท้าย ใช้ดังนี้
		- arr[2:5]
		- arr[2:]
	- ยกตัวอย่าง
		a := [5]string{"A", "B", "C", "D", "E"}
		slice := a[0:3]
		slice[2] = "Z"
		fmt.Println(a, slice) //output [A B Z D E][A B Z]
	- การสร้าง slice
		s := []string{"A", "B", "C", "D", "E"} = add ptr(0x1111),len(5),cap(5)
		- การ append หรือ เพิ่มข้อมูลเข้าไปใน arr 
			s2 := append(s,"F") = add ptr(0x3333),len(6),cap(10)
			- go จะเพิ่ม capacity เผื่อไว้ให้
	- append,delete
		s := []string{"A", "B", "C", "D", "E"}
		// s = append(s, "F") // Add array
		// s = append(s[:2], s[3:]...) // 1 delete array
		s = slices.Delete(s, 2, 3) // 2 delete array
		fmt.Println(s)
การใช้งาน make
	- make(สิ่งที่ต้องการจะสร้าง)
	s1 := make([]int, 0, 5) = ระบุชนิดข้อมูล, len, cap
	s1 = append(s1, 2, 2, 2, 2)
	// s1 = append(s1, 2, 2, 2, 2, 2)
	s2 := append(s1, 1)
	s2[0] = 1
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
การใช้งาน Empty slice และ NIL slice
	- Empty slice
		s := []string{}
		s2 := make([]string, 0) //ระบุ len,cap เป็น 0
		output json >> []
	- NIL slice
		var s []string
		output json >> null
การใช้งาน map
	- คล้ายๆกับ arr แตกต่างคือ index ไม่จำเป็นจำต้องเป็น int
	- ถ้าระบุ index ที่ไม่มีอยู่ใน map มันจะกลายเป็นค่า zero value ของชนิดข้อมูลนั้นๆ
	- map[ชนิดข้อมูลของ index]ชนิดของ value{} ยกตัวอย่างเช่น
	- credits := map[string]int{
		"Java": 3,
		"C++": 2,
		"Python": 4,
	}
	result,checkIndex := credits["test"] // output 0,false เนื่อจาก checkIndex เป็นตัวเช็คว่ามี index นั้นอยู๋หรือไม่
การใช้งาน for range
	- for ตัวแปร 1,ตัวแปร 2 := range ตัวแปรที่ต้องการ{}
	- วนลูปได้ 3 รูปแบบ 
		1. Array
			cou := [5]string{"A", "B", "C", "D", "E"}
			for k, v := range cou {
				fmt.Println(k, v)
			}
		2. Slice
			cou := []string{"Java", "C", "Python"}
			for k, v := range cou {
				fmt.Println(k, v)
			}
		3. map
			credits := map[string]int{
				"Java":   3,
				"C++":    2,
				"Python": 4,
			}
			for k, v := range credits {
				fmt.Println(k, v)
			}
####################################### End.. Array,Slice และ Map #################################################

####################################### Start.. string,[]byte, rune #################################################
byte slice
	- string = byte() = uint8 (0-255)
	- string เป็น slice ของ byte
	- string สามารถอ่่านค่าได้เท่านั้น ไม่สามารถเปลี่ยนค่าได้
	str := "Hello word"
	firstLetter := str[0]
	fmt.Println(firstLetter)// output ASCII ของ H
	fmt.Println(string(firstLetter)) // output H
	fmt.Println(len(str)) // output 9
rune
	- rune = int32
	- ถ้าอยากใช้ภาษาไทยต้อง convert จาก byte เป็น rune ก่อน เนื่องจาก byte ปกติไม่สามารถแปลงเลข ASCII ของอักษรภาษาไทยได้เนื่องจากมีตัวเลขแค่ 0-255
	str := "ไงพวก"
	word := []rune(str)
การใช้งาน strings
	result := strings.Contains("BabelCopder", "bel") // output true >> Contains จะเช็คว่า bel มัอยู่ในข้อมูลที่ถูกตรวจสอบหรือไม่
	result := strings.Count("สวัสดีและโชคดี", "ดี") // output 2 >> Count จะนับว่ามีคำว่า ดี กี่คำในข้อมูลที่ถูกตรวจสอบ
	result := strings.HasPrefix("Hello world", "Hell") // output true >> HasPrefix จะดูว่ามีคำที่ขึ้นต้นด้วย Hell ไหม
	result := strings.HasSuffix("Hello world", "ld") // output true >> HasSuffix จะดูว่ามีคำที่ลงท้ายด้วย ld ไหม
	result := strings.Join([]string{"Hello", "world"}, "-") // output Hello-world >> Join จะเป็นการต่อคำ
	result := strings.ToUpper("Hello world") // output HELLO WORLD
	result := strings.ToLower("Hello world") // output hello world
การทำงาน number parsing และ strconv
	// a, err1 := strconv.ParseFloat("2.14", 64)
	// fmt.Println(a, err1) // 2.14

	a, err1 := strconv.ParseFloat("xxxxx", 64)
	if err1 != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(a)
	}

	e, err2 := strconv.ParseInt("0110", 2, 64)
	fmt.Println(e, err2) // 6

	i, err3 := strconv.ParseUint("123", 10, 64)
	fmt.Println(i, err3) // 6

	o, err4 := strconv.Atoi("65")
	fmt.Println(o, err4) // 6

	w := strconv.Itoa(65)
	fmt.Println(w) // 6
####################################### End.. string,[]byte, rune #################################################

####################################### Start.. Structs, Interface, Go Modules #################################################
การสร้งาและใช้งาน struct
	type user struct {
	user string
	age  int64
	}
	type book struct {
		name   string
		title  string
		author user
	}
	func main() {
		// user := user{user: "john"}
		test := book{name: "name", title: "title", author: user{user: "john"}}
		fmt.Println(test.name)
	}
การนิยาม Methods และ Method Receiver
	- ตัวอย่างการใช้งาน 
	type user struct {
		name   string
		course []string
	}

	func (e *user) addCourse(course string) { // เหตุผลที่ต้องใช่ pointer เพราะว่าการที่เรา test.addCourse("C") จะเป็นการสำเนาข้อมูลขึ้นมาใหม่ ซึ่งจะทำให้ค่าไม่ได้เปลี่ยนไป
		e.course = append(e.course, course)
	}

	func main() {
		test := user{name: "John", course: []string{"A", "B"}}
		test.addCourse("C")
		fmt.Println(test)
	}
การใช้งาน Interface
	- สิ่งที่ใช้ระบุว่าสิ่งที่เราสนใจในพฤติกรรมนั้น คืออะไร
	- ยกตัวอย่างเช่น
	type generator interface {
		generate()
	}
	type user struct {
		name   string
		course []string
	}
	func (u user) generate() {
		fmt.Println("Generate...")
	}
	func main() {
		var test generator
		test = user{name: "John", course: []string{"A", "B"}}
		test.generate()
	}
กฏของ Method Sets
	- ถ้า Receiver แบบมี pointer
		- จะรับ Value ที่มี pointer ได้เท่านั้น
	- ถ้า Receiver แบบไม่มี pointer
		- จะรับ Value ที่มี pointer และไม่มี pointer
Polymorphism
	- เกิดจากการใช้งาน interface 
type assertions
	- ยกตัวอย่างถ้าโค้ดเป็นดังนี้
	type generator interface {
		generate()
	}
	type user struct {
		name   string
		course []string
	}
	func (u user) generate() {
		fmt.Println("Generate...")
	}
	func main() {
		var test generator
		test = user{name: "John", course: []string{"A", "B"}}
		test.generate()
	}
	- เราจะไม่สามารถใช้ test.user ได้ เนื่องจากไม่ใช่ type ของ struct แต่เป็น type ของ interface
		จึงต้องใช้ type assertions โดยการ กำหนดตัวแปรใหม่แล้วให้เก็บเป็น type assertions ดังนี้
		test1 := test.(user)
			และสามารถรับแต่แปรอีก 1 ตัวแปรที่ไว้เช็คได้ว่า type ที่รับเข้ามาว่ามี type ตรงกับที่ถูกกำหนดไว้ไหมดังนี้
			test,checkType := test.(user) // output ของ checkType จะออกมาเป็น true หรือ false ก็ได้
		และเรียกใช้แบบปกติได้
		test.name
	- ตัวอย่างการแก้ไขแล้ว
		type generator interface {
			generate()
		}
		type user struct {
			name   string
			course []string
		}

		func (u user) generate() {
			fmt.Println("Generate...")
		}
		func main() {
			var test generator
			test = user{name: "John", course: []string{"A", "B"}}
			test1 := test.(user)
			fmt.Println(test1.name)
			test.generate()
		}

type switch 
	- ยกตัวอย่างการใช้งาน
		type generator interface {
			generate()
		}
		type user struct {
			name   string
			course []string
		}

		type people struct {
			firstname string
			lastname  string
		}

		func (u user) generate() {
			fmt.Println("Generate...")
		}

		func (u people) generate() {
			fmt.Println("Generate...")
		}

		func main() {
			var test generator
			test = user{name: "John", course: []string{"A", "B"}}
			switch test.(type) {
			case user:
				fmt.Println("user")
			case people:
				fmt.Println("people")
			}
		}
Empty interface
	- ตัวอย่างการใช้งาน ไว้สำหรับสร้างข้อมูลที่ไม่ต้องระบุว่าจะเก็บข้อมูลเป็น type อะไร ซึ่งแต่เดิมต้องใช้ []string{"A", "B"} จะเปลี่ยนไปใช้ดังตัวอย่าง
		func main() {
			interfaceTest := []interface{}{"A", false, 2}
			fmt.Println(interfaceTest)
		}
package and Modules
	- go module คือ การรวมหลายๆ package อยู่ใน project เดียว
	- go mod init ชื่อที่ต้องการ
		จะสร้างไฟล์ package ที่ชื่อว่า go.mod ขึ้นมา
	- ถ้าแยก package ออกมาแล้ว ถ้าเราต้องการใช้งานที่เราสร้างไว้ ใน package จะต้องกำหนดอักษรตัวหน้าของ type และ ตัวแปร เป็นตัวพิมพ์ใหญ่เท่านั้น
		ถ้าขึ้นต้นด้วยตัวเล็ก จะเห็นเฉพาะใน package ถ้าเป็นตัวใหญ่ สามารถใช้งานนอก package ได้
	- go mod tidy
		- ลบส่วนที่ไม่ได้ใช้ออกไป
####################################### End.. Structs, Interface, Go Modules #################################################

####################################### Start.. Defer,Error,Panic และ Recover #################################################
Defer 
	- การใส่ defer ตัวที่ถูกใส่ จะถูกทำก็ต่อเมื่อ function ทำงานเสร็จ
	- การทำงานของ defer จะถูกทำจากล่างขึ้นบน
	- ตัวอย่างโค้ด
	func CopyFile(dstName, srcName string)(written int64, err error){
		src,err := os.Open(srcName)
		if err!=nil{
			return 0,err
		}
		defer src.close()

		dst,err:=os.Create(dstName)
		if err!=nil{
			return 0,err
		}
		defer dst.close()

		return io.Copy(dst,src)
	}
กาจัดการ error
	func findIndex(s []int, num int) (int, error) {
		for i, d := range s {
			if d == num {
				return i, nil
			}
		}
		return 0, errors.New("Number not found")
	}

	func main() {
		num := []int{
			1,
			2,
			3,
		}
		index, error := findIndex(num, 2)
		if error != nil {
			fmt.Println(error)
		} else {
			fmt.Println(index)
		}
	}
คุณสมบัติของ Panic และการใช้งาน
	- ถ้าเกิดผิดพลาดขึ้นมาแล้วจะไม่สามารถทำงาน application ต่อไปได้
	- ยกตัวอย่างเช่น เชื่อมต่อฐานข้อมูลไม่ได้ ควรจะหยุดแอพไปเลย
	panic("Fail")
การดักจับ Panic ด้วย Recover
	func e() {
		panic("Panic na ja")
	}
	func d() {
		e()
	}
	func c() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r, "test")
			}
		}()
		d()
	}
	func b() {
		c()
	}
	func a() {
		b()
	}

	func main() {
		a()
	}
####################################### End.. Defer,Error,Panic และ Recover #################################################

####################################### Start.. Goroutines #################################################
Concurrentcy VS Palallelism
	ตัวอย่างถ้า CPU มี 4 Core แล้วโปรแกรมทำงานอยู่แค่ 1 Core จะทำงานได้ไม่มีประสิทธิภาพ จึงต้องมีตัวช่วย
	Concurrentcy
		- CPU มีแค่ 1 Core
		- เป็นการหั่นชิ้นงานแต่ละชิ้น เพื่อให้ทำงานใน core แบบตามลำดับที่เรากำหนด
	Palallelism
		- จะเกิด Palallelism ได้ ต้องการเกิดหั่นงานแบบ Concurrentcy ก่อน
		- การทำงานแบบขนาน
		- CPU มีหลาย Core
		- เป็นการหั่นชิ้นงานแต่ละชิ้น เพื่อให้ทำงานในหลายๆ core ที่กำหนด
การใช้งาน Goroutines
	var wg sync.WaitGroup

	func printAndSleep(num int) {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println(num)
	}

	func main() {
		wg.Add(10)

		for i := 1; i <= 10; i++ {
			go printAndSleep(i)
		}

		wg.Wait()
	}
Channels
	- เปรียบเสมือนท่อที่มีข้อมูลไหลไปไหลมา
	- การสร้าง channels ใช้ make(chan ชนิดข้อมูลที่ต้องการ) เช่น
		ch := make(chan int)
	- ตัวอย่างโค้ด
	import (
		"fmt"
	)

	func sum(ch chan int, num []int) {
		result := 0
		for _, v := range num {
			result += v
		}
		ch <- result
	}

	func main() {
		result := 0
		ch := make(chan int)
		num := []int{
			1, 2, 3, 4, 5,
			1, 2, 3, 4, 5,
			1, 2, 3, 4, 5,
			1, 2, 3, 4, 5,
			1, 2, 3, 4, 5,
			1, 2, 3, 4, 5,
		}
		go sum(ch, num[0:5])
		go sum(ch, num[5:10])
		go sum(ch, num[10:15])
		go sum(ch, num[15:20])
		go sum(ch, num[20:25])
		go sum(ch, num[25:30])

		for i := 0; i < 6; i++ {
			result += <-ch
		}

		fmt.Println(result)
	}
####################################### End.. Goroutines #################################################
    package main

    import "fmt"
    import "math/rand"
    import "time"

    type building struct {
    group1 chan int
    group2 chan int
    group3 chan int
    }

    func (p *building) town() {
    for {
    time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

    p.group1 <- rand.Intn(10)

    }
    }

    func (p *building) town2() {
    for {
    time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

    p.group2 <- rand.Intn(10)

    }
    }

    func (p *building) town3() {
    for {
    time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

    p.group3 <- rand.Intn(10)

    }
    }

    func main() {
    work := time.Duration(rand.Intn(9))

    null := time.Duration(-1)
    second := time.Duration(1)

    read := make(chan int, 10)

    all:=0
    prod := building{group1: read}

    read2 := make(chan int, 10)
    prod2 := building{group2: read2}

    read3 := make(chan int, 10)
    prod3 := building{group3: read3}

    go prod.town()
    go prod2.town2()
    go prod3.town3()

    for null < work{


    null = null + second
    select {
    case i1 := <-read:
    all = all + i1
    fmt.Println("1 группа рабочих произвела количество выполненых работ", i1, "Всего работ группами выполнено ", all)

    case i2 := <-read2:
    all = all + i2
    fmt.Println("2 группа рабочих произвела количество выполненых работ", i2, "Всего работ группами выполнено ", all)

    case i3 := <-read3:
    all = all + i3
    fmt.Println("3 группа рабочих произвела количество выполненых работ", i3, "Всего работ группами выполнено ", all)

    }
    }
    fmt.Println("Рабочий день закончен в", null)

    if all >= 15 {
    fmt.Println("Прораб проверил работы выполненые группой и остался довольным")

    }else{
    fmt.Println("Прораб проверил работы выполненые группой и остался не довольным")
    }
    }


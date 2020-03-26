package main

    import "fmt"
    import "math/rand"
    import "time"

    type building struct {
    group1 chan int
    
    }

    func (p *building) town() {
    for {
    time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

    p.group1 <- rand.Intn(10)

    }
    }

    func main() {

	rand.Seed(time.Now().UnixNano())
	work := time.Duration(rand.Intn(3))
    all:= 0
    null := time.Duration(-1)
    second := time.Duration(1)

    read := make(chan int, 10)

    
    prod := building{group1: read}


    go prod.town()

    for null < work{


    null = null + second
    select {
    case i1 := <-read:
    all = all + i1
    fmt.Println("1 group complete work ", i1, "all works comlete ", all)

    }
    }
    fmt.Println("Work comlete in time ", null)

    if all >= 15 {
    fmt.Println("Big builder check work, he was happy")

    }else{
    fmt.Println("Big builder check work, he was angry")
    }
    }


	
    
    



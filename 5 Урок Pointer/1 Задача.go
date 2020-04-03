package main

import "fmt"
import "regexp"


func check (f string)(w bool){
q := ""
s1 := 0
s2 := 0
s3 := 0

w = false
reg := regexp.MustCompile("[^],/[,(,),{,}]+")

q = reg.ReplaceAllString(f, "")


words := regexp.MustCompile("").Split(q, -1)
fmt.Println(words)
fmt.Println(len(words))
 
for i:= 0; i<len(words); i++{
  if words[i] == "("{
	s1 = i + 1

      if words[s1] == ")"{
	w = true

	}

    }

  if words[i] == "["{
	s2 = i + 1

      if words[s2] == "]"{
	w = true

	}
    }

  if words[i] == "{"{
	s3 = i + 1

      if words[s3] == "}"{
	w = true

	}
    }
}		

return w
}


func main(){

f:= "()[]{}"

fmt.Println(check(f))

}

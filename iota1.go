package main

const(
	a=iota //0
	b=iota //1
	c=iota //2
)

// special const
// can be modified by complier

// iota will be reset to 0 when const appears;
// iota+=1 when add new line in const def;
// iota is the line number of const


const(
	d=iota
	e
	f
)

func main(){
	println(a)
	println(b)
	println(c)
	
	println(d)
	println(e)
	println(f)
}

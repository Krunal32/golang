// +build !solution

// Leave an empty line above this comment.
package lab4
type CspStack interface {
Push(string, interface{})
Pop(string)
Len() int
}
type commandData struct {
action commandAction
value interface{}
result chan<- interface{}
top int
}
type cspStack chan commandData
type commandAction int

const (
pop commandAction = iota
push
length
)

func NewCspStack() cspStack {
cs := make(cspStack) // type safeMap chan commandData
go cs.run()
return cs


}

func (cs cspStack) Len() int {
	reply := make(chan interface{})
cs <- commandData{action: length, result: reply}
return (<-reply).(int)

}

func (cs cspStack) Push(value interface{}) {
cs<- commandData{action: push, value: value}

}

func (cs cspStack) Pop() (value interface{}) {
reply := make(chan interface{})
 cs <- commandData{action: pop, result: reply}
rep:= (<-reply).(interface{})
if rep!=-198{
value=rep
return
}else {
 return nil
}


}

func (cs cspStack) run() {
store:= make([]interface{},100)
 top:=-1
for command := range cs {


switch command.action {
        case push: 
        top++
        if top == len(store) {
	tmp := make([]interface{}, len(store)*2)
	copy(tmp,store)
		store = tmp
	}

	store[top] =command.value
               
        case pop:
        
        if top > -1 {
		value:=store[top]
                top--
                command.result <-value
               
	}else {
               command.result <- -198
        } 
 case length: command.result <- top+1



}
}
}


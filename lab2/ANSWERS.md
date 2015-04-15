##Important Instructions

Please preserve the structure of this file, as it will subjected to *partial*
automatic analysis. **Only insert your answers by replacing the text `YOUR
ANSWER HERE`; do not delete anything else.**

Please use [markdown](https://help.github.com/articles/markdown-basics)
formating to typeset code and Unix commands with the backtick character, for
example, `ls -la`, or if you want to write code blocks, each line should be
indented with four spaces, as done in the code below:

    #include <stdio.h>
    
    int main(void) {
    	printf("Hello, world!\n");
    	return 0;
    }


##Exercises from the Go tour.golang.org

###Exercise 25

**Answer:** see code: gotour/gotour25.go

###Exercise 38

**Answer:** see code: gotour/gotour38.go

###Exercise 43

**Answer:** see code: gotour/gotour43.go

###Exercise 58

**Answer:**  see code: gotour/gotour58.go

###Exercise 60

**Answer:**  see code: gotour/gotour60.go

###Exercise 63

**Answer:** see code: gotour/gotour43.go

##Go Language Questions

1. Write a loop that repeats exactly n times. (Please inline your code below, as long as it is just a few line code snippet).
  - **Answer:** for i:=0;i<n;i++{}
1. What is the value of ok in the following example:
  - **Answer:** 
1. Object-oriented programming: How can you define a “class” Message with two attributes Sender and Content in Go?
  - **Answer:** type  Message struct{ Sender, Content string}
1. How can a method CheckForError be defined on the above Message “class”, that returns an error? (The function body can be empty.)
  - **Answer:** func (msg Message )CheckForError ()err{}


##Go Exercises

Submit completed Go exercise source files as part of a Pull request.

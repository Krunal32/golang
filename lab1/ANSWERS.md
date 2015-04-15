##Important Instructions

Please preserve the structure of this file, as it will subjected to *partial*
automatic analysis. **Only insert your answers by replacing the text `YOUR ANSWER HERE`; do not delete anything else.** 

Please use [markdown](https://help.github.com/articles/markdown-basics) formating to typeset code and Unix commands with the backtick character, for example, `ls -la`, or if you want to write code blocks, each line should be indented with four spaces, as done in the code below:

    #include <stdio.h>
    
    int main(void) {
    	printf("Hello, world!\n");
    	return 0;
    }


##Exercises from the online Unix tutorial

###Exercise 1a

Make another directory inside the `unixstuff` directory called `backups`

**Answer:** mkdir backups

###Exercise 1b

Use the commands `cd`, `ls` and `pwd` to explore the file system.

(Remember, if you get lost, type `cd` by itself to return to your home-directory)

**Answer:** done

###Exercise 2a

Create a backup of your `science.txt` file by copying it to a file called `science.bak`

**Answer:** cp science.txt science.bak

###Exercise 2b

Create a directory called `tempstuff` using `mkdir`, then remove it using the `rmdir` command.

**Answer:** * mkdir tempstuff   .   rmdir tempstuff

###Exercise 3a

Using the above method, create another file called `list2` containing the following fruit: orange, plum, mango, grapefruit. Read the contents of `list2`.

**Answer:**  mkdir list2   .   mkdir  orange, plum, mango, grapefruit ,  cd .. ls -l list2

###Exercise 3b

Using pipes, display all lines of `list1` and `list2` containing the letter 'p', and sort the result.

**Answer:** ls -l list list2 |grep "p" | sort
###Exercise 5a

Try changing access permissions on the file `science.txt` and on the directory `backups`.

Use `ls -l` to check that the permissions have changed.

**Answer:** chmod o=rwx science.txt       chmod 0=rwx backups    ls -l science.txt            ls -l backups

##Shell questions

1. What option with the command `rm` is required to remove a directory?
  - **Answer:**   -r or rf
1. What is the command used to display the manual pages for any command?
  - **Answer:** man
1. What command will show the first 5 lines of an input file?
  - **Answer:**  head -5 <file name> 
1. What command can be used to rename a file?
  - **Answer:** mv [old file name] [new file name]
1. What option can we given to `ls` to show the hidden files?
  - **Answer:**  -all
1. What will the command `cat -n file` do?
  - **Answer:**  cat -n splits file into multiple lines. -n numbers all output lines.
1. What will the command `echo -n hello` do?
  - **Answer:**  write "hello" . Won't append new line.
1. What command will display s list of the users who currently logged in in the system?
  - **Answer:** w
1. How do you change password on your account?
  - **Answer:** passwd
1. How can you list a file in reverse order?
  - **Answer:** *YOUR ANSWER HERE*
1. What does the `less` command do?
  - **Answer:** opens files in simple text editor.
1. With `less` how do you navigate?
  - **Answer:** arrow keys
1. What command will display the running processes of the current user?
  - **Answer:**  ps -u <username>
1. What command can be used to find the process(es) consuming the most CPU?
  - **Answer:**  ps aux --sort -%cpu   (sorts cpu usage from highest to lowest)

##vi questions
1. How do we save a file in `vi` and continue working?
  - **Answer:**  :w
1. What command/key is used to start entering text?
  - **Answer:** i
1. What are the different modes the editor can be in?
  - **Answer:** insert/command
1. What command can be used to place the cursor at the beginning of line 4?
  - **Answer:**  :4
1. What will `dd` command do (in command-mode)?
  - **Answer:** delete current line
1. How do you undo the most recent changes?
  - **Answer:** :u
1. How do you move back one word?
  - **Answer:** dw

##The C Language and Make tool Questions

1. How do you use `gcc` to only produce the `.o` file?  What is the difference between generating only the `.o` file, and building the `hello` executable done in the previous compilation above?
  - **Answer:** gcc <filename.c> -o <filename.o>   .o files are assembled files that haven't been linked together by the linker.
1. Give the command for compiling with `debug` enabled instead of normal compilation for the two examples shown in Listing 2 and Listing 3. Explain how to turn debugging on/off for the two cases.
  - **Answer:** Command: use the flag “-g”. Example 1 use the flag -DDEBUG . Example 2 use the flag -ddebug=1
1. Give a brief pros and cons discussion for the two methods to add debug code shown in Listing 2 and Listing 3.
  - **Answer:** First example uses macros. The pros are that you can apply logic and conditional checks quite
easily using the the keywords below. You can define them to execute code under certain conditions
and use conditional operators. Con is that you have to use the #ifdef #endif for every block of code
that you want to debug. The other method is less sophisticated but easier and more convinient to
implement. I'm guessing it's also better to use with threads.

1. Provide the command for generating the *map* file. Which of the `gcc` tools is responsible for producing a *map* file?
  - **Answer:** gcc -o filename.exe -Wl,-Map,filename.map filename.c. They are produced by linker.
1. What is the content of each of the sections in a *map* file. Explain briefly.
  - **Answer:** rodata: constants (initialized), .data: variables (initialized), .bss: variables
(uninitilized).
1. Rewrite `hello.c` to produce entries in the *map* file for `.data`, `.bss`, and `.rodata`. Hint: This can be done by adding one variable for each type to the file.
  - **Answer:**  see code in helloworld folder
1. Add the following function to `hello.c`: `double multiply(double x1, double x2)`, which returns `x1*x2`. Use `gcc` to generate an assembly code listing for the program, and examine the assembly code. What assembly instructions are used to do this? Repeat this task, but now replace `double` with `float`. Explain!
  - **Answer:** float version uses only l (long) suffix (ends with l, like fstpl). Uses fstpl and fldl to do float operations.
1. How does `make` know if a file must be recompiled?
  - **Answer:**  uses timestamps.
1. Provide a `make` command to use a file named `mymakefile` instead of the default `makefile`.
  - **Answer:** *YOUR ANSWER HERE*
1. How do you implement an *include guard*, and why is it needed?
  - **Answer:** #ifndef <name of header> #define <name of header> used to prevent multiple inclusion

##Library Task
See: folder "lab1/sorting"
Insert your code between the brackets `{}`:

    int main( int argc, char *argv[] )
    {
    if(argc>1)tsize=strtod(argv[1],NULL);
    if (argc==2)min=strtod(argv[2],NULL); 
   
    else if (argc==3)max= strtod(argv[3],NULL); 
    table=malloc(tsize*sizeof(double));  
    int i=0;
    printf("\nNumbers before sorting:  \n\n");
    for(;i<tsize;i++){
    table[i]=randomfunction(min,max); //call random function
    printf("%0.3f %s",table[i],i<tsize-1?",":"\n"); 
    }
    tab_sort(table,tsize);  //sort table 
    printf("Numbers after sorting:  \n\n");
    i=0;
    for(;i<tsize;i++){
    printf("%0.3f %s",table[i],i<tsize-1?",":"\n"); 
    }
    return 0;
    }
    double randomfunction(){
    double diff=max-min;
    double div=(double)RAND_MAX + min;
    return ( (double)rand() * diff) / div;
    }
    }
   void tab_sort( double *tab, int tab_size )
   {
   int i, j;
   double index;
   for (i = 1; i < tab_size; ++i)
   {
   index = tab[i];
   for (j = i; j > 0 && tab[j-1] > index; j--){
   tab[j] = tab[j-1];
   }
   tab[j] = index;
   }
}




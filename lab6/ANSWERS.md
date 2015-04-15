##Instructions

Follow the instructions in the pdf file and populate this file with your answers.

##Questions

---------------------------------------------------------------------------------------------------------
QUESTIONS 3.4. 
1.                                                                                                       
You can find the “kernel.h” and the printk.h files under: /usr/src/linux-headers-3.13.0-30/include/linux/
The make files KDIR := /lib/modules/$(shell uname -r)/build where uname -r gives the linnux header
version. In PWD := $(shell pwd) the shell command pwd gives the current directory. If he kernel tree is
configured properly this should tell the makefile where the include files are located.
2. 
The kernel uses the printk function because functions from the standard library are not availible in
kernel space. Printk also has some additional features like log levels that are useful for kernel
 debugging.

3.  The  makefiles location is at /usr/src/linux-headers-3.13.0-30-generic

4. If I remove the MODULE_LICENCE line the following warning is produced by the compiler:
  "WARNING: modpost: missing MODULE_LICENSE() in /home/andreasj/shared/simp_lko_drive/simp_lko.o"
  It still produces a kernel object file.
----------------------------------------------------------------------------------------------------------
3.7

a) see: lab6/3.7/devread_a.go
b) see: lab6/3.7/simp_read.c 
c) see: lab6/3.7/devread_c.go
----------------------------------------------------------------------------------------------------------
3.9
a) -----------------
b) I have chosen to do b). See: lab6/3.9/simp_rw.c
c) See: lab6/3.9/echo_cat.go
d) By doing bound checking. If the position *ppos or the size of the message is larger than the capasity
   of the "inbox", the program will return. After a successful copy to kernel space the position and the 
   index(variable that keeps track on the last index that has been written to) are increased by the string 
   length of the message.
-----------------------------------------------------------------------------------------------------------
4.3

|Where memory is    |Tried to deref.    |txt pointer|Result                                               |
|-------------------|-------------------|-----------|-----------------------------------------------------|
|Kernel             |Userspace          |0x97c3880  |Segmentation fault(core dumped)                      |
|Userspace          |Kernel  		    |0x8ee4170  |Dereferencing is successful. Receives orginal message|
|Userspace process 1|Userspace process2 |0x894e008  |Segmentation fault(core dumped)                      | 
|Userspace  thread1 |Userspace thread2  |0xb6400468 |Dereferencing is successful. Receives orginal message| 
|Kernel driver 1    |Kernel driver2     |0xe958a680 |Pointer does not contain anything                    |
         

e) Kernel and user space are different, an adress that is valid in user space could point to data in kernel 
space or point nowhere. To check whether we can write/read to memory to/from the kernel we use these 
functions. They check whether the address is within use kernel space and the appropriate action. It's important 
that userspace programs do not have direct access to kernel space which could lead to -severe errors or kernel 
crash. 

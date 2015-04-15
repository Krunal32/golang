#Makefile for simp_read
obj-m := simp_read.o
	KDIR := /lib/modules/$(shell uname -r)/build
	PWD := $(shell pwd)
all:
	$(MAKE) -C $(KDIR) M=$(PWD) modules 
clean:
	$(MAKE) -C $(KDIR) M=$(PWD) clean

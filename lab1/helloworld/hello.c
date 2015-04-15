#include <stdio.h>
float multiply(float x1, float x2);
int main(void){
double x=10.0,y=12.0;
const int rodata_type= 0x6666;
int data_type   = 0x8888;
int bss_type;
printf("Hello World\n");
printf("%f and %f multiplied is:  %f \n",x,y,multiply(x,y));
return 0;
}
float multiply(float x1, float x2){
 return x1*x2;
}

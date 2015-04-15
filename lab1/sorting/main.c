#include <stdio.h>
#include  <stdlib.h>
#include "l1.h"
static double min=0,max=100,tsize=10;
double *table;
double randomfunction();
int main( int argc, char *argv[] )
{
//Make a table of size give by an argument on the command line.
if(argc>1)tsize=strtod(argv[1],NULL); //table size, first argument
if (argc==2)min=strtod(argv[2],NULL);  // min value, second argument
else if (argc==3)max= strtod(argv[3],NULL); // max value, third argument
table=malloc(tsize*sizeof(double));  //allocates memory for array 

int i=0;
printf("\nNumbers before sorting:  \n\n");
for(;i<tsize;i++){
 table[i]=randomfunction(min,max); //call random function
 printf("%0f %s",table[i],i<tsize-1?",":"\n"); 
}
tab_sort(table,tsize);  //sort table 
printf("Numbers after sorting:  \n\n");

i=0;

for(;i<tsize;i++){
 printf("%0f %s",table[i],i<tsize-1?",":"\n");//print 1 decimal precision floats from array
}

return 0;
}
double randomfunction(){

double diff=max-min;
  double div=(double)RAND_MAX + min;
  return ( (double)rand() * diff) / div;

}

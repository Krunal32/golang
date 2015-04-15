#include "l1.h"
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



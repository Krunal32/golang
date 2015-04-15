#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <string.h>
void *run_processes (void *arg);
int num_threads=99,num_accesses=9999,sharedvalue=0, mode=0;
pthread_mutex_t mutexA = PTHREAD_MUTEX_INITIALIZER;
int main(int argc,char* argv[]) {

    if(strcmp("--with-protection",argv[1])==0) {  // runs without-protection mode by default
        mode=1;
        pthread_mutex_init(&mutexA, NULL);
    }
    else if(strcmp("--assembly",argv[1])==0) {
        mode=2;
    }
    num_threads= atoi(strtok(argv[2],"--threads="));
    num_accesses= atoi(strtok(argv[3],"--accesses="));
    int i;
    pthread_t threads[num_threads];
    for (i=0; i<num_threads; i++) {
        pthread_create(&threads[i],NULL,run_processes ,NULL);
        pthread_join(threads[i],NULL);
    }
    printf("\n Expeced value: %d  VS. value:  %d  \n",num_accesses*num_threads,sharedvalue);

    return 0;
}

void *run_processes (void *arg) {
    printf("Thread %ld  \n", pthread_self());
    int i;
    for ( i=0; i<num_accesses; i++) {
        if(mode==0) sharedvalue++;
        else if (mode==1) {
            pthread_mutex_lock(&mutexA);
            sharedvalue++;
            pthread_mutex_unlock(&mutexA);
        }
        else if (mode==2) {
            asm("spin: lock btr $0, sharedvalue ");
            asm(" jnc spin ");
            sharedvalue++;
            asm(" bts $0, sharedvalue ");
        }

    }
    return NULL;
}


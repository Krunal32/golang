##Instructions

Follow the instructions in the pdf file.


##Questions



1. Inline Assembly Pros: can express low level, more control, less memory costs more robust. 
  Cons: harder to code/implement, lower scalability, harder to debug/more error prone . 
 Examples of OS locks are mutexes, semaphores. Their pros are con a pretty much the opposite of inline assembly.

2. Yes. Code example:
	int intLock()
	{
	pthread_mutex_lock(&interrupt_mutex);
	return 0;
	}
	void intUnlock(int level)
	{
	pthread_mutex_unlock(&interrupt_mutex);
	return 0;
	}
   Intlock may however lock threads in o ther applicatons.

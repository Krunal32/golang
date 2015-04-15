d.
A: Some channels have negative viewer numbers. This is because all channels start with 0 viewers when we start logging (because we do not
know the state of all viewers before that). From then channels either lose or gain viewers - the total will always be 0. Some channels
will have negative viewers if there is alot of viewers leaving when we start recording  (f.ex commercial break) but could still be the 
channel with the most viewers.

2 a) 
1. The zapserver receives streams of bytes which it parses into ChanZap structs with respective fields. 
There is 1 write to the map of channel viewersfor every zap.  Since the viewers is updated immediately on every click, a viewer computation is just a lookup from the map. A top10 list is copying that map into a slice that can be sorted. 
Since there can be more than 1 zap written to an ip index in the map the number of reads used to compute top10 is lower than the writes required to keep it updated. The rpc server will do len(ztore) reads from the pointer (zapstore) every n seconds on m clients. With enough clients running the read would be higher than the write. 

2. 
To prevent malformed statistics: check for errors generating STBS, bounds checking, checking lenghts, checking for nil pointers and parsing errors. Basically ensure the integrity of each channelviewer and chzap and ensure that they are properly stored and processed. I'm using channels  to write to go routines for concurrency. My first approach was reading  directly from the pointer. I'm using mutexes as well to protect r/w errors. 

2 c) I ran the profiler and did the following changes: 

-storing SubScriptions in a map with ip address as key instead of slice. Slightly better performance.
- added unsubscribe method to call from the client when it receives control signals
- - cleaned up some uneccessary method calls. Split channel updaters into two routines because they would sometimes block and halt performance.
- 
-Added option to unsubscribe which a client will do when closed with ctrl+c 
- Simplified the subscription call back to the client. 
I tested with 30 clients running. Also the client  would sometimes not receive the statistics or the server would hang. Fixed address bind issues by running a clean up routine fixing error handling. 

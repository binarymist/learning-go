# Examples I used

https://github.com/Becivells/ftpserver/tree/master was forked from https://github.com/fclairamb/ftpserver

You can see down the [bottom](https://github.com/Becivells/ftpserver/blob/master/README.md#history-of-the-project) that:

* [go-ftp](https://github.com/micahhausler/go-ftp) is a minimalistic implementation
* [ftpd.go](https://github.com/shenfeng/ftpd.go) is a very basic and 4 years old

I started using the simplest example (ftpd.go) which had a channel for commands and a channel for data (how FTP is supposed to be done).  
Then looked at go-ftp to apply some more richness.  
Then looked at ftpserver to apply some more richness.  
Then realised that the scope of this exercise was far beyond what was reasonable for an exercise. Learning the FTP sequence of events is non-trivial. My ftpserver project was taken from [ftpd.go](https://github.com/shenfeng/ftpd.go) and extended, I got to the point of `ls` functionality, then gave up, this is far too much work for an exercise that doesn't interest me, and has no relevance to Goroutines and Channels. `cd` and `get` is also not implemented.


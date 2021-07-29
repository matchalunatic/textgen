textgen / generate chunks of dictionary-generated data

I'm running experiments on code and code compression.

It is not easy to generate large volumes of code.

It is easy to generate large series of words.

They share properties when it comes to being easy to compress.

They suck less than, say, /dev/urandom which is not compressible.

Therefore, I need that.

My python implementation sucks and is slow.

This generates 300 MB of data, written to a disk, in 10 seconds.

This is good enough for testing my junk.

Do WTF you want with that. Give it a dictionary and let it work.

# dictionary format
one item per line
lines are \n separated

# precision

chunk size is loose, chunks can be bigger. the size of the majoration is no more than the size of tthe longest element in the list, because I don't check the size of an element before appending it. Because: speed.

# generate dictionary

just run fetch-dictionary.sh. You will need curl and awk and /bin/sh and sort.

# building

I just use

```
go build .
```

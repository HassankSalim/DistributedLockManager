### Distributed Lock Manager

## Implementation Detail [Redis RedLock](https://redis.io/topics/distlock)

This repo is implementation of redis redlock according to my understanding of the algorithm

## TODO 
* Do `DelIfKeyHasVal` in transaction if possible in `Redis`
* Test the algo with more than 2 redis nodes
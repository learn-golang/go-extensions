Go concurrency versus Python asyncIO
====================================


Python asyncIO coroutines
-------------------------

In `test.py` you can find two examples of coroutines utilization:
 * event loop runs two sleepers through tasks API (further reference [1])
 * event loop runs coroutine that runs two sleepers (further reference [2])
 
And here are some results with respect description above:

`time python3 coroutines.py`

    real    0m5.158s
    user    0m0.094s
    sys     0m0.026s
    
`time python3 coroutines.py`

    real    0m7.103s
    user    0m0.083s
    sys     0m0.015s

So, it appears that sequenced coroutines are running faster than a coroutine that awaits other coroutines.


Go
---

Same examples as for Python we have in Go implemented using goroutines concept.

`time go run goroutines.go`

    real    0m5.113s
    user    0m0.073s
    sys     0m0.050s

`time go run goroutines.go`
    
    real    0m5.119s
    user    0m0.075s
    sys     0m0.051s

But in Go code goroutines are not explicitly for developer interaction with scheduler(event loop).
Also, goroutines are being synchronized through WaitGroup and each goroutine reports its execution 
state through WaitGroup API within `defer` section, so it is an additional overhead, but nevertheless,
runs faster than similar code in Python.

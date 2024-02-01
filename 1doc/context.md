# context

### What for context is mostly used?
For cancellation, deadlines in concurrent application

### Explain how context works?
Central object in context is a channel 
```go
Done() <-chan struct{}
```
While this channel is open child function can continue executing context, but when it is closed -> it should 
imideatly return.

### When context is cancelled? 
- when cancel function invoked
- when deadline or timeout exceeds 
- when parent context was cancelled 

### Why it is important to cancel functions in defer?
If execution went through a normal flow we should invoke cancel function to stop child goroutines.
Otherwise, they will continue running until timeout, deadline or cancellation of a parent context.

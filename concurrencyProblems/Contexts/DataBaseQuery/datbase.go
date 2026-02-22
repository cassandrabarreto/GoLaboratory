/* 

Timeout database query
Simulate a slow database query that takes 3 seconds to respond. But your context times out after 1 second, cancelling the query before it finishes.
Constraints:

Use context.WithTimeout this time instead of context.WithCancel
Create a queryDB function that receives a context and simulates a slow query using time.Sleep(3 * time.Second) inside a goroutine
Use select to either return the query result or cancel if ctx.Done() fires
If cancelled, print "query timed out"
If successful, print "query result: 42"
Call defer cancel() in main right after creating the context
*/


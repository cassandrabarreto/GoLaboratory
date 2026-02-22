/* Dynamic multiplier
In main, create a slice of multipliers: []int{2, 3, 4}. For each multiplier, launch an anonymous goroutine that multiplies it by 10 and sends the result into a results channel.
Constraints:

Use go func(m int) { ... }(multiplier) — notice this anonymous function takes an argument to avoid the classic
 loop variable capture bug
Use a results channel to collect the 3 results
Close the results channel after all goroutines finish using the anonymous goroutine pattern you learned: go func() { wg.Wait(); close(resultsChan) }()
Range over results in main and print each one
Use wg.Go() won't work here since the func takes arguments, so use wg.Add(1) and defer wg.Done() inside
*/

package main

func multiplier(){

}

func main(){
	numbers := []int{}
	numbers = append(numbers, 2,3,4)
}
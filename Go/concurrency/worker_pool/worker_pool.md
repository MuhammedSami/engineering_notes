Question

Can you write a simple worker-pool in Go where:

You have N workers,

Workers read integers from a jobs channel,

Each worker multiplies the number by 2,

Sends the result into a results channel,

And the program waits for all workers to finish before closing the results channel?

Please also return all results in the same slice order as the number of inputs.
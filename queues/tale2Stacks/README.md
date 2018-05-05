**Queues: A Tale of Two Stacks**

A queue is an abstract data type that maintains the order in which elements were added to it, allowing the oldest elements to be removed from the front and new elements to be added to the rear. This is called a First-In-First-Out (FIFO) data structure because the first element added to the queue (i.e., the one that has been waiting the longest) is always the first one to be removed.

A basic queue has the following operations:

- Enqueue: add a new element to the end of the queue.
- Dequeue: remove the element from the front of the queue and return it.

In this challenge, you must first implement a queue using two stacks. 
Then process _q_ queries, where each query is one of the following 3 types:
1. Enqueue element _x_ into the end of the queue.
2. Dequeue the element at the front of the queue.
3. Print the element at the front of the queue.

Input Format

The first line contains a single integer, _q_, denoting the number of queries. 
Each line _i_ of the _q_ subsequent lines contains a single query in the form described in the problem statement above.
 All three queries start with an integer denoting the query _type_, but only query 1 is followed by an additional 
 space-separated value, _x_, denoting the value to be enqueued.

Constraints
- 1 <= _q_ <= 10^5
- 1 <= _type_ <= 3
- 1 <= abs(_x_) <= 10^9
- It is guaranteed that a valid answer always exists for each query of type 3.

Output Format

For each query of type 3, print the value of the element at the front of the queue on a new line.

Sample Input
```
10
1 42
2
1 14
3
1 28
3
1 60
1 78
2
2
```

Sample Output
```
14
14
```


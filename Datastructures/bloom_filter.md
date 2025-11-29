Bloom filters are a way to check for existence of an item in a collection.

We can store items within a map with some keywords per item but map is a overhead and not space friendly.

What we usually do we create bloom filter arrays.

We create an array for 20 items and we set all positions(indexes) to 0.
When we receive an item we hash few times, this usually depends on how much we want to prevent false positives.

Lets say we hash it and we get a big number in return like for item called "apple" ex: 1238687643773

So we get modulo of that number to the size of our array

So on each hash function we will get a different number we will get modulo of 20 per each hash result

hash 1 Result = 1238687643773 % 20 = x
hash 2 Result = 9872349872348 % 20 = y
hash 3 Result = 12082392834 % 20 = z

then in our array we will simply say turn those positions OurBloomFilterArray[x],OurBloomFilterArray[y],OurBloomFilterArray[z] to 1

now when the same item came in we just hash it and we check if the positions are 1 and usually bloom filter response is like => the item might be in the list.

but think of it like a book store each time you put a new book on your shelf you will go and say this item is in my library..

and your bloom filter will do the formula of hash and turn the found positions into 1 then again and again.

PROBLEM:

Problem with this approach of using only bits is that we might get same positions turned into 1 for different items and the probability of giving false positives is too high.

So we can use a counter bloom filter each time we got a request to turn into one we can just increment the position OurBloomFilterArray[x]++
and this way we know how many items got turned the position into 1.

ISSUES:

If we keep our array size big, we probably will get less false positives but the more the size the more we get memory usage.




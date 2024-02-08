import heapq

ls = (35,6,6,76,54,33,43,6,7,6,786,54,32)

heapq.heapify(list(ls))

print(heapq.nsmallest(1, ls))
print(heapq.nlargest(1, ls))


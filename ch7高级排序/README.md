##  7.1 归并排序

1. 算法思路：

   > 以整个数据列为对象执行`mergeSort`。
   >
   > `mergeSort`如下所示：
   >
   > 1. 将给定的包含n个元素的局部数组”分割“成两个局部数组， 每个数组各包含n/2个元素。
   > 2.  对局部两个数组执行mergeSort
   > 3. 通过merge合并两个排序完毕的数组。
   
2. 算法描述：

   ```
   mergeSort(A, left, right)
   	if left + 1 < right
   		mid = (left + right) / 2
   		mergeSort(A, left, mid)
   		mergeSort(A, mid, right)
   		merge(A, left, mid, right)
   		
   merge(A, left, mid, right)
   	n1 = mid -left 
   	n2 = right - mid
   	creat L[0:n1], R[0:n2]
   	for i = 0 to n1 - 1
   		L[i] = A[left + i]
   	for i = 0 to n2 - 1
   		R[i] = A[mid + 1]
   	L[n1] = INFTY
   	R[n2] = INFTY
   	i = 0
   	j = 0
   	for k = left to right - 1
   		if L[i] <= R[j]
   			A[k] = L[i]
   			i ++
   		else
   			A[k] = R[j]
   			j ++
               
   ```

   
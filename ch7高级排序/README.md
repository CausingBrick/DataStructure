##  7.1 归并排序

1. 算法思路：

   > 以整个数据列为对象执行`mergeSort`。
   >
   > `mergeSort`如下所示：
   >
   > 1. 将给定的包含n个元素的局部数组”分割“成两个局部数组， 每个数组各包含n/2个元素。
   > 2.  对局部两个数组执行`mergeSort`
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
   	creat L[0:n1], R[0:n2] .
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

   

## 7.2 分割

1. 算法描述：

   ```
   partition(A, p, r)
   	x = A[r]
   	i = p - 1
       for j = p to r-1
       	if A[i] <= x
       	i = i+1
       	exchage A[j], A[i]
       exchage A[i + 1]， A[r]
       return i + 1
   ```



##  7.3 快速排序

1. 算法描述：

   ```
   quickSort（A, p, r)
   	if p < r
   		q = partition(A, p, r)
   		quickSort（A, p, q-1)
   		quickSort（A, q+1, r)
   ```

2. 算法分析：

   基于分治法， 在分割的时候会交换不相邻的元素， 属于不稳定排序是一种快速排序（内部排序）。



##  7.4  计数排序

1. 算法描述:

   ```
   //A为待排序的数组, k为A中的最大数值, B为排序之后的数组.
   CountingSort(A, B, K)
   	for i = 0 to k 
   		c[i] = 0
   	//在C[i]中记录i的出现次数
   	for j = 1 to n
   		c[A[j]]++
   	//排序处理
	n = 0
   	for i = 0 to k
   		for j = 0 to C[i]
   			B[n] = i
   			n++
   ```
   
   



## 7.5 利用标准库排序

### 7.5.1 STL sort排序

​	STL 的sort基于快速排序,是一种复杂度为O(nlogn)的高效排序算法,对最坏的情况添加了应对机制, 克服了最坏的情况下复杂度高达O(n^2),属于不稳定排序.

可以选择稳定的复杂度相同的stable_sort,但是比Sort的内存更多, 速度稍慢.



## 7.6 逆序数

1. 问题描述:

   ![1567412138582](/home/causingbrick/workspace/go/src/DataStructre/ch7高级排序/README.assets/1567412138582.png)

2. 问题分析:

      可以暴力求解但是复杂度为0(n^2), 可以考虑使用归并排序来查找逆序(因为是稳定排序), 即排序中交换元素的次数.

3. 算法描述:

      ```
      mergeSort(A, left, right)
      	if left + 1 < right
      		mid = (left + right)/2
      		v1 = mergeSort(A, left, right)
      		v2 = mergeSort(A, left, right)
      		v3 = merge(A, left, mid, right)
      		return v1 + v2 + v3
      	else
      		return 0
      		
      merge(A, left, mid, right)
      	n1 = mid -left 
      	n2 = right - mid
      	creat L[0:n1], R[0:n2] .
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
      	return i + j
      ```

      

##  7.7   最小成本排序

1. 问题描述:

   ![1567516980646](/home/causingbrick/workspace/go/src/DataStructre/ch7高级排序/README.assets/1567516980646.png)

2. 算法分析:

   ![1567517201239](/home/causingbrick/workspace/go/src/DataStructre/ch7高级排序/README.assets/1567517201239.png)

   ![1567517212803](/home/causingbrick/workspace/go/src/DataStructre/ch7高级排序/README.assets/1567517212803.png)

   ![1567517246573](/home/causingbrick/workspace/go/src/DataStructre/ch7高级排序/README.assets/1567517246573.png)

3. 算法描述:

   * 借整体最小元素的成本算法:

   ![1567517999649](/home/causingbrick/workspace/go/src/DataStructre/ch7高级排序/README.assets/1567517999649.png)

   * 不接元素的最小成本算法:

     ![1567518072526](/home/causingbrick/workspace/go/src/DataStructre/ch7高级排序/README.assets/1567518072526.png)

     对两种结果比较取较小值.
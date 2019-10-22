

##  7.1 归并排序

1. 算法思路：

   > 以整个数据列为对象执行`mergeSort`。
   >
   > `mergeSort`如下所示：
   >
   > 1. 将给定的包含n个元素的局部数组”分割“成两个局部数组， 每个数组各包含n/2个元素。
   > 2.  对局部两个数组执行`mergeSort`
   > 3. 通过merge合并两个排序完毕的数组。
   
   在归并排序里, merge是最重要的基础, 它将已经排好序的片段合并为一个, 并不是简单的合并在一起在用排序算法排序, 而是结合两个数组的排序的特点排序.
   归并操作的工作原理如下：  
	- 第一步：申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列  
	- 第二步：设定两个指针，最初位置分别为两个已经排序序列的起始位置  
	- 第三步：比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置  
	- 重复步骤3直到某一指针超出序列尾  
	- 将另一序列剩下的所有元素直接复制到合并序列尾  


2. 算法描述：

   ```
   mergeSort(A, left, right)
   	if left + 1 < right
   		mid = (left + right) / 2
   		mergeSort(A, left, mid)
   		mergeSort(A, mid, right)
   		merge(A, left, mid, right)
   		
   merge(A, left, mid, right)
		creat L[mid-l+1], R[r-mid]
		copy(al, arr[l:mid+1])
		copy(ar, arr[mid+1:r+1])
		while i < len(al) && j < len(ar) 
			if al[i] <= ar[j] 
				arr[l+j+i] = al[i]
				i++
			 else 
				arr[l+j+i] = ar[j]
				j++
		while i < len(al)
			arr[l+j+i] = al[i]
			i++
		while j < len(ar) 
			arr[l+j+i] = ar[j]
			j++
		
   ```

- [点击查看源码](/ch7高级排序/merge/main.go)

   

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

   ![1567412138582](/ch7高级排序/README.assets/1567412138582.png)

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

   ![1567516980646](/ch7高级排序/README.assets/1567516980646.png)

2. 算法分析:

   ![1567517201239](/ch7高级排序/README.assets/1567517201239.png)

   ![1567517212803](/ch7高级排序/README.assets/1567517212803.png)

   ![1567517246573](/ch7高级排序/README.assets/1567517246573.png)

3. 算法描述:

   * 借整体最小元素的成本算法:

   ![1567517999649](/ch7高级排序/README.assets/1567517999649.png)

   * 不接元素的最小成本算法:

     ![1567518072526](/ch7高级排序/README.assets/1567518072526.png)

     对两种结果比较取较小值.
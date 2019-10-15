##  6.1 递归与分治

递归函数是自己调用自己的函数.

**Example 1**:下面是递归实现$n$的阶层的函数:

```
factorial(n)
	if n == 1
		return 1
	return n * factorial(n -1)
```
- [点击查看源码](/ch6递归与分治法/factorial/main.go)

通过求解局部性的小问题来解决原本的问题,这种技巧称为分治法. 实现步骤如下:

> 分治法
>
> 1. 将问题分割成局部问题(Divide).
> 2. 递归求解局部问题(Solve).
> 3. 将局部问题整合,解决原问题(Conquer)

**Example 2**:分治法查找数组中最大的元素:

```
//数组A的l到r范围内查找最大值
findMaximum(A, l, r)
	m = (1 + r) / 2
	if l == r - 1
		return A[l]
	else
		u = findMaximum(A, l, m)
		v = findMaximum(A, m, r)
		x = max(u, v)
	return x
```

- [点击查看源码](/ch6递归与分治法/findMax/main.go)

##  6.2 穷竭搜索

1. 问题描述: 

   > 对于长度为n的数列A和整数m, 判断A中任意的几个元素相加是否能得到m. A中每个元素只能用一次 .

2. 算法描述：

   ```
   //列举组合的递归函数
   makeCombination()
   	for i from 0 to n-1
   		S[i] = 0
   	rec(0)
   	
   rec(i)
   	if i == n
       	print(S)
       	return
       rec(i+1)
       S[i] = 1
       rec(i+1)
       S[i] = 0
   
   //判断是否得出某个整数的递归函数
   solve（i, m)
   	if m == 0
   		return true
   	if i >= n
   		return false
   	res = solve(i + 1, m) || solve(i+1, m -A[i])
   	return res
   ```

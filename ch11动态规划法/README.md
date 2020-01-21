# 11.动态规划法

将算式的结果存储在内存之中, 在需要时调用结果而无需重复计算,提高处理效率.动态规划就是这类处理手法之一.

## 11.1 斐波那契数列

斐波那契数列：1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...

如果设F(n）为该数列的第n项（n∈N*），那么这句话可以写成如下形式：:F(n)=F(n-1)+F(n-2).

在每次运算之前先判断要运算的值是否已经存储了, 若存储了那就直接调用不然计算之后进行存储, 这样可以大大提高运算效率.

算法描述:

```
fibonacci(n)
    if n == 1 || n == 0
        if data[n] not exist
            data[n] = 1
        return data[n]
  
    if data[n] not exist
         data[n] = fibonacci(n-1) + fibonacci(n-2)

    return data[n]
```

*[源码实现](./fibonacci.go)*

## 11.2 最长公共子序列

1. 相关定义:
    - 公共子序列:若序列z为X和Y的子序列, 那么Z就能称为两数列的公共子序列
    - 最长公共子序列(Longest Common Subesequence): 若z在X和Y中的公共子序列中长度最长, 则z为X和Y的最长公共子序列.
    
2. 思路解析:
    对于两个序列序列$X_i, Y_j$, 求出两个序列的最长公共子序列, 用二维数组$c[m+1][n+1]$代表$X_i, Y_j$的LCS的长度, 则$c[i][j]$的值可从下列推导公式求得.
    $$
    c[i][j]=\left\{
     \begin{array}{rcl}
     0    &   & {if\ i=0 \ or\ j=0}\\
     c[i-1][j-1]+1   &   & {if\ i,j>0\ and\ X_i = Y_j}\\
     max(c[i][j-1], c[i-1][j])   &   & {if\ i,j>0\ and\ X_i \ne Y_j}\\
     \end{array} \right.
    $$
    
    - $当X_m = Y_n时在X_m-1与Y_n-1后面加上1就是LCS$
    - $当X_m = Y_n时在X_m-1与Y_n-1后面加上1就是LCS$

3. 算法描述

   ```
   lcs(X,Y)
   	m = X.length
   	n = Y.length
   	for i = 1 to m
   		c[i][0] = 0
   	for j = 1 to n 
   		c[0][j] = 0
   	for i = 1 to m
   		for j = 1 to n
   			if X[i] == Y[j]
   				c[i][j] = c[i-1][j-1]+1
   			else if c[i-1 ][j] >= c[i][j-1]
   				c[i][j] = c[i-1][j]
   			else
   				c[i][j] = c[i][j-1]
   		
   ```

[源码实现](./lcs/lcs.go)
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
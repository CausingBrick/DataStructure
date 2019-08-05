##  5.1 搜索

搜索指从数据集合里找出目标的处理.基本的搜索算法有如下三种:

1. 线性搜索.

   线性搜索指从数据集合的开头顺次访问各元素, 若找到就返回该元素或位置,并结束搜索,若无则返回一个特殊值表明没有搜索到值.  线性搜索效率低, 但是适用于各种数据形式

2. 二分搜索

   也称为二分查找法, 在基于已经排序好的数据集合基础上进行的搜索,搜索思路如下:

   1. 先把整个集合区域作为搜索范围(数据集合升序排列)
   2. 检查数据集合中央的元素,若与目标关键字一致则结束,若小于关键字则对前半部分做第二部的操作.
   3. 若检查到数据返回数据或者数据位置,否则返回特殊值表明未查询到值.

3. 散列法

   散列是一种数据结构也是一种散列表的算法.此算法只需要将元素的关键字带入函数即可找出结果.



##  5.2 线性搜索

* 问题描述: 

![1563364078115](/home/causingbrick/workspace/go/src/DataStructre/ch5搜索/README.assets/1563364078115.png)

* 算法描述:

  ```
  linearSearcn (key, A)
  	for i from 0 to n-1
  		if A[i] == key
           return i
  	return NOT_FOUND
  ```

  将上面的算法引入标记之后可以将效率提高数倍:

  ```
  linearSearch (key, A)
  	i = 0
  	A[n] = key
  	while A[i] != key 
  		i++
  	if i = n
  		return NOT_FOUND
  	retrun i
  ```

* 算法实现:

  [LinearSearch](linear/main.go)

* 算法分析:

  ![1563412211146](/home/causingbrick/workspace/go/src/DataStructre/ch5搜索/README.assets/1563412211146.png)

##  5.3 二分搜索 

1. 问题描述: 对于两个数列`sequence1`,`sequnece2`,找出他们并集元素的个数.

2. 算法描述:

   ```
   //先决条件: seq为升序
   binarySearch(key, seq) 
   	left = 0
   	right = seq.length
   	while left < right
   		mid = (left + right)/2
   		if key = seq[mid] 
   			return mid
   		else if key < mid 
   			right = mid
   		else
   			left = mid
   	return NOT_FOUND
   ```

3. 算法实现


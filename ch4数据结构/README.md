[TOC]

## 4.1 挑战问题之前--什么是数据结构

1. ###### 数据结构通常有三个概念组合而成:

   1. 数据集合.通过数据的本体来保存数据集合.(比如数组和结构体等基本数据结构)
   2.  规则.保证数据集合按照一定规则进行正确操作, 管理保存的规则.
   3. 操作.插入和取出元素的集合的操作.

2. ###### 栈

   栈`stack`: 后入先出(`Last In First out`)

   * 操作:
     * push(x):栈顶添加元素x
     * pop():从栈顶取出元素
     * isEmpty():检查栈顶是否为空
     * isFull(): 检查栈是否已满
     * 引用栈顶元素
     * 检查栈中是否含有指定数据
   * 规则: 数据中最后加入的元素先被取出.

3. ###### 队列

   队列(`queue`)是一个处理等待的行列,按照先入先出(`First In First out`)的规则管理数据.

   * 操作:
     * enqueue(x):队列末尾添加元素.
     * dequeue():从队列开头取出元素
     * isEmpty():检查队列是否为空
     * isFull(): 检查队是否已满
     * 引用队头元素
     * 检查队列中是否还有指定数据
   * 规则: 最先进入队列的元素先被取出.

4. ###### 表

   能让数据保持一定顺序, 有可以在特定的位置进行插入或删除操作的数据结构.



##  4. 2 栈 Stack

> 栈`stack`: 后入先出(`Last In First out`)
>
> - 操作:
>   - push(x):栈顶添加元素x
>   - pop():从栈顶取出元素
>   - isEmpty():检查栈顶是否为空
>   - isFull(): 检查栈是否已满
>   - 引用栈顶元素
>   - 检查栈中是否含有指定数据
> - 规则: 数据中最后加入的元素先被取出.

* 算法描述

  [go Slice实现栈](/ch4数据结构/stack/stack.go)

* 算法分析:`Pop()`和`Push()`复杂度为$O(1)$



## 4.3 队列

> 队列(`queue`)是一个处理等待的行列,按照先入先出(`First In First out`)的规则管理数据.
>
> - 操作:
>   - enqueue(x):队列末尾添加元素.
>   - dequeue():从队列开头取出元素
>   - isEmpty():检查队列是否为空
>   - isFull(): 检查队是否已满
>   - 引用队头元素
>   - 检查队列中是否还有指定数据
> - 规则: 最先进入队列的元素先被取出.

* 算法描述:
  * [Slice 实现队列](ch4数据结构/queue/queue.go)
* 复杂度分析:
  * 使用环形缓冲区,复杂度可以为O(1)

## 4.4 链表

1. 链表是用一组任意的存储单元存储线性表线性表的数据元素.（这组存储单元可以是连续的，也可以是不连续的）.数据元素称为节点.

2. 链表的构成单位为节点. 每个节点除了有保存其本身的信息之外.还要存储一个指示其直接后继的信息.对于双链表而言, 还需要存储一个指示其直接后继的信息.

   ```go
   //双链表节点
   type node struct {
   	prev, next *node
   	data       interface{}
   }
   ```

3. **[链表的操作代码描述](ch4数据结构/list/doublyLinkedList.go)**

4. 复杂度分析:

   * 往链表中插入元素时, 只需要更爱几个指针的指向,复杂度为O(1)​.
   * 搜索节点需要从节点的开头搜索, 复杂度为$O(N)$.

## 4.5 标准库的数据结构

### 	4.5.1 C++标准库

​	C++库以提供"模板"为主. 所谓模板,是指不必预先制定类型的函数或类. C++的标准库的核心为`STL`, 即`Standard Template Library`.

* [STL数据结构代码示例](ch4数据结构/STL)

### 4.5.2 [stack](STL/stack.cpp)

队列的主要操作包括： 
+ 入队（push） 
+ 出队（pop） 
+ 判断队列是否为空（empty） 
+ 统计队列元素的个数（size） 
+ 访问队头元素（front） 
+ 访问队尾元素（back）

### 4.5.3 [queue](STL/queue.cpp)

队列的主要操作包括： 
+ 入队（push） 
+ 出队（pop） 
+ 判断队列是否为空（empty） 
+ 统计队列元素的个数（size） 
+ 访问队头元素（front） 
+ 访问队尾元素（back）



###  4.5.4 [vector](STL/vector.cpp)

1. 为c++中可变长度的动态数组或可变长的数组

2. C++中直接构造一个vector的语句为:

   ```cpp
   vector<T>vec;
   ```

   这样我们定义了一个名为 vec 的储存 T 类型数据的动态数组。其中 T 是我们要储存的数据类型，可以是 int、float、double 或者其他自定义的数据类型等等。初始的时候 vec 是空的。

3. 插入元素
   C++中通过 push_back ( ) 方法在数组最后面插入一个新的元素。

   ```cpp
   #include <vector>
   using namespace std;
   int main() {
       vector<int> vec;  // []
       vec.push_back(1); // [1]
       vec.push_back(2); // [1, 2]
       vec.push_back(3); // [1, 2, 3]
       return 0;
   }
   ```

4. 获取长度并且访问元素
   C++ 中通过 size ( ) 方法获取 vector 的长度，通过 [ ] 操作直接访问 vector 中的元素，这一点和数组是一样的。

   ```cpp
   #include <vector>
   #include <stdio.h>
   using namespace std;
   int main() {
       vector<int> vec;  // []
       vec.push_back(1); // [1]
       vec.push_back(2); // [1, 2]
       vec.push_back(3); // [1, 2, 3]
       for (int i = 0; i < vec.size(); ++i) {
           printf("%d\n", vec[i]);
       }
       return 0;
   }
   ```

5. 修改元素
C++ 中修改 vector 中某个元素很简单，只需要用 = 给它赋值就好了，比如 vec[1]=3。

    ```cpp
    #include <vector>
    #include <stdio.h>
    using namespace std;
    int main() {
    vector<int> vec;  // []
    vec.push_back(1); // [1]
    vec.push_back(2); // [1, 2]
    vec.push_back(3); // [1, 2, 3]
    vec[1] = 3; // [1, 3, 3]
    vec[2] = 2; // [1, 3, 2]
    for (int i = 0; i < vec.size(); ++i) {
    printf("%d\n", vec[i]);
    }
    return 0;
    }
    ```

6.  清空
C++需调用 clear( ) 方法就可以清空 vector 。 
C++中 vector 的 clear( ) 只是清空 vector ，并不会清空开的内存。用一种方法可以清空 vector 的内存。先定义一个空的 vector x，然后用需要清空的 vector 和 x 交换，因为 x 是局部变量，所以会被系统回收内存（注意：大括号一定不能去掉）。

    ```cpp
    vector<int> v;
    {
    vector<int> x;
    v.swap(x);
    }
    ```



### 4.5.5 [list](STL/list.cpp)

> list是双向链表，与向量相比，它允许快读的插入和删除，但是随机访问比较慢
> Lst1.assign() 给list赋值 
> Lst1.back() 返回最后一个元素 
> Lst1.begin() 返回指向第一个元素的迭代器 
> Lst1.clear() 删除所有元素 
> Lst1.empty() 如果list是空的则返回true 
> Lst1.end() 返回末尾的迭代器 
> Lst1.erase() 删除一个元素 
> Lst1.front() 返回第一个元素 
> Lst1.get_allocator() 返回list的配置器 
> Lst1.insert() 插入一个元素到list中 
> Lst1.max_size() 返回list能容纳的最大元素数量 
> Lst1.merge() 合并两个list 
> Lst1.pop_back() 删除最后一个元素 
> Lst1.pop_front() 删除第一个元素 
> Lst1.push_back() 在list的末尾添加一个元素 
> Lst1.push_front() 在list的头部添加一个元素 
> Lst1.rbegin() 返回指向第一个元素的逆向迭代器 
> Lst1.remove() 从list删除元素 
> Lst1.remove_if() 按指定条件删除元素 
> Lst1.rend() 指向list末尾的逆向迭代器 
> Lst1.resize() 改变list的大小 
> Lst1.reverse() 把list的元素倒转 
> Lst1.size() 返回list中的元素个数 
> Lst1.sort() 给list排序 
> Lst1.splice() 合并两个list 
> Lst1.swap() 交换两个list 
> Lst1.unique() 删除list中重复的元素
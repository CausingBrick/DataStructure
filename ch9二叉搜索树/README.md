# 二叉搜索树

搜索树是一种可以进行插入, 搜索, 删除等操作的数据结构了可以用做字典优先级队列. 

### 1. 树的节点

二叉搜索树是最基本的搜索树, 且各节点均有键值.一个典型的节点结构如下

```
struct Node{
	value
	left, right ,parent *Node
}
```

- `value`: 为每个节点存储的信息
- `left`, `right`: 左右子节点的指针地址
- `parent`:节点的父节点



### 2. 树的遍历

与二叉树相同的思路, 有三种排序方式:

> head 为搜索树的根节点

1. 前序遍历:

   ```
   preorder(head)
   	if head == null
   		return
   	preorder(head.left)
   	print head.value
   	preorder(head.right)
   ```

2. 中序遍历:

   ```
   inorder(head)
   	if head == null
   		return
   	inorder(head.left)
   	print head.value
   	inorder(head.right)
   ```

3. 后序遍历:

   ```
   postorder(head)
   	if head == null
   		return
   	postorder(head.left)
   	postorder(head.right)
   	print head.value
   ```

   
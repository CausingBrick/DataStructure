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

### 3.插入

> head 为树的头结点, node为待插入节点

```
insert(head, node)
	if head == null
		return
	
	if head.value < node.value //将节点插入head右侧
		if head.right == null
			head.right = node
			node.right = head
			return
		insert(head.left, node)
	else //将节点插入head左侧
		if head.right == null
			head.right = node
			node.parent = head
			return
		insert(head.left, node)
```

### 4.查找

查找要在二叉搜索树中找出含有指定键值的节点

> head 为树的头结点, val为指定的键值

```
find(head, val)
    point = head
	while point != null && point.value != val
		if point.value < val
			point = point.right
		else
			point = point.left
	return point
```

### 5.删除

删除包含指定键值的节点

> head 为树的头结点, node为待删除

```
delete(head, val)
	point = head
	if point.left != null && point.right != null
		point.value = point.left.value
		delete(point.left,point.value)
		return
	if point.left == nil && point.right == null 
		if point.parent.left == point
			point.parent.left = null
			point = null
		else
			point.parent.right = null
			point.parent = null 
		return
```

### 6. [代码实现](./binarySearchTree.go)
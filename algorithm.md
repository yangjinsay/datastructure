### hash table & hash function & hash collisions
* hash 表
* hash 函数
* hash 碰撞
### map vs set
### hashamp, hashSet,treemap, treeSet

### PriorityQueue - 优先队列
* 正常入，按照优先级出
* 实现机制
  * Heap(Binary, Binomial, Fibonacci)
    * Min Heap
    * Max Heap
  * Binary Search Tree

### Tree
* root 节点: 最头上节点
* 左子树
* 右子树
* 子节点
* 兄弟节点
* 层
* Binary tree: 完全二叉树
* 二叉搜索树
  * 左子树上所有节点的值均小于它的根结点的值
  * 右子树上所有节点的值均大于它的根结点的值
  * Recurisively 左,右子树也分别为二叉搜索树
* Graph
* 二叉树的遍历
  * Pre-order: root-> left -> right
  * In-older: left -> root -> right
  * Post-order: left -> right -> root

### Recursion 递归 - 循环
* 通过函数体来进行循环

### 贪心算法 Greedy algorithm
> 适用贪心算法的场景  
* 问题能够分解为子问题来解决,子问题的最优解能递推到最终问题的最优解.
这种子问题最优解成为最优子结构。
* 贪心算法与动态规划的不同在于它对每个子问题的解决方案都作出选择，不能回退。动态规划则会保存以前的运算结果，并根觉以前的结果对当前进行选择，有回退功能。


### BFS 广度优先算法
* 如果是树不会出现重复访问
* 如果是图可能会重复访问
```python
def BFS(graph, start, end):
  queue = []
  queue.append([start])
  # visited 几何，用来记录已经被访问的节点
  visited.add(start)

  while queue:
    node = queue.pop()
    # 标记node已经被访问
    visited.add(node)

    process(node)

    # 找node的后继节点并判断是否已经被访问过
    nodes = generate_related_nodes(node)
    queue.push(nodes)

    # others processing work
    ...
```
### DFS 深度优先算法
* 先走一条路，再返回查看是否有遗漏
```python
visited = set()
def dfs(node, visited):
  visited.add(node)
  # process current node here.
  ...
  for next_node in node.children():
    if not next_node in visited:
      dfs(next_node, visited)
```

#### 非递归写法
```python
def DFS(self, tree):
  if tree.root is None:
    return []
  visited, stack = [], [tree.root]

  while stack:
    node = stack.pop()
    visited.add(node)

    process(node)
    nodes = generate_related_nodes(node)
    stack.push(nodes)

    # other processing work
    ...
```

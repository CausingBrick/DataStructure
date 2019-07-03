#include <queue>
#include <iostream>
using namespace std;
int main()
{
    queue<int> q; // 声明一个装 int 类型数据的队列
    q.push(1);    // 入队
    q.push(2);
    q.push(3);
    cout << q.size() << endl; // 输出队列元素个数
    while (!q.empty())
    {                              // 判断队列是否为空
        cout << q.front() << endl; // 访问队首元素
        q.pop();                   // 出队
    }
    return 0;
}
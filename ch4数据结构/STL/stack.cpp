#include <iostream>
#include <stack>

using namespace std;

int main()
{
    stack<string> S;

    S.push("monkey");
    cout << S.top() << endl;
    cout << S.empty() << endl;
    cout <<  S.size() << endl;

    return 0;
}
#include <iostream>
#include <list>
#include <algorithm>
#include <iterator>
using namespace std;
int main()
{
    list<int> mylist;
    cout << "size = " << mylist.size() << endl; //0

    mylist.push_back(0);
    mylist.push_back(1);
    mylist.push_back(2);
    mylist.push_back(3);
    mylist.push_back(4);
    cout << "size = " << mylist.size() << endl; //5
    for (list<int>::iterator ite = mylist.begin(); ite != mylist.end(); ite++)
        cout << *ite << "  ";
    cout << endl; //0 1 2 3 4

    list<int>::iterator ite = find(mylist.begin(), mylist.end(), 3);
    if (*ite != 0)
        mylist.insert(ite, 99);

    cout << "size = " << mylist.size() << endl; //6
    cout << *ite << endl;                       //3

    for (list<int>::iterator ite = mylist.begin(); ite != mylist.end(); ite++)
        cout << *ite << "  "; //0 1 2 99 3 4

    ite = find(mylist.begin(), mylist.end(), 1);
    if (*ite != 0)
        cout << *(mylist.erase(ite)) << endl;

    for (list<int>::iterator ite = mylist.begin(); ite != mylist.end(); ite++)
        cout << *ite << "  ";
    cout << endl; //0  2 99 3 4

    int iv[5] = {5, 6, 7, 8, 9};
    list<int> mylist1(iv, iv + 5);

    ite = find(mylist.begin(), mylist.end(), 99);
    mylist.splice(ite, mylist1);
    for (list<int>::iterator ite = mylist.begin(); ite != mylist.end(); ite++)
        cout << *ite << "  ";
    cout << endl; //0 2 5 6 7 8 9 99 3 4

    mylist.reverse();
    for (list<int>::iterator ite = mylist.begin(); ite != mylist.end(); ite++)
        cout << *ite << "  ";
    cout << endl; //4 3 99 9 8 7 6 5 2 0

    mylist.sort();
    for (list<int>::iterator ite = mylist.begin(); ite != mylist.end(); ite++)
        cout << *ite << "  ";
    cout << endl; //0 2 3 4 5 6 7 8 9 99
    return 0;
}
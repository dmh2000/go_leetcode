#include <iostream>
#include <vector>
#include <deque>
#include <cstdlib>
#include <cassert>

using namespace std;

void rotateRight(deque<int> &a) 
{
    int v;
    v = a[a.size()-1];
    a.pop_back();
    a.push_front(v);
}

int main()
{
    int n;
    int k;
    int q;
    deque<int> na;
    vector<int> qa;
    int v;
    cin >> n;
    cin >> k;
    cin >> q;
    for(int i=0;i<n;++i) {
        cin >> v;
        na.push_back(v);
    }
    for(int i=0;i<q;++i) {
        cin >> v;
        qa.push_back(v);
    }
    for(int i=0;i<k;++i) {
        rotateRight(na);
    }
    for(int i=0;i<(int)qa.size();i++) {
        cout << na[qa[i]] << endl;
    }
    return 0;
}
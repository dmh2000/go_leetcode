#include <iostream>
#include <vector>
#include <vector>
#include <cstdlib>
#include <cassert>

using namespace std;

int partition(vector <int> &ar,int lo,int hi)
{
    int p;
    int i;
    int j;
    int e;

    p = ar[lo];
    i = lo;
    j = hi+1;
    for(;;) {
        i++;
        while(ar[i] < p) {
            if(i == hi) {
                break;
            }
            i++;
        }
        j--;
        while(p < ar[j]) {
            if(j == lo) {
                break;
            }
            j--;
        }
        if(i >= j) {
            break;
        }
        e = ar[i];
        ar[i] = ar[j];
        ar[j] = e;
    }
    e = ar[j];
    ar[j] = ar[lo];
    ar[lo] = e;

    return j;
}

int main()
{
    int n;
    vector<int> a;
    int v;
    cin >> n;
    for(int i=0;i<n;++i) {
        cin >> v;
        a.push_back(v);
    }

    int lo = 1;
    int hi = (int)a.size()-1;
    int m  = (int)a.size()/2;
    for(;;) {
        int p = partition(a,lo,hi);
        int d = p-lo+1;
        if (d == m) {
            cout << a[p] << endl;
            break;
        }
        if (m < d) {
            hi = p-1;
        }
        else {
            m  = m-d;
            lo = p+1;
        }
    }

    return 0;
}
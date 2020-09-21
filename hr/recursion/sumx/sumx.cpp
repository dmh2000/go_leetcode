#include <iostream>
#include <vector>
#include <cmath>
#include <climits>

using namespace std;

int countx = 0;
vector<int> b;
void compute(int v, vector<int> &a,int lo,int hi,int sum)
{
    sum += a[lo];
    b.push_back(a[lo]);
    if (sum == v) {
        countx++;
        // cout << sum << ',' << countx << endl;
        return;
    }
    else if (sum > v) {
        return;
    }
    else if (lo >= hi) {
         return;
    }
    else {
    }

    for (int i=lo+1;i<=hi;i++) {
        compute(v, a, i, hi,sum);
        b.pop_back();
    }
}
int main(int argc,char *argv[])
{
    int v = 10;
    int n = 2;
    vector<int> a;

    
    int i = 1;
    int x = static_cast<int>(pow(i,n));
    while(x <= v) {
       a.push_back(x);
       i++;
       x = static_cast<int>(pow(i,n));
    }

    int lo = 0;
    int hi = a.size()-1;
    for (int i=lo;i<=hi;i++) {
        b.clear();
        compute(v,a,i,hi,0);
    }
    cout << countx << endl;
    return 0;
}
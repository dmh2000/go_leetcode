#include <vector>
#include <cstdlib>
#include <cassert>

using namespace std;

bool isSorted(vector<int> &a,int lo,int hi)
{
    for(lo;lo<=hi;lo++) {
        if(a[lo] > a[hi]) {
            return false;
        }
    }
    return true;
}

vector<int> mergeSort(vector<int> a);
vector<int> quickSort(vector<int> a);
vector<int> insertionSort(vector <int>  a);
int main(void)
{
    vector<int> a;
    vector<int> b;
    for(int i = 0;i<1000;++i) {
        a.clear();
        for(int j=0;j<i;j++) {
            a.push_back(rand());
        }
        b = insertionSort(a);
        assert(isSorted(b,0,(int)b.size()-1));
        b = quickSort(a);
        assert(isSorted(b,0,(int)b.size()-1));
        b = mergeSort(a);
        assert(isSorted(b,0,(int)b.size()-1));
    }

    return 0;
}


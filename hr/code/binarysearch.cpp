#include <vector>
using namespace std;

// recursive
int binarySearchR(vector<int> a,int key,int lo,int hi)
{
    if(lo == hi) {
        // terminate, check if value
        if(a[lo] == key) {
            return lo;
        }
        else {
            return -1;
        }
    }

    int m = ((hi-lo) / 2) + lo;
    if(a[m] == key) {
        // midpoint
        return m;
    }
    else if(key < a[m]) {
        // search lower half
        return binarySearchR(a,key,lo,m-1);
    }
    else {
        // search upper half
        return binarySearchR(a,key,m+1,hi);
    }
}

// iterative
int binarySearchI(vector<int> a,int key,int lo,int hi)
{
    for(;;) {
        if(lo == hi) {
            // terminate, check if value
            if(a[lo] == key) {
                return lo;
            }
            else {
                return -1;
            }
        }

        int m = ((hi-lo) / 2) + lo;
        if(a[m] == key) {
            // midpoint
            return m;
        }
        else if(key < a[m]) {
            lo = lo;
            hi = m-1;
        }
        else {
            lo = m+1;
            hi = hi;
        }
    }
}

int binarySearch(vector<int> a,int key)
{
    return binarySearchI(a,key,0,(int)a.size()-1);
}

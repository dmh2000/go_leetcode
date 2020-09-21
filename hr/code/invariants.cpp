#include "../hr.h"

bool isSorted(const vector<int> &a,int lo,int hi)
{
    for(lo;lo<=hi;lo++) {
        if(a[lo] > a[hi]) {
            return false;
        }
    }
    return true;
}

bool isPartitioned(const vector<int> &a,int lo,int hi,int m)
{
    for(int i=lo;i<m;i++) {
        if(a[i] > a[m]) {
            return false;
        }
    }
	for (int i = m+1; i <= hi; i++) {
		if (a[i] < a[m]) {
			return false;
		}
	}
	return true;
}
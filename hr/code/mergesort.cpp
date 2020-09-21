#include <cstdio>
#include <cstdlib>
#include <cstdint>
#include <cassert>
#include <iostream>
#include <vector>
#include "../hr.h"

using namespace std;

void merge(vector<int> &a,vector<int> &b,int lo,int m,int hi)
{
    int i = lo;
    int j = m;
    int k;

	// high,low,middle are correct
	assert((m - lo) > (hi - m));

	// if lo and high equal, done
    if(i == j) {
        // nothing to swap
        return;
    }
	// if one apart, and lo > hi, swap them
    else if((hi-lo) == 1) {
        // only two to swap
        if (a[lo] > a[hi]) {
            int t = a[hi];
            a[hi] = a[lo];
            a[lo] = t;
        }
        return;
    }

	// low to middle-1 is sorted
    assert(isSorted(a,lo,m-1));
	// middle to high is sorted
    assert(isSorted(a,m,hi));

    // copy a to b
    for(i=lo;i<=hi;i++) {
        b[i] = a[i];
    }
    // merge
    i = lo;
    j = m;
    k = lo;
    for(;;) {
        // move smallest of b[i],b[j] to a[k]
        if(b[i] > b[j]) {
            a[k] = b[j];
            j++;
            k++;
        }
        else {
            a[k] = b[i];
            i++;
            k++;
        }

        // check for done
        // i goes from lo to m-1
        if(i == m) {
            while(j <= hi) {
                a[k] = b[j];
                j++;
                k++;
            }
            break;
        }

        // j goes from m to hi inclusive
        if(j > hi) {
            while(i < m) {
                a[k] = b[i];
                i++;
                k++;
            }
            break;
        }
    }
}
void mergeSortR(vector<int> &a,vector<int> &b,int lo,int hi)
{
    if ((hi - lo) <= 0) {
        return;
    }

    int m = ((hi - lo) / 2) + lo;

    // lower half
    mergeSortR(a,b,lo,m);

	// upper half
    mergeSortR(a,b,m+1,hi);

    // lo,m+1,hi as limits are to make sure lo to m is equal to or greater than m+1 to hi
    merge(a,b,lo,m+1,hi);
}

vector<int> mergeSort(vector<int> a)
{
    vector<int> b(a.size());

    mergeSortR(a,b,0,(int)a.size()-1);

    assert(isSorted(a,0,(int)a.size()-1));

    return a;
}


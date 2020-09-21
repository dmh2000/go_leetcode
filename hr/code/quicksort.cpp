#include <vector>
#include <iostream>
#include <cstdlib>
#include <cassert>
#include "../hr.h"

using namespace std;

void print(vector<int> &ar, int lo, int hi,int p)
{
	for (auto v : ar) {
		cout << v << ' ';
	}
	cout << ':' << p << endl;
}

void print(size_t len,int i, int j)
{
	for (size_t k = 0; k < len; k++) {
		if (k == i) {
			cout << "l ";
		}
		else if (k == j) {
			cout << "h ";
		}
		else {
			cout << "  ";
		}
	}
	cout << endl;
}

int partition(vector <int> &ar,int lo,int hi)
{
    int pivot;
    int i;
    int j;
    int e;

	// pick the partition  pivot
	// use random, median
	pivot = ar[lo];
    i = lo;
    j = hi+1;
	// continue until low index 'i' crosses over 
	// high index 'j'
	for(;;) {
		// while value less than pivot, skip over it
		i++;
        while(ar[i] < pivot) {
            if(i == hi) {
                break;
            }
            i++;
        }
		// while pivot less than value, skip over it
        j--;
        while(pivot < ar[j]) {
            if(j == lo) {
                break;
            }
            j--;
        }
		// if crossed over, it is done
        if(i >= j) {
            break;
        }
		// low side 'i' is >= pivot
		// high side 'j' is <= pivot
		// swap the two
		//print(ar.size(),i, j);
        e = ar[i];
        ar[i] = ar[j];
        ar[j] = e;
		//print(ar, lo, hi, pivot);
	}
	// swap the pivot with the 'j'th index
	//print(ar.size(),lo, j);
    e = ar[j];
    ar[j] = ar[lo];
    ar[lo] = e;
	//print(ar, lo, hi, pivot);

    return j;
}

void xsort(vector<int> &a,int lo,int hi)
{

    if(hi <= lo) {
        return;
    }

    int p = partition(a,lo,hi);
    assert(isPartitioned(a,lo,hi,p));
    xsort(a,lo,p-1);
    xsort(a,p+1,hi);
}

vector<int> quickSort(vector<int> a)
{
    xsort(a,0,(int)a.size()-1);

    return a;
}

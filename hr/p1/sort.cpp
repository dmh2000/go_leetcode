#include <cstdlib>
#include <cassert>
#include <algorithm>
#include "../hr.h"

using namespace std;


int main(void)
{
    vector<int> a;
    vector<int> b;
    for(int i = 10;i<1000;++i) {
        a.clear();
        for(int j=0;j<i;j++) {
            a.push_back(j);
        }
		std::random_shuffle(a.begin(), a.end());

        //b = insertionSort(a);
        //assert(isSorted(b,0,(int)b.size()-1));
        //b = quickSort(a);
        //assert(isSorted(b,0,(int)b.size()-1));
        b = mergeSort(a);
        assert(isSorted(b,0,(int)b.size()-1));
    }

    return 0;
}


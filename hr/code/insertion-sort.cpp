#include <cstdio>
#include <vector>

using namespace std;

vector<int> insertionSort(vector <int>  ar)
{
    int e;

    for(int j=1;j<ar.size();j++) {
        int i = j;
        while((i > 0)&&(ar[i] < ar[i-1])) {
            e = ar[i];
            ar[i] = ar[i-1];
            ar[i-1] = e;
            i--;
        }
    }
    return ar;
}

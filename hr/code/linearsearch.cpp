#include <vector>
using namespace std;

int linearSearch(vector<int> a,int key)
{
    for(int i=0;i<a.size();i++) {
        if(a[i] == key) {
            return i;
        }
    }

    return -1;
}
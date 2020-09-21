#include <cstdio>
#include <cstdint>
#include <cassert>
#include "../hr.h"

using namespace std;

int linearSearch(vector<int> a,int key);
int binarySearch(vector<int> a,int key);

int main(int argc,char *argv[])
{

    vector<int> a;
    int key;
    int index;

    // create a sorted array (linear search doesn't care, binary search requires sorted)
    for(int i=0;i<101;++i) {
        a.push_back(i);
    }

    key = 1;
    index = linearSearch(a,key);
    assert(a[index] == key);
    index = binarySearch(a,key);
    assert(a[index] == key);


    key = 25;
    index = linearSearch(a,key);
    assert(a[index] == key);
    index = binarySearch(a,key);
    assert(a[index] == key);

    key = 50;
    index = linearSearch(a,key);
    assert(a[index] == key);
    index = binarySearch(a,key);
    assert(a[index] == key);

    key = 99;
    index = linearSearch(a,key);
    assert(a[index] == key);
    index = binarySearch(a,key);
    assert(a[index] == key);

    key = 1000;
    index = linearSearch(a,key);
    assert(index == -1);
    index = binarySearch(a,key);
    assert(index == -1);

    return 0;
}

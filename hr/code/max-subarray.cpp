#include <iostream>
#include <algorithm>
#include <vector>
#include <deque>
#include <cstdlib>
#include <cassert>
#include <cstdint>

using namespace std;

int64_t max_contiguous(const vector<int64_t> &a)
{
    int64_t m;
    int64_t m2;
    m  = a[0];
    m2 = a[0];
    for (size_t i=1;i<a.size();i++) {
        int64_t x = m + a[i];
        int64_t y = a[i];
        if (x > y) {
            m = x;
        }
        else {
            m = y;
        }
        if (m > m2) {
            m2 = m;
        }
    }
    return m2;
}
int64_t max_noncontiguous(vector<int64_t> &a)
{
    vector<int64_t> b = a;
    int64_t sum;
    // sort the array
    sort(b.begin(),b.end());

    // find first positive
    int32_t p = -1;
    for (size_t i=0;i<b.size();i++) {
        if (b[i] > 0) {
            p = (int32_t)i;
            break;
        }
    }
    if (p < 0) {
        // no positive values, return largest negative
        sum = b[b.size()-1];
    }
    else {
        sum = 0;
        for (size_t i=p;i<(int32_t)b.size();i++) {
            sum += b[i];
        }
    }

    return sum;
}

void max_subarray(vector<int64_t> &a)
{
    int64_t mc;
    int64_t mu;

    // max contiguous
    mc = max_contiguous(a);
    mu = max_noncontiguous(a);
    cout << mc << ' ' << mu << endl;
}
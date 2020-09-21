#include <iostream>
#include <algorithm>
#include <vector>
#include <deque>
#include <cstdlib>
#include <cassert>
#include <cstdint>

using namespace std;

int64_t calc(int64_t count)
{
    int64_t sum;
    if(count == 0) return 0;
    if(count == 1) return 0;
    if(count == 2) return 2;
    sum = 0;
    for(int64_t i=1;i<count;++i) {
        sum += i;
    }
    return sum*2;
}

void sherlock_pairs(vector<int64_t> &a)
{
    int64_t count;
    int64_t sum;
    int state;

    // sort a
    sort(a.begin(),a.end());
    // count where a[i] == a[i+1]
    count = 0;
    state = 0;
    sum   = 0;
    for(size_t k=0;k<a.size()-1;k++) {
        switch(state) {
        case 0:
            if(a[k] == a[k+1]) {
                count += 2;
                state = 1;
            }
            break;
        case 1:
            if(a[k] == a[k+1]) {
                count += 1;
            }
            else {
                sum += calc(count);
                count = 0;
                state = 0;
            }
            break;
        default:
            break;
        }
    }
    if(count > 0) {
        sum += calc(count);
    }
    cout << sum << endl;

}

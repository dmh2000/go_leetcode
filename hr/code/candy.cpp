#include <iostream>
#include <algorithm>
#include <unordered_map>
#include <vector>
#include <deque>
#include <cstdlib>
#include <cassert>
#include <cstdint>
#include <string>

using namespace std;

int64_t solvex(vector<int64_t> rating)
{
    vector<int64_t> candy;


    // everyone gets 1
    for (int i=0;i<(int32_t)rating.size();i++) {
        candy.push_back(1);
    }

    // fixup forward order
    for (int i=1;i<(int32_t)rating.size();i++) {
        if ((rating[i] > rating[i-1])&&(candy[i] <= candy[i-1])) {
            candy[i] += abs(candy[i-1] - candy[i]) + 1;
        }
    }

    // fixup reverse ordering
    for (int i=(int32_t)rating.size()-1;i>0;i--) {
        if ((rating[i-1] > rating[i])&&(candy[i-1] <= candy[i])) {
            candy[i-1] += abs(candy[i-1] - candy[i]) + 1;
        }
    }

    int64_t count = 0;
    for (int i=0;i<(int32_t)candy.size();i++) {
        count += candy[i];
    }

    for (int i=0;i<(int32_t)candy.size();i++) {
        // cout << rating[i] << ' ' << candy[i] << endl;
    }

    return count;
}


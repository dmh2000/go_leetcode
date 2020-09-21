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

struct store_list_t
{
    int64_t n;
    int64_t lo;
    int64_t hi;
    int64_t count;
};

std::string store_key(store_list_t &f)
{
    char b[256];
    sprintf_s(b,"%08llx%08llx%08llx",f.n,f.lo,f.hi);
    return b;
}

typedef unordered_map<string,store_list_t> store_map_t;
store_map_t store_map;

int64_t solveR(vector<int64_t> &a,int64_t n,int64_t lo,int64_t hi)
{
    int64_t count = 0;
    int64_t k;
    int64_t d;
    store_list_t f;

    // lo and hi converged?
    if (lo == hi) {
        // is it an even divisor?
        if ((n % a[hi]) == 0) {
            // yes
            f ={n,lo,hi,1};
            store_map.insert({store_key(f),f});
            return 1;
        }
        else {
            // no
            return 0;
        }
    }
    else if (lo > hi) {
        // nothing left to check
        return 0;
    }

    d = n / a[lo];
    k = a[lo];
    for (int64_t i=0;i<d;++i) {
        n -= k;
        if (n == 0) {
            f ={n,lo,hi,1};
            store_map.insert({store_key(f),f});
            count++;
        }
        else {
            for (int64_t m=1;m<=hi;m++) {
                int64_t r;
                store_map_t::iterator g;
                // check if already been serviced
                f ={n,lo+m,hi,0};
                g = store_map.find(store_key(f));

                // not already solved
                if (g == store_map.end()) {
                    // count it
                    r = solveR(a,n,lo+m,hi);
                    count  += r;
                    // insert it
                    f.count = r;
                    store_map.insert({store_key(f),f});
                }
                else {
                    count += g->second.count;
                }
            }
        }
    }
    return count;
}

int64_t solve(vector<int64_t> &a,int64_t n)
{
    int64_t count;
    sort(a.begin(),a.end());

    count = 0;
    for (int i=0;i<(int)a.size();i++) {
        count += solveR(a,n,i,(int64_t)a.size()-1);
    }
    return count;
}

#include <map>
#include <vector>
#include <iostream>

struct kval {
    int d;
    std::vector<int> p;

    kval() : d(0) {

    }

    bool operator() (const int& lhs, const int& rhs) {
        return lhs < rhs;
    }
};

int distance(std::vector<int> p) {
    return p[0] * p[0] + p[1] * p[1];
}

class Solution {
public:
    std::vector<std::vector<int>> kClosest(std::vector<std::vector<int>>& points, int K) {
        std::vector<std::vector<int>> a;
        std::multimap<int, kval> h;

        // map the points
        for (size_t i = 0; i < points.size(); i++ ) {
            kval kv;
            kv.d = distance(points[i]);
            kv.p = points[i];
            h.insert(std::pair<int, kval>{kv.d, kv});
        }

        // get K values in ascending order
        std::multimap<int,kval>::const_iterator v = h.cbegin();
        for (int i = 0; i < K; i++) {
            kval kv = v->second;
            a.push_back(kv.p);
            v++;
        }

        return a;
    }
};

void run_kclose() {
    {
        std::vector < std::vector<int>> points = { {3,3},{5,-1},{-2,4} };
        std::vector < std::vector<int>> r;
        Solution s;


        r = s.kClosest(points, 2);

        for (size_t i = 0; i < r.size(); i++) {
            std::cout << r[i][0] << "," << r[i][1] << "\n";
        }
    };

    {
        std::vector < std::vector<int>> points = { {1,0},{0,1} };
        std::vector < std::vector<int>> r;
        Solution s;

        r = s.kClosest(points, 2);

    for (size_t i = 0; i < r.size(); i++) {
        std::cout << r[i][0] << "," << r[i][1] << "\n";
    };
}

}
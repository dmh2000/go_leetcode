#include <iostream>
#include <algorithm>
#include <unordered_map>
#include <vector>
#include <list>
#include <queue>
#include <stack>
#include <cstdlib>
#include <cassert>
#include <cstdint>
#include <string>

using namespace std;

typedef std::pair<int64_t,int64_t> node_pair_t;
typedef std::vector<node_pair_t>   node_array_t;
typedef std::list<int64_t>         node_list_t;
typedef node_list_t::iterator      node_iter_t;
typedef std::vector<node_list_t>   graph_t;

void bfs(int64_t start,int64_t nodes,node_array_t a)
{
    // find the largest value in the vector
    graph_t         graph(nodes+1);
    vector<bool>    visited(nodes+1);
    vector<int64_t> edge(nodes+1);
    queue<int64_t>  bq;
    int64_t         n;
    int64_t         m;
    stack<int64_t>  st;

    for (auto i=a.begin();i!=a.end();++i) {
        graph[i->first].push_back(i->second);
        graph[i->second].push_back(i->first);
    }

    for (int64_t i=1;i<nodes+1;++i) {
        visited[i] = false;
        edge[i]    = -1;
    }

    visited[start] = true;
    bq.push(start);
    while (!bq.empty()) {
        n = bq.front();
        bq.pop();

        for (node_iter_t i=graph[n].begin();i!= graph[n].end();i++) {
            m = *i;
            if (!visited[m]) {
                edge[m] = n;
                visited[m] = true;
                bq.push(m);
            }
        }
    }

    // for all nodes except 'start'
    int64_t len;
    for (int i=1;i<nodes+1;i++) {
        // ski pstart
        if (i == start) continue;
        // compute the path length to start
        len = 6;
        m   = edge[i];
        for (;;) {
            if (m == -1) {
                cout << -1 << ' ';
                break;
            }
            else if (m == start) {
                cout << len << ' ';
                break;
            }
            else {
                len += 6;
            }
            m = edge[m];
        }
    }
    cout << endl;
}

int bfs_main()
{
    int64_t t;
    int64_t nodes;
    int64_t edges;
    int64_t start;
    vector<node_pair_t> a;

    cin >> t;

    for (int i=0;i<t;++i) {
        cin >> nodes;
        cin >> edges;
        a.clear();
        for (int j=0;j<edges;j++) {
            int64_t p;
            int64_t q;
            cin >> p;
            cin >> q;
            a.push_back({p,q});
        }
        cin >> start;
        bfs(start,nodes,a);

    }

    return 0;
}
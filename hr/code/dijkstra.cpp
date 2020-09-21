#include <iostream>
#include <algorithm>
#include <vector>
#include <list>
#include <queue>
#include <deque>
#include <cstdlib>
#include <cstdint>
#include <climits>
#include <string>

using namespace std;

// for input
struct node_t {
    int64_t p;
    int64_t q;
    int64_t len;
    node_t(int64_t xp,int64_t xq,int64_t xlen): p(xp),q(xq),len(xlen) {}
};

struct edge_t {
    int64_t from;
    int64_t to;
    int64_t len;
    edge_t() {}
    edge_t(int64_t f,int64_t t,int64_t l): from(f),to(t),len(l) {}
};

typedef std::list<edge_t>           node_list_t;
typedef node_list_t::iterator       node_iter_t;

struct graph_t {
    std::vector<node_list_t>  node;
    std::vector<edge_t>       edge;

    graph_t(int64_t n): node(n+1) {}

    int64_t edges() { return (int64_t)edge.size(); }
};


struct minq_t {
    int64_t node;
    int64_t len;
    minq_t(int64_t n,int64_t l): node(n),len(l) {}
    minq_t(): node(-1),len(-1) {}
};

typedef std::vector<node_t>                 node_array_t;

// edge + edge length


minq_t minq_get(deque<minq_t> &q)
{
    minq_t i;
    stable_sort(q.begin(),q.end(),[](minq_t a,minq_t b) {return a.len < b.len;});
    i = q.front();
    q.pop_front();
    return i;
}

int64_t minq_find(deque<minq_t> &q,int64_t n)
{
    for (int i=0;i<q.size();i++) {
        if (q[i].node == n) {
            return i;
        }
    }
    return -1;
}

void minq_put(deque<minq_t> &q,int64_t n,int64_t len)
{
    int64_t i = minq_find(q,n);
    if (i == -1) {
        q.push_back(minq_t(n,len));
    }
    else {
        q[i].len = len;
    }
}

graph_t *graphInstance(int64_t nodes,node_array_t &a)
{
    int64_t edges;
    graph_t *graph = new graph_t(nodes);

    // create the adjacency list and the edge list
    edges = 0;
    edge_t e;
    for (auto i=a.begin();i!=a.end();++i) {
        e = edge_t(i->p,i->q,i->len);
        graph->node[i->p].push_back(e);
        graph->edge.push_back(e);
        edges++;

        e = edge_t(i->q,i->p,i->len);
        graph->node[i->q].push_back(edge_t(i->q,i->p,i->len));
        graph->edge.push_back(e);
        edges++;
    }

    return graph;
}

void dijkstra(int64_t start,int64_t nodes,node_array_t a)
{
    // find the largest value in the vector
    graph_t         *graph;

    // traversal info
    vector<int64_t> dist(nodes+1);
    deque<minq_t>   minq;

    // instantiate a graph
    graph = graphInstance(nodes,a);

    for (int64_t i=1;i<nodes+1;++i) {
        dist[i]    = INT_MAX;
    }

    dist[start]    = 0;
    minq_put(minq,start,0);
    while (!minq.empty()) {
        minq_t mq = minq_get(minq);

        int64_t v = mq.node;
        int64_t d = mq.len;
        for (node_iter_t e=graph->node[mq.node].begin();e != graph->node[mq.node].end();e++) {
            int64_t w = e->to;
            if (dist[w] > (dist[v] + e->len)) {
                dist[w] = dist[v] + e->len;
                graph->edge[w] = *e;
                minq_put(minq,w,dist[w]);
            }
        }
    }
    for (int i=1;i<nodes+1;++i) {
        if (i == start) continue;
        if (dist[i] == INT_MAX) {
            dist[i] = -1;
        }
        cout << dist[i] << ' ';
    }
    cout << endl;
    cout << endl;

    // done with graph
    delete graph;
}

int main_dijkstra()
{
    int64_t t;
    int64_t nodes;
    int64_t edges;
    int64_t start;
    vector<node_t> a;

    cin >> t;

    for (int i=0;i<t;++i) {
        cin >> nodes;
        cin >> edges;
        a.clear();
        for (int j=0;j<edges;j++) {
            int64_t p;
            int64_t q;
            int64_t len;

            cin >> p;
            cin >> q;
            cin >> len;
            a.push_back(node_t(p,q,len));
        }
        cin >> start;
        dijkstra(start,nodes,a);

    }

    return 0;
}
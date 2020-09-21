#include <vector>
#include <array>
#include <cassert>
#include "../hr.h"

using namespace std;

typedef vector<int> adj_t;

struct ugraph_t : public undirected_graph_t
{
private:
    vector<vector<int> > m_vlist;
    int m_edges;
    int m_vertices;

public:
    ugraph_t(int vertices) : m_vlist(vertices)
    {
        m_vertices = vertices;
        m_edges    = 0;
    }

    ~ugraph_t(){}

    // number of vertices
    int vertices() override
    {
        return m_vertices;
    }

    // number of edges
    int edges() override
    {
        return m_edges;
    }

    // add an edge from v to w
    void addEdge(int v,int w, int weight) override
    {
        assert(v < static_cast<int>(m_vlist.size()));
        assert(w < static_cast<int>(m_vlist.size()));

        // count new vertices
        if (m_vlist[v].empty())
        {
            m_vertices++;
        }
        // put w on list for v
        m_vlist[v].push_back(w);

        // count edges
        m_edges++;

        // put v on list for w
        m_vlist[w].push_back(v);

        // count edges
        m_edges++;
    }

    // return list of adjacent vertices
    vector<int> adjacent(int v) override
    {
        assert(v < static_cast<int>(m_vlist.size()));

        // return a copy of the list for 'v'
        return m_vlist[v]; // (v);
    }
};

undirected_graph_t *undirected_graph_t::getInstance(int size)
{
    return new ugraph_t(size);
}

#include <iostream>
int ugraph_main(int argc,char *argv[])
{
    undirected_graph_t *g;

    g = undirected_graph_t::getInstance(5);

    g->addEdge(0,1);
    g->addEdge(0,2);
    g->addEdge(1,3);
    g->addEdge(2,4);

    for(int i=0;i<5;i++) {
        std::vector<int> v = g->adjacent(i);
        std::cout << i << " : ";
        std::for_each(v.begin(),v.end(),[](int e) { std::cout << e << ",";});
        std::cout << std::endl;

    }

    std::cout << g->vertices() << "," << g->edges() << std::endl;

    return 0;
}

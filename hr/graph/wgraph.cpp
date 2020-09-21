#include "../hr.h"

struct edge_weighted_graph_impl_t : public edge_weighted_graph_t
{
	int m_V;
	int m_E;
	vector<vector<edge_t *> > m_adj;

	edge_weighted_graph_impl_t(int V) : m_V(V), m_E(0) {

	}

	virtual int V() {
		return m_V;
	}

	virtual int E() {
		return m_E;
	}

	virtual void addEdge(edge_t *edge) {

	}

	// adjacency list
	vector<edge_t *> adj(int v) {
		return m_adj[v];
	}

	// all edges
	vector<edge_t *> edges() {
		vector<edge_t *> e;
		for (vector<edge_t *> a : m_adj) {
			for (edge_t *edge : a) {
				e.push_back(edge);
			}
		}
		return e;
	}

	static edge_weighted_graph_t *getInstance(int V) {
		return new edge_weighted_graph_impl_t(V);
	}
};

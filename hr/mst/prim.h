#ifndef PRIM_H__
#define PRIM_H__

#include "mst.h"

class Prim {
private:
	int m_V;
	int m_E;
	double m_weight;
	bool *m_marked;
	edge_weighted_graph_t &m_graph;
	set<edge_t> m_minq;
	vector<edge_t> m_edges;

	void visit(int v) {
		m_marked[v] = true;
		for (edge_t e : m_graph.adj(v)) {
			if (!m_marked[e.other(v)]) {
				m_minq.insert(e);
			}
		}
	}

public:
	Prim(edge_weighted_graph_t &graph) : m_graph(graph) {
		m_weight = 0.0;
		m_marked = new bool[m_graph.V()];
		for (int i = 0; i < m_graph.V(); i++) {
			m_marked[i] = false;
		}

		visit(0);
		while (!m_minq.empty()) {
			// get smallest edge 
			edge_t e = *m_minq.begin();
			m_minq.erase(m_minq.begin());

			// get the vertex indices for this edge
			int v = e.either();
			int w = e.other(v);

			// if both are marked, skip the rest of the loop and continue
			if (m_marked[v] && m_marked[w]) {
				continue;
			}

			// save the edge
			m_edges.push_back(e);
			m_weight += e.weight();

			// process the vertices for this edge
			if (!m_marked[v]) {
				visit(v);
			}
			if (!m_marked[w]) {
				visit(w);
			}
		}
	}

	~Prim() {
		delete[] m_marked;
	}

	vector<edge_t> &edges() {
		return m_edges;
	}

	double weight() {
		return m_weight;
	}


};

#endif

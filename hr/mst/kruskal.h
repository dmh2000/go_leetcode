#ifndef KRUSKAL_H__
#define KRUSKAL_H__

#include "mst.h"
#include "union-find.h"

class Kruskal {
private:
	vector<edge_t> m_edges;
	double m_weight;
	edge_weighted_graph_t &m_graph;
public:
	Kruskal(edge_weighted_graph_t &graph) : m_graph(graph) {
		set<edge_t> pq;
		for (edge_t e : graph.edges()) {
			pq.insert(e);
		}

		UF uf(graph.V());

		while (!pq.empty() && (m_edges.size() < (graph.V() - 1))) {
			edge_t e = *pq.begin();
			pq.erase(pq.begin());

			int v = e.either();
			int w = e.other(v);

			if (uf.connected(v, w)) {
				continue;
			}

			uf.union_(v, w);

			m_edges.push_back(e);

			m_weight += e.weight();
		}
	}

	vector<edge_t> &edges() {
		return m_edges;
	}

	double weight() {
		return m_weight;
	}
};

#endif

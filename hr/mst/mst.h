

#ifndef MST_H__
#define MST_H__

#include <algorithm>
#include <vector>
#include <set>
#include <memory>
#include <string>
#include <cassert>

using namespace std;
class edge_t
{
private:
	int m_v;
	int m_w;
	double m_weight;

public:
	// constructor
	edge_t(int v, int w, double weight) : m_v(v), m_w(w), m_weight(weight) {

	}

	edge_t() {
		m_v = 0;
		m_w = 0;
		m_weight = 0.0;
	}

	double weight() const { return m_weight; }

	int either() const { return m_v; }

	int other(int v) const {
		if (v == m_v) {
			return m_w;
		}
		else {
			return m_v;
		}
	}

	int compareTo(const edge_t &edge) const {
		if (m_weight < edge.weight()) {
			return -1;
		}
		else if (m_weight > edge.weight()) {
			return 1;
		}
		return 0;
	}

	string toString() const {
		stringstream ss;

		ss << '(' << m_v << ',' << m_w << ' ' << m_weight << ')';

		return ss.str();
	}

	edge_t getInstance(int v, int w, double weight) {
		return edge_t(v, w, weight);
	}
};

bool operator< (const edge_t &a, const edge_t &b) {
	return a.weight() < b.weight();
}

class edge_weighted_graph_t
{
private:
	int m_V;
	int m_E;

	unique_ptr<vector<edge_t>[]> m_adj;

public:

	edge_weighted_graph_t(int V, vector<edge_t> &edges) : m_V(V) {
		m_E = static_cast<int>(edges.size());
		m_adj = make_unique<vector<edge_t>[]>(V);
		for (edge_t edge : edges) {
			addEdge(edge);
		}
		for (int i = 0; i < V; i++) {
			int v = m_adj[i][0].either();
			int w = m_adj[i][0].other(v);
			double wgt = m_adj[i][0].weight();
			cout << v << ',' << w << ',' << wgt << endl;
		}
		cout << endl;
	}

	int V() {
		return m_V;
	}

	int E() {
		return m_E;
	}

	void addEdge(edge_t edge) {
		int v = edge.either();
		int w = edge.other(v);
		m_adj[v].push_back(edge);
		m_adj[w].push_back(edge);
		cout << v << '.' << w << endl;
	}

	// adjacency list
	vector<edge_t> adj(int v) {
		return m_adj[v];
	}

	// all edges
	vector<edge_t> edges() {
		vector<edge_t> e;
		for (int i = 0; i<m_V; i++) {
			for (edge_t edge : m_adj[i]) {
				e.push_back(edge);
			}
		}
		return e;
	}
};

#endif


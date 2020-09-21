#include <iostream>
#include <algorithm>
#include <vector>
#include <set>
#include <sstream>
#include <cstdlib>
#include <cstdint>
#include <climits>
#include <string>


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

	double weight() const { return m_weight; }

	int from() const { return m_v; }

	int to(int v) const {
		if (v == m_v) {
			return m_w;
		}
		else {
			return m_v;
		}
	}

	bool operator()(const edge_t &a, const edge_t &b) {
		return a.weight() < b.weight();
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

class edge_weighted_graph_t 
{
private:
	int m_V;
	int m_E;

	vector< vector<edge_t> > m_adj;

public:

	edge_weighted_graph_t(vector<edge_t> &edges) {
		for (edge_t edge : edges) {
			addEdge(edge);
		}
	}

	 int V() {
		return m_V;
	}

	 int E() {
		return m_E;
	}

	 void addEdge(edge_t edge) {
		 int v = edge.from();
		 int w = edge.to(v);
		 m_adj[v].push_back(edge);
		 m_adj[w].push_back(edge);
	}

	// adjacency list
	vector<edge_t> adj(int v) {
		return m_adj[v];
	}

	// all edges
	vector<edge_t> edges() {
		vector<edge_t> e;
		for (vector<edge_t> a : m_adj) {
			for (edge_t edge : a) {
				e.push_back(edge);
			}
		}
		return e;
	}
};

class Prim {
private:
	int m_V;
	int m_E;
	bool *m_marked;
	edge_weighted_graph_t &m_graph;
	set<edge_t> m_minq;

public:
	Prim(edge_weighted_graph_t &graph) : m_graph(graph) {
		m_marked = new bool[m_graph.V()];
	}
};

class Kruskal {

};
struct mst_t
{
	static mst_t *getInstance(edge_weighted_graph_t &g);

	virtual vector<edge_t> edges() = 0;

	virtual double weight() = 0;

};

int main(int argc, char *argv[])
{
	int t;
	int nodes;
	int edges;
	vector<edge_t> a;

	cin >> t;

	for (int i = 0; i<t; ++i) {
		cin >> nodes;
		cin >> edges;
		a.clear();
		for (int j = 0; j<edges; j++) {
			int v;
			int w;
			double weight;

			cin >> v;
			cin >> w;
			cin >> weight;
			a.push_back(edge_t(v, w, weight));
		}

		edge_weighted_graph_t graph{ a };

		mst_t *mst = mst_t::getInstance(graph);
	}

	return 0;
}
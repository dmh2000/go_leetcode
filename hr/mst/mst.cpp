#include <iostream>
#include <algorithm>
#include <vector>
#include <set>
#include <memory>
#include <sstream>
#include <cstdlib>
#include <cstdint>
#include <climits>
#include <string>
#include <cassert>

using namespace std;
class edge_t
{
private:
	int m_v;
	int m_w;
	int m_weight;

public:
	// constructor
	edge_t(int v, int w, int weight) : m_v(v), m_w(w), m_weight(weight) {

	}

	edge_t() {
		m_v = 0;
		m_w = 0;
		m_weight = 0;
	}

	int weight() const { return m_weight; }

	int either() const { return m_v; }

	int other(int v) const {
		if (v == m_v) {
			return m_w;
		}
		else  if (v == m_w) {
			return m_v;
		}
		else {
			assert(0);
			return -1;
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

	edge_t getInstance(int v, int w, int weight) {
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

	vector<edge_t> *m_adj;


public:

	edge_weighted_graph_t(int V, vector<edge_t> &edges) : m_V(V) {
		m_E = static_cast<int>(edges.size());
		m_adj = new vector<edge_t>[V];
		
		for (edge_t edge : edges) {
			addEdge(edge);
		}
		for (int i = 0; i < V; i++) {
			int v = m_adj[i][0].either();
			int w = m_adj[i][0].other(v);
			int wgt = m_adj[i][0].weight();
			//cout << v << ',' << w << ',' << wgt << endl;
		}
		//cout << endl;
	}
	~edge_weighted_graph_t() {
		delete[] m_adj;
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
		//cout << v << '.' << w << endl;
	}

	// adjacency list
	vector<edge_t> adj(int v) {
		return m_adj[v];
	}

	// all edges
	vector<edge_t> edges() {
		vector<edge_t> b;
		for (int i = 0; i<m_V; i++) {
			for (edge_t e : m_adj[i]) {
				if (e.other(i) > i) {
					b.push_back(e);
				}
			}
		}
		return b;
	}
};

class Prim {
private:
	int m_V;
	int m_E;
	int m_weight;
	bool *m_marked;
	edge_weighted_graph_t &m_graph;
	multiset<edge_t> m_minq;
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
	Prim(edge_weighted_graph_t &graph, int start) : m_graph(graph) {
		m_weight = 0;
		m_marked = new bool[m_graph.V()];
		for (int i = 0; i < m_graph.V(); i++) {
			m_marked[i] = false;
		}

		visit(start);

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

	int weight() {
		return m_weight;
	}
};


class UF {
private:
	int m_n;
	int m_count;
	vector<int> m_id;
	vector<int> m_sz;

public:
	UF(int n) : m_n(n), m_count(0) {

		for (int i = 0; i < n; i++) {
			m_id.push_back(i);
		}

		for (int i = 0; i < n; i++) {
			m_sz.push_back(1);
		}

	}

	void union_(int p, int q)
	{
		int pid = find(p);
		int qid = find(q);

		if (pid == qid) {
			return;
		}

		if (m_sz[pid] < m_sz[qid]) {
			m_id[pid] = qid;
			m_sz[qid] += m_sz[pid];
		}
		else {
			m_id[qid] = pid;
			m_sz[pid] += m_sz[qid];
		}

		m_count--;
	}

	int find(int p) {
		while (p != m_id[p]) {
			p = m_id[p];
		}
		return p;
	}

	bool connected(int p, int q) {
		return find(p) == find(q);
	}

	int count() {
		return m_count;
	}
};

class Kruskal {
private:
	vector<edge_t> m_edges;
	int m_weight;
	edge_weighted_graph_t &m_graph;
public:
	Kruskal(edge_weighted_graph_t &graph) : m_graph(graph) {
		m_weight = 0;

		multiset<edge_t> pq;
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

	int weight() {
		return m_weight;
	}
};

int main(int argc, char *argv[])
{
	int vertices;
	int edges;
	int start;
	vector<edge_t> a;

	cin >> vertices;
	cin >> edges;
	a.clear();
	for (int j = 0; j<edges; j++) {
		int v;
		int w;
		int weight;

		cin >> v;
		cin >> w;
		cin >> weight;
		a.push_back(edge_t(v-1, w-1, weight));
	}


	edge_weighted_graph_t graph{ vertices, a };

#define PRIM
#ifdef PRIM
	cin >> start;
	start -= 1;
	Prim mst{ graph, start };
#else
	Kruskal mst{ graph };
#endif

	/*for (edge_t e : mst.edges()) {
		cout << '(' << e.either() << ',' << e.other(e.either()) << ',' << e.weight() << ')' << endl;
	}*/
	cout << mst.weight() << endl;

	return 0;
}


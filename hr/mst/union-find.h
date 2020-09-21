#ifndef UNION_FIND_H__
#define UNION_FIND_H__

#include <vector>

using namespace std;

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
#endif


#include <string>
#include <sstream>
#include "../hr.h"

using namespace std;

struct edge_impl_t : public edge_t
{
	int m_v;
	int m_w;
	double m_weight;

	// construcyor
	edge_impl_t(int v, int w, double weight) : m_v(v), m_w(w), m_weight(weight) {

	}

	// destructor
	virtual ~edge_impl_t() {}

	virtual double weight() const { return m_weight; }


	virtual int either() const { return m_v; }

	virtual int other(int v) const  {
		if (v == m_v) {
			return m_w;
		}
		if (v == m_w) {
			return m_v;
		}
	}

	virtual int compareTo(const edge_t &edge) const {
		if (m_weight < edge.weight()) {
			return -1;
		}
		else if (m_weight > edge.weight()) {
			return 1;
		}
		return 0;
	}

	virtual string toString() const {
		stringstream ss;

		ss << '(' << m_v << ',' << m_w << ' ' << m_weight << ')';

		return ss.str();
	}

     edge_t *getInstance(int v, int w, double weight) {
		return new edge_impl_t(v, w, weight);
	}
};

#ifndef HR_H__
#define HR_H__
#include <vector>
using namespace std;

// invariants

// the vector from lo to hi inclusive is sorted
bool isSorted(const vector<int> &a,int lo,int hi);

// the vector from lo to hi inclusive is all less than m
bool isPartitioned(const vector<int> &a,int lo,int hi,int m);

// sorting
vector<int> mergeSort(vector<int> a);
vector<int> quickSort(vector<int> a);
vector<int> insertionSort(vector <int>  a);

// searching
int linearSearch(vector<int> a,int key);
int binarySearch(vector<int> a,int key);


// symbol table
struct key_value_t
{
    virtual ~key_value_t() {}

    virtual void insert(int key,int value) = 0;
    virtual int  find(int key) = 0;

    static key_value_t *getInstanceRBTree();
    static key_value_t *getInstanceHash(int size);
};

struct edge_t
{
	virtual ~edge_t() {}

	virtual double weight() const = 0;

	virtual int either() const = 0;

	virtual int other(int v) const  = 0;

	virtual int compareTo(const edge_t &that) const = 0;

	virtual string toString() const = 0;

	static edge_t *getInstance(int v, int w, double weight);
};

struct undirected_graph_t
{
    virtual ~undirected_graph_t(){}

    // number of vertices
    virtual int vertices() = 0;

    // number of edges
    virtual int edges() = 0;

    // add an edge from v to w
    virtual void addEdge(int v,int w, int weight=0) = 0;

    // return list of adjacent vertices
    virtual vector<int> adjacent(int v) = 0;

    // get an instance of the graph
    static undirected_graph_t *getInstance(int size);
};

struct edge_weighted_graph_t
{
	static edge_weighted_graph_t *getInstance(int V);

	virtual int V() = 0;

	virtual int E() = 0;

	virtual void addEdge(edge_t *e) = 0;

	// adjacency list
	virtual vector<edge_t *> adj(int v) = 0;

	// all edges
	virtual vector<edge_t *> edges() = 0;
};

struct mst_t
{
	static mst_t *getInstance(edge_weighted_graph_t *g);

	virtual vector<edge_t *> edges() = 0;

	virtual double weight() = 0;

};
#endif // HR_H__

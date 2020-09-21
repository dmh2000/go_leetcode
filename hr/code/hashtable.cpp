#include <cstdio>
#include <cassert>
#include <list>
#include <vector>
#include <algorithm>
#include "../hr.h"

using namespace std;

struct hashNode_t
{
    int m_key;
    int m_val;

    hashNode_t(int key,int val) : m_key(key),m_val(val)  {}
};

bool operator==(const hashNode_t& a,const hashNode_t &b)
{
    return a.m_key == b.m_key;
}

typedef list<hashNode_t>     hashList_t;
typedef hashList_t::iterator hashIter_t;

struct hashTable_t: public key_value_t
{
    int               m_size;
    vector<hashList_t>  *m_ht;

    hashTable_t(int size) : m_size(size) {
        // create the hashlist vector of the specified size
        m_ht = new vector<hashList_t>(size);
    }

    virtual ~hashTable_t()  {}

    // derived key_value_t function
    void insert(int key,int value) override
    {
        // get the hash of the key
        int h = hash(key);
        assert(h < m_ht->size());

        // get the list for this hash
        hashList_t &hl = m_ht->at(h);

        // add it to the list
        hl.push_back(hashNode_t(key,value));
    }

    // derived key_value_t function
    int find(int key) override
    {
        int h = hash(key);
        
        // get the list for this hash
        hashList_t &hl = m_ht->at(h);

        // try to find the key in the hash list
        hashIter_t v = std::find(hl.begin(),hl.end(),hashNode_t(key,0));

        if(v == hl.end()) {
            // not found
            return -1;
        }
        else {
            // found
            return v->m_val;
        }
    }

    // private functions

    // use FowlerNollVo FNV has for strings
    int hash(int key)
    {
        // compute a simple hash
        int h = key % m_size;
        assert(h < m_size);

        return h;
    }
};


// class factory for red-black tree
key_value_t *key_value_t::getInstanceHash(int size)
{
    return new hashTable_t(size);
}

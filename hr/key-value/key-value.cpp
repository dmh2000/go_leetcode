#include <vector>
#include <cassert>
#include "../hr.h"

using namespace std;


int main(int argc,char *argv[])
{
    key_value_t *rbt;
    key_value_t *hsh;
    int key;
    int val;
    int res;
    rbt = key_value_t::getInstanceRBTree();
    hsh = key_value_t::getInstanceHash(100);

    for(int i=0;i<1000;i++) {
        key = i;
        val = i*100;
        rbt->insert(key,val);
        hsh->insert(key,val);
    }

    for(int i=999;i>=0;i--) {
        key = i;
        val = i*100;
        res = rbt->find(key);
        assert(res == val);
        res = hsh->find(key);
        assert(res == val);
    }

    res = rbt->find(100000);
    assert(res == -1);
    res = hsh->find(100000);
    assert(res == -1);

    return 0;
}

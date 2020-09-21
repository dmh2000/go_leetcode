#include <cstdio>
#include <cassert>
#include "../hr.h"

/**
 * this code is based on the Java implementation from Algorithms by Sedgewick algs4.cs.princeton.edu/code
 * its a subset of that code base
 */
const int RBT_RED   = 0;
const int RBT_BLACK = 1;

struct rb_node_t
{
    int   key;      // search key
    int   val;      // value
    rb_node_t *left;     // left subtree
    rb_node_t *right;    // right subtree
    int   n;        // nodes in subtree
    int   color;    // color of link from parent

    rb_node_t() {
        assert(0);
    }

    rb_node_t(int key,int val,int n,int color) :
        key(key),
        val(val),
        left(nullptr),
        right(nullptr),
        n(n),
        color(color)
    {
    }
};

struct rbtree: public key_value_t
{
    rb_node_t *root;

    rbtree(): root(nullptr) {}

    virtual ~rbtree()  {}

    // derived key_value_t function
    void insert(int key,int value) override
    {
        // insert new rb_node_t starting at root, may return a new root
        root = insert(key,value,root);
        root->color = RBT_BLACK;
    }

    // derived key_value_t function
    int find(int key) override
    {
        return find(key,root);
    }

    // =========================================
    // private functions
    // =========================================
    int find(int key,rb_node_t *h)
    {
        if(h == nullptr) {
            return -1;
        }

        // found
        if(key == h->key) {
            return h->val;
        }

        // search less than
        if (key < h->key) {
            return find(key,h->left);
        }

        // search greater than
        if(key > h->key) {
            return find(key,h->right);
        }

        // won't get here
        assert(0);
        return -1;
    }

    bool isRed(rb_node_t *h)
    {
        if(h == nullptr) {
            return false;
        }
        return (h->color == RBT_RED);
    }

    int size(rb_node_t *h)
    {
        if(h == nullptr) {
            return 0;
        }
        return h->n;
    }


    // before
    //              E
    //      black       red
    //       <E          S
    //                ES   >S
    //      
    // after
    //              S
    //       red         blk
    //        E          >S
    //    <E     ES
    // 
    rb_node_t *rotateLeft(rb_node_t *h)
    {
        rb_node_t *x  = h->right;
        h->right = x->left;
        x->left  = h;
        x->color = h->color;
        h->color = RBT_RED;
        x->n     = h->n;
        h->n     = 1 + size(h->left) + size(h->right);
        return x;
    }

    // before
    //             S
    //      red          blk
    //      E            >S
    //  <E    ES
    //  
    // after
    //            E
    //     blk          red
    //     <E            S
    //               ES     >S
    rb_node_t *rotateRight(rb_node_t *h)
    {
        rb_node_t *x  = h->left;
        h->left = x->right;
        x->right  = h;
        x->color = h->color;
        h->color = RBT_RED;
        x->n     = h->n;
        h->n     = 1 + size(h->right) + size(h->left);
        return x;
    }

    rb_node_t *insert(int key,int value,rb_node_t *h)
    {
        // new tree
        if(h == nullptr) {
            return new rb_node_t(key,value,1,RBT_BLACK);
        }

        // recursive insert
        if(key < h->key) {
            h->left = insert(key,value,h->left);
        }
        else if (key > h->key) {
            h->right = insert(key,value,h->right);
        }
        else {
            h->val = value;
        }

        // fixup

        // right is red and left is not red
        // rotate left to make left red and right black
        if(isRed(h->right) && !isRed(h->left)) {
            h = rotateLeft(h);
        }
        // left child is red and left child child is red
        // rotate right to make left child black
        if(isRed(h->left) && isRed(h->left->left)) {
            h = rotateRight(h);
        }
        // both left and right children are red
        // make this one red, both children black
        if(isRed(h->left)&&isRed(h->right)) {
            flipColors(h);
        }
        h->n = size(h->left) + size(h->right) + 1;

        return h;
    }

    void flipColors(rb_node_t *h)
    {
        h->color        = RBT_RED;
        h->left->color  = RBT_BLACK;
        h->right->color = RBT_BLACK;
    }

};


// class factory for red-black tree
key_value_t *key_value_t::getInstanceRBTree()
{
    return new rbtree();
}

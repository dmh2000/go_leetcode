#include <Windows.h>
#include <iostream>
#include <mutex>
#include <algorithm>
#include <functional>
#include <list>


static void printNumber(int n) {
    std::cout << n ;
}

class ZeroEvenOdd {
private:
    int n;
    std::list<int> a;
    std::mutex mtx;
    int front() {
        int i;
        i =  a.front();
        return i;
    }

    int get() {
        int i;
        i = a.front();
        a.pop_front();
        return i;
    }

    bool empty() {
        bool b;
        b = a.empty();
        return b;
    }

public:
    ZeroEvenOdd(int n) {
        this->n = n;
        for (int i = 1; i <= n; i++) {
            a.push_back(0);
            a.push_back(i);
        }
    }

    // printNumber(x) outputs "x", where x is an integer.
    void zero(std::function<void(int)> printNumber) {

        int i;
        while (!empty()) {
            std::unique_lock<std::mutex> lk(mtx);
            if (!empty() && (front() == 0)) {
                i = get();
                printNumber(i);
            }
            lk.unlock();
            std::this_thread::yield();
          }

    }

    void even(std::function<void(int)> printNumber) {
        while (!empty()) {
            std::unique_lock<std::mutex> lk(mtx);
            if (!empty() && ((front() % 2) == 0)) {
                int i = get();
                printNumber(i);
            }
            lk.unlock();
            std::this_thread::yield();
        }
    }

    void odd(std::function<void(int)> printNumber) {
        while (!empty()) {
            std::unique_lock<std::mutex> lk(mtx);
            if (!empty() && ((front()%2) == 1)) {
                int i = get();
                printNumber(i);
            }
            lk.unlock();
            std::this_thread::yield();
        }
    }
};




static void exec_zeo(int i, ZeroEvenOdd* zeo) {
    switch (i) {
    case 1:
        zeo->zero(printNumber);
        break;
    case 2:
        zeo->odd(printNumber);
        break;
    case 3:
        zeo->even(printNumber);
        break;
    }
}


void run_evenodd1(int n) {
    ZeroEvenOdd zeo(5);

    // start foo
    std::thread t2(exec_zeo, 2, &zeo);
    std::thread t1(exec_zeo, 1, &zeo);
    std::thread t3(exec_zeo, 3, &zeo);


    // wait for it to finish
    t1.join();
    t2.join();
    t3.join();
    std::cout << "\n";
}



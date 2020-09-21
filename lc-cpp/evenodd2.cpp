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
    std::condition_variable cv;

    int get() {
        int i;
        i = a.front();
        a.pop_front();
        return i;
    }

    int get_zero() {
        int r = 0;
        std::unique_lock<std::mutex> lk(mtx);

        for (;;) {
            if (a.empty()) {
                r = -1;
                break;
            }
            if (a.front() == 0) {
                r = get();
                break;
            }
            else {
                cv.wait(lk);
            }
        }
        lk.unlock();

        return r;
    }

    int get_odd() {
        int r = 0;
        std::unique_lock<std::mutex> lk(mtx);

        for (;;) {
            if (a.empty()) {
                r = -1;
                break;
            }

            if ((a.front() % 2) == 1) {
                r = get();
                break;
            }
            else {
                cv.wait(lk);
            }
        }

        lk.unlock();

        return r;
    }

    int get_even() {
        int r = 0;
        std::unique_lock<std::mutex> lk(mtx);

        for (;;) {
            if (a.empty()) {
                r = -1;
                break;
            }

            if ((a.front() != 0) && ((a.front() % 2) == 0)) {
                r = get();
                break;
            }
            else {
                cv.wait(lk);
            }
        }

        lk.unlock();

        return r;
    }

    void notify() {
        cv.notify_all();
    }

public:
    ZeroEvenOdd(int n) {
        this->n = n;

        // fill the queue with the output in order
        for (int i = 1; i <= n; i++) {
            a.push_back(0);
            a.push_back(i);
        }
    }

    // printNumber(x) outputs "x", where x is an integer.
    void zero(std::function<void(int)> printNumber) {

        int i;
        for (;;) {
            i = get_zero();
            if (i < 0) {
                break;
            }
            printNumber(i);
            notify();
        }
    }

    void even(std::function<void(int)> printNumber) {
        int i;
        for (;;) {
            i = get_even();
            if (i < 0) {
                break;
            }
            printNumber(i);
            notify();
        }
    }

    void odd(std::function<void(int)> printNumber) {
        int i;
        for (;;) {
            i = get_odd();
            if (i < 0) {
                break;
            }
            printNumber(i);
            notify();
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


void run_evenodd(int n) {
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



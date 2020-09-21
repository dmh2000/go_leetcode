#include <iostream>
#include <mutex>
#include <algorithm>
#include <functional>

void printFoo() {
    std::cout << "Foo";
}

void printBar() {
    std::cout << "Bar\n";
}

class FooBar {
private:
    int n;

    int x;
    std::mutex mtx;
    std::condition_variable cv;

    void next(int n,int m) {
        // lock the mutex
        std::unique_lock<std::mutex> lk(mtx);

        for (;;) {
            // wait on cv
            if (n == x) {

                // set m to next value
                x = m;

                // notify waiters
                cv.notify_one();

                // leave the loop
                break;
            }
            else {
                cv.wait(lk);
            }
        }

        // unlock the mutex
        lk.unlock();
    }

public:
    FooBar(int n) {
        this->n = n;
        this->x = 1;
        
    }


    void foo(std::function<void()> printFoo) {

        for (int i = 0; i < n; i++) {
            next(1,2);

            // printFoo() outputs "foo". Do not change or remove this line.
            printFoo();

            next(2,3);
        }
    }

    void bar(std::function<void()> printBar) {

        for (int i = 0; i < n; i++) {
            next(3,4);

            // printBar() outputs "bar". Do not change or remove this line.
            printBar();

            next(4,1);
        }
    }
};


void exec_foobar(int i, FooBar* foobar) {
    switch (i) {
    case 1:
        foobar->foo(printFoo);
        break;
    case 2:
        foobar->bar(printBar);
        break;
    }
}

void run_foobar() {
    FooBar foobar(10);

    // start foo
    std::thread t2(exec_foobar, 2, &foobar);
    std::thread t3(exec_foobar, 1, &foobar);
    // wait for it to finish

    t2.join();
    t3.join();
}

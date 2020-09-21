#include <iostream>
#include <mutex>
#include <algorithm>
#include <functional>


void printFirst() {
    std::cout << "first\n";
}

void printSecond() {
    std::cout << "second\n";
}

void printThird() {
    std::cout << "third\n";
}

class Foo {
private:

    std::mutex mtx;
    std::condition_variable cv;
    int  f;



public:
    Foo() :  f(1) {
        
    }


    void next(int i) {
        // lock the mutex
        std::unique_lock<std::mutex> lk(mtx);

        for (;;) {
            // wait on cv
            if (i == f) {
                // increment f for next function
                f++;

                // notify waiters
                cv.notify_all();

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

    void first(std::function<void()> printFirst) {

        // printFirst() outputs "first". Do not change or remove this line.
        next(1);
        printFirst();
        next(2);
    }

    void second(std::function<void()> printSecond) {

        // printSecond() outputs "second". Do not change or remove this line.
        next(3);
        printSecond();
        next(4);

    }

    void third(std::function<void()> printThird) {

        // printThird() outputs "third". Do not change or remove this line.
         next(5);
        printThird();

    }
};

void exec_foo(int i, Foo *foo) {
    switch (i) {
    case 1:
        foo->first(printFirst);
            break;
    case 2:
        foo->second(printSecond);
        break;
    case 3:
        foo->third(printThird);
        break;
    }
}

void run_foo() {
    Foo foo;

    // start foo
    std::thread t3(exec_foo, 3, &foo);
    std::thread t2(exec_foo, 2, &foo);
    std::thread t1(exec_foo, 1, &foo);
        // wait for it to finish

    t1.join();
    t2.join();
    t3.join();
}

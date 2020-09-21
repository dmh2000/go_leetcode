#include <string>
#include <iostream>
#include <mutex>
#include <algorithm>
#include <functional>
#include <list>


class H2O {
    std::mutex mtx_hyd;
    std::condition_variable cv_hyd;
    std::mutex mtx_oxy;
    std::condition_variable cv_oxy;
    int hyd_out;
    int oxy_out;

    // release 2 hydrogens then wait for an oxygen
    void hyd(std::function<void()> releaseHydrogen) {
        std::unique_lock<std::mutex> lk(mtx_hyd);
        for (;;) {
            if (hyd_out > 0) {
                releaseHydrogen();
                hyd_out--;

                if ((hyd_out == 0) && (oxy_out == 0)) {
                    hyd_out = 2;
                    oxy_out = 1;
                }
                break;
            }
            else {
                // wait for an oxygen
                cv_hyd.wait(lk);
            }
        }
        lk.unlock();

        // notify oxygen
        cv_oxy.notify_all();
    }

    // release one oxygen then wait for 2 hydrogens
    void oxy(std::function<void()> releaseOxygen) {
        std::unique_lock<std::mutex> lk(mtx_oxy);
        for (;;) {
            if (oxy_out > 0) {
                releaseOxygen();
                oxy_out--;
                if ((hyd_out == 0) && (oxy_out == 0)) {
                    hyd_out = 2;
                    oxy_out = 1;
                }
                break;
            }
            else {
                // wait for a hydrogen
                cv_oxy.wait(lk);
            }
        }
        lk.unlock();

        // notify the hydrogens
        cv_hyd.notify_all();
    }


public:
    H2O() {
        hyd_out = 2;
        oxy_out = 1;
    }

    void hydrogen(std::function<void()> releaseHydrogen) {
        // releaseHydrogen() outputs "H". Do not change or remove this line.
        hyd(releaseHydrogen);
    }

    void oxygen(std::function<void()> releaseOxygen) {
        oxy(releaseOxygen);
    }
};

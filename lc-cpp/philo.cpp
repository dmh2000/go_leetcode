#include <Windows.h>
#include <iostream>
#include <mutex>
#include <algorithm>
#include <functional>
#include <list>


class BinarySemaphore {
private:
	int count;
	std::condition_variable cv;
	std::mutex mtx;

public:

	BinarySemaphore(int v) {
		count = v;
	}

	void put() {
		std::unique_lock<std::mutex> lk(mtx);
		// only notify on the first put
		if (count == 0) {
			count = 1;
		}
		else {
			std::cout << "put error\n";
		}
		lk.unlock();

		cv.notify_one();
	}

	void get() {
		std::unique_lock<std::mutex> lk(mtx);

		for (;;) {
			if (count == 1) {
				count = 0;
				break;
			}
			else {
				cv.wait(lk);
			}
		}
		lk.unlock();

		cv.notify_one();
	}
};


class Philosopher {
private:
	std::function<void()> f_pickLeft;
	std::function<void()> f_pickRight;
	std::function<void()> f_eat;
	std::function<void()> f_putLeft;
	std::function<void()> f_putRight;


public:
	bool set;

	Philosopher() {
		set = false;
		f_pickLeft = nullptr;
		f_pickRight = nullptr;
		f_eat = nullptr;
		f_putLeft = nullptr;
		f_putRight = nullptr;
	}

	void update(
		std::function<void()> pickLeftFork,
		std::function<void()> pickRightFork,
		std::function<void()> eat,
		std::function<void()> putLeftFork,
		std::function<void()> putRightFork
	) {
		f_pickLeft  = pickLeftFork;
		f_pickRight = pickRightFork;
		f_eat = eat;
		f_putLeft = putLeftFork;
		f_putRight = putRightFork;
		set = true;
	}

	void get() {
		f_pickLeft();
		f_pickRight();
		f_eat();
	}

	void release() {
		f_putLeft();
		f_putRight();
	}
};

class BinarySemaphore {
private:
	int count;
	std::condition_variable cv;
	std::mutex mtx;

public:

	BinarySemaphore() {
		count = 1;
	}

	void put() {
		std::unique_lock<std::mutex> lk(mtx);
		// only notify on the first put
		if (count == 0) {
			count = 1;
		}
		else {
			std::cout << "put error\n";
		}
		lk.unlock();

		cv.notify_one();
	}

	void get() {
		std::unique_lock<std::mutex> lk(mtx);

		for (;;) {
			if (count == 1) {
				count = 0;
				break;
			}
			else {
				cv.wait(lk);
			}
		}
		lk.unlock();

		cv.notify_one();
	}
};



class DiningPhilosophers {
private:
	// one at a time
	/*
		if none ar eating
		allow one to pick up forks
		if one is eating
		release current forks
		new one picks up forks
	
	*/
	BinarySemaphore bsem;
	
	bool eating[5];
	BinarySemaphore fork[5];

public:
	DiningPhilosophers() : bsem() {
	}

	void wantsToEat(int philosopher,
		std::function<void()> pickLeftFork,
		std::function<void()> pickRightFork,
		std::function<void()> eat,
		std::function<void()> putLeftFork,
		std::function<void()> putRightFork) {

		//bsem.get();

		//pickLeftFork();
		//pickRightFork();
		//eat();
		//putLeftFork();
		//putRightFork();

		//bsem.put();

		bsem.get();
		if (!eating[philosopher]){
			fork[philosopher].get();
			pickLeftFork();
			fork[philosopher+1].get();
			pickRightFork();
			eat();
			eating[philosopher]= true;
		}
		else {
			putLeftFork();
			fork[philosopher].put();
			putRightFork();
			fork[philosopher + 1].put();
			eating[philosopher] = false;
		}
		bsem.put();
	}
};


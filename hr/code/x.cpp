#include <iostream>
#include <algorithm>
#include <vector>
#include <cstdlib>

using namespace std;

int kaden(const vector<int> &a)
{
    int m1;
    int m2;
    m1 = m2 = a[0];
    for (int i=1;i<a.size();++i) {
        int x = a[i];
        int y = m1 + a[i];
        m1 = std::max(x,y);
        m2 = std::max(m1,m2);
    }
    return m2;
}

int dm(const vector<int> &a)
{
    int m;
    int *srt = new int[a.size()];
    int *end = new int[a.size()];
    m = a[0];
    int j = 0;
    srt[j] = j;
    end[j] = j;
    for (int i=1;i<a.size();i++) {
        int x = m + a[i];
        int y = m;
        if (x > y) {
            m = x;
            end[j] = i;
        }
        else {
            j++;
            m = 0;
            srt[j] = i+1;
            end[j] = i+1;
        }
    }
    cout << srt[j] << ' ' << end[j] << endl;

    delete[] srt;
    delete[] end;

    return m;
}

int main(int argc,char *argv[])
{
    vector<int> a;

    a.push_back(1);
    a.push_back(1);
    a.push_back(-2);
    a.push_back(1);
    a.push_back(2);


    int m;

    m = dm(a);
    cout << m << endl;

    m = kaden(a);
    cout << m << endl;

    return 0;
}

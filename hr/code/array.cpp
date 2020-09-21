#include <cstdio>
#include <cstdlib>
#include <array>


using namespace std;

int main(void)
{
    array<int,20> a;

    for(array<int,20>::iterator i=a.begin();i!= a.end();i++) {
        *i = rand();
        printf("%d\n",*i);
    }



    return 0;
}


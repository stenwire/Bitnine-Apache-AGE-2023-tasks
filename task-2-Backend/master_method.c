#include<stdio.h>

int piecewise(int i)
{
    if(i == 0) return 0;
    if(i == 1) return 1;
    if(i == 2) return 2;

    return piecewise(i-3) + piecewise(i-2);
}

int main()
{
    int i;

    for (int i = 0; i < 20; i++)
    {
        printf("%d\t", piecewise(i));
    }

    return 0;
}

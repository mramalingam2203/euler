#include <stdio.h>
#define MAX_NUM 4000000
int main() {
    long unsigned int i, n, t1 = 0, t2 = 1, nextTerm;
    printf("Enter the number of terms: ");
    scanf("%ld", &n);
    printf("Fibonacci Series: ");

    for (i = 1; i <= n; ++i) {
        
        if (t1 > MAX_NUM){
            printf("%ld, ", t1);
            break;
        } 
        nextTerm = t1 + t2;
        t1 = t2;
        t2 = nextTerm;
    }

    return 0;
}
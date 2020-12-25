#include <stdio.h>
int main() {
   long int n;
   n=600851475143;
   long int div=2, ans = 0, maxFact;
   while(n!=0) {
      if(n % div !=0)
         div = div + 1;
      else {
         maxFact = n;
         n = n / div;
         if(n == 1) {
            printf("%ld is the largest prime factor !\n",maxFact);
            ans = 1;
            break;
         }
      }
   }
   return 0;
}

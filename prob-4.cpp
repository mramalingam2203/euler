/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

#include <stdio.h>
#define MAX_NUM 999
using namespace std;
void check_palindrome(int);

int main(){

	for (int i=100; i<=MAX_NUM; i++)
		for (int j=100; j <= MAX_NUM ; j++)
			//printf("%d\n",i*j);
			check_palindrome(i*j)
	
	return 0;
}

void check_palindrome(int number){

	while (number > 0)
	{
	    int digit = number%10;
	    number /= 10;
	    printf("%d",digit);
	}

	stack<int> sd;

	while (number > 0)
	{
	    int digit = number%10;
	    number /= 10;
	    sd.push(digit);
	}

	while (!sd.empty())
	{
	    int digit = sd.top();
	    sd.pop();
	    printf("%d",digit);
	}

}
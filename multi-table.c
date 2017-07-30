/*
* @Author: suifengtec
* @Date:   2017-07-30 13:21:46
* @Last Modified by:   suifengtec
* @Last Modified time: 2017-07-30 13:40:13
*/

/*
gcc main.c && a
 */
#include <stdio.h>

void multiTable_v1(){

	int i;
	int j;

	for(i=1;i<=9;i++){


		for (j=1;j<=9;j++){

			printf("%d*%d=%2d\t",i,j,i*j);
		}

		printf("\n");
	}


}

void multiTable_v2(){

	int i;
	int j;

	for(i=1;i<=9;i++){


		for (j=1;j<=9;j++){

			if(j<i){

				printf("        ");
			}else{

				printf("%d*%d=%2d\t",i,j,i*j);
			}
			
		}

		printf("\n");
	}


}

void multiTable_v3(){

	int i;
	int j;

	for(i=1;i<=9;i++){


		for (j=1;j<=9;j++){

			if(j>i){

				printf("        ");
			}else{

				printf("%d*%d=%2d\t",i,j,i*j);
			}
			
		}

		printf("\n");
	}


}


void multiTable_v4(){

	int i, j,n;

	for(i=1;i<=9;i++){
			for (n=1;n<=9-i;n++){

				printf("        ");
			}

		for (j=1;j<=9;j++){

				printf("%d*%d=%2d\t",i,j,i*j);

			
		}

		printf("\n");
	}


}

int main() {

	printf("\nVer1\n");
	printf("\n********\n");

	multiTable_v1();


	printf("\nVer2\n");
	printf("\n********\n");

	multiTable_v2();

	printf("\nVer 3\n");
	printf("\n********\n");

	multiTable_v3();

	printf("\nVer 4\n");
	printf("\n********\n");


	multiTable_v4();

    return 0;
}
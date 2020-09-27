#include <stdio.h>
#include <unistd.h>
#include <string.h>


/* Bakend-to-database */

int setup();
int add();


int main()
{   
    char fname[8] = "master.dlb";
    if( access( fname, F_OK ) != -1 ) 
    {
    // file exists
    return 1;
    } 
    else 
    {
    // file doesn't exist
    setup();
    printf("\n Creating file... \n");
    return 0;
    }

}

int setup() 
{
    FILE *fp;

    fp = fopen("master.dib", "w+");
    fprintf(fp, "This is testing for fprintf...\n");
    fputs("This is testing for fputs...\n", fp);
    fclose(fp);
    return 0;
}


int add(key) 
{
    return 0;
}

#include <stdio.h>
#include <unistd.h>
#include <string.h>


/* Bakend-to-database */
int checkdb();
int createfile();

int add(char doc);
int read_doc(char doc);
int delete(char doc);
int update(char doc);


char fname[8] = "main.dlb";

int main()
{
    int c;
    checkdb();
    printf( "Enter a function : (-r -a -d -u)  +  doc ");
    c = getchar( );
    printf( "\nYou entered: ");
    putchar( c );


}

int checkdb() 
{
    if( access(fname, F_OK) != 1 )
    {
        start();
        return true;
    }
    else 
    {
        // file do not exist
        printf("\n Creating database");
        createfile();
    }
}

// Create init database file 

int createfile() 
{
    FILE *fp;

    fp = fopen(fname, "w+");
    fprintf(fp, "This is testing for fprintf...\n");
    fputs("This is testing for fputs...\n", fp);
    fclose(fp);
    return 0;
}

// Add to database

int add(char doc) 
{
    return 0;
}

// Read database

int read_doc(char doc)
{
    return 0;
}

// Delete document or collection

int delete(char doc)
{
    return 0;
}

// Update codument

int update(char  doc) 
{
    return 0;
}



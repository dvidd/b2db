
#include <stdio.h>

int setup() {
    
    FILE *fp;

    fp = fopen("db.json", "w+");
    fprintf(fp, "This is testing for fprintf...\n");
    fputs("This is testing for fputs...\n", fp);
    fclose(fp);
}

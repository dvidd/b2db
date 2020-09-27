#include <stdio.h>

int auth(int key);
int delete_doc(int index);

int  delete(int key,int index) 
{
    if (auth(key) == 0)
    {
        if(delete_doc(index) == 0)
        {
            printf("Doc deleted correctly");
            return 0;
        }
    } else 
    {   
        printf("Not valid key");
        return 1;
    }
}

int auth(int key) 
{
    return 0;
}

int delete_doc(int index) 
{
    return 0;
}

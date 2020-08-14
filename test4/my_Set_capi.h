#include <stdbool.h>
typedef struct myset_T myset_T; 
myset_T* newmyset();
void insertMyset(myset_T* p,void* i);
bool findMyset(myset_T* p,void* i);
void eraseMyset(myset_T* p,void* i);
int sizeMyset(myset_T* p);

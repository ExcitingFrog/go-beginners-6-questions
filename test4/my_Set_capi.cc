#include "my_Set.h" 
extern "C"{
  #include "my_Set_capi.h"
}
struct myset_T:myset{};
myset_T* newmyset(){
  auto p = new myset_T();
  return p;
}
void insertMyset(myset_T* p,void* i){
  p->Insert(i);
}
bool findMyset(myset_T* p,void* i){
  return p->Find(i);
}
int sizeMyset(myset_T* p){
  return p->Size();
}
void eraseMyset(myset_T* p,void* i){
  return p->Erase(i);
}

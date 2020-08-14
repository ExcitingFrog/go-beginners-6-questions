#include <set>
struct myset{
  std::set<void*> s;
  void  Insert(void* i) {
		this->s.insert(i); 
  }
  bool Find(void* i){
		if(this->s.find(i) != this->s.end())
			return true;
		return false;
  }

  int Size(){
  	return this->s.size();
  }
  void Erase(void* i){
  	this->s.erase(i);
  }
};
// int main(){
  // auto nset = new()
  // auto find =
// }

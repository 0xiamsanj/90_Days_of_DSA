#include<bits/stdc++.h>
using namespace std; 
int main (int argc, char *argv[])
{
  int n;
  cin>>n;

  for(int i=n;i>0;i--){
    for(int j=i-1;j>0;j--){
      cout<<"-";
    }
    for(int k=n;k>=i;k--){
      cout<<k;
    }
    for(int k=n;k>i;k--){
      cout<<k;
    }
    cout<<endl;
  }
  return 0;
}

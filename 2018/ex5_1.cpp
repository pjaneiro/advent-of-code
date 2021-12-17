#include <cstdlib>
#include <cstdio>
#include <cstring>

#define INPUT_SIZE 50001
#define CAP_DIF 32

char input[INPUT_SIZE] = {};
bool found;
int i, j, cur_size;

int main(int argc, char* argv[]) {
  scanf("%s", &input);
  do{
    found = false;
    cur_size = strlen(input);
    for(i=1; i<cur_size; i++) {
      if(input[i] == input[i-1]+CAP_DIF || input[i] == input[i-1]-CAP_DIF ) {
        found = true;
        for(j = i-1; j<=cur_size-2; j++) {
          input[j] = input[j+2];
        }
        cur_size -= 2;
      }
    }
  } while(found);
  printf("%d\n", cur_size);
  return 0;
}

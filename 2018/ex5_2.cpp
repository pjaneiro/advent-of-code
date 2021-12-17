#include <cstdlib>
#include <cstdio>
#include <cstring>
#include <climits>

#define INPUT_SIZE 50001
#define CAP_DIF 32

char input[INPUT_SIZE] = {};
char original_input[INPUT_SIZE] = {};
bool found;
int i, j, k, cur_size, best_size = INT_MAX;
char letter, best_letter;

int main(int argc, char* argv[]) {
  scanf("%s", &original_input);
  for(letter = 65; letter<91; letter++) {
    strcpy(input, original_input);
    cur_size = strlen(input);
    for(i=0; i<cur_size; i++) {
      if(input[i] == letter || input[i] == letter+CAP_DIF) {
        for(j=i; j<cur_size; j++) {
          input[j] = input[j+1];
        }
        cur_size--;
        i--;
      }
    }

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

    if(cur_size < best_size) {
      best_size = cur_size;
    }
  }
  printf("%d\n", best_size);
  return 0;
}

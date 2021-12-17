#include <cstdlib>
#include <cstdio>
#include <cstring>

#define INPUT_SIZE 101
#define ALPHA_SIZE 26
#define FIRST_LETTER 65

int adj_matrix[ALPHA_SIZE][ALPHA_SIZE] = {0};
int already_allocated[ALPHA_SIZE] = {0};
char final_string[ALPHA_SIZE+1] = {0};
int in_place = 0, i, j, cur_letter, final_index=0;
char from, to;
bool clear;

int main(int argc, char* argv[]) {
  for(i=0; i<INPUT_SIZE; i++) {
    scanf("Step %c must be finished before step %c can begin.\n", &from, &to);
    adj_matrix[from - FIRST_LETTER][to - FIRST_LETTER] = 1;
  }

  while(in_place < ALPHA_SIZE) {
    for(cur_letter=0; cur_letter<ALPHA_SIZE; cur_letter++) {
      clear = true;
      for(i=0; i<ALPHA_SIZE; i++) {
        if(adj_matrix[i][cur_letter] != 0) {
          clear = false;
        }
      }
      if(clear && already_allocated[cur_letter] == 0) {
        final_string[in_place] = cur_letter + FIRST_LETTER;
        in_place++;
        already_allocated[cur_letter] = 1;
        for(j=0; j<ALPHA_SIZE; j++) {
          adj_matrix[cur_letter][j] = 0;
        }
        cur_letter = -1;
      }
    }
  }

  printf("%s\n", final_string);

  return 0;
}

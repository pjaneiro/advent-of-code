#include <cstdlib>
#include <cstdio>

#define INPUT_SIZE 101
#define ALPHA_SIZE 26
#define FIRST_LETTER 65
#define BASE_TIME 60
#define NUMBER_OF_WORKERS 5
#define ARBITRARILY_LARGE_NUMBER 2000

int timeline[ARBITRARILY_LARGE_NUMBER][NUMBER_OF_WORKERS] = {0};
int adj_matrix[ALPHA_SIZE][ALPHA_SIZE] = {0};
bool already_allocated[ALPHA_SIZE] = {false};
bool already_finished[ALPHA_SIZE] = {false};
int dependencies[ALPHA_SIZE] = {0};
char from, to;
int second, worker, letter, i, j, in_place, max;

int main(int argc, char* argv[]) {
  for(i=0; i<INPUT_SIZE; i++) {
    scanf("Step %c must be finished before step %c can begin.\n", &from, &to);
    adj_matrix[from - FIRST_LETTER][to - FIRST_LETTER] = 1;
    dependencies[to - FIRST_LETTER]++;
  }

  for(second=0; second<ARBITRARILY_LARGE_NUMBER && in_place < ALPHA_SIZE; second++) {
    if(second != 0) {
      for(worker=0; worker<NUMBER_OF_WORKERS; worker++) {
        if(timeline[second][worker] == 0 && timeline[second-1][worker] != 0) {
          letter = timeline[second-1][worker] - FIRST_LETTER;
          already_finished[letter] = true;
          for(i=0; i<ALPHA_SIZE; i++) {
            if(adj_matrix[letter][i] != 0) {
              adj_matrix[letter][i] = 0;
              dependencies[i]--;
            }
          }
        }
      }
    }
    for(worker=0; worker<NUMBER_OF_WORKERS; worker++) {
      if(timeline[second][worker] == 0) {
        for(i=0; i<ALPHA_SIZE; i++) {
          if(!already_allocated[i] && dependencies[i] == 0) {
            already_allocated[i] = true;
            in_place++;
            for(j = second; j < second + BASE_TIME + i + 1; j++) {
              timeline[j][worker] = i + FIRST_LETTER;
              if(j > max) {
                max = j;
              }
            }
            break;
          }
        }
      }
    }
  }
  printf("%d\n", max+1);
  return 0;
}

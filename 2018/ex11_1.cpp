#include <cstdlib>
#include <cstdio>

#define SERIAL_NUMBER 9435
#define SQUARE_SIZE 300

int points[SQUARE_SIZE][SQUARE_SIZE];
int best_x, best_y, i, j, k, l, power, rack_id, tmp_power, max_power;

int main(int argc, char* argv[]) {
  for(i=0; i<SQUARE_SIZE; i++) {
    for(j=0; j<SQUARE_SIZE; j++) {
      rack_id = i + 11;
      power = rack_id * (j+1);
      power+= SERIAL_NUMBER;
      power *= rack_id;
      power /= 100;
      power %= 10;
      power -= 5;
      points[i][j] = power;
    }
  }
  for(i=0; i<SQUARE_SIZE-2; i++) {
    for(j=0; j<SQUARE_SIZE-2; j++) {
      tmp_power = 0;
      for(k=i; k<i+3; k++) {
        for(l=j; l<j+3; l++) {
          tmp_power += points[k][l];
        }
      }
      if(tmp_power > max_power) {
        max_power = tmp_power;
        best_x = i;
        best_y = j;
      }
    }
  }
  printf("%d,%d\n", best_x+1, best_y+1);
  return 0;
}

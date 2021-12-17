#include <cstdlib>
#include <cstdio>
#include <cstring>
#include <climits>

#define INPUT_SIZE 50
#define MAX 355
#define TARGET 10000

struct point {
  int x;
  int y;
};

struct point points[INPUT_SIZE] = {};
int map[MAX][MAX] = {};
int i, j, k, cur_distance, count=0;

int main(int argc, char* argv[]) {
  for(i=0; i<INPUT_SIZE; i++) {
    scanf("%d, %d\n", &points[i].x, &points[i].y);
  }
  for(i=0; i<MAX; i++) {
    for(j=0; j<MAX; j++) {
      for(k=0; k<INPUT_SIZE; k++) {
        map[i][j] += abs(i - points[k].x) + abs(j - points[k].y);
      }
      if(map[i][j] < TARGET) {
        count++;
      }
    }
  }
  printf("%d\n", count);
  return 0;
}

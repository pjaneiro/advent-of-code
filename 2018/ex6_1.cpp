#include <cstdlib>
#include <cstdio>
#include <cstring>
#include <climits>

#define INPUT_SIZE 50
#define MAX 355

struct point {
  int x;
  int y;
  int area;
  bool finite;
};

struct point points[INPUT_SIZE] = {};
int map[MAX][MAX] = {};
int i, j, k, best_distance, cur_distance, best_area;

int main(int argc, char* argv[]) {
  for(i=0; i<INPUT_SIZE; i++) {
    scanf("%d, %d\n", &points[i].x, &points[i].y);
    points[i].finite = true;
    points[i].area = 0;
  }
  for(i = 0; i<MAX; i++) {
    for(j = 0; j<MAX; j++) {
      best_distance = INT_MAX;
      for(k = 0; k<INPUT_SIZE; k++) {
        cur_distance = abs( points[k].x - i ) + abs( points[k].y - j );
        if(cur_distance < best_distance) {
          best_distance = cur_distance;
          map[i][j] = k;
        } else if(cur_distance == best_distance) {
          map[i][j] = -1;
        }
      }
    }
  }
  for(j=0; j<MAX; j++) {
    if(map[0][j] != -1) {
      points[map[0][j]].finite = false;
    }
    if(map[MAX-1][j] != -1) {
      points[map[MAX-1][j]].finite = false;
    }
  }
  for(i=0; i<MAX; i++) {
    if(map[i][0] != -1) {
      points[map[i][0]].finite = false;
    }
    if(map[i][MAX-1] != -1) {
      points[map[i][MAX-1]].finite = false;
    }
  }
  for(i = 0; i<MAX; i++) {
    for(j = 0; j<MAX; j++) {
      if(map[i][j] != -1) {
        points[map[i][j]].area++;
      }
    }
  }
  for(i=0; i<INPUT_SIZE; i++) {
    if(points[i].finite && points[i].area > best_area) {
      best_area = points[i].area;
    }
  }
  printf("%d\n", best_area);
  return 0;
}

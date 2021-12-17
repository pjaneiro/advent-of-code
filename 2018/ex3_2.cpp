#include <cstdlib>
#include <cstdio>

#define FABRIC_SIZE 1001
#define INPUT_SIZE 1350

int fabric[FABRIC_SIZE][FABRIC_SIZE] = {};
int overlaps[INPUT_SIZE] = {};
int id, x, y, w, h, i, j;

int main(int argc, const char* argv[]) {
  while(scanf("#%d @ %d,%d: %dx%d\n", &id, &x, &y, &w, &h) == 5) {
    printf("%d @ %d:%d:%d:%d\n", id, x, y, w, h);
    for(i = x; i < x+w; i++) {
      for(j = y; j < y+h; j++) {
        if(fabric[i][j] != 0) {
          overlaps[fabric[i][j]] = 1;
          overlaps[id] = 1;
        }
        fabric[i][j] = id;
      }
    }
  }
  for(i = 1; i < INPUT_SIZE ; i++) {
    if(overlaps[i] == 0) {
      printf("%d\n", i);
    }
  }
  return 0;
}

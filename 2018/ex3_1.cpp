#include <cstdlib>
#include <cstdio>

#define FABRIC_SIZE 1001

int fabric[FABRIC_SIZE][FABRIC_SIZE] = {};
int id, x, y, w, h, i, j;
int count = 0;

int main(int argc, const char* argv[]) {
  while(scanf("#%d @ %d,%d: %dx%d\n", &id, &x, &y, &w, &h) == 5) {
    printf("%d @ %d:%d:%d:%d\n", id, x, y, w, h);
    for(i = x; i < x+w; i++) {
      for(j = y; j < y+h; j++) {
        fabric[i][j]++;
        if(fabric[i][j] == 2) {
          count++;
        }
      }
    }
  }
  printf("%d\n", count);
  return 0;
}

#include <cstdlib>
#include <cstdio>

#define INPUT_SIZE 15908

int input[INPUT_SIZE];
int idx, count=0;

int read_node(int);

int main(int argc, char* argv[]) {
  for(idx=0; idx<INPUT_SIZE; idx++) {
    scanf("%d", input+idx);
  }
  read_node(0);
  printf("%d\n", count);

  return 0;
}

int read_node(int pos) {
  int num_children = *(input+pos);
  int num_metadata_entries = *(input+pos+1);
  int next_node = pos+2;
  for(int i=0; i<num_children; i++) {
    next_node = read_node(next_node);
  }
  for(int i=0; i<num_metadata_entries; i++) {
    count+=*(input+next_node);
    next_node++;
  }
  return next_node;
}

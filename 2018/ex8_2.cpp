#include <cstdlib>
#include <cstdio>

#define INPUT_SIZE 15908

int input[INPUT_SIZE];
int idx, count=0;

int read_node(int, bool);

int main(int argc, char* argv[]) {
  for(idx=0; idx<INPUT_SIZE; idx++) {
    scanf("%d", input+idx);
  }
  read_node(0, true);
  printf("%d\n", count);

  return 0;
}

int read_node(int pos, bool mode) {
  int num_children = *(input+pos);
  int* children = (int*)calloc(num_children, sizeof(int));
  int num_metadata_entries = *(input+pos+1);
  int next_node = pos+2;
  int next_metadata_entry;
  for(int i=0; i<num_children; i++) {
    *(children + i) = next_node;
    next_node = read_node(next_node, false);
  }
  if(num_children == 0) {
    for(int i=0; i<num_metadata_entries; i++) {
      if(mode) {
        count+=*(input+next_node);
      }
      next_node++;
    }
  } else {
    for(int i=0; i<num_metadata_entries; i++) {
      next_metadata_entry = *(input+next_node);
      if(next_metadata_entry <=num_children) {
        read_node(*(children+next_metadata_entry-1), mode);
      }
      next_node++;
    }
  }
  free(children);
  return next_node;
}

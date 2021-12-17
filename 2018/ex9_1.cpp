#include <cstdlib>
#include <cstdio>

#define NUM_PLAYERS 465
#define NUM_MARBLES 7194000

typedef struct node{
  struct node* next;
  struct node* prev;
  int value;
} node;

int circle[NUM_MARBLES] = {0};
unsigned int scores[NUM_PLAYERS] = {};
int i, j, cur_player = 1;
unsigned int max_points;

node* insert_node(node*, int);
node* insert_23_node(node*, int);

int main(int argc, char* argv[]){
  node* current_node = new node;
  node* zero_node = current_node;
  node* aux_node = current_node;
  current_node->next = current_node;
  current_node->prev = current_node;
  current_node->value = 0;

  for(i=1; i<=NUM_MARBLES; i++) {
    if(i % 23 != 0) {
      current_node = insert_node(current_node, i);
    } else {
      current_node = insert_23_node(current_node, i);
      scores[cur_player-1] += i;
    }
    cur_player = 1 + (cur_player % NUM_PLAYERS);
  }

  for(i = 0; i<NUM_PLAYERS; i++) {
    if(scores[i] > max_points) {
      max_points = scores[i];
    }
  }
  printf("%u\n", max_points);

  return 0;
}

node* insert_node(node* current_node, int value) {
  node* new_node = new node;
  new_node->value = value;

  node* left_node = current_node->next;
  node* right_node = current_node->next->next;

  new_node->next = right_node;
  new_node->prev = left_node;
  left_node->next = new_node;
  right_node->prev = new_node;

  return new_node;
}

node* insert_23_node(node* current_node, int value) {
  node* old_node = current_node->prev->prev->prev->prev->prev->prev->prev;
  node* left_node = old_node->prev;
  node* right_node = old_node->next;

  left_node->next = right_node;
  right_node->prev = left_node;

  scores[cur_player-1] += old_node->value;
  delete old_node;

  return right_node;
}

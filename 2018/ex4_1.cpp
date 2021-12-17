#include <cstdlib>
#include <cstdio>
#include <cstring>

#define INPUT_SIZE 1134
#define ARBITRARILY_LARGE_NUMBER 3100

enum op_types {wake_up, fall_asleep, begin_shift};

struct event {
  int year, month, day;
  int hour, minute;
  enum op_types operation;
  int guard_id;
};

struct event input[INPUT_SIZE] = {};
int i, j, guard_id, winning_guard, max_minutes, best_minute, best_minute_details;
char action[256] = {};
int guard_sleep_total_time[ARBITRARILY_LARGE_NUMBER] = {};
int guard_sleep_minutes[ARBITRARILY_LARGE_NUMBER][60] = {};

int compare(const void*, const void*);
const char* string_from_op_type(int);

int main(int argc, char* argv[]) {
  i = 0;
  while(scanf("[%d-%d-%d %d:%d] %[^\n]s", &input[i].year, &input[i].month, &input[i].day, &input[i].hour, &input[i].minute, &action) == 6) {
    getchar();
    if(strcmp(action, "wakes up") == 0) {
      input[i].operation = wake_up;
    } else if(strcmp(action, "falls asleep") == 0) {
      input[i].operation = fall_asleep;
    } else {
      input[i].operation = begin_shift;
      sscanf(action, "Guard #%d begins shift", &guard_id);
      input[i].guard_id = guard_id;
    }
    i++;
  }

  qsort(input, INPUT_SIZE, sizeof(struct event), compare);
  for(i = 1; i<INPUT_SIZE; i++) {
    if(input[i].guard_id == 0) {
      input[i].guard_id = input[i-1].guard_id;
    }

    if(input[i].operation == wake_up) {
      guard_sleep_total_time[input[i].guard_id] += (input[i].minute - input[i-1].minute);
      for(j = input[i-1].minute ; j < input[i].minute ; j++) {
        guard_sleep_minutes[input[i].guard_id][j]++;
      }
    }
  }

  for(i = 0; i<ARBITRARILY_LARGE_NUMBER; i++) {
    if(guard_sleep_total_time[i] > max_minutes) {
      max_minutes = guard_sleep_total_time[i];
      winning_guard = i;
    }
  }

  best_minute_details = 0;
  for(i = 0; i<60; i++) {
    if(guard_sleep_minutes[winning_guard][i] > best_minute_details) {
      best_minute_details = guard_sleep_minutes[winning_guard][i];
      best_minute = i;
    }
  }

  printf("%d\n", winning_guard * best_minute);

  return 0;
}

int compare(const void* a, const void* b) {
  struct event e_a = *(struct event*)a;
  struct event e_b = *(struct event*)b;

  if(e_a.year == e_b.year) {
    if(e_a.month == e_b.month) {
      if(e_a.day == e_b.day) {
        if(e_a.hour == e_b.hour) {
          return e_a.minute - e_b.minute;
        } else {
          return e_a.hour - e_b.hour;
        }
      } else {
        return e_a.day - e_b.day;
      }
    } else {
      return e_a.month - e_b.month;
    }
  } else {
    return e_a.year - e_b.year;
  }
}

const char* string_from_op_type(int op_type) {
  switch(op_type) {
    case 0:
      return "wake up";
    case 1:
      return "fall asleep";
    case 2:
      return "begin shift";
    default:
      return "break this";
  }
}

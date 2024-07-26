#include "all.h"
#include <vector>

char get_symbol(int value) {
  switch (value) {
  case 0:
    return '-';
  case 1:
    return 'X';
  case 2:
    return 'O';
  default:
    return ' ';
  }
}

void print_table(std::vector<std::vector<int>> table) {
  for (int i = 0; i < table.size(); i++) {
    auto row = table.at(i);
    for (int j = 0; j < row.size(); j++) {
      std::cout << get_symbol(row.at(j));
    }
    std::cout << std::endl;
  }
}

int get_winner(std::vector<std::vector<int>> table) {
  for (int i = 0; i < table.at(0).size(); i++) {
    auto row = table.at(i);
    if (row.at(0) != 0 && row.at(0) == row.at(1) && row.at(1) == row.at(2)) {
      return row.at(0);
    }
  }

  for (int i = 0; i < table.size(); i++) {
    if (table.at(0).at(i) != 0 && table.at(0).at(i) == table.at(1).at(i) &&
        table.at(1).at(i) == table.at(2).at(i)) {
      return table.at(0).at(i);
    }
  }

  if (table.at(0).at(0) != 0 && table.at(0).at(0) == table.at(1).at(1) &&
      table.at(1).at(1) == table.at(2).at(2)) {
    return table.at(0).at(0);
  }

  if (table.at(0).at(2) != 0 && table.at(0).at(2) == table.at(1).at(1) &&
      table.at(1).at(1) == table.at(0).at(2)) {
    return table.at(0).at(2);
  }

  return 0;
}

int main() {
  std::vector<std::vector<int>> table = {
      {0, 0, 0},
      {0, 0, 0},
      {0, 0, 0},
  };

  print_table(table);
  std::cout << get_winner(table);
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node389_0 {
    int value;
    struct Node389_0 *next;
};

int length389_0(const struct Node389_0 *head) {
    int count = 0;
    const struct Node389_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
enum Kind389_1 { KIND_EMPTY389_1 = 0, KIND_FULL389_1 = 25 };

union Cell389_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind389_1(union Cell389_1 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL389_1 : KIND_EMPTY389_1;
}
int gcd389_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale389_2(int value) {
    return gcd389_2(value, 9) * 32;
}
enum Kind389_3 { KIND_EMPTY389_3 = 0, KIND_FULL389_3 = 39 };

union Cell389_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind389_3(union Cell389_3 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL389_3 : KIND_EMPTY389_3;
}
#define MASK389_4(v) (((v) << 4) | ((v) & 0x0F))

int classify389_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK389_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind389_5 { KIND_EMPTY389_5 = 0, KIND_FULL389_5 = 53 };

union Cell389_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind389_5(union Cell389_5 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL389_5 : KIND_EMPTY389_5;
}
int scan389_6(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 60;
}
enum Kind389_7 { KIND_EMPTY389_7 = 0, KIND_FULL389_7 = 67 };

union Cell389_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind389_7(union Cell389_7 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL389_7 : KIND_EMPTY389_7;
}

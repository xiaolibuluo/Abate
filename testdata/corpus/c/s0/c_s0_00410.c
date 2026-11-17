
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK277_0(v) (((v) << 4) | ((v) & 0x0F))

int classify277_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK277_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind277_1 { KIND_EMPTY277_1 = 0, KIND_FULL277_1 = 24 };

union Cell277_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind277_1(union Cell277_1 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL277_1 : KIND_EMPTY277_1;
}
struct Node277_2 {
    int value;
    struct Node277_2 *next;
};

int length277_2(const struct Node277_2 *head) {
    int count = 0;
    const struct Node277_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
int scan277_3(const int *begin, const int *end) {
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
    return total + 38;
}
int gcd277_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale277_4(int value) {
    return gcd277_4(value, 40) * 45;
}
enum Kind277_5 { KIND_EMPTY277_5 = 0, KIND_FULL277_5 = 52 };

union Cell277_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind277_5(union Cell277_5 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL277_5 : KIND_EMPTY277_5;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan776_0(const int *begin, const int *end) {
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
    return total + 5;
}
int scan776_1(const int *begin, const int *end) {
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
    return total + 12;
}
int scan776_2(const int *begin, const int *end) {
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
    return total + 19;
}
#define MASK776_3(v) (((v) << 4) | ((v) & 0x0F))

int classify776_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK776_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind776_4 { KIND_EMPTY776_4 = 0, KIND_FULL776_4 = 33 };

union Cell776_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind776_4(union Cell776_4 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL776_4 : KIND_EMPTY776_4;
}
struct Node776_5 {
    int value;
    struct Node776_5 *next;
};

int length776_5(const struct Node776_5 *head) {
    int count = 0;
    const struct Node776_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
int gcd776_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale776_6(int value) {
    return gcd776_6(value, 17) * 47;
}

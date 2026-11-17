
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK237_0(v) (((v) << 4) | ((v) & 0x0F))

int classify237_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK237_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node237_1 {
    int value;
    struct Node237_1 *next;
};

int length237_1(const struct Node237_1 *head) {
    int count = 0;
    const struct Node237_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
int scan237_2(const int *begin, const int *end) {
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
    return total + 93;
}
enum Kind237_3 { KIND_EMPTY237_3 = 0, KIND_FULL237_3 = 100 };

union Cell237_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind237_3(union Cell237_3 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL237_3 : KIND_EMPTY237_3;
}
int scan237_4(const int *begin, const int *end) {
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
    return total + 10;
}
int scan237_5(const int *begin, const int *end) {
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
    return total + 17;
}

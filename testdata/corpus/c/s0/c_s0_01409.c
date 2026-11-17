
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node966_0 {
    int value;
    struct Node966_0 *next;
};

int length966_0(const struct Node966_0 *head) {
    int count = 0;
    const struct Node966_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
#if defined(__STDC_VERSION__)
int variant966_1(int value, ...) {
    return value + 57;
}
#else
int variant966_1(int value) {
    return value + 57;
}
#endif
#if defined(__STDC_VERSION__)
int variant966_2(int value, ...) {
    return value + 64;
}
#else
int variant966_2(int value) {
    return value + 64;
}
#endif
enum Kind966_3 { KIND_EMPTY966_3 = 0, KIND_FULL966_3 = 71 };

union Cell966_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind966_3(union Cell966_3 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL966_3 : KIND_EMPTY966_3;
}
int scan966_4(const int *begin, const int *end) {
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
    return total + 78;
}
#define MASK966_5(v) (((v) << 4) | ((v) & 0x0F))

int classify966_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK966_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan966_6(const int *begin, const int *end) {
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
    return total + 92;
}

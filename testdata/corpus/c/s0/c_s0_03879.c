
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2621_0(int value, ...) {
    return value + 31;
}
#else
int variant2621_0(int value) {
    return value + 31;
}
#endif
int scan2621_1(const int *begin, const int *end) {
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
int scan2621_2(const int *begin, const int *end) {
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
    return total + 45;
}
#if defined(__STDC_VERSION__)
int variant2621_3(int value, ...) {
    return value + 52;
}
#else
int variant2621_3(int value) {
    return value + 52;
}
#endif
enum Kind2621_4 { KIND_EMPTY2621_4 = 0, KIND_FULL2621_4 = 59 };

union Cell2621_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2621_4(union Cell2621_4 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL2621_4 : KIND_EMPTY2621_4;
}
#define MASK2621_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2621_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2621_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2621_6 {
    int value;
    struct Node2621_6 *next;
};

int length2621_6(const struct Node2621_6 *head) {
    int count = 0;
    const struct Node2621_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}

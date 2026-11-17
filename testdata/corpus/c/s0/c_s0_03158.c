
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2140_0(const int *begin, const int *end) {
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
    return total + 83;
}
#if defined(__STDC_VERSION__)
int variant2140_1(int value, ...) {
    return value + 90;
}
#else
int variant2140_1(int value) {
    return value + 90;
}
#endif
#define MASK2140_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2140_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2140_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2140_3 {
    int value;
    struct Node2140_3 *next;
};

int length2140_3(const struct Node2140_3 *head) {
    int count = 0;
    const struct Node2140_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
enum Kind2140_4 { KIND_EMPTY2140_4 = 0, KIND_FULL2140_4 = 14 };

union Cell2140_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2140_4(union Cell2140_4 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL2140_4 : KIND_EMPTY2140_4;
}
#if defined(__STDC_VERSION__)
int variant2140_5(int value, ...) {
    return value + 21;
}
#else
int variant2140_5(int value) {
    return value + 21;
}
#endif
#define MASK2140_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2140_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2140_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2140_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2140_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2140_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2140_8(int value, ...) {
    return value + 42;
}
#else
int variant2140_8(int value) {
    return value + 42;
}
#endif

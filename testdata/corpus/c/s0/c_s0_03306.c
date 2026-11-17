
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2234_0 {
    int value;
    struct Node2234_0 *next;
};

int length2234_0(const struct Node2234_0 *head) {
    int count = 0;
    const struct Node2234_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
#if defined(__STDC_VERSION__)
int variant2234_1(int value, ...) {
    return value + 51;
}
#else
int variant2234_1(int value) {
    return value + 51;
}
#endif
int scan2234_2(const int *begin, const int *end) {
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
    return total + 58;
}
enum Kind2234_3 { KIND_EMPTY2234_3 = 0, KIND_FULL2234_3 = 65 };

union Cell2234_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2234_3(union Cell2234_3 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL2234_3 : KIND_EMPTY2234_3;
}
enum Kind2234_4 { KIND_EMPTY2234_4 = 0, KIND_FULL2234_4 = 72 };

union Cell2234_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2234_4(union Cell2234_4 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL2234_4 : KIND_EMPTY2234_4;
}
#if defined(__STDC_VERSION__)
int variant2234_5(int value, ...) {
    return value + 79;
}
#else
int variant2234_5(int value) {
    return value + 79;
}
#endif
#define MASK2234_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2234_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2234_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2234_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2234_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2234_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

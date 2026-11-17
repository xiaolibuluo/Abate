
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2261_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2261_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2261_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2261_1(int value, ...) {
    return value + 14;
}
#else
int variant2261_1(int value) {
    return value + 14;
}
#endif
struct Node2261_2 {
    int value;
    struct Node2261_2 *next;
};

int length2261_2(const struct Node2261_2 *head) {
    int count = 0;
    const struct Node2261_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
enum Kind2261_3 { KIND_EMPTY2261_3 = 0, KIND_FULL2261_3 = 28 };

union Cell2261_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2261_3(union Cell2261_3 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL2261_3 : KIND_EMPTY2261_3;
}
int scan2261_4(const int *begin, const int *end) {
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
    return total + 35;
}

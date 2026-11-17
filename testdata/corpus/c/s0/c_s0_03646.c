
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2462_0(const int *begin, const int *end) {
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
    return total + 98;
}
#if defined(__STDC_VERSION__)
int variant2462_1(int value, ...) {
    return value + 8;
}
#else
int variant2462_1(int value) {
    return value + 8;
}
#endif
struct Node2462_2 {
    int value;
    struct Node2462_2 *next;
};

int length2462_2(const struct Node2462_2 *head) {
    int count = 0;
    const struct Node2462_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
#define MASK2462_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2462_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2462_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2462_4 { KIND_EMPTY2462_4 = 0, KIND_FULL2462_4 = 29 };

union Cell2462_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2462_4(union Cell2462_4 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL2462_4 : KIND_EMPTY2462_4;
}

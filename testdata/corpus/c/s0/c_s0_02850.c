
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1933_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1933_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1933_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1933_1(int value, ...) {
    return value + 18;
}
#else
int variant1933_1(int value) {
    return value + 18;
}
#endif
#define MASK1933_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1933_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1933_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1933_3 {
    int value;
    struct Node1933_3 *next;
};

int length1933_3(const struct Node1933_3 *head) {
    int count = 0;
    const struct Node1933_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
enum Kind1933_4 { KIND_EMPTY1933_4 = 0, KIND_FULL1933_4 = 39 };

union Cell1933_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1933_4(union Cell1933_4 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1933_4 : KIND_EMPTY1933_4;
}
int scan1933_5(const int *begin, const int *end) {
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
    return total + 46;
}

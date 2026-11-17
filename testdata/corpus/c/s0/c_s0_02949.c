
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1994_0(int value, ...) {
    return value + 28;
}
#else
int variant1994_0(int value) {
    return value + 28;
}
#endif
int scan1994_1(const int *begin, const int *end) {
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
enum Kind1994_2 { KIND_EMPTY1994_2 = 0, KIND_FULL1994_2 = 42 };

union Cell1994_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1994_2(union Cell1994_2 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL1994_2 : KIND_EMPTY1994_2;
}
int scan1994_3(const int *begin, const int *end) {
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
    return total + 49;
}
#define MASK1994_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1994_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1994_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1994_5(int value, ...) {
    return value + 63;
}
#else
int variant1994_5(int value) {
    return value + 63;
}
#endif
#define MASK1994_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1994_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1994_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1994_7 {
    int value;
    struct Node1994_7 *next;
};

int length1994_7(const struct Node1994_7 *head) {
    int count = 0;
    const struct Node1994_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 77;
}

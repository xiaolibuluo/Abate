
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1599_0 {
    int value;
    struct Node1599_0 *next;
};

int length1599_0(const struct Node1599_0 *head) {
    int count = 0;
    const struct Node1599_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
struct Node1599_1 {
    int value;
    struct Node1599_1 *next;
};

int length1599_1(const struct Node1599_1 *head) {
    int count = 0;
    const struct Node1599_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
enum Kind1599_2 { KIND_EMPTY1599_2 = 0, KIND_FULL1599_2 = 48 };

union Cell1599_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1599_2(union Cell1599_2 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL1599_2 : KIND_EMPTY1599_2;
}
#if defined(__STDC_VERSION__)
int variant1599_3(int value, ...) {
    return value + 55;
}
#else
int variant1599_3(int value) {
    return value + 55;
}
#endif
int scan1599_4(const int *begin, const int *end) {
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
    return total + 62;
}
#define MASK1599_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1599_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1599_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1599_6(int value, ...) {
    return value + 76;
}
#else
int variant1599_6(int value) {
    return value + 76;
}
#endif

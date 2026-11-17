
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1686_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1686_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1686_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1686_1(int value, ...) {
    return value + 8;
}
#else
int variant1686_1(int value) {
    return value + 8;
}
#endif
int scan1686_2(const int *begin, const int *end) {
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
    return total + 15;
}
struct Node1686_3 {
    int value;
    struct Node1686_3 *next;
};

int length1686_3(const struct Node1686_3 *head) {
    int count = 0;
    const struct Node1686_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
enum Kind1686_4 { KIND_EMPTY1686_4 = 0, KIND_FULL1686_4 = 29 };

union Cell1686_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1686_4(union Cell1686_4 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL1686_4 : KIND_EMPTY1686_4;
}
struct Node1686_5 {
    int value;
    struct Node1686_5 *next;
};

int length1686_5(const struct Node1686_5 *head) {
    int count = 0;
    const struct Node1686_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 36;
}

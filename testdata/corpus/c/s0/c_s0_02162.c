
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1453_0(const int *begin, const int *end) {
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
    return total + 76;
}
#if defined(__STDC_VERSION__)
int variant1453_1(int value, ...) {
    return value + 83;
}
#else
int variant1453_1(int value) {
    return value + 83;
}
#endif
#if defined(__STDC_VERSION__)
int variant1453_2(int value, ...) {
    return value + 90;
}
#else
int variant1453_2(int value) {
    return value + 90;
}
#endif
struct Node1453_3 {
    int value;
    struct Node1453_3 *next;
};

int length1453_3(const struct Node1453_3 *head) {
    int count = 0;
    const struct Node1453_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
#define MASK1453_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1453_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1453_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1453_5 { KIND_EMPTY1453_5 = 0, KIND_FULL1453_5 = 14 };

union Cell1453_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1453_5(union Cell1453_5 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL1453_5 : KIND_EMPTY1453_5;
}
enum Kind1453_6 { KIND_EMPTY1453_6 = 0, KIND_FULL1453_6 = 21 };

union Cell1453_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1453_6(union Cell1453_6 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1453_6 : KIND_EMPTY1453_6;
}

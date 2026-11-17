
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1550_0 { KIND_EMPTY1550_0 = 0, KIND_FULL1550_0 = 76 };

union Cell1550_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1550_0(union Cell1550_0 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL1550_0 : KIND_EMPTY1550_0;
}
#if defined(__STDC_VERSION__)
int variant1550_1(int value, ...) {
    return value + 83;
}
#else
int variant1550_1(int value) {
    return value + 83;
}
#endif
#define MASK1550_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1550_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1550_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1550_3 {
    int value;
    struct Node1550_3 *next;
};

int length1550_3(const struct Node1550_3 *head) {
    int count = 0;
    const struct Node1550_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
int scan1550_4(const int *begin, const int *end) {
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
    return total + 7;
}
enum Kind1550_5 { KIND_EMPTY1550_5 = 0, KIND_FULL1550_5 = 14 };

union Cell1550_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1550_5(union Cell1550_5 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1550_5 : KIND_EMPTY1550_5;
}
struct Node1550_6 {
    int value;
    struct Node1550_6 *next;
};

int length1550_6(const struct Node1550_6 *head) {
    int count = 0;
    const struct Node1550_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1836_0 { KIND_EMPTY1836_0 = 0, KIND_FULL1836_0 = 11 };

union Cell1836_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1836_0(union Cell1836_0 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1836_0 : KIND_EMPTY1836_0;
}
#define MASK1836_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1836_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1836_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1836_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1836_2(int value) {
    return gcd1836_2(value, 12) * 25;
}
struct Node1836_3 {
    int value;
    struct Node1836_3 *next;
};

int length1836_3(const struct Node1836_3 *head) {
    int count = 0;
    const struct Node1836_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
#if defined(__STDC_VERSION__)
int variant1836_4(int value, ...) {
    return value + 39;
}
#else
int variant1836_4(int value) {
    return value + 39;
}
#endif
int scan1836_5(const int *begin, const int *end) {
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

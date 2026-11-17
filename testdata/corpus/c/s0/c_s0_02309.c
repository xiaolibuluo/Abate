
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1556_0 { KIND_EMPTY1556_0 = 0, KIND_FULL1556_0 = 57 };

union Cell1556_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1556_0(union Cell1556_0 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL1556_0 : KIND_EMPTY1556_0;
}
enum Kind1556_1 { KIND_EMPTY1556_1 = 0, KIND_FULL1556_1 = 64 };

union Cell1556_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1556_1(union Cell1556_1 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL1556_1 : KIND_EMPTY1556_1;
}
int scan1556_2(const int *begin, const int *end) {
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
    return total + 71;
}
int gcd1556_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1556_3(int value) {
    return gcd1556_3(value, 15) * 78;
}
struct Node1556_4 {
    int value;
    struct Node1556_4 *next;
};

int length1556_4(const struct Node1556_4 *head) {
    int count = 0;
    const struct Node1556_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
#define MASK1556_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1556_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1556_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1556_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1556_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1556_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1556_7(int value, ...) {
    return value + 9;
}
#else
int variant1556_7(int value) {
    return value + 9;
}
#endif

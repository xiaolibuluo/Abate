
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1896_0(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant1896_1(int value, ...) {
    return value + 22;
}
#else
int variant1896_1(int value) {
    return value + 22;
}
#endif
struct Node1896_2 {
    int value;
    struct Node1896_2 *next;
};

int length1896_2(const struct Node1896_2 *head) {
    int count = 0;
    const struct Node1896_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 29;
}
int gcd1896_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1896_3(int value) {
    return gcd1896_3(value, 16) * 36;
}
int gcd1896_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1896_4(int value) {
    return gcd1896_4(value, 27) * 43;
}
enum Kind1896_5 { KIND_EMPTY1896_5 = 0, KIND_FULL1896_5 = 50 };

union Cell1896_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1896_5(union Cell1896_5 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL1896_5 : KIND_EMPTY1896_5;
}
#define MASK1896_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1896_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1896_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1896_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1896_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1896_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

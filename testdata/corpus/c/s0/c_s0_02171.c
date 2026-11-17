
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1458_0(const int *begin, const int *end) {
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
    return total + 44;
}
#define MASK1458_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1458_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1458_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1458_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1458_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1458_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1458_3 {
    int value;
    struct Node1458_3 *next;
};

int length1458_3(const struct Node1458_3 *head) {
    int count = 0;
    const struct Node1458_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
int gcd1458_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1458_4(int value) {
    return gcd1458_4(value, 11) * 72;
}
enum Kind1458_5 { KIND_EMPTY1458_5 = 0, KIND_FULL1458_5 = 79 };

union Cell1458_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1458_5(union Cell1458_5 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1458_5 : KIND_EMPTY1458_5;
}
#if defined(__STDC_VERSION__)
int variant1458_6(int value, ...) {
    return value + 86;
}
#else
int variant1458_6(int value) {
    return value + 86;
}
#endif
int gcd1458_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1458_7(int value) {
    return gcd1458_7(value, 44) * 93;
}

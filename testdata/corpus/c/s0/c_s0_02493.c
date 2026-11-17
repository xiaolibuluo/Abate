
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1695_0 { KIND_EMPTY1695_0 = 0, KIND_FULL1695_0 = 21 };

union Cell1695_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1695_0(union Cell1695_0 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL1695_0 : KIND_EMPTY1695_0;
}
int gcd1695_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1695_1(int value) {
    return gcd1695_1(value, 51) * 28;
}
int scan1695_2(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant1695_3(int value, ...) {
    return value + 42;
}
#else
int variant1695_3(int value) {
    return value + 42;
}
#endif
#define MASK1695_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1695_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1695_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1695_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1695_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1695_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1695_6 {
    int value;
    struct Node1695_6 *next;
};

int length1695_6(const struct Node1695_6 *head) {
    int count = 0;
    const struct Node1695_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}

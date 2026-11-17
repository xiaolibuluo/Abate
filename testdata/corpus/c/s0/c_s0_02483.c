
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1688_0 {
    int value;
    struct Node1688_0 *next;
};

int length1688_0(const struct Node1688_0 *head) {
    int count = 0;
    const struct Node1688_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
enum Kind1688_1 { KIND_EMPTY1688_1 = 0, KIND_FULL1688_1 = 34 };

union Cell1688_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1688_1(union Cell1688_1 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL1688_1 : KIND_EMPTY1688_1;
}
#define MASK1688_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1688_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1688_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1688_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1688_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1688_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1688_4(int value, ...) {
    return value + 55;
}
#else
int variant1688_4(int value) {
    return value + 55;
}
#endif
#if defined(__STDC_VERSION__)
int variant1688_5(int value, ...) {
    return value + 62;
}
#else
int variant1688_5(int value) {
    return value + 62;
}
#endif
int gcd1688_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1688_6(int value) {
    return gcd1688_6(value, 57) * 69;
}

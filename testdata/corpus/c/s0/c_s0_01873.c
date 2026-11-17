
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1256_0 {
    int value;
    struct Node1256_0 *next;
};

int length1256_0(const struct Node1256_0 *head) {
    int count = 0;
    const struct Node1256_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
enum Kind1256_1 { KIND_EMPTY1256_1 = 0, KIND_FULL1256_1 = 44 };

union Cell1256_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1256_1(union Cell1256_1 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL1256_1 : KIND_EMPTY1256_1;
}
int gcd1256_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1256_2(int value) {
    return gcd1256_2(value, 39) * 51;
}
#define MASK1256_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1256_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1256_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1256_4 { KIND_EMPTY1256_4 = 0, KIND_FULL1256_4 = 65 };

union Cell1256_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1256_4(union Cell1256_4 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL1256_4 : KIND_EMPTY1256_4;
}
#if defined(__STDC_VERSION__)
int variant1256_5(int value, ...) {
    return value + 72;
}
#else
int variant1256_5(int value) {
    return value + 72;
}
#endif
#define MASK1256_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1256_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1256_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1541_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1541_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1541_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1541_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1541_1(int value) {
    return gcd1541_1(value, 10) * 63;
}
#if defined(__STDC_VERSION__)
int variant1541_2(int value, ...) {
    return value + 70;
}
#else
int variant1541_2(int value) {
    return value + 70;
}
#endif
int gcd1541_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1541_3(int value) {
    return gcd1541_3(value, 32) * 77;
}
enum Kind1541_4 { KIND_EMPTY1541_4 = 0, KIND_FULL1541_4 = 84 };

union Cell1541_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1541_4(union Cell1541_4 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL1541_4 : KIND_EMPTY1541_4;
}
struct Node1541_5 {
    int value;
    struct Node1541_5 *next;
};

int length1541_5(const struct Node1541_5 *head) {
    int count = 0;
    const struct Node1541_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1010_0 {
    int value;
    struct Node1010_0 *next;
};

int length1010_0(const struct Node1010_0 *head) {
    int count = 0;
    const struct Node1010_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
int gcd1010_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1010_1(int value) {
    return gcd1010_1(value, 14) * 47;
}
#define MASK1010_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1010_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1010_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1010_3 { KIND_EMPTY1010_3 = 0, KIND_FULL1010_3 = 61 };

union Cell1010_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1010_3(union Cell1010_3 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL1010_3 : KIND_EMPTY1010_3;
}
int gcd1010_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1010_4(int value) {
    return gcd1010_4(value, 47) * 68;
}
#if defined(__STDC_VERSION__)
int variant1010_5(int value, ...) {
    return value + 75;
}
#else
int variant1010_5(int value) {
    return value + 75;
}
#endif
int gcd1010_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1010_6(int value) {
    return gcd1010_6(value, 69) * 82;
}

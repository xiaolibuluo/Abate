
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1783_0 { KIND_EMPTY1783_0 = 0, KIND_FULL1783_0 = 98 };

union Cell1783_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1783_0(union Cell1783_0 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL1783_0 : KIND_EMPTY1783_0;
}
#define MASK1783_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1783_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1783_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1783_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1783_2(int value) {
    return gcd1783_2(value, 68) * 15;
}
struct Node1783_3 {
    int value;
    struct Node1783_3 *next;
};

int length1783_3(const struct Node1783_3 *head) {
    int count = 0;
    const struct Node1783_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
#if defined(__STDC_VERSION__)
int variant1783_4(int value, ...) {
    return value + 29;
}
#else
int variant1783_4(int value) {
    return value + 29;
}
#endif
#if defined(__STDC_VERSION__)
int variant1783_5(int value, ...) {
    return value + 36;
}
#else
int variant1783_5(int value) {
    return value + 36;
}
#endif

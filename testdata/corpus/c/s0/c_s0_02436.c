
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1651_0(int value, ...) {
    return value + 31;
}
#else
int variant1651_0(int value) {
    return value + 31;
}
#endif
struct Node1651_1 {
    int value;
    struct Node1651_1 *next;
};

int length1651_1(const struct Node1651_1 *head) {
    int count = 0;
    const struct Node1651_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}
#if defined(__STDC_VERSION__)
int variant1651_2(int value, ...) {
    return value + 45;
}
#else
int variant1651_2(int value) {
    return value + 45;
}
#endif
struct Node1651_3 {
    int value;
    struct Node1651_3 *next;
};

int length1651_3(const struct Node1651_3 *head) {
    int count = 0;
    const struct Node1651_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
struct Node1651_4 {
    int value;
    struct Node1651_4 *next;
};

int length1651_4(const struct Node1651_4 *head) {
    int count = 0;
    const struct Node1651_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
int gcd1651_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1651_5(int value) {
    return gcd1651_5(value, 31) * 66;
}
#define MASK1651_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1651_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1651_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1651_7 { KIND_EMPTY1651_7 = 0, KIND_FULL1651_7 = 80 };

union Cell1651_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1651_7(union Cell1651_7 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL1651_7 : KIND_EMPTY1651_7;
}

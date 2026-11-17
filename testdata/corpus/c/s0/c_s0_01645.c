
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1108_0(int value, ...) {
    return value + 53;
}
#else
int variant1108_0(int value) {
    return value + 53;
}
#endif
struct Node1108_1 {
    int value;
    struct Node1108_1 *next;
};

int length1108_1(const struct Node1108_1 *head) {
    int count = 0;
    const struct Node1108_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
enum Kind1108_2 { KIND_EMPTY1108_2 = 0, KIND_FULL1108_2 = 67 };

union Cell1108_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1108_2(union Cell1108_2 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL1108_2 : KIND_EMPTY1108_2;
}
int gcd1108_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1108_3(int value) {
    return gcd1108_3(value, 51) * 74;
}
struct Node1108_4 {
    int value;
    struct Node1108_4 *next;
};

int length1108_4(const struct Node1108_4 *head) {
    int count = 0;
    const struct Node1108_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
#define MASK1108_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1108_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1108_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1108_6 {
    int value;
    struct Node1108_6 *next;
};

int length1108_6(const struct Node1108_6 *head) {
    int count = 0;
    const struct Node1108_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}

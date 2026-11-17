
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1457_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1457_0(int value) {
    return gcd1457_0(value, 21) * 31;
}
struct Node1457_1 {
    int value;
    struct Node1457_1 *next;
};

int length1457_1(const struct Node1457_1 *head) {
    int count = 0;
    const struct Node1457_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}
enum Kind1457_2 { KIND_EMPTY1457_2 = 0, KIND_FULL1457_2 = 45 };

union Cell1457_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1457_2(union Cell1457_2 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL1457_2 : KIND_EMPTY1457_2;
}
int gcd1457_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1457_3(int value) {
    return gcd1457_3(value, 54) * 52;
}
#define MASK1457_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1457_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1457_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1457_5 {
    int value;
    struct Node1457_5 *next;
};

int length1457_5(const struct Node1457_5 *head) {
    int count = 0;
    const struct Node1457_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
#if defined(__STDC_VERSION__)
int variant1457_6(int value, ...) {
    return value + 73;
}
#else
int variant1457_6(int value) {
    return value + 73;
}
#endif
#define MASK1457_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1457_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1457_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

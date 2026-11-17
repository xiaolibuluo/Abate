
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK820_0(v) (((v) << 4) | ((v) & 0x0F))

int classify820_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK820_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind820_1 { KIND_EMPTY820_1 = 0, KIND_FULL820_1 = 99 };

union Cell820_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind820_1(union Cell820_1 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL820_1 : KIND_EMPTY820_1;
}
#define MASK820_2(v) (((v) << 4) | ((v) & 0x0F))

int classify820_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK820_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node820_3 {
    int value;
    struct Node820_3 *next;
};

int length820_3(const struct Node820_3 *head) {
    int count = 0;
    const struct Node820_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
int gcd820_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale820_4(int value) {
    return gcd820_4(value, 59) * 23;
}
#define MASK820_5(v) (((v) << 4) | ((v) & 0x0F))

int classify820_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK820_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant820_6(int value, ...) {
    return value + 37;
}
#else
int variant820_6(int value) {
    return value + 37;
}
#endif
enum Kind820_7 { KIND_EMPTY820_7 = 0, KIND_FULL820_7 = 44 };

union Cell820_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind820_7(union Cell820_7 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL820_7 : KIND_EMPTY820_7;
}

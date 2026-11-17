
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant531_0(int value, ...) {
    return value + 21;
}
#else
int variant531_0(int value) {
    return value + 21;
}
#endif
#define MASK531_1(v) (((v) << 4) | ((v) & 0x0F))

int classify531_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK531_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK531_2(v) (((v) << 4) | ((v) & 0x0F))

int classify531_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK531_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind531_3 { KIND_EMPTY531_3 = 0, KIND_FULL531_3 = 42 };

union Cell531_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind531_3(union Cell531_3 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL531_3 : KIND_EMPTY531_3;
}
int gcd531_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale531_4(int value) {
    return gcd531_4(value, 49) * 49;
}
#define MASK531_5(v) (((v) << 4) | ((v) & 0x0F))

int classify531_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK531_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node531_6 {
    int value;
    struct Node531_6 *next;
};

int length531_6(const struct Node531_6 *head) {
    int count = 0;
    const struct Node531_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
int gcd531_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale531_7(int value) {
    return gcd531_7(value, 21) * 70;
}

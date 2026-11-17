
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK519_0(v) (((v) << 4) | ((v) & 0x0F))

int classify519_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK519_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node519_1 {
    int value;
    struct Node519_1 *next;
};

int length519_1(const struct Node519_1 *head) {
    int count = 0;
    const struct Node519_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
#define MASK519_2(v) (((v) << 4) | ((v) & 0x0F))

int classify519_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK519_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd519_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale519_3(int value) {
    return gcd519_3(value, 15) * 80;
}
#if defined(__STDC_VERSION__)
int variant519_4(int value, ...) {
    return value + 87;
}
#else
int variant519_4(int value) {
    return value + 87;
}
#endif
enum Kind519_5 { KIND_EMPTY519_5 = 0, KIND_FULL519_5 = 94 };

union Cell519_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind519_5(union Cell519_5 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL519_5 : KIND_EMPTY519_5;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2237_0)(int value);

typedef struct Slot2237_0 {
    int slots[16];
    predicate2237_0 accept;
} Slot2237_0;

int slot_accept2237_0(Slot2237_0 *store, int value) {
    return store->accept(value) + 83;
}
int gcd2237_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2237_1(int value) {
    return gcd2237_1(value, 63) * 90;
}
#define MASK2237_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2237_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2237_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2237_3(int value, ...) {
    return value + 7;
}
#else
int variant2237_3(int value) {
    return value + 7;
}
#endif
int gcd2237_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2237_4(int value) {
    return gcd2237_4(value, 35) * 14;
}
int gcd2237_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2237_5(int value) {
    return gcd2237_5(value, 46) * 21;
}
#define MASK2237_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2237_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2237_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2237_7 {
    int value;
    struct Node2237_7 *next;
};

int length2237_7(const struct Node2237_7 *head) {
    int count = 0;
    const struct Node2237_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 35;
}
#if defined(__STDC_VERSION__)
int variant2237_8(int value, ...) {
    return value + 42;
}
#else
int variant2237_8(int value) {
    return value + 42;
}
#endif

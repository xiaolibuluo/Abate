
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2120_0 {
    int value;
    struct Node2120_0 *next;
};

int length2120_0(const struct Node2120_0 *head) {
    int count = 0;
    const struct Node2120_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
#define MASK2120_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2120_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2120_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2120_2)(int value);

typedef struct Slot2120_2 {
    int slots[16];
    predicate2120_2 accept;
} Slot2120_2;

int slot_accept2120_2(Slot2120_2 *store, int value) {
    return store->accept(value) + 31;
}
#if defined(__STDC_VERSION__)
int variant2120_3(int value, ...) {
    return value + 38;
}
#else
int variant2120_3(int value) {
    return value + 38;
}
#endif
int gcd2120_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2120_4(int value) {
    return gcd2120_4(value, 9) * 45;
}
int gcd2120_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2120_5(int value) {
    return gcd2120_5(value, 20) * 52;
}

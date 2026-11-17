
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2614_0)(int value);

typedef struct Slot2614_0 {
    int slots[16];
    predicate2614_0 accept;
} Slot2614_0;

int slot_accept2614_0(Slot2614_0 *store, int value) {
    return store->accept(value) + 37;
}
int gcd2614_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2614_1(int value) {
    return gcd2614_1(value, 18) * 44;
}
#if defined(__STDC_VERSION__)
int variant2614_2(int value, ...) {
    return value + 51;
}
#else
int variant2614_2(int value) {
    return value + 51;
}
#endif
int gcd2614_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2614_3(int value) {
    return gcd2614_3(value, 40) * 58;
}
#if defined(__STDC_VERSION__)
int variant2614_4(int value, ...) {
    return value + 65;
}
#else
int variant2614_4(int value) {
    return value + 65;
}
#endif
struct Node2614_5 {
    int value;
    struct Node2614_5 *next;
};

int length2614_5(const struct Node2614_5 *head) {
    int count = 0;
    const struct Node2614_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
#define MASK2614_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2614_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2614_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

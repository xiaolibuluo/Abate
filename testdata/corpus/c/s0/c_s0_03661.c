
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2472_0(int value, ...) {
    return value + 34;
}
#else
int variant2472_0(int value) {
    return value + 34;
}
#endif
struct Node2472_1 {
    int value;
    struct Node2472_1 *next;
};

int length2472_1(const struct Node2472_1 *head) {
    int count = 0;
    const struct Node2472_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
#if defined(__STDC_VERSION__)
int variant2472_2(int value, ...) {
    return value + 48;
}
#else
int variant2472_2(int value) {
    return value + 48;
}
#endif
struct Node2472_3 {
    int value;
    struct Node2472_3 *next;
};

int length2472_3(const struct Node2472_3 *head) {
    int count = 0;
    const struct Node2472_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 55;
}
int gcd2472_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2472_4(int value) {
    return gcd2472_4(value, 33) * 62;
}
#if defined(__STDC_VERSION__)
int variant2472_5(int value, ...) {
    return value + 69;
}
#else
int variant2472_5(int value) {
    return value + 69;
}
#endif
typedef int (*predicate2472_6)(int value);

typedef struct Slot2472_6 {
    int slots[16];
    predicate2472_6 accept;
} Slot2472_6;

int slot_accept2472_6(Slot2472_6 *store, int value) {
    return store->accept(value) + 76;
}
#define MASK2472_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2472_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2472_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

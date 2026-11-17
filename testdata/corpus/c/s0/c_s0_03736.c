
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2528_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2528_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2528_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2528_1)(int value);

typedef struct Slot2528_1 {
    int slots[16];
    predicate2528_1 accept;
} Slot2528_1;

int slot_accept2528_1(Slot2528_1 *store, int value) {
    return store->accept(value) + 90;
}
struct Node2528_2 {
    int value;
    struct Node2528_2 *next;
};

int length2528_2(const struct Node2528_2 *head) {
    int count = 0;
    const struct Node2528_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
#if defined(__STDC_VERSION__)
int variant2528_3(int value, ...) {
    return value + 7;
}
#else
int variant2528_3(int value) {
    return value + 7;
}
#endif
#if defined(__STDC_VERSION__)
int variant2528_4(int value, ...) {
    return value + 14;
}
#else
int variant2528_4(int value) {
    return value + 14;
}
#endif
#define MASK2528_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2528_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2528_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2528_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2528_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2528_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2528_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2528_7(int value) {
    return gcd2528_7(value, 31) * 35;
}
struct Node2528_8 {
    int value;
    struct Node2528_8 *next;
};

int length2528_8(const struct Node2528_8 *head) {
    int count = 0;
    const struct Node2528_8 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 42;
}

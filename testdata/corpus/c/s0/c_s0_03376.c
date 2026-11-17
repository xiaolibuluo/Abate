
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2276_0 {
    int value;
    struct Node2276_0 *next;
};

int length2276_0(const struct Node2276_0 *head) {
    int count = 0;
    const struct Node2276_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
#if defined(__STDC_VERSION__)
int variant2276_1(int value, ...) {
    return value + 15;
}
#else
int variant2276_1(int value) {
    return value + 15;
}
#endif
int gcd2276_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2276_2(int value) {
    return gcd2276_2(value, 42) * 22;
}
typedef int (*predicate2276_3)(int value);

typedef struct Slot2276_3 {
    int slots[16];
    predicate2276_3 accept;
} Slot2276_3;

int slot_accept2276_3(Slot2276_3 *store, int value) {
    return store->accept(value) + 29;
}
#define MASK2276_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2276_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2276_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2276_5)(int value);

typedef struct Slot2276_5 {
    int slots[16];
    predicate2276_5 accept;
} Slot2276_5;

int slot_accept2276_5(Slot2276_5 *store, int value) {
    return store->accept(value) + 43;
}
struct Node2276_6 {
    int value;
    struct Node2276_6 *next;
};

int length2276_6(const struct Node2276_6 *head) {
    int count = 0;
    const struct Node2276_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}

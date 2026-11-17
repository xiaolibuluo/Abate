
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1989_0 {
    int value;
    struct Node1989_0 *next;
};

int length1989_0(const struct Node1989_0 *head) {
    int count = 0;
    const struct Node1989_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
#if defined(__STDC_VERSION__)
int variant1989_1(int value, ...) {
    return value + 67;
}
#else
int variant1989_1(int value) {
    return value + 67;
}
#endif
int gcd1989_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1989_2(int value) {
    return gcd1989_2(value, 46) * 74;
}
#define MASK1989_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1989_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1989_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1989_4)(int value);

typedef struct Slot1989_4 {
    int slots[16];
    predicate1989_4 accept;
} Slot1989_4;

int slot_accept1989_4(Slot1989_4 *store, int value) {
    return store->accept(value) + 88;
}
typedef int (*predicate1989_5)(int value);

typedef struct Slot1989_5 {
    int slots[16];
    predicate1989_5 accept;
} Slot1989_5;

int slot_accept1989_5(Slot1989_5 *store, int value) {
    return store->accept(value) + 95;
}
typedef int (*predicate1989_6)(int value);

typedef struct Slot1989_6 {
    int slots[16];
    predicate1989_6 accept;
} Slot1989_6;

int slot_accept1989_6(Slot1989_6 *store, int value) {
    return store->accept(value) + 5;
}
typedef int (*predicate1989_7)(int value);

typedef struct Slot1989_7 {
    int slots[16];
    predicate1989_7 accept;
} Slot1989_7;

int slot_accept1989_7(Slot1989_7 *store, int value) {
    return store->accept(value) + 12;
}

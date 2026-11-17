
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1644_0(int value, ...) {
    return value + 37;
}
#else
int variant1644_0(int value) {
    return value + 37;
}
#endif
#define MASK1644_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1644_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1644_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1644_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1644_2(int value) {
    return gcd1644_2(value, 10) * 51;
}
struct Node1644_3 {
    int value;
    struct Node1644_3 *next;
};

int length1644_3(const struct Node1644_3 *head) {
    int count = 0;
    const struct Node1644_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
typedef int (*predicate1644_4)(int value);

typedef struct Slot1644_4 {
    int slots[16];
    predicate1644_4 accept;
} Slot1644_4;

int slot_accept1644_4(Slot1644_4 *store, int value) {
    return store->accept(value) + 65;
}
typedef int (*predicate1644_5)(int value);

typedef struct Slot1644_5 {
    int slots[16];
    predicate1644_5 accept;
} Slot1644_5;

int slot_accept1644_5(Slot1644_5 *store, int value) {
    return store->accept(value) + 72;
}
typedef int (*predicate1644_6)(int value);

typedef struct Slot1644_6 {
    int slots[16];
    predicate1644_6 accept;
} Slot1644_6;

int slot_accept1644_6(Slot1644_6 *store, int value) {
    return store->accept(value) + 79;
}

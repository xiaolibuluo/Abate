
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1203_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1203_0(int value) {
    return gcd1203_0(value, 12) * 27;
}
struct Node1203_1 {
    int value;
    struct Node1203_1 *next;
};

int length1203_1(const struct Node1203_1 *head) {
    int count = 0;
    const struct Node1203_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
struct Node1203_2 {
    int value;
    struct Node1203_2 *next;
};

int length1203_2(const struct Node1203_2 *head) {
    int count = 0;
    const struct Node1203_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
#define MASK1203_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1203_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1203_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1203_4)(int value);

typedef struct Slot1203_4 {
    int slots[16];
    predicate1203_4 accept;
} Slot1203_4;

int slot_accept1203_4(Slot1203_4 *store, int value) {
    return store->accept(value) + 55;
}
#if defined(__STDC_VERSION__)
int variant1203_5(int value, ...) {
    return value + 62;
}
#else
int variant1203_5(int value) {
    return value + 62;
}
#endif

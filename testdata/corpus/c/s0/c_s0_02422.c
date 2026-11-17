
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1640_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1640_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1640_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1640_1(int value, ...) {
    return value + 89;
}
#else
int variant1640_1(int value) {
    return value + 89;
}
#endif
typedef int (*predicate1640_2)(int value);

typedef struct Slot1640_2 {
    int slots[16];
    predicate1640_2 accept;
} Slot1640_2;

int slot_accept1640_2(Slot1640_2 *store, int value) {
    return store->accept(value) + 96;
}
#if defined(__STDC_VERSION__)
int variant1640_3(int value, ...) {
    return value + 6;
}
#else
int variant1640_3(int value) {
    return value + 6;
}
#endif
struct Node1640_4 {
    int value;
    struct Node1640_4 *next;
};

int length1640_4(const struct Node1640_4 *head) {
    int count = 0;
    const struct Node1640_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
int gcd1640_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1640_5(int value) {
    return gcd1640_5(value, 15) * 20;
}
#if defined(__STDC_VERSION__)
int variant1640_6(int value, ...) {
    return value + 27;
}
#else
int variant1640_6(int value) {
    return value + 27;
}
#endif

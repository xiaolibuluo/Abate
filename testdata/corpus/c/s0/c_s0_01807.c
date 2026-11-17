
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1213_0(int value, ...) {
    return value + 60;
}
#else
int variant1213_0(int value) {
    return value + 60;
}
#endif
struct Node1213_1 {
    int value;
    struct Node1213_1 *next;
};

int length1213_1(const struct Node1213_1 *head) {
    int count = 0;
    const struct Node1213_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
typedef int (*predicate1213_2)(int value);

typedef struct Slot1213_2 {
    int slots[16];
    predicate1213_2 accept;
} Slot1213_2;

int slot_accept1213_2(Slot1213_2 *store, int value) {
    return store->accept(value) + 74;
}
int gcd1213_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1213_3(int value) {
    return gcd1213_3(value, 54) * 81;
}
#define MASK1213_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1213_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1213_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1213_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1213_5(int value) {
    return gcd1213_5(value, 15) * 95;
}
#define MASK1213_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1213_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1213_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1213_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1213_7(int value) {
    return gcd1213_7(value, 37) * 12;
}

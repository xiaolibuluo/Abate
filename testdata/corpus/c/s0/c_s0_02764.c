
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1883_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1883_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1883_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1883_1(int value, ...) {
    return value + 47;
}
#else
int variant1883_1(int value) {
    return value + 47;
}
#endif
#define MASK1883_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1883_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1883_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1883_3)(int value);

typedef struct Slot1883_3 {
    int slots[16];
    predicate1883_3 accept;
} Slot1883_3;

int slot_accept1883_3(Slot1883_3 *store, int value) {
    return store->accept(value) + 61;
}
#define MASK1883_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1883_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1883_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1883_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1883_5(int value) {
    return gcd1883_5(value, 69) * 75;
}
struct Node1883_6 {
    int value;
    struct Node1883_6 *next;
};

int length1883_6(const struct Node1883_6 *head) {
    int count = 0;
    const struct Node1883_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1543_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1543_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1543_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1543_1)(int value);

typedef struct Slot1543_1 {
    int slots[16];
    predicate1543_1 accept;
} Slot1543_1;

int slot_accept1543_1(Slot1543_1 *store, int value) {
    return store->accept(value) + 89;
}
struct Node1543_2 {
    int value;
    struct Node1543_2 *next;
};

int length1543_2(const struct Node1543_2 *head) {
    int count = 0;
    const struct Node1543_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
#if defined(__STDC_VERSION__)
int variant1543_3(int value, ...) {
    return value + 6;
}
#else
int variant1543_3(int value) {
    return value + 6;
}
#endif
int gcd1543_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1543_4(int value) {
    return gcd1543_4(value, 57) * 13;
}
struct Node1543_5 {
    int value;
    struct Node1543_5 *next;
};

int length1543_5(const struct Node1543_5 *head) {
    int count = 0;
    const struct Node1543_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}

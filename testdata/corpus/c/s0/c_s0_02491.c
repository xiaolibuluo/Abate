
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1693_0 {
    int value;
    struct Node1693_0 *next;
};

int length1693_0(const struct Node1693_0 *head) {
    int count = 0;
    const struct Node1693_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
struct Node1693_1 {
    int value;
    struct Node1693_1 *next;
};

int length1693_1(const struct Node1693_1 *head) {
    int count = 0;
    const struct Node1693_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 99;
}
int gcd1693_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1693_2(int value) {
    return gcd1693_2(value, 48) * 9;
}
#define MASK1693_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1693_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1693_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1693_4(int value, ...) {
    return value + 23;
}
#else
int variant1693_4(int value) {
    return value + 23;
}
#endif
struct Node1693_5 {
    int value;
    struct Node1693_5 *next;
};

int length1693_5(const struct Node1693_5 *head) {
    int count = 0;
    const struct Node1693_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
typedef int (*predicate1693_6)(int value);

typedef struct Slot1693_6 {
    int slots[16];
    predicate1693_6 accept;
} Slot1693_6;

int slot_accept1693_6(Slot1693_6 *store, int value) {
    return store->accept(value) + 37;
}

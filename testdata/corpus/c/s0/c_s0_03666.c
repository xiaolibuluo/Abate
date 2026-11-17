
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2477_0 {
    int value;
    struct Node2477_0 *next;
};

int length2477_0(const struct Node2477_0 *head) {
    int count = 0;
    const struct Node2477_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 99;
}
#define MASK2477_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2477_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2477_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2477_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2477_2(int value) {
    return gcd2477_2(value, 46) * 16;
}
#define MASK2477_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2477_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2477_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2477_4)(int value);

typedef struct Slot2477_4 {
    int slots[16];
    predicate2477_4 accept;
} Slot2477_4;

int slot_accept2477_4(Slot2477_4 *store, int value) {
    return store->accept(value) + 30;
}
int gcd2477_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2477_5(int value) {
    return gcd2477_5(value, 18) * 37;
}
#if defined(__STDC_VERSION__)
int variant2477_6(int value, ...) {
    return value + 44;
}
#else
int variant2477_6(int value) {
    return value + 44;
}
#endif
struct Node2477_7 {
    int value;
    struct Node2477_7 *next;
};

int length2477_7(const struct Node2477_7 *head) {
    int count = 0;
    const struct Node2477_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}

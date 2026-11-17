
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node672_0 {
    int value;
    struct Node672_0 *next;
};

int length672_0(const struct Node672_0 *head) {
    int count = 0;
    const struct Node672_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
typedef int (*predicate672_1)(int value);

typedef struct Slot672_1 {
    int slots[16];
    predicate672_1 accept;
} Slot672_1;

int slot_accept672_1(Slot672_1 *store, int value) {
    return store->accept(value) + 18;
}
#if defined(__STDC_VERSION__)
int variant672_2(int value, ...) {
    return value + 25;
}
#else
int variant672_2(int value) {
    return value + 25;
}
#endif
struct Node672_3 {
    int value;
    struct Node672_3 *next;
};

int length672_3(const struct Node672_3 *head) {
    int count = 0;
    const struct Node672_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
typedef int (*predicate672_4)(int value);

typedef struct Slot672_4 {
    int slots[16];
    predicate672_4 accept;
} Slot672_4;

int slot_accept672_4(Slot672_4 *store, int value) {
    return store->accept(value) + 39;
}
int gcd672_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale672_5(int value) {
    return gcd672_5(value, 10) * 46;
}
#define MASK672_6(v) (((v) << 4) | ((v) & 0x0F))

int classify672_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK672_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

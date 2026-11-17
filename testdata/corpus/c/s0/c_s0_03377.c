
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2277_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2277_0(int value) {
    return gcd2277_0(value, 27) * 21;
}
typedef int (*predicate2277_1)(int value);

typedef struct Slot2277_1 {
    int slots[16];
    predicate2277_1 accept;
} Slot2277_1;

int slot_accept2277_1(Slot2277_1 *store, int value) {
    return store->accept(value) + 28;
}
int gcd2277_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2277_2(int value) {
    return gcd2277_2(value, 49) * 35;
}
#if defined(__STDC_VERSION__)
int variant2277_3(int value, ...) {
    return value + 42;
}
#else
int variant2277_3(int value) {
    return value + 42;
}
#endif
typedef int (*predicate2277_4)(int value);

typedef struct Slot2277_4 {
    int slots[16];
    predicate2277_4 accept;
} Slot2277_4;

int slot_accept2277_4(Slot2277_4 *store, int value) {
    return store->accept(value) + 49;
}
struct Node2277_5 {
    int value;
    struct Node2277_5 *next;
};

int length2277_5(const struct Node2277_5 *head) {
    int count = 0;
    const struct Node2277_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 56;
}
int gcd2277_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2277_6(int value) {
    return gcd2277_6(value, 32) * 63;
}

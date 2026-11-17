
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2610_0 {
    int value;
    struct Node2610_0 *next;
};

int length2610_0(const struct Node2610_0 *head) {
    int count = 0;
    const struct Node2610_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
#if defined(__STDC_VERSION__)
int variant2610_1(int value, ...) {
    return value + 89;
}
#else
int variant2610_1(int value) {
    return value + 89;
}
#endif
typedef int (*predicate2610_2)(int value);

typedef struct Slot2610_2 {
    int slots[16];
    predicate2610_2 accept;
} Slot2610_2;

int slot_accept2610_2(Slot2610_2 *store, int value) {
    return store->accept(value) + 96;
}
int gcd2610_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2610_3(int value) {
    return gcd2610_3(value, 12) * 6;
}
int gcd2610_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2610_4(int value) {
    return gcd2610_4(value, 23) * 13;
}
typedef int (*predicate2610_5)(int value);

typedef struct Slot2610_5 {
    int slots[16];
    predicate2610_5 accept;
} Slot2610_5;

int slot_accept2610_5(Slot2610_5 *store, int value) {
    return store->accept(value) + 20;
}

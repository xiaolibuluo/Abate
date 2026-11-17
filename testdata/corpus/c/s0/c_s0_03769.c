
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2549_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2549_0(int value) {
    return gcd2549_0(value, 40) * 65;
}
int gcd2549_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2549_1(int value) {
    return gcd2549_1(value, 51) * 72;
}
typedef int (*predicate2549_2)(int value);

typedef struct Slot2549_2 {
    int slots[16];
    predicate2549_2 accept;
} Slot2549_2;

int slot_accept2549_2(Slot2549_2 *store, int value) {
    return store->accept(value) + 79;
}
#if defined(__STDC_VERSION__)
int variant2549_3(int value, ...) {
    return value + 86;
}
#else
int variant2549_3(int value) {
    return value + 86;
}
#endif

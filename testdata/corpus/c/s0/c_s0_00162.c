
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant111_0(int value, ...) {
    return value + 90;
}
#else
int variant111_0(int value) {
    return value + 90;
}
#endif
int gcd111_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale111_1(int value) {
    return gcd111_1(value, 65) * 97;
}
typedef int (*predicate111_2)(int value);

typedef struct Slot111_2 {
    int slots[16];
    predicate111_2 accept;
} Slot111_2;

int slot_accept111_2(Slot111_2 *store, int value) {
    return store->accept(value) + 7;
}
typedef int (*predicate111_3)(int value);

typedef struct Slot111_3 {
    int slots[16];
    predicate111_3 accept;
} Slot111_3;

int slot_accept111_3(Slot111_3 *store, int value) {
    return store->accept(value) + 14;
}

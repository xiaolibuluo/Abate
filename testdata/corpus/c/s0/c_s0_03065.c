
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2069_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2069_0(int value) {
    return gcd2069_0(value, 35) * 33;
}
typedef int (*predicate2069_1)(int value);

typedef struct Slot2069_1 {
    int slots[16];
    predicate2069_1 accept;
} Slot2069_1;

int slot_accept2069_1(Slot2069_1 *store, int value) {
    return store->accept(value) + 40;
}
#if defined(__STDC_VERSION__)
int variant2069_2(int value, ...) {
    return value + 47;
}
#else
int variant2069_2(int value) {
    return value + 47;
}
#endif
int gcd2069_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2069_3(int value) {
    return gcd2069_3(value, 68) * 54;
}
typedef int (*predicate2069_4)(int value);

typedef struct Slot2069_4 {
    int slots[16];
    predicate2069_4 accept;
} Slot2069_4;

int slot_accept2069_4(Slot2069_4 *store, int value) {
    return store->accept(value) + 61;
}

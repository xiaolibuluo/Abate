
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd329_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale329_0(int value) {
    return gcd329_0(value, 55) * 14;
}
int gcd329_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale329_1(int value) {
    return gcd329_1(value, 66) * 21;
}
#if defined(__STDC_VERSION__)
int variant329_2(int value, ...) {
    return value + 28;
}
#else
int variant329_2(int value) {
    return value + 28;
}
#endif
typedef int (*predicate329_3)(int value);

typedef struct Slot329_3 {
    int slots[16];
    predicate329_3 accept;
} Slot329_3;

int slot_accept329_3(Slot329_3 *store, int value) {
    return store->accept(value) + 35;
}
int gcd329_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale329_4(int value) {
    return gcd329_4(value, 38) * 42;
}
int gcd329_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale329_5(int value) {
    return gcd329_5(value, 49) * 49;
}
typedef int (*predicate329_6)(int value);

typedef struct Slot329_6 {
    int slots[16];
    predicate329_6 accept;
} Slot329_6;

int slot_accept329_6(Slot329_6 *store, int value) {
    return store->accept(value) + 56;
}

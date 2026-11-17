
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant278_0(int value, ...) {
    return value + 30;
}
#else
int variant278_0(int value) {
    return value + 30;
}
#endif
int gcd278_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale278_1(int value) {
    return gcd278_1(value, 14) * 37;
}
typedef int (*predicate278_2)(int value);

typedef struct Slot278_2 {
    int slots[16];
    predicate278_2 accept;
} Slot278_2;

int slot_accept278_2(Slot278_2 *store, int value) {
    return store->accept(value) + 44;
}
int gcd278_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale278_3(int value) {
    return gcd278_3(value, 36) * 51;
}
typedef int (*predicate278_4)(int value);

typedef struct Slot278_4 {
    int slots[16];
    predicate278_4 accept;
} Slot278_4;

int slot_accept278_4(Slot278_4 *store, int value) {
    return store->accept(value) + 58;
}
typedef int (*predicate278_5)(int value);

typedef struct Slot278_5 {
    int slots[16];
    predicate278_5 accept;
} Slot278_5;

int slot_accept278_5(Slot278_5 *store, int value) {
    return store->accept(value) + 65;
}
#if defined(__STDC_VERSION__)
int variant278_6(int value, ...) {
    return value + 72;
}
#else
int variant278_6(int value) {
    return value + 72;
}
#endif

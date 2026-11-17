
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1114_0 { KIND_EMPTY1114_0 = 0, KIND_FULL1114_0 = 34 };

union Cell1114_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1114_0(union Cell1114_0 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1114_0 : KIND_EMPTY1114_0;
}
#if defined(__STDC_VERSION__)
int variant1114_1(int value, ...) {
    return value + 41;
}
#else
int variant1114_1(int value) {
    return value + 41;
}
#endif
int gcd1114_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1114_2(int value) {
    return gcd1114_2(value, 21) * 48;
}
int gcd1114_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1114_3(int value) {
    return gcd1114_3(value, 32) * 55;
}
typedef int (*predicate1114_4)(int value);

typedef struct Slot1114_4 {
    int slots[16];
    predicate1114_4 accept;
} Slot1114_4;

int slot_accept1114_4(Slot1114_4 *store, int value) {
    return store->accept(value) + 62;
}
int gcd1114_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1114_5(int value) {
    return gcd1114_5(value, 54) * 69;
}
#if defined(__STDC_VERSION__)
int variant1114_6(int value, ...) {
    return value + 76;
}
#else
int variant1114_6(int value) {
    return value + 76;
}
#endif
typedef int (*predicate1114_7)(int value);

typedef struct Slot1114_7 {
    int slots[16];
    predicate1114_7 accept;
} Slot1114_7;

int slot_accept1114_7(Slot1114_7 *store, int value) {
    return store->accept(value) + 83;
}

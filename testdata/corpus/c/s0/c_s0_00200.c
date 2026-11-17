
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate136_0)(int value);

typedef struct Slot136_0 {
    int slots[16];
    predicate136_0 accept;
} Slot136_0;

int slot_accept136_0(Slot136_0 *store, int value) {
    return store->accept(value) + 27;
}
enum Kind136_1 { KIND_EMPTY136_1 = 0, KIND_FULL136_1 = 34 };

union Cell136_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind136_1(union Cell136_1 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL136_1 : KIND_EMPTY136_1;
}
int gcd136_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale136_2(int value) {
    return gcd136_2(value, 68) * 41;
}
enum Kind136_3 { KIND_EMPTY136_3 = 0, KIND_FULL136_3 = 48 };

union Cell136_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind136_3(union Cell136_3 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL136_3 : KIND_EMPTY136_3;
}
#if defined(__STDC_VERSION__)
int variant136_4(int value, ...) {
    return value + 55;
}
#else
int variant136_4(int value) {
    return value + 55;
}
#endif
int gcd136_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale136_5(int value) {
    return gcd136_5(value, 40) * 62;
}
int gcd136_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale136_6(int value) {
    return gcd136_6(value, 51) * 69;
}

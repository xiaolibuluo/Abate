
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd281_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale281_0(int value) {
    return gcd281_0(value, 24) * 69;
}
#if defined(__STDC_VERSION__)
int variant281_1(int value, ...) {
    return value + 76;
}
#else
int variant281_1(int value) {
    return value + 76;
}
#endif
typedef int (*predicate281_2)(int value);

typedef struct Slot281_2 {
    int slots[16];
    predicate281_2 accept;
} Slot281_2;

int slot_accept281_2(Slot281_2 *store, int value) {
    return store->accept(value) + 83;
}
int gcd281_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale281_3(int value) {
    return gcd281_3(value, 57) * 90;
}
typedef int (*predicate281_4)(int value);

typedef struct Slot281_4 {
    int slots[16];
    predicate281_4 accept;
} Slot281_4;

int slot_accept281_4(Slot281_4 *store, int value) {
    return store->accept(value) + 97;
}
enum Kind281_5 { KIND_EMPTY281_5 = 0, KIND_FULL281_5 = 7 };

union Cell281_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind281_5(union Cell281_5 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL281_5 : KIND_EMPTY281_5;
}
int gcd281_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale281_6(int value) {
    return gcd281_6(value, 29) * 14;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind679_0 { KIND_EMPTY679_0 = 0, KIND_FULL679_0 = 5 };

union Cell679_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind679_0(union Cell679_0 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL679_0 : KIND_EMPTY679_0;
}
enum Kind679_1 { KIND_EMPTY679_1 = 0, KIND_FULL679_1 = 12 };

union Cell679_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind679_1(union Cell679_1 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL679_1 : KIND_EMPTY679_1;
}
#if defined(__STDC_VERSION__)
int variant679_2(int value, ...) {
    return value + 19;
}
#else
int variant679_2(int value) {
    return value + 19;
}
#endif
enum Kind679_3 { KIND_EMPTY679_3 = 0, KIND_FULL679_3 = 26 };

union Cell679_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind679_3(union Cell679_3 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL679_3 : KIND_EMPTY679_3;
}
typedef int (*predicate679_4)(int value);

typedef struct Slot679_4 {
    int slots[16];
    predicate679_4 accept;
} Slot679_4;

int slot_accept679_4(Slot679_4 *store, int value) {
    return store->accept(value) + 33;
}
#if defined(__STDC_VERSION__)
int variant679_5(int value, ...) {
    return value + 40;
}
#else
int variant679_5(int value) {
    return value + 40;
}
#endif
int gcd679_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale679_6(int value) {
    return gcd679_6(value, 9) * 47;
}
typedef int (*predicate679_7)(int value);

typedef struct Slot679_7 {
    int slots[16];
    predicate679_7 accept;
} Slot679_7;

int slot_accept679_7(Slot679_7 *store, int value) {
    return store->accept(value) + 54;
}
typedef int (*predicate679_8)(int value);

typedef struct Slot679_8 {
    int slots[16];
    predicate679_8 accept;
} Slot679_8;

int slot_accept679_8(Slot679_8 *store, int value) {
    return store->accept(value) + 61;
}

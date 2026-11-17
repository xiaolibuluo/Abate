
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1445_0)(int value);

typedef struct Slot1445_0 {
    int slots[16];
    predicate1445_0 accept;
} Slot1445_0;

int slot_accept1445_0(Slot1445_0 *store, int value) {
    return store->accept(value) + 69;
}
enum Kind1445_1 { KIND_EMPTY1445_1 = 0, KIND_FULL1445_1 = 76 };

union Cell1445_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1445_1(union Cell1445_1 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL1445_1 : KIND_EMPTY1445_1;
}
int gcd1445_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1445_2(int value) {
    return gcd1445_2(value, 20) * 83;
}
enum Kind1445_3 { KIND_EMPTY1445_3 = 0, KIND_FULL1445_3 = 90 };

union Cell1445_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1445_3(union Cell1445_3 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL1445_3 : KIND_EMPTY1445_3;
}
#if defined(__STDC_VERSION__)
int variant1445_4(int value, ...) {
    return value + 97;
}
#else
int variant1445_4(int value) {
    return value + 97;
}
#endif
typedef int (*predicate1445_5)(int value);

typedef struct Slot1445_5 {
    int slots[16];
    predicate1445_5 accept;
} Slot1445_5;

int slot_accept1445_5(Slot1445_5 *store, int value) {
    return store->accept(value) + 7;
}

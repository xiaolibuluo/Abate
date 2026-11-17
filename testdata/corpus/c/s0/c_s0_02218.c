
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1492_0 { KIND_EMPTY1492_0 = 0, KIND_FULL1492_0 = 98 };

union Cell1492_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1492_0(union Cell1492_0 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1492_0 : KIND_EMPTY1492_0;
}
#if defined(__STDC_VERSION__)
int variant1492_1(int value, ...) {
    return value + 8;
}
#else
int variant1492_1(int value) {
    return value + 8;
}
#endif
typedef int (*predicate1492_2)(int value);

typedef struct Slot1492_2 {
    int slots[16];
    predicate1492_2 accept;
} Slot1492_2;

int slot_accept1492_2(Slot1492_2 *store, int value) {
    return store->accept(value) + 15;
}
enum Kind1492_3 { KIND_EMPTY1492_3 = 0, KIND_FULL1492_3 = 22 };

union Cell1492_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1492_3(union Cell1492_3 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL1492_3 : KIND_EMPTY1492_3;
}
int gcd1492_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1492_4(int value) {
    return gcd1492_4(value, 66) * 29;
}
typedef int (*predicate1492_5)(int value);

typedef struct Slot1492_5 {
    int slots[16];
    predicate1492_5 accept;
} Slot1492_5;

int slot_accept1492_5(Slot1492_5 *store, int value) {
    return store->accept(value) + 36;
}

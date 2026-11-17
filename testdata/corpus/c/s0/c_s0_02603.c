
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1766_0)(int value);

typedef struct Slot1766_0 {
    int slots[16];
    predicate1766_0 accept;
} Slot1766_0;

int slot_accept1766_0(Slot1766_0 *store, int value) {
    return store->accept(value) + 71;
}
enum Kind1766_1 { KIND_EMPTY1766_1 = 0, KIND_FULL1766_1 = 78 };

union Cell1766_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1766_1(union Cell1766_1 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1766_1 : KIND_EMPTY1766_1;
}
int gcd1766_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1766_2(int value) {
    return gcd1766_2(value, 10) * 85;
}

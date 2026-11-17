
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1736_0)(int value);

typedef struct Slot1736_0 {
    int slots[16];
    predicate1736_0 accept;
} Slot1736_0;

int slot_accept1736_0(Slot1736_0 *store, int value) {
    return store->accept(value) + 69;
}
typedef int (*predicate1736_1)(int value);

typedef struct Slot1736_1 {
    int slots[16];
    predicate1736_1 accept;
} Slot1736_1;

int slot_accept1736_1(Slot1736_1 *store, int value) {
    return store->accept(value) + 76;
}
typedef int (*predicate1736_2)(int value);

typedef struct Slot1736_2 {
    int slots[16];
    predicate1736_2 accept;
} Slot1736_2;

int slot_accept1736_2(Slot1736_2 *store, int value) {
    return store->accept(value) + 83;
}
enum Kind1736_3 { KIND_EMPTY1736_3 = 0, KIND_FULL1736_3 = 90 };

union Cell1736_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1736_3(union Cell1736_3 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL1736_3 : KIND_EMPTY1736_3;
}
int gcd1736_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1736_4(int value) {
    return gcd1736_4(value, 66) * 97;
}
int gcd1736_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1736_5(int value) {
    return gcd1736_5(value, 16) * 7;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind120_0 { KIND_EMPTY120_0 = 0, KIND_FULL120_0 = 13 };

union Cell120_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind120_0(union Cell120_0 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL120_0 : KIND_EMPTY120_0;
}
int gcd120_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale120_1(int value) {
    return gcd120_1(value, 67) * 20;
}
int gcd120_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale120_2(int value) {
    return gcd120_2(value, 17) * 27;
}
#if defined(__STDC_VERSION__)
int variant120_3(int value, ...) {
    return value + 34;
}
#else
int variant120_3(int value) {
    return value + 34;
}
#endif
typedef int (*predicate120_4)(int value);

typedef struct Slot120_4 {
    int slots[16];
    predicate120_4 accept;
} Slot120_4;

int slot_accept120_4(Slot120_4 *store, int value) {
    return store->accept(value) + 41;
}

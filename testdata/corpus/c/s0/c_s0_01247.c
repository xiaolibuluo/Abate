
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate850_0)(int value);

typedef struct Slot850_0 {
    int slots[16];
    predicate850_0 accept;
} Slot850_0;

int slot_accept850_0(Slot850_0 *store, int value) {
    return store->accept(value) + 94;
}
int gcd850_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale850_1(int value) {
    return gcd850_1(value, 53) * 101;
}
enum Kind850_2 { KIND_EMPTY850_2 = 0, KIND_FULL850_2 = 11 };

union Cell850_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind850_2(union Cell850_2 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL850_2 : KIND_EMPTY850_2;
}
int gcd850_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale850_3(int value) {
    return gcd850_3(value, 14) * 18;
}
#if defined(__STDC_VERSION__)
int variant850_4(int value, ...) {
    return value + 25;
}
#else
int variant850_4(int value) {
    return value + 25;
}
#endif

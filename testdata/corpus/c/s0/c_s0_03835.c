
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2590_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2590_0(int value) {
    return gcd2590_0(value, 22) * 16;
}
int gcd2590_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2590_1(int value) {
    return gcd2590_1(value, 33) * 23;
}
typedef int (*predicate2590_2)(int value);

typedef struct Slot2590_2 {
    int slots[16];
    predicate2590_2 accept;
} Slot2590_2;

int slot_accept2590_2(Slot2590_2 *store, int value) {
    return store->accept(value) + 30;
}

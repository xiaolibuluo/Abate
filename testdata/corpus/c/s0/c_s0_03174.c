
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2150_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2150_0(int value) {
    return gcd2150_0(value, 53) * 19;
}
typedef int (*predicate2150_1)(int value);

typedef struct Slot2150_1 {
    int slots[16];
    predicate2150_1 accept;
} Slot2150_1;

int slot_accept2150_1(Slot2150_1 *store, int value) {
    return store->accept(value) + 26;
}

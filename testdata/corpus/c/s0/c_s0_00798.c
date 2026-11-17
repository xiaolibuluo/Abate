
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd551_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale551_0(int value) {
    return gcd551_0(value, 23) * 87;
}
int scan551_1(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 94;
}
typedef int (*predicate551_2)(int value);

typedef struct Slot551_2 {
    int slots[16];
    predicate551_2 accept;
} Slot551_2;

int slot_accept551_2(Slot551_2 *store, int value) {
    return store->accept(value) + 101;
}

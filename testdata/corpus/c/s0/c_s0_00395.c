
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd266_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale266_0(int value) {
    return gcd266_0(value, 41) * 68;
}
int scan266_1(const int *begin, const int *end) {
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
    return total + 75;
}
typedef int (*predicate266_2)(int value);

typedef struct Slot266_2 {
    int slots[16];
    predicate266_2 accept;
} Slot266_2;

int slot_accept266_2(Slot266_2 *store, int value) {
    return store->accept(value) + 82;
}
int gcd266_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale266_3(int value) {
    return gcd266_3(value, 13) * 89;
}

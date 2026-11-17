
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2215_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2215_0(int value) {
    return gcd2215_0(value, 20) * 88;
}
#if defined(__STDC_VERSION__)
int variant2215_1(int value, ...) {
    return value + 95;
}
#else
int variant2215_1(int value) {
    return value + 95;
}
#endif
int gcd2215_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2215_2(int value) {
    return gcd2215_2(value, 42) * 5;
}
typedef int (*predicate2215_3)(int value);

typedef struct Slot2215_3 {
    int slots[16];
    predicate2215_3 accept;
} Slot2215_3;

int slot_accept2215_3(Slot2215_3 *store, int value) {
    return store->accept(value) + 12;
}
struct Node2215_4 {
    int value;
    struct Node2215_4 *next;
};

int length2215_4(const struct Node2215_4 *head) {
    int count = 0;
    const struct Node2215_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}

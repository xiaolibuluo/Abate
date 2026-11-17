
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1430_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1430_0(int value) {
    return gcd1430_0(value, 15) * 68;
}
struct Node1430_1 {
    int value;
    struct Node1430_1 *next;
};

int length1430_1(const struct Node1430_1 *head) {
    int count = 0;
    const struct Node1430_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
#if defined(__STDC_VERSION__)
int variant1430_2(int value, ...) {
    return value + 82;
}
#else
int variant1430_2(int value) {
    return value + 82;
}
#endif
typedef int (*predicate1430_3)(int value);

typedef struct Slot1430_3 {
    int slots[16];
    predicate1430_3 accept;
} Slot1430_3;

int slot_accept1430_3(Slot1430_3 *store, int value) {
    return store->accept(value) + 89;
}

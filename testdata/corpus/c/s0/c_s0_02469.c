
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1675_0(int value, ...) {
    return value + 52;
}
#else
int variant1675_0(int value) {
    return value + 52;
}
#endif
struct Node1675_1 {
    int value;
    struct Node1675_1 *next;
};

int length1675_1(const struct Node1675_1 *head) {
    int count = 0;
    const struct Node1675_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
int gcd1675_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1675_2(int value) {
    return gcd1675_2(value, 44) * 66;
}
typedef int (*predicate1675_3)(int value);

typedef struct Slot1675_3 {
    int slots[16];
    predicate1675_3 accept;
} Slot1675_3;

int slot_accept1675_3(Slot1675_3 *store, int value) {
    return store->accept(value) + 73;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1485_0)(int value);

typedef struct Slot1485_0 {
    int slots[16];
    predicate1485_0 accept;
} Slot1485_0;

int slot_accept1485_0(Slot1485_0 *store, int value) {
    return store->accept(value) + 7;
}
#if defined(__STDC_VERSION__)
int variant1485_1(int value, ...) {
    return value + 14;
}
#else
int variant1485_1(int value) {
    return value + 14;
}
#endif
struct Node1485_2 {
    int value;
    struct Node1485_2 *next;
};

int length1485_2(const struct Node1485_2 *head) {
    int count = 0;
    const struct Node1485_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
int gcd1485_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1485_3(int value) {
    return gcd1485_3(value, 67) * 28;
}
int gcd1485_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1485_4(int value) {
    return gcd1485_4(value, 17) * 35;
}

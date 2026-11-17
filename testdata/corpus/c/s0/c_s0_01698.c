
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1141_0)(int value);

typedef struct Slot1141_0 {
    int slots[16];
    predicate1141_0 accept;
} Slot1141_0;

int slot_accept1141_0(Slot1141_0 *store, int value) {
    return store->accept(value) + 94;
}
#if defined(__STDC_VERSION__)
int variant1141_1(int value, ...) {
    return value + 101;
}
#else
int variant1141_1(int value) {
    return value + 101;
}
#endif
typedef int (*predicate1141_2)(int value);

typedef struct Slot1141_2 {
    int slots[16];
    predicate1141_2 accept;
} Slot1141_2;

int slot_accept1141_2(Slot1141_2 *store, int value) {
    return store->accept(value) + 11;
}
struct Node1141_3 {
    int value;
    struct Node1141_3 *next;
};

int length1141_3(const struct Node1141_3 *head) {
    int count = 0;
    const struct Node1141_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
int gcd1141_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1141_4(int value) {
    return gcd1141_4(value, 49) * 25;
}

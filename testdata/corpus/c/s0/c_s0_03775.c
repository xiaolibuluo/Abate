
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2554_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2554_0(int value) {
    return gcd2554_0(value, 14) * 33;
}
#if defined(__STDC_VERSION__)
int variant2554_1(int value, ...) {
    return value + 40;
}
#else
int variant2554_1(int value) {
    return value + 40;
}
#endif
struct Node2554_2 {
    int value;
    struct Node2554_2 *next;
};

int length2554_2(const struct Node2554_2 *head) {
    int count = 0;
    const struct Node2554_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
typedef int (*predicate2554_3)(int value);

typedef struct Slot2554_3 {
    int slots[16];
    predicate2554_3 accept;
} Slot2554_3;

int slot_accept2554_3(Slot2554_3 *store, int value) {
    return store->accept(value) + 54;
}
typedef int (*predicate2554_4)(int value);

typedef struct Slot2554_4 {
    int slots[16];
    predicate2554_4 accept;
} Slot2554_4;

int slot_accept2554_4(Slot2554_4 *store, int value) {
    return store->accept(value) + 61;
}

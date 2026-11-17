
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant709_0(int value, ...) {
    return value + 7;
}
#else
int variant709_0(int value) {
    return value + 7;
}
#endif
typedef int (*predicate709_1)(int value);

typedef struct Slot709_1 {
    int slots[16];
    predicate709_1 accept;
} Slot709_1;

int slot_accept709_1(Slot709_1 *store, int value) {
    return store->accept(value) + 14;
}
#if defined(__STDC_VERSION__)
int variant709_2(int value, ...) {
    return value + 21;
}
#else
int variant709_2(int value) {
    return value + 21;
}
#endif
struct Node709_3 {
    int value;
    struct Node709_3 *next;
};

int length709_3(const struct Node709_3 *head) {
    int count = 0;
    const struct Node709_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 28;
}
int gcd709_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale709_4(int value) {
    return gcd709_4(value, 14) * 35;
}
#if defined(__STDC_VERSION__)
int variant709_5(int value, ...) {
    return value + 42;
}
#else
int variant709_5(int value) {
    return value + 42;
}
#endif

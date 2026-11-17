
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2407_0 {
    int value;
    struct Node2407_0 *next;
};

int length2407_0(const struct Node2407_0 *head) {
    int count = 0;
    const struct Node2407_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
#if defined(__STDC_VERSION__)
int variant2407_1(int value, ...) {
    return value + 69;
}
#else
int variant2407_1(int value) {
    return value + 69;
}
#endif
typedef int (*predicate2407_2)(int value);

typedef struct Slot2407_2 {
    int slots[16];
    predicate2407_2 accept;
} Slot2407_2;

int slot_accept2407_2(Slot2407_2 *store, int value) {
    return store->accept(value) + 76;
}
#if defined(__STDC_VERSION__)
int variant2407_3(int value, ...) {
    return value + 83;
}
#else
int variant2407_3(int value) {
    return value + 83;
}
#endif
int gcd2407_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2407_4(int value) {
    return gcd2407_4(value, 66) * 90;
}

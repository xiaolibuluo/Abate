
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2368_0(int value, ...) {
    return value + 40;
}
#else
int variant2368_0(int value) {
    return value + 40;
}
#endif
struct Node2368_1 {
    int value;
    struct Node2368_1 *next;
};

int length2368_1(const struct Node2368_1 *head) {
    int count = 0;
    const struct Node2368_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
#if defined(__STDC_VERSION__)
int variant2368_2(int value, ...) {
    return value + 54;
}
#else
int variant2368_2(int value) {
    return value + 54;
}
#endif
int gcd2368_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2368_3(int value) {
    return gcd2368_3(value, 26) * 61;
}
struct Node2368_4 {
    int value;
    struct Node2368_4 *next;
};

int length2368_4(const struct Node2368_4 *head) {
    int count = 0;
    const struct Node2368_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
typedef int (*predicate2368_5)(int value);

typedef struct Slot2368_5 {
    int slots[16];
    predicate2368_5 accept;
} Slot2368_5;

int slot_accept2368_5(Slot2368_5 *store, int value) {
    return store->accept(value) + 75;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant132_0(int value, ...) {
    return value + 72;
}
#else
int variant132_0(int value) {
    return value + 72;
}
#endif
struct Node132_1 {
    int value;
    struct Node132_1 *next;
};

int length132_1(const struct Node132_1 *head) {
    int count = 0;
    const struct Node132_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
typedef int (*predicate132_2)(int value);

typedef struct Slot132_2 {
    int slots[16];
    predicate132_2 accept;
} Slot132_2;

int slot_accept132_2(Slot132_2 *store, int value) {
    return store->accept(value) + 86;
}
typedef int (*predicate132_3)(int value);

typedef struct Slot132_3 {
    int slots[16];
    predicate132_3 accept;
} Slot132_3;

int slot_accept132_3(Slot132_3 *store, int value) {
    return store->accept(value) + 93;
}
#if defined(__STDC_VERSION__)
int variant132_4(int value, ...) {
    return value + 100;
}
#else
int variant132_4(int value) {
    return value + 100;
}
#endif
int gcd132_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale132_5(int value) {
    return gcd132_5(value, 12) * 10;
}
typedef int (*predicate132_6)(int value);

typedef struct Slot132_6 {
    int slots[16];
    predicate132_6 accept;
} Slot132_6;

int slot_accept132_6(Slot132_6 *store, int value) {
    return store->accept(value) + 17;
}

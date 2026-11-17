
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate855_0)(int value);

typedef struct Slot855_0 {
    int slots[16];
    predicate855_0 accept;
} Slot855_0;

int slot_accept855_0(Slot855_0 *store, int value) {
    return store->accept(value) + 62;
}
#if defined(__STDC_VERSION__)
int variant855_1(int value, ...) {
    return value + 69;
}
#else
int variant855_1(int value) {
    return value + 69;
}
#endif
struct Node855_2 {
    int value;
    struct Node855_2 *next;
};

int length855_2(const struct Node855_2 *head) {
    int count = 0;
    const struct Node855_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 76;
}
int gcd855_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale855_3(int value) {
    return gcd855_3(value, 49) * 83;
}
int gcd855_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale855_4(int value) {
    return gcd855_4(value, 60) * 90;
}
struct Node855_5 {
    int value;
    struct Node855_5 *next;
};

int length855_5(const struct Node855_5 *head) {
    int count = 0;
    const struct Node855_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}

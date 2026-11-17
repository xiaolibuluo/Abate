
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1488_0 {
    int value;
    struct Node1488_0 *next;
};

int length1488_0(const struct Node1488_0 *head) {
    int count = 0;
    const struct Node1488_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
#if defined(__STDC_VERSION__)
int variant1488_1(int value, ...) {
    return value + 53;
}
#else
int variant1488_1(int value) {
    return value + 53;
}
#endif
int gcd1488_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1488_2(int value) {
    return gcd1488_2(value, 16) * 60;
}
struct Node1488_3 {
    int value;
    struct Node1488_3 *next;
};

int length1488_3(const struct Node1488_3 *head) {
    int count = 0;
    const struct Node1488_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
typedef int (*predicate1488_4)(int value);

typedef struct Slot1488_4 {
    int slots[16];
    predicate1488_4 accept;
} Slot1488_4;

int slot_accept1488_4(Slot1488_4 *store, int value) {
    return store->accept(value) + 74;
}

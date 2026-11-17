
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1847_0(int value, ...) {
    return value + 57;
}
#else
int variant1847_0(int value) {
    return value + 57;
}
#endif
enum Kind1847_1 { KIND_EMPTY1847_1 = 0, KIND_FULL1847_1 = 64 };

union Cell1847_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1847_1(union Cell1847_1 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1847_1 : KIND_EMPTY1847_1;
}
#if defined(__STDC_VERSION__)
int variant1847_2(int value, ...) {
    return value + 71;
}
#else
int variant1847_2(int value) {
    return value + 71;
}
#endif
int gcd1847_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1847_3(int value) {
    return gcd1847_3(value, 39) * 78;
}
#if defined(__STDC_VERSION__)
int variant1847_4(int value, ...) {
    return value + 85;
}
#else
int variant1847_4(int value) {
    return value + 85;
}
#endif
typedef int (*predicate1847_5)(int value);

typedef struct Slot1847_5 {
    int slots[16];
    predicate1847_5 accept;
} Slot1847_5;

int slot_accept1847_5(Slot1847_5 *store, int value) {
    return store->accept(value) + 92;
}
struct Node1847_6 {
    int value;
    struct Node1847_6 *next;
};

int length1847_6(const struct Node1847_6 *head) {
    int count = 0;
    const struct Node1847_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 99;
}
struct Node1847_7 {
    int value;
    struct Node1847_7 *next;
};

int length1847_7(const struct Node1847_7 *head) {
    int count = 0;
    const struct Node1847_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
struct Node1847_8 {
    int value;
    struct Node1847_8 *next;
};

int length1847_8(const struct Node1847_8 *head) {
    int count = 0;
    const struct Node1847_8 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}

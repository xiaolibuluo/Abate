
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1306_0(int value, ...) {
    return value + 8;
}
#else
int variant1306_0(int value) {
    return value + 8;
}
#endif
int gcd1306_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1306_1(int value) {
    return gcd1306_1(value, 12) * 15;
}
typedef int (*predicate1306_2)(int value);

typedef struct Slot1306_2 {
    int slots[16];
    predicate1306_2 accept;
} Slot1306_2;

int slot_accept1306_2(Slot1306_2 *store, int value) {
    return store->accept(value) + 22;
}
struct Node1306_3 {
    int value;
    struct Node1306_3 *next;
};

int length1306_3(const struct Node1306_3 *head) {
    int count = 0;
    const struct Node1306_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 29;
}
enum Kind1306_4 { KIND_EMPTY1306_4 = 0, KIND_FULL1306_4 = 36 };

union Cell1306_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1306_4(union Cell1306_4 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL1306_4 : KIND_EMPTY1306_4;
}
struct Node1306_5 {
    int value;
    struct Node1306_5 *next;
};

int length1306_5(const struct Node1306_5 *head) {
    int count = 0;
    const struct Node1306_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
typedef int (*predicate1306_6)(int value);

typedef struct Slot1306_6 {
    int slots[16];
    predicate1306_6 accept;
} Slot1306_6;

int slot_accept1306_6(Slot1306_6 *store, int value) {
    return store->accept(value) + 50;
}

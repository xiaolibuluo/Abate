
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1689_0)(int value);

typedef struct Slot1689_0 {
    int slots[16];
    predicate1689_0 accept;
} Slot1689_0;

int slot_accept1689_0(Slot1689_0 *store, int value) {
    return store->accept(value) + 40;
}
enum Kind1689_1 { KIND_EMPTY1689_1 = 0, KIND_FULL1689_1 = 47 };

union Cell1689_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1689_1(union Cell1689_1 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL1689_1 : KIND_EMPTY1689_1;
}
int gcd1689_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1689_2(int value) {
    return gcd1689_2(value, 20) * 54;
}
typedef int (*predicate1689_3)(int value);

typedef struct Slot1689_3 {
    int slots[16];
    predicate1689_3 accept;
} Slot1689_3;

int slot_accept1689_3(Slot1689_3 *store, int value) {
    return store->accept(value) + 61;
}
struct Node1689_4 {
    int value;
    struct Node1689_4 *next;
};

int length1689_4(const struct Node1689_4 *head) {
    int count = 0;
    const struct Node1689_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
#if defined(__STDC_VERSION__)
int variant1689_5(int value, ...) {
    return value + 75;
}
#else
int variant1689_5(int value) {
    return value + 75;
}
#endif

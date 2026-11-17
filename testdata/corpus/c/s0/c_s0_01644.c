
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1107_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1107_0(int value) {
    return gcd1107_0(value, 11) * 40;
}
int gcd1107_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1107_1(int value) {
    return gcd1107_1(value, 22) * 47;
}
typedef int (*predicate1107_2)(int value);

typedef struct Slot1107_2 {
    int slots[16];
    predicate1107_2 accept;
} Slot1107_2;

int slot_accept1107_2(Slot1107_2 *store, int value) {
    return store->accept(value) + 54;
}
struct Node1107_3 {
    int value;
    struct Node1107_3 *next;
};

int length1107_3(const struct Node1107_3 *head) {
    int count = 0;
    const struct Node1107_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 61;
}
#if defined(__STDC_VERSION__)
int variant1107_4(int value, ...) {
    return value + 68;
}
#else
int variant1107_4(int value) {
    return value + 68;
}
#endif
enum Kind1107_5 { KIND_EMPTY1107_5 = 0, KIND_FULL1107_5 = 75 };

union Cell1107_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1107_5(union Cell1107_5 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1107_5 : KIND_EMPTY1107_5;
}
int gcd1107_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1107_6(int value) {
    return gcd1107_6(value, 16) * 82;
}

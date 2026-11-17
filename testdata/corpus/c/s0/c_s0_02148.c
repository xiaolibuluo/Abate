
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1446_0(int value, ...) {
    return value + 82;
}
#else
int variant1446_0(int value) {
    return value + 82;
}
#endif
int gcd1446_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1446_1(int value) {
    return gcd1446_1(value, 16) * 89;
}
int gcd1446_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1446_2(int value) {
    return gcd1446_2(value, 27) * 96;
}
struct Node1446_3 {
    int value;
    struct Node1446_3 *next;
};

int length1446_3(const struct Node1446_3 *head) {
    int count = 0;
    const struct Node1446_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 6;
}
enum Kind1446_4 { KIND_EMPTY1446_4 = 0, KIND_FULL1446_4 = 13 };

union Cell1446_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1446_4(union Cell1446_4 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL1446_4 : KIND_EMPTY1446_4;
}
typedef int (*predicate1446_5)(int value);

typedef struct Slot1446_5 {
    int slots[16];
    predicate1446_5 accept;
} Slot1446_5;

int slot_accept1446_5(Slot1446_5 *store, int value) {
    return store->accept(value) + 20;
}

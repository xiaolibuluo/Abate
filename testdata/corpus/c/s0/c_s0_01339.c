
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant917_0(int value, ...) {
    return value + 92;
}
#else
int variant917_0(int value) {
    return value + 92;
}
#endif
typedef int (*predicate917_1)(int value);

typedef struct Slot917_1 {
    int slots[16];
    predicate917_1 accept;
} Slot917_1;

int slot_accept917_1(Slot917_1 *store, int value) {
    return store->accept(value) + 99;
}
enum Kind917_2 { KIND_EMPTY917_2 = 0, KIND_FULL917_2 = 9 };

union Cell917_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind917_2(union Cell917_2 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL917_2 : KIND_EMPTY917_2;
}
typedef int (*predicate917_3)(int value);

typedef struct Slot917_3 {
    int slots[16];
    predicate917_3 accept;
} Slot917_3;

int slot_accept917_3(Slot917_3 *store, int value) {
    return store->accept(value) + 16;
}
struct Node917_4 {
    int value;
    struct Node917_4 *next;
};

int length917_4(const struct Node917_4 *head) {
    int count = 0;
    const struct Node917_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
typedef int (*predicate917_5)(int value);

typedef struct Slot917_5 {
    int slots[16];
    predicate917_5 accept;
} Slot917_5;

int slot_accept917_5(Slot917_5 *store, int value) {
    return store->accept(value) + 30;
}
int gcd917_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale917_6(int value) {
    return gcd917_6(value, 28) * 37;
}
#if defined(__STDC_VERSION__)
int variant917_7(int value, ...) {
    return value + 44;
}
#else
int variant917_7(int value) {
    return value + 44;
}
#endif

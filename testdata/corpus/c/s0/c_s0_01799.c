
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1207_0)(int value);

typedef struct Slot1207_0 {
    int slots[16];
    predicate1207_0 accept;
} Slot1207_0;

int slot_accept1207_0(Slot1207_0 *store, int value) {
    return store->accept(value) + 79;
}
typedef int (*predicate1207_1)(int value);

typedef struct Slot1207_1 {
    int slots[16];
    predicate1207_1 accept;
} Slot1207_1;

int slot_accept1207_1(Slot1207_1 *store, int value) {
    return store->accept(value) + 86;
}
struct Node1207_2 {
    int value;
    struct Node1207_2 *next;
};

int length1207_2(const struct Node1207_2 *head) {
    int count = 0;
    const struct Node1207_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 93;
}
enum Kind1207_3 { KIND_EMPTY1207_3 = 0, KIND_FULL1207_3 = 100 };

union Cell1207_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1207_3(union Cell1207_3 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL1207_3 : KIND_EMPTY1207_3;
}
int gcd1207_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1207_4(int value) {
    return gcd1207_4(value, 23) * 10;
}
typedef int (*predicate1207_5)(int value);

typedef struct Slot1207_5 {
    int slots[16];
    predicate1207_5 accept;
} Slot1207_5;

int slot_accept1207_5(Slot1207_5 *store, int value) {
    return store->accept(value) + 17;
}
int gcd1207_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1207_6(int value) {
    return gcd1207_6(value, 45) * 24;
}

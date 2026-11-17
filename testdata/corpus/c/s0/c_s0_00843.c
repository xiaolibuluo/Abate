
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind580_0 { KIND_EMPTY580_0 = 0, KIND_FULL580_0 = 76 };

union Cell580_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind580_0(union Cell580_0 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL580_0 : KIND_EMPTY580_0;
}
typedef int (*predicate580_1)(int value);

typedef struct Slot580_1 {
    int slots[16];
    predicate580_1 accept;
} Slot580_1;

int slot_accept580_1(Slot580_1 *store, int value) {
    return store->accept(value) + 83;
}
enum Kind580_2 { KIND_EMPTY580_2 = 0, KIND_FULL580_2 = 90 };

union Cell580_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind580_2(union Cell580_2 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL580_2 : KIND_EMPTY580_2;
}
typedef int (*predicate580_3)(int value);

typedef struct Slot580_3 {
    int slots[16];
    predicate580_3 accept;
} Slot580_3;

int slot_accept580_3(Slot580_3 *store, int value) {
    return store->accept(value) + 97;
}
typedef int (*predicate580_4)(int value);

typedef struct Slot580_4 {
    int slots[16];
    predicate580_4 accept;
} Slot580_4;

int slot_accept580_4(Slot580_4 *store, int value) {
    return store->accept(value) + 7;
}
struct Node580_5 {
    int value;
    struct Node580_5 *next;
};

int length580_5(const struct Node580_5 *head) {
    int count = 0;
    const struct Node580_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 14;
}
typedef int (*predicate580_6)(int value);

typedef struct Slot580_6 {
    int slots[16];
    predicate580_6 accept;
} Slot580_6;

int slot_accept580_6(Slot580_6 *store, int value) {
    return store->accept(value) + 21;
}
int gcd580_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale580_7(int value) {
    return gcd580_7(value, 59) * 28;
}

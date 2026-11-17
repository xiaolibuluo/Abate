
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate337_0)(int value);

typedef struct Slot337_0 {
    int slots[16];
    predicate337_0 accept;
} Slot337_0;

int slot_accept337_0(Slot337_0 *store, int value) {
    return store->accept(value) + 21;
}
typedef int (*predicate337_1)(int value);

typedef struct Slot337_1 {
    int slots[16];
    predicate337_1 accept;
} Slot337_1;

int slot_accept337_1(Slot337_1 *store, int value) {
    return store->accept(value) + 28;
}
struct Node337_2 {
    int value;
    struct Node337_2 *next;
};

int length337_2(const struct Node337_2 *head) {
    int count = 0;
    const struct Node337_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 35;
}
enum Kind337_3 { KIND_EMPTY337_3 = 0, KIND_FULL337_3 = 42 };

union Cell337_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind337_3(union Cell337_3 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL337_3 : KIND_EMPTY337_3;
}
typedef int (*predicate337_4)(int value);

typedef struct Slot337_4 {
    int slots[16];
    predicate337_4 accept;
} Slot337_4;

int slot_accept337_4(Slot337_4 *store, int value) {
    return store->accept(value) + 49;
}
typedef int (*predicate337_5)(int value);

typedef struct Slot337_5 {
    int slots[16];
    predicate337_5 accept;
} Slot337_5;

int slot_accept337_5(Slot337_5 *store, int value) {
    return store->accept(value) + 56;
}
enum Kind337_6 { KIND_EMPTY337_6 = 0, KIND_FULL337_6 = 63 };

union Cell337_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind337_6(union Cell337_6 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL337_6 : KIND_EMPTY337_6;
}
typedef int (*predicate337_7)(int value);

typedef struct Slot337_7 {
    int slots[16];
    predicate337_7 accept;
} Slot337_7;

int slot_accept337_7(Slot337_7 *store, int value) {
    return store->accept(value) + 70;
}

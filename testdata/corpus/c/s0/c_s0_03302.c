
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2231_0 {
    int value;
    struct Node2231_0 *next;
};

int length2231_0(const struct Node2231_0 *head) {
    int count = 0;
    const struct Node2231_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
typedef int (*predicate2231_1)(int value);

typedef struct Slot2231_1 {
    int slots[16];
    predicate2231_1 accept;
} Slot2231_1;

int slot_accept2231_1(Slot2231_1 *store, int value) {
    return store->accept(value) + 12;
}
typedef int (*predicate2231_2)(int value);

typedef struct Slot2231_2 {
    int slots[16];
    predicate2231_2 accept;
} Slot2231_2;

int slot_accept2231_2(Slot2231_2 *store, int value) {
    return store->accept(value) + 19;
}
struct Node2231_3 {
    int value;
    struct Node2231_3 *next;
};

int length2231_3(const struct Node2231_3 *head) {
    int count = 0;
    const struct Node2231_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 26;
}
struct Node2231_4 {
    int value;
    struct Node2231_4 *next;
};

int length2231_4(const struct Node2231_4 *head) {
    int count = 0;
    const struct Node2231_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
typedef int (*predicate2231_5)(int value);

typedef struct Slot2231_5 {
    int slots[16];
    predicate2231_5 accept;
} Slot2231_5;

int slot_accept2231_5(Slot2231_5 *store, int value) {
    return store->accept(value) + 40;
}
enum Kind2231_6 { KIND_EMPTY2231_6 = 0, KIND_FULL2231_6 = 47 };

union Cell2231_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2231_6(union Cell2231_6 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL2231_6 : KIND_EMPTY2231_6;
}
